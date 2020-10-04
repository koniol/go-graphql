package graph

import (
	"context"
	"graphqltest/graph/generated"
	"graphqltest/graph/model"
	"graphqltest/models"
)

func (r *userResolver) Documents(ctx context.Context, obj *models.User) ([]*models.Document, error) {
	return r.UserRepo.GetDocumentsByUserId(obj.ID)
}

func (r *queryResolver) GetUserByID(ctx context.Context, input model.GetUser) (*models.User, error) {
	return r.UserRepo.GetUserByID(input)
}

func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type userResolver struct{ *Resolver }
