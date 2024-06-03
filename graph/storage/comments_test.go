package storage

import (
	"context"
	"testing"

	"github.com/eeQuillibrium/posts/graph/model"
	"github.com/stretchr/testify/assert"
)

func TestCreateComment(t *testing.T) {
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
	for i := 0; i < len(users); i++ {
		st.Users.CreateUser(ctx, users[i])
	}
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
		comment *model.NewComment
		outID   int
	}{
		{
			comment: &model.NewComment{
				Text:     "some text",
				Level:    1,
				UserID:   1,
				PostID:   2,
				ParentID: nil,
			},
			outID: 1,
		},
		{
			comment: &model.NewComment{
				Text:     "texttext",
				Level:    2,
				UserID:   1,
				PostID:   2,
				ParentID: &[]int{1}[0],
			},
			outID: 2,
		},
	}

	for _, test := range table {
		outID, err := st.CreateComment(ctx, test.comment)
		if err != nil {
			t.Error(err)
		}
		assert.Equal(t, test.outID, outID)
	}
}

func TestGetPostComments(t *testing.T) {
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
	for i := 0; i < len(users); i++ {
		st.Users.CreateUser(ctx, users[i])
	}
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
	comments := []*model.NewComment{
		{
			Text:     "some text",
			Level:    1,
			UserID:   1,
			PostID:   2,
			ParentID: nil,
		},
		{
			Text:     "texttext",
			Level:    2,
			UserID:   1,
			PostID:   2,
			ParentID: &[]int{1}[0],
		},
		{
			Text:     "supertext",
			Level:    3,
			UserID:   2,
			PostID:   2,
			ParentID: &[]int{2}[0],
		},
		{
			Text:     "bambam",
			Level:    1,
			UserID:   2,
			PostID:   1,
			ParentID: nil,
		},
	}
	for i := 0; i < len(comments); i++ {
		st.Comments.CreateComment(ctx, comments[i])
	}
	table := []struct {
		postID   int
		expected []*model.Comment
	}{
		{
			postID: 2,
			expected: []*model.Comment{
				{
					Text:     "some text",
					Level:    1,
					UserID:   1,
					PostID:   2,
					ParentID: nil,
				},
				{
					Text:     "texttext",
					Level:    2,
					UserID:   1,
					PostID:   2,
					ParentID: &[]int{1}[0],
				},
				{
					Text:     "supertext",
					Level:    3,
					UserID:   2,
					PostID:   2,
					ParentID: &[]int{2}[0],
				},
			},
		},
		{
			postID: 1,
			expected: []*model.Comment{
				{
					Text:     "bambam",
					Level:    1,
					UserID:   2,
					PostID:   1,
					ParentID: nil,
				},
			},
		},
	}

	for _, test := range table {
		out, err := st.GetPostComments(ctx, test.postID)
		if err != nil {
			t.Error(err)
		}
		if len(out) != len(test.expected) {
			t.Log(out, test.expected)
			t.Errorf("len")
		}
		for i := 0; i < len(out); i++ {
			assert.Equal(t, test.expected[i].Text, out[i].Text)
			assert.Equal(t, test.expected[i].Level, out[i].Level)
			assert.Equal(t, test.expected[i].UserID, out[i].UserID)
			assert.Equal(t, test.expected[i].PostID, out[i].PostID)
			assert.Equal(t, test.expected[i].ParentID, out[i].ParentID)
		}
	}
}

func TestGetChildLevel(t *testing.T) {
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
	for i := 0; i < len(users); i++ {
		st.Users.CreateUser(ctx, users[i])
	}
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
	comments := []*model.NewComment{
		{
			Text:     "some text",
			Level:    1,
			UserID:   1,
			PostID:   2,
			ParentID: nil,
		},
		{
			Text:     "texttext",
			Level:    2,
			UserID:   1,
			PostID:   2,
			ParentID: &[]int{1}[0],
		},
		{
			Text:     "text2",
			Level:    2,
			UserID:   1,
			PostID:   2,
			ParentID: &[]int{1}[0],
		},
		{
			Text:     "supertext",
			Level:    3,
			UserID:   2,
			PostID:   2,
			ParentID: &[]int{2}[0],
		},
		{
			Text:     "bambam",
			Level:    1,
			UserID:   2,
			PostID:   1,
			ParentID: nil,
		},
		{
			Text:     "bimibm",
			Level:    2,
			UserID:   1,
			PostID:   1,
			ParentID: &[]int{5}[0],
		},
	}
	for i := 0; i < len(comments); i++ {
		st.Comments.CreateComment(ctx, comments[i])
	}

	table := []struct {
		commentID int
		expected  []*model.Comment
	}{
		{
			commentID: 1,
			expected: []*model.Comment{
				{
					Text:     "texttext",
					Level:    2,
					UserID:   1,
					PostID:   2,
					ParentID: &[]int{1}[0],
				},
				{
					Text:     "text2",
					Level:    2,
					UserID:   1,
					PostID:   2,
					ParentID: &[]int{1}[0],
				},
			},
		},
		{
			commentID: 5,
			expected: []*model.Comment{
				{
					Text:     "bimibm",
					Level:    2,
					UserID:   1,
					PostID:   1,
					ParentID: &[]int{5}[0],
				},
			},
		},
	}

	for _, test := range table {
		out, err := st.Comments.GetChildLevel(ctx, test.commentID)
		if err != nil {
			t.Error(err)
		}
		if len(out) != len(test.expected) {
			t.Log(out)
			t.Log(test.expected)
			t.Errorf("len")
		}
		for i := 0; i < len(out); i++ {
			assert.Equal(t, test.expected[i].Text, out[i].Text)
			assert.Equal(t, test.expected[i].Level, out[i].Level)
			assert.Equal(t, test.expected[i].UserID, out[i].UserID)
			assert.Equal(t, test.expected[i].PostID, out[i].PostID)
			assert.Equal(t, test.expected[i].ParentID, out[i].ParentID)
		}
	}
}
