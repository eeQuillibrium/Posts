package graph

import (
	"testing"

	"github.com/99designs/gqlgen/client"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/eeQuillibrium/posts/graph/model"
	"github.com/eeQuillibrium/posts/internal/service"
	service_mocks "github.com/eeQuillibrium/posts/internal/service/mocks"
	"github.com/eeQuillibrium/posts/pkg/logger"
	"github.com/stretchr/testify/require"
	//"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {
	type respAccessor struct {
		CreateUser int `json:"createUser"`
	}
	tests := []struct {
		name     string
		q        string
		expected int
		output   *respAccessor
	}{
		{
			name:     "OK",
			expected: 1,
			q: `
			mutation createUser {
				createUser(input: {login: "equillibrium", password:"qwerty123", name:"nikita"})
			}
			`,
			output: &respAccessor{1},
		},
		{
			name:     "OK",
			expected: 1,
			q: `
			mutation createUser {
				createUser(input: {login: "balora", password:"qwerty123", name:"matvey"})
			}
			`,
			output: &respAccessor{1},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			repo := &service_mocks.MockAuthRepository{}

			services := &service.Service{Auth: repo}
			notifyChan := make(chan *model.Notification)
			defer close(notifyChan)

			cl := client.New(handler.NewDefaultServer(NewExecutableSchema(Config{Resolvers: NewResolver(services, logger.NewLogger(), notifyChan)})))

			cl.MustPost(test.q, &test.output)

			require.Equal(t, test.expected, test.output.CreateUser)
		})
	}
}
func TestCreatePost(t *testing.T) {
	type respAccessor struct {
		CreatePost int `json:"createPost"`
	}

	tests := []struct {
		name     string
		q        string
		expected int
		output   *respAccessor
	}{
		{
			name:     "OK",
			expected: 1,
			q: `
			mutation createPost {
				createPost(input:{header: "how to write", text: "simple", userId: 1})  
			}
			`,
			output: &respAccessor{1},
		},
		{
			name:     "OK",
			expected: 1,
			q: `
			mutation createPost {
				createPost(input:{header: "how to write", text: "simple", userId: 1})  
			}
			`,
			output: &respAccessor{1},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			repo := &service_mocks.MockPostsRepository{}

			services := &service.Service{Posts: repo}
			notifyChan := make(chan *model.Notification)
			defer close(notifyChan)

			cl := client.New(handler.NewDefaultServer(NewExecutableSchema(Config{Resolvers: NewResolver(services, logger.NewLogger(), notifyChan)})))

			cl.MustPost(test.q, &test.output)

			require.Equal(t, test.expected, test.output.CreatePost)
		})
	}
}

func TestCreateComment(t *testing.T) {
	type respAccessor struct {
		CreateComment int `json:"createComment"`
	}

	tests := []struct {
		name     string
		q        string
		expected int
		output   *respAccessor
	}{
		{
			name:     "OK",
			expected: 1,
			q: `
			mutation createComment {
				createComment(input:{text:"cool", userId:1, postId:1, level: 1})  
			}
			`,
			output: &respAccessor{1},
		},
		{
			name:     "OK",
			expected: 1,
			q: `
			mutation createComment {
				createComment(input:{text:"bright", userId:1, postId:1, level: 2, parentId: 1})
			}
			`,
			output: &respAccessor{1},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			repo := &service_mocks.MockCommentsRepository{}

			services := &service.Service{Comments: repo}
			notifyChan := make(chan *model.Notification)

			cl := client.New(handler.NewDefaultServer(NewExecutableSchema(Config{Resolvers: NewResolver(services, logger.NewLogger(), notifyChan)})))

			cl.MustPost(test.q, &test.output)

			require.Equal(t, test.expected, test.output.CreateComment)
		})
	}
}
