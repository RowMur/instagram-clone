package helpers

import (
	"fmt"
	"github/rowmur/insta-clone/internal/database"
	"github/rowmur/insta-clone/internal/graph/model"

	"github.com/google/uuid"
)

func DBUserToGqlCurrentUser(user database.User) model.CurrentUser {
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

func DBUserToGqlUser(user database.User) model.User {
	return model.User{
		ID:        user.ID.String(),
		Name:      user.Name,
		CreatedAt: user.CreatedAt.String(),
		UpdatedAt: user.UpdatedAt.String(),
	}
}

func DBUsersToGqlUsers(dbUsers []database.User) []*model.User {
	users := []*model.User{}
	for _, dbUser := range dbUsers {
		user := DBUserToGqlUser(dbUser)
		users = append(users, &user)
	}

	return users
}

func DBPostToGqlPost(dbPost database.Post, dbUser database.User) model.Post {
	user := DBUserToGqlUser(dbUser)
	return model.Post{
		ID:        dbPost.ID.String(),
		CreatedAt: dbPost.CreatedAt.String(),
		UpdatedAt: dbPost.UpdatedAt.String(),
		User:      &user,
		Text:      dbPost.PostText,
	}
}

func StringsToUUIDs(strings []string) ([]uuid.UUID, error) {
	UUIDs := []uuid.UUID{}
	for _, s := range strings {
		userUUID, err := uuid.Parse(s)
		if err != nil {
			return nil, fmt.Errorf("invalid user ID")
		}
		UUIDs = append(UUIDs, userUUID)
	}

	return UUIDs, nil
}
