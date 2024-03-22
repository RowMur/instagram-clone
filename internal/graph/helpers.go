package graph

import (
	"github/rowmur/insta-clone/internal/database"
	"github/rowmur/insta-clone/internal/graph/model"
)

func dbUserToGqlCurrentUser(user database.User) model.CurrentUser {
	return model.CurrentUser{
		User: &model.User{
			ID:        user.ID.String(),
			Name:      user.Name,
			CreatedAt: user.CreatedAt.String(),
			UpdatedAt: user.UpdatedAt.String(),
		},
		APIKey: user.ApiKey,
	}
}

func dbUserToGqlUser(user database.User) model.User {
	return model.User{
		ID:        user.ID.String(),
		Name:      user.Name,
		CreatedAt: user.CreatedAt.String(),
		UpdatedAt: user.UpdatedAt.String(),
	}
}

func dbPostToGqlPost(dbPost database.Post, dbUser database.User) model.Post {
	user := dbUserToGqlUser(dbUser)
	return model.Post{
		ID:        dbPost.ID.String(),
		CreatedAt: dbPost.CreatedAt.String(),
		UpdatedAt: dbPost.UpdatedAt.String(),
		User:      &user,
		Text:      dbPost.PostText,
	}
}
