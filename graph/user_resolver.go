package graph

import (
	"context"
	"graphqltest/graph/generated"
	"graphqltest/models"
)

func (r *userResolver) Documents(ctx context.Context, obj *models.User) ([]*models.Document, error) {
	return r.UserRepo.GetDocumentsByUserId(obj.ID)
}

func (r *resolvers.queryResolver) GetUserByID(ctx context.Context, input string) (*models.User, error) {
	return r.UserRepo.GetUserByID(input)
}

func (r *graph.Resolver) User() generated.UserResolver { return &userResolver{r} }

type userResolver struct{ *graph.Resolver }
