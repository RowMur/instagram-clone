package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.45

import (
	"context"
	"github/rowmur/insta-clone/internal/database"
	"github/rowmur/insta-clone/internal/graph/model"
	"time"

	"github.com/google/uuid"
)

// CreateUser is the resolver for the CreateUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, name string) (*model.User, error) {
	currentTime := time.Now()
	dbUser, err := r.DBQueries.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: currentTime,
		Name:      name,
	})
	if err != nil {
		return nil, err
	}

	user := dbUserToGqlUser(dbUser)
	return &user, nil
}

// Users is the resolver for the Users field.
func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	dbUsers, err := r.DBQueries.GetUsers(context.Background())
	if err != nil {
		return nil, err
	}

	users := []*model.User{}
	for _, dbUser := range dbUsers {
		user := dbUserToGqlUser(dbUser)
		users = append(users, &user)
	}

	return users, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
