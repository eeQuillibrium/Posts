package storage

import (
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
func (s *PostStorage) PaginationComments(postID int, offset int, limit int) []*model.Comment {
	s.mu.Lock()
	defer s.mu.Unlock()

	if offset > len(s.pc[postID]) {
		return nil
	}

	if len(s.pc[postID]) > offset+limit {
		return s.pc[postID][offset:limit]
	}

	return s.pc[postID][offset:]
}
