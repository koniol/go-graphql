package graph

import (
	"context"
	"fmt"
	"graphqltest/auth"
	"graphqltest/graph/generated"
	"graphqltest/graph/model"
	"graphqltest/models"
	"strconv"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.Auth, error) {
	var authData *model.Auth
	user, err := r.UserRepo.CreateUser(input)
	if err != nil {
		fmt.Println("Cant create user ", err)
		return authData, err
	}
	id, convError := strconv.Atoi(user.ID)

	if convError != nil {
		fmt.Println("Cant convert to int", err)
		return authData, err
	}

	aUser := auth.UserToken{
		Id:      id,
		Email:   user.Email,
		IsAdmin: false,
	}

	token, tokenError := auth.CreateToken(&aUser)
	if tokenError != nil {
		fmt.Println("Error create token", err)
	}
	authData = &model.Auth{
		User: user,
		Token: token,
	}

	return authData, err
}

func (r *mutationResolver) DeleteUser(ctx context.Context, input int) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateUser(ctx context.Context, input *model.NewUser) (*models.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct {
	*Resolver
}
