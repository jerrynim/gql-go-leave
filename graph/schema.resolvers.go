package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/jerrynim/gql-leave/graph/generated"
	"github.com/jerrynim/gql-leave/graph/model"
)

func (r *mutationResolver) SignUp(ctx context.Context, email string, password string, name string, bio *string, department string, position string, workSpace string, contact string, birthday string, enteredDate string, remainLeaves int) (*model.AuthResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) Login(ctx context.Context, email string, password string) (*model.AuthResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) Me(ctx context.Context) (*model.AuthResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) MakeLeaveHistory(ctx context.Context, date string, reason *string, typeArg model.LeaveType) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) ChangeLeaveStatus(ctx context.Context, leaveID int, status model.LeaveStatus) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetMyLeaves(ctx context.Context) ([]*model.LeaveHistory, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetAppliedLeaves(ctx context.Context) ([]*model.LeaveHistory, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetUsers(ctx context.Context) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *mutationResolver) ApproveLeave(ctx context.Context, leaveID int) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}
