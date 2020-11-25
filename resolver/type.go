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
//! 여기가 에러난다면 Mutation 중 타입이 잘못 된게 있다는 뜻
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
//! 여기가 에러난다면 Query 중 타입이 잘못 된게 있다는 뜻
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
