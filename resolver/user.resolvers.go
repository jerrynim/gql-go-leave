package resolver

import (
	"context"
	"fmt"

	"github.com/jerrynim/gql-leave/graph/model"
)

func (r *mutationResolver) SignUp(ctx context.Context, email string, password string, name string, bio *string, profileImage string, birthday *string) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetUsers(ctx context.Context) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}