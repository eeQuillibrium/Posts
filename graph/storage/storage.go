package storage

import (
	"errors"
	"fmt"
	"sync"

	"github.com/eeQuillibrium/posts/graph/model"
)

type PostStorage struct {
	mu sync.Mutex
	pc map[int][]*model.Comment
}

func NewPostStorage() *PostStorage {
	return &PostStorage{
		pc: make(map[int][]*model.Comment),
	}
}

func (s *PostStorage) LoadPost(comments []*model.Comment, postID int) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.pc[postID] = comments
}
func (s *PostStorage) PaginationComments(postID int, offset int, limit int) ([]*model.Comment, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.pc[postID]; !ok { //если не подгружен пост
		return nil, errors.New("PostStorage.PaginationComments(): post with id:" +
			fmt.Sprintf("%d", postID) + " didn't loaded, execute the queryResolver.PostStorage.LoadPost() before")
	}
	if offset > len(s.pc[postID]) {
		return nil, nil
	}

	if len(s.pc[postID]) > offset+limit {
		return s.pc[postID][offset:limit], nil
	}

	return s.pc[postID][offset:], nil
}
