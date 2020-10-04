package resolvers

import (
	"context"
	"fmt"
	"graphqltest/graph/generated"
	"graphqltest/graph/model"
	"graphqltest/models"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*models.User, error) {
	return r.UserRepo.CreateUser(input)
}

func (r *mutationResolver) DeleteUser(ctx context.Context, input int) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateUser(ctx context.Context, input *model.NewUser) (*models.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *graph.Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct {
	*graph.Resolver
}
