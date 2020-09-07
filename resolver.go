package main

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

import (
	"context"
	"graphqltest/graph/generated"
	"graphqltest/graph/model"
	"graphqltest/models"
)

type Resolver struct{}

func (r *documentResolver) Description(ctx context.Context, obj *models.Document) (string, error) {
	panic("not implemented")
}

func (r *documentResolver) User(ctx context.Context, obj *models.Document) (*models.User, error) {
	panic("not implemented")
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*models.User, error) {
	panic("not implemented")
}

func (r *mutationResolver) DeleteUser(ctx context.Context, input int) (bool, error) {
	panic("not implemented")
}

func (r *mutationResolver) UpdateUser(ctx context.Context, input *model.NewUser) (*models.User, error) {
	panic("not implemented")
}

func (r *queryResolver) Users(ctx context.Context) ([]*models.User, error) {
	panic("not implemented")
}

func (r *queryResolver) GetUserByID(ctx context.Context, input string) (*models.User, error) {
	panic("not implemented")
}

func (r *queryResolver) Documents(ctx context.Context) ([]*models.Document, error) {
	panic("not implemented")
}

func (r *userResolver) Documents(ctx context.Context, obj *models.User) ([]*models.Document, error) {
	panic("not implemented")
}

// Document returns generated.DocumentResolver implementation.
func (r *Resolver) Document() generated.DocumentResolver { return &documentResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type documentResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
