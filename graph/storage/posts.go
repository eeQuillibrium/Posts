package storage

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/eeQuillibrium/posts/graph/model"
)

const (
	preparedPostsVolume = 10000
	maxLimit            = 100
)

var (
	errNoPost = errors.New("post with this id doesn't exist")
)

type postsStorage struct {
	posts    map[int]*model.Post
	mu       *sync.Mutex
	idSerial int
	mediator Mediator
}

func NewPostsStorage() *postsStorage {
	return &postsStorage{
		posts:    make(map[int]*model.Post, preparedPostsVolume),
		idSerial: 1,
		mu:       &sync.Mutex{},
	}
}

func (ps *postsStorage) CreatePost(
	ctx context.Context,
	newPost *model.NewPost,
) (int, error) {

	if err := ps.mediator.Notify(
		eventIsUserExist + eventSeparator +
			strconv.Itoa(newPost.UserID)); err != nil {
		return 0, fmt.Errorf("postsStorage.CreatePost():\n%w", err)
	}

	ps.mu.Lock()
	defer ps.mu.Unlock()

	post := &model.Post{
		ID:        ps.idSerial,
		UserID:    newPost.UserID,
		Text:      newPost.Text,
		Header:    newPost.Header,
		CreatedAt: time.Now().GoString(),
		Comments:  []*model.Comment{},
		IsClosed:  false,
	}
	ps.posts[ps.idSerial] = post

	ps.idSerial++

	return post.ID, nil
}

func (ps *postsStorage) ClosePost(
	ctx context.Context,
	postID int,
) (bool, error) {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	if _, ok := ps.posts[postID]; !ok {
		return false, fmt.Errorf("postsStorage.ClosePost(): %w: id = %d", errNoPost, postID)
	}

	ps.posts[postID].IsClosed = true

	return true, nil
}
func (ps *postsStorage) GetPosts(
	ctx context.Context,
	offset int,
	limit int,
) ([]*model.Post, error) { // don't return err
	if limit > maxLimit {
		return nil, nil
	}

	ps.mu.Lock()
	defer ps.mu.Unlock()

	i := ps.idSerial - offset
	fn := i + limit
	posts := make([]*model.Post, 0, limit)

	if i < 0 {
		return nil, nil
	}

	for ; i < fn; i++ {
		post, ok := ps.posts[i]
		if !ok {
			break
		}
		posts = append(posts, post)
	}

	return posts, nil
}
func (ps *postsStorage) GetPost(
	ctx context.Context,
	postID int,
) (*model.Post, error) {
	post, ok := ps.posts[postID]
	if !ok {
		return nil, fmt.Errorf("GetPost(): %w", errNoPost)
	}
	return post, nil
}

func (ps *postsStorage) SetMediator(mediator Mediator) {
	ps.mediator = mediator
}

func (ps *postsStorage) isPostExist(postID int) error {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	if _, ok := ps.posts[postID]; !ok {
		return fmt.Errorf("isPostExist(): %w", errNoPost)
	}

	return nil
}
