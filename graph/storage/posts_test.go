package storage

import (
	"context"
	"testing"

	"github.com/eeQuillibrium/posts/graph/model"
	"github.com/stretchr/testify/assert"
)

func TestCreatePost(t *testing.T) {
	ctx := context.Background()

	st := NewStorage()
	users := []*model.NewUser{
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
	st.CreateUser(ctx, users[0])
	st.CreateUser(ctx, users[0])

	table := []struct {
		newPost *model.NewPost
		want    int
	}{
		{
			newPost: &model.NewPost{
				Header: "Hello",
				Text:   "my name is Nikita",
				UserID: 1,
			},
			want: 1,
		},
		{
			newPost: &model.NewPost{
				Header: "Bye",
				Text:   "my name is Andrew",
				UserID: 1,
			},
			want: 2,
		},
	}

	for _, test := range table {
		out, err := st.Posts.CreatePost(ctx, test.newPost)
		if err != nil {
			t.Error(err)
		}
		assert.Equal(t, test.want, out)
	}

}

func TestClosePost(t *testing.T) {
	ctx := context.Background()

	st := NewStorage()
	users := []*model.NewUser{
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
	st.CreateUser(ctx, users[0])
	st.CreateUser(ctx, users[1])
	st.CreatePost(ctx, &model.NewPost{
		Header: "Hello",
		Text:   "my name is Nikita",
		UserID: 1,
	})
	st.CreatePost(ctx, &model.NewPost{
		Header: "Bye",
		Text:   "my name is Andrew",
		UserID: 1,
	})

	table := []struct {
		postID int
		want   bool
	}{
		{
			postID: 1,
			want:   true,
		},
		{
			postID: 2,
			want:   true,
		},
	}

	for _, test := range table {
		out, err := st.Posts.ClosePost(ctx, test.postID)
		if err != nil {
			t.Error(err)
		}
		assert.Equal(t, test.want, out)
	}
}

func TestGetPost(t *testing.T) {
	ctx := context.Background()

	st := NewStorage()
	users := []*model.NewUser{
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
	st.CreateUser(ctx, users[0])
	st.CreateUser(ctx, users[1])

	posts := []*model.NewPost{
		{
			Header: "Hello",
			Text:   "my name is Nikita",
			UserID: 1,
		},
		{
			Header: "Bye",
			Text:   "my name is Andrew",
			UserID: 1,
		},
	}

	st.CreatePost(ctx, posts[0])
	st.CreatePost(ctx, posts[1])

	table := []struct {
		postID int
		want   model.Post
	}{
		{
			postID: 1,
			want: model.Post{
				ID:     1,
				UserID: 1,
				Text:   posts[0].Text,
				Header: posts[0].Header,
			},
		},
		{
			postID: 2,
			want: model.Post{
				ID:     2,
				UserID: 1,
				Text:   posts[1].Text,
				Header: posts[1].Header,
			},
		},
	}

	for _, test := range table {
		out, err := st.Posts.GetPost(ctx, test.postID)
		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, test.want.ID, out.ID)
		assert.Equal(t, test.want.UserID, out.UserID)
		assert.Equal(t, test.want.Text, out.Text)
		assert.Equal(t, test.want.Header, out.Header)
	}
}
func TestGetPosts(t *testing.T) {
	ctx := context.Background()

	st := NewStorage()
	users := []*model.NewUser{
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
	st.Users.CreateUser(ctx, users[0])
	st.Users.CreateUser(ctx, users[1])

	posts := []*model.NewPost{
		{
			Header: "Hello",
			Text:   "my name is Nikita",
			UserID: 1,
		},
		{
			Header: "Bye",
			Text:   "my name is Andrew",
			UserID: 1,
		},
		{
			Header: "Alright",
			Text:   "my name is Mikhail",
			UserID: 1,
		},
	}

	for i := 0; i < len(posts); i++ {
		st.Posts.CreatePost(ctx, posts[i])
	}

	table := []struct {
		limit  int
		offset int
		want   []model.Post
	}{
		{
			offset: 1,
			limit:  2,
			want: []model.Post{
				{
					Header: "Alright",
					Text:   "my name is Mikhail",
					UserID: 1,
				},
			},
		},
		{
			offset: 2,
			limit:  10,
			want: []model.Post{
				{
					Header: "Bye",
					Text:   "my name is Andrew",
					UserID: 1,
				},
				{
					Header: "Alright",
					Text:   "my name is Mikhail",
					UserID: 1,
				},
			},
		},
	}

	for _, test := range table {
		out, err := st.Posts.GetPosts(ctx, test.offset, test.limit)
		if err != nil {
			t.Error(err)
		}
		if len(out) != len(test.want) {
			t.Errorf("Not compitable: %v", err)
		}
		t.Log(test.want, out[0])
		for i, post := range out {
			assert.Equal(t, test.want[i].UserID, post.UserID)
			assert.Equal(t, test.want[i].Text, post.Text)
			assert.Equal(t, test.want[i].Header, post.Header)
		}

	}
}
