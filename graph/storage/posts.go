package storage

import (
	"context"
	"errors"
	"strconv"
	"sync"
	"time"

	"github.com/eeQuillibrium/posts/graph/model"
)

const preparedPostsVolume = 10000

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
	post *model.NewPost,
) (int, error) {
	if err := ps.mediator.Notify(
		eventIsUserExist + eventSeparator +
			strconv.Itoa(ps.idSerial)); err != nil {
		return 0, errors.New("postsStorage.CreatePost():\n" + err.Error())
	}

	ps.mu.Lock()
	defer ps.mu.Unlock()

	ps.posts[ps.idSerial] = &model.Post{
		ID:        ps.idSerial,
		UserID:    post.UserID,
		Text:      post.Text,
		Header:    post.Header,
		CreatedAt: time.Now().GoString(),
		Comments:  []*model.Comment{},
		IsClosed:  false,
	}

	ps.idSerial++

	return ps.idSerial, nil
}

func (ps *postsStorage) ClosePost(
	ctx context.Context,
	postID int,
) (bool, error) {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	if _, ok := ps.posts[postID]; !ok {
		return false, errors.New("postsStorage.ClosePost():" + "post with this id doesn't exits")
	}

	ps.posts[postID].IsClosed = true

	return true, nil
}
func (ps *postsStorage) GetPosts(
	ctx context.Context,
	offset int,
	limit int,
) ([]*model.Post, error) {
	posts := make([]*model.Post, offset+limit)

	i := offset + 1
	for ; i < offset+limit; i++ {
		ps.mu.Lock()
		post, ok := ps.posts[i]
		if !ok {
			ps.mu.Unlock()
			break
		}
		posts = append(posts, post)
		ps.mu.Unlock()
	}

	return posts[:i-offset], nil
}
func (ps *postsStorage) GetPost(
	ctx context.Context,
	postID int,
) (*model.Post, error) {
	return nil, nil
}

func (ps *postsStorage) SetMediator(mediator Mediator) {
	ps.mediator = mediator
}

func (ps *postsStorage) isPostExist(postID int) error {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	if _, ok := ps.posts[postID]; !ok {
		return errors.New("isPostExist(): post with this postID doesn't exist")
	}

	return nil
}
