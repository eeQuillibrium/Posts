package storage

import (
	"errors"
	"strconv"
	"sync"

	"github.com/eeQuillibrium/posts/graph/model"
)

// for postgres caching
type PostCacheStorage struct {
	mu sync.Mutex
	pc map[int][]*model.Comment
}

func NewPostCacheStorage() *PostCacheStorage {
	return &PostCacheStorage{
		pc: make(map[int][]*model.Comment),
	}
}

func (pcs *PostCacheStorage) LoadPost(comments []*model.Comment, postID int) {
	pcs.mu.Lock()
	defer pcs.mu.Unlock()

	pcs.pc[postID] = comments
}
func (pcs *PostCacheStorage) PaginationComments(postID int, offset int, limit int) ([]*model.Comment, error) {
	pcs.mu.Lock()
	defer pcs.mu.Unlock()

	if _, ok := pcs.pc[postID]; !ok { //если не подгружен пост
		return nil, errors.New("PostStorage.PaginationComments(): post with id:" +
			strconv.Itoa(postID) + " didn't load, execute the queryResolver.PostStorage.LoadPost() before")
	}
	
	if offset > len(pcs.pc[postID]) {
		return nil, nil
	}

	if len(pcs.pc[postID]) > offset+limit {
		return pcs.pc[postID][offset:limit], nil
	}

	return pcs.pc[postID][offset:], nil
}
