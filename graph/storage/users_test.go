package storage

import (
	"context"
	"testing"

	"github.com/eeQuillibrium/posts/graph/model"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	st := NewStorage()
	ctx := context.Background()

	table := []struct {
		in   *model.NewUser
		want int
	}{
		{
			in: &model.NewUser{
				Name:     "Nikita",
				Login:    "eqillibrium",
				Password: "qwerty",
			},
			want: 1,
		},
		{
			in: &model.NewUser{
				Name:     "Andrew",
				Login:    "Hlebushek",
				Password: "qwerty123",
			},
			want: 2,
		},
	}

	for _, test := range table {
		out, err := st.Users.CreateUser(ctx, test.in)
		if err != nil {
			t.Error(err)
		}
		assert.Equal(t, test.want, out)
	}

}
