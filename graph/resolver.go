package graph

import "github.com/jerrynim/gql-leave/graph/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	users []*model.User
	leaveHistories  []*model.LeaveHistory
}
