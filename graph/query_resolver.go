package graph

import (
	"context"
	"fmt"
	"graphqltest/auth"
	"graphqltest/graph/generated"
	"graphqltest/models"
)

func (r *queryResolver) Documents(ctx context.Context) ([]*models.Document, error) {
	return r.DocumentRepo.GetDocuments()
}

func (r *queryResolver) Users(ctx context.Context) ([]*models.User, error) {
	return r.UserRepo.GetUsers()
}

func (r *queryResolver) ValidateToken(ctx context.Context, input string) (bool, error) {
	err := auth.IsValid(input)

	if err != nil {
		fmt.Println("Token is not valid", err)
		return false, err
	}
	return true, nil
}

func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }


