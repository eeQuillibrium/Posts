package service_mocks

import (
	"context"
	"testing"

	"github.com/eeQuillibrium/posts/graph/model"
)

type MockAuthRepository struct{}

func (r *MockAuthRepository) Register(
	ctx context.Context,
	user *model.NewUser,
) (int, error) {
	return 1, nil
}
func TestRegister(t *testing.T) {
}
