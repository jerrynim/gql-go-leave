package resolver

import (
	"github.com/jerrynim/gql-leave/graph/generated"
	"github.com/jerrynim/gql-leave/graph/model"
)

type Resolver struct {
	users []*model.User
	leaveHistories  []*model.LeaveHistory
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
