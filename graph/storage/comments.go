package storage

import (
	"context"
	"errors"
	"strconv"
	"sync"
	"time"

	"github.com/eeQuillibrium/posts/graph/model"
)

const (
	preparedCommentsVolume = 50000
	preparedLevelVolume    = 50
)

// комментарии по level
type commentNode struct {
	child    *commentNode
	comments map[int][]*model.Comment
}

type commentsStorage struct {
	comments     map[int]*model.Comment
	commentsPost map[int][]*model.Comment
	idSerial     int
	root         *commentNode
	mu           *sync.Mutex
	mediator     Mediator
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
		return 0, errors.New("commentsStorage.CreateComment():\n" + err.Error())
	}
	if newComment.ParentID != nil {
		if _, ok := cs.comments[*newComment.ParentID]; !ok {
			return 0, errors.New("commentsStorage.CreateComment(): " + "comment with this parentID doesn't exist")
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

	if comment.Level == 1 {
		cs.root.comments[0] = append(cs.root.comments[0], comment)
		cs.idSerial++
		return cs.idSerial, nil
	}

	cs.commentsPost[comment.PostID] = append(cs.commentsPost[comment.PostID], comment)

	curr := cs.root
	for i := 1; i < comment.Level; i++ {
		curr = curr.child
	}

	if curr == nil {
		curr = &commentNode{
			comments: make(map[int][]*model.Comment, preparedLevelVolume),
		}
		curr.comments[*comment.ParentID] = []*model.Comment{comment}
		return cs.idSerial, nil
	}

	(*curr).comments[*comment.ParentID] = append(curr.comments[*comment.ParentID], comment)

	return cs.idSerial, nil
}

func (cs *commentsStorage) GetComments(
	ctx context.Context,
	postID int,
) ([]*model.Comment, error) {
	if err := cs.mediator.Notify(eventIsPostExist + eventSeparator +
		strconv.Itoa(postID)); err != nil {
		return nil, errors.New("commentsStorage.CreateComment():\n" + err.Error())
	}
	return cs.commentsPost[postID], nil
}

func (cs *commentsStorage) GetByComment(
	ctx context.Context,
	commentID int,
) ([]*model.Comment, error) {
	cs.mu.Lock()
	defer cs.mu.Unlock()

	curr := cs.root

	for i := 0; i < cs.comments[commentID].Level; i++ {
		curr = curr.child
	}

	if curr == nil {
		return nil, nil
	}

	return curr.comments[commentID], nil
}


