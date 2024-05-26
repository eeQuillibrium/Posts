package graph

import (
	"testing"

	"github.com/99designs/gqlgen/client"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/eeQuillibrium/posts/internal/service"
	service_mocks "github.com/eeQuillibrium/posts/internal/service/mocks"
	//"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {

	tests := []struct {
		name     string
		q        string
		expected int
	}{
		{
			name:     "OK",
			expected: 1,
			q: `
			mutation createUser {
				createUser(input: {login: "equillibrium", password:"qwerty123", name:"nikita"})
			}
			`,
		},
		{
			name:     "OK",
			expected: 2,
			q: `
			mutation createUser {
				createUser(input: {login: "balora", password:"qwerty123", name:"matvey"})
			}
			`,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			repo := &service_mocks.MockAuthRepository{}

			services := &service.Service{Auth: repo}
			resolvers := Resolver{service: services}

			cl := client.New(handler.NewDefaultServer(NewExecutableSchema(Config{Resolvers: &resolvers})))

			var id int
			cl.MustPost(test.q, &id)
			//require.Equal(t, test.expected, resp)
		})
	}
}
