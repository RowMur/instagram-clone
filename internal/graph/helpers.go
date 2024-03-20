package graph

import (
	"github/rowmur/insta-clone/internal/database"
	"github/rowmur/insta-clone/internal/graph/model"
)

func dbUserToGqlUser(user database.User) model.User {
	return model.User{
		ID:        user.ID.String(),
		Name:      user.Name,
		CreatedAt: user.CreatedAt.String(),
		UpdatedAt: user.UpdatedAt.String(),
		APIKey:    user.ApiKey,
	}
}