package graph

import (
	"github/rowmur/insta-clone/internal/database"
	"github/rowmur/insta-clone/internal/graph/model"
)

func dbUserToGqlCurrentUser(user database.User) model.CurrentUser {
	return model.CurrentUser{
		ID:        user.ID.String(),
		Name:      user.Name,
		CreatedAt: user.CreatedAt.String(),
		UpdatedAt: user.UpdatedAt.String(),
		APIKey:    user.ApiKey,
	}
}

func dbUserToGqlUser(user database.User) model.OtherUser {
	return model.OtherUser{
		ID:        user.ID.String(),
		Name:      user.Name,
		CreatedAt: user.CreatedAt.String(),
		UpdatedAt: user.UpdatedAt.String(),
	}
}
