package graph

import (
	"context"
	"graphqltest/graph/generated"
	"graphqltest/models"
)

func (r *queryResolver) Documents(ctx context.Context) ([]*models.Document, error) {
	return r.DocumentRepo.GetDocuments()
}

func (r *queryResolver) Users(ctx context.Context) ([]*models.User, error) {
	return r.UserRepo.GetUsers()
}

func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
