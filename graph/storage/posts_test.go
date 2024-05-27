package storage

import (
	"context"
	"testing"

	"github.com/eeQuillibrium/posts/graph/model"
	"github.com/stretchr/testify/assert"
)

func setupStorage(ctx context.Context) *Storage {
	st := NewStorage()
	setup := []*model.NewUser{
		{
			Name:     "Nikita",
			Login:    "eqillibrium",
			Password: "qwerty",
		},
		{
			Name:     "Andrew",
			Login:    "krutoi",
			Password: "da",
		},
	}
	for i := 0; i < len(setup); i++ {
		st.Users.CreateUser(ctx, setup[i])
	}
	return st
}

func TestCreatePost(t *testing.T) {
	ctx := context.Background()

	st := setupStorage(ctx)

	table := []struct {
		in   *model.NewPost
		want int
	}{
		{
			in: &model.NewPost{
				Header: "Hello",
				Text:   "my name is Nikita",
				UserID: 1,
			},
			want: 1,
		},
		{
			in: &model.NewPost{
				Header: "Test",
				Text:   "today i test posts",
				UserID: 1,
			},
			want: 2,
		},
	}

	for _, test := range table {
		out, err := st.Posts.CreatePost(ctx, test.in)
		if err != nil {
			t.Error(err)
		}
		assert.Equal(t, test.want, out)
	}

}

func TestClosePost(t *testing.T) {
	ctx := context.Background()

	st := setupStorage(ctx)

	st.CreatePost(ctx, &model.NewPost{
		Header: "Hello",
		Text:   "my name is Nikita",
		UserID: 1,
	})

	table := []struct {
		in   int
		want bool
	}{
		{
			in: 1,
			want: true,
		},
		{
			in: 2,
			want: false,
		},
	}

	for _, test := range table {
		out, err := st.Posts.ClosePost(ctx, test.in)
		if err != nil {
			t.Error(err)
		}
		assert.Equal(t, test.want, out)
	}
}
