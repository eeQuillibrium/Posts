package storage

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"
	"sync"
	"time"

	"github.com/eeQuillibrium/posts/graph/model"
)

const (
	preparedCommentsVolume = 50000
	preparedLevelVolume    = 50
)

var (
	errNoParent = errors.New("comment with this parentID doesn't exist")
)

// комментарии по level
type commentNode struct {
	child    *commentNode
	comments map[int][]*model.Comment //parentID-childs
}

type commentsStorage struct {
	comments     map[int]*model.Comment   //commentID-comment
	commentsPost map[int][]*model.Comment //postID-[]comment
	idSerial     int
	root         *commentNode
	mu           *sync.Mutex
	mediator     Mediator
}

func (cs *commentsStorage) Print() {
	log.Println(cs.comments)
	log.Println(cs.commentsPost)
	curr := cs.root
	for curr != nil {
		log.Println(curr.comments)
		curr = curr.child
	}
}
func NewCommentsStorage() *commentsStorage {
	return &commentsStorage{
		mu:           &sync.Mutex{},
		root:         &commentNode{child: nil, comments: make(map[int][]*model.Comment, preparedLevelVolume)},
		commentsPost: make(map[int][]*model.Comment, preparedLevelVolume),
		comments:     make(map[int]*model.Comment, preparedCommentsVolume),
		idSerial:     1,
	}
}

func (cs *commentsStorage) SetMediator(mediator Mediator) {
	cs.mediator = mediator
}

func (cs *commentsStorage) CreateComment(
	ctx context.Context,
	newComment *model.NewComment,
) (int, error) {
	if err := cs.mediator.Notify(eventIsPostExist + eventSeparator +
		strconv.Itoa(newComment.PostID)); err != nil {
		return 0, fmt.Errorf("commentsStorage.CreateComment():\n%w", err)
	}
	if newComment.ParentID != nil {
		if _, ok := cs.comments[*newComment.ParentID]; !ok {
			return 0, fmt.Errorf("commentsStorage.CreateComment(): %w", errNoParent)
		}
	}

	cs.mu.Lock()
	defer cs.mu.Unlock()

	comment := &model.Comment{
		ID:        cs.idSerial,
		PostID:    newComment.PostID,
		ParentID:  newComment.ParentID,
		UserID:    newComment.UserID,
		Text:      newComment.Text,
		Level:     newComment.Level,
		CreatedAt: time.Now().GoString(),
		Comments:  []*model.Comment{},
	}

	cs.comments[cs.idSerial] = comment

	cs.idSerial++

	cs.commentsPost[comment.PostID] = append(cs.commentsPost[comment.PostID], comment)

	if comment.Level == 1 {
		cs.root.comments[0] = append(cs.root.comments[0], comment)
		return comment.ID, nil
	}

	curr := cs.root
	for i := 1; i < comment.Level-1; i++ {
		curr = curr.child
	}

	if curr.child == nil {
		curr.child = &commentNode{
			comments: make(map[int][]*model.Comment, preparedLevelVolume),
		}
		curr.child.comments[*comment.ParentID] = append(curr.child.comments[*comment.ParentID], comment)
		return comment.ID, nil
	}

	curr.child.comments[*comment.ParentID] = append(curr.child.comments[*comment.ParentID], comment)

	return comment.ID, nil
}

func (cs *commentsStorage) GetPostComments(
	ctx context.Context,
	postID int,
) ([]*model.Comment, error) {
	if err := cs.mediator.Notify(eventIsPostExist + eventSeparator +
		strconv.Itoa(postID)); err != nil {
		return nil, fmt.Errorf("commentsStorage.CreateComment():\n%w", err)
	}
	return cs.commentsPost[postID], nil
}

func (cs *commentsStorage) GetChildLevel(
	ctx context.Context,
	parentID int,
) ([]*model.Comment, error) {
	cs.mu.Lock()
	defer cs.mu.Unlock()

	curr := cs.root

	for i := 0; i < cs.comments[parentID].Level; i++ {
		curr = curr.child
	}

	if curr == nil {
		return nil, nil
	}

	return curr.comments[parentID], nil
}

func (cs *commentsStorage) PaginationComment(
	ctx context.Context,
	postID int,
	offset int,
	limit int,
) ([]*model.Comment, error) {
	defer cs.mu.Unlock()
	cs.mu.Lock()

	if offset > len(cs.commentsPost[postID]) {
		return nil, nil
	}

	if len(cs.commentsPost[postID]) > offset+limit {
		return cs.commentsPost[postID][offset:limit], nil
	}

	return cs.commentsPost[postID][offset:], nil
}
