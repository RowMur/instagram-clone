package loaders

import (
	"context"
	"github/rowmur/insta-clone/internal/database"
	"github/rowmur/insta-clone/internal/graph/model"
	"github/rowmur/insta-clone/internal/helpers"
)

type userReader struct {
	dbQueries *database.Queries
}

func (u *userReader) getUsers(ctx context.Context, userIDs []string) ([]*model.User, []error) {
	userUUIDs, err := helpers.StringsToUUIDs(userIDs)
	if err != nil {
		return nil, []error{err}
	}

	dbUsers, err := u.dbQueries.GetUsersByIds(ctx, userUUIDs)
	if err != nil {
		return nil, []error{err}
	}

	users := make([]*model.User, 0, len(userIDs))
	for _, dbUser := range dbUsers {
		user := helpers.DBUserToGqlUser(dbUser)
		users = append(users, &user)
	}

	return users, nil
}

func GetUser(ctx context.Context, userID string) (*model.User, error) {
	loaders := ForContext(ctx)
	return loaders.UserLoader.Load(ctx, userID)
}

func GetUsers(ctx context.Context, userIDs []string) ([]*model.User, error) {
	loaders := ForContext(ctx)
	return loaders.UserLoader.LoadAll(ctx, userIDs)
}
