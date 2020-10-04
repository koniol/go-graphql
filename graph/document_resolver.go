package graph

import (
	"context"
	"graphqltest/graph/generated"
	"graphqltest/models"
)

func (r *documentResolver) User(ctx context.Context, obj *models.Document) (*models.User, error) {

	return GetUserLoader(ctx).Load(obj.UserId)
}

func (d documentResolver) Description(ctx context.Context, obj *models.Document) (string, error) {
	panic("implement me")
}

func (r *Resolver) Document() generated.DocumentResolver { return &documentResolver{r} }

type documentResolver struct{ *Resolver }
