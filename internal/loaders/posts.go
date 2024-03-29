package loaders

import (
	"context"
	"github/rowmur/insta-clone/internal/database"
	"github/rowmur/insta-clone/internal/graph/model"
	"github/rowmur/insta-clone/internal/helpers"
)

type userPostsReader struct {
	dbQueries *database.Queries
}

func (up *userPostsReader) getPostsByUserIDs(ctx context.Context, userIDs []string) ([][]*model.Post, []error) {
	userUUIDs, err := helpers.StringsToUUIDs(userIDs)
	if err != nil {
		return nil, []error{err}
	}

	dbPosts, err := up.dbQueries.GetPostsFromUsers(ctx, userUUIDs)
	if err != nil {
		return nil, []error{err}
	}

	userIDsMapToIndex := make(map[string]int)
	for index, userID := range userIDs {
		userIDsMapToIndex[userID] = index
	}

	userPosts := make([][]*model.Post, len(userIDs))
	for index := range userPosts {
		userPosts[index] = []*model.Post{}
	}

	for _, dbPost := range dbPosts {
		userIndex := userIDsMapToIndex[dbPost.UserID.String()]
		post := helpers.DBPostToGqlPost(dbPost)
		userPosts[userIndex] = append(userPosts[userIndex], &post)
	}

	return userPosts, nil
}

func GetUserPosts(ctx context.Context, userID string) ([]*model.Post, error) {
	loaders := ForContext(ctx)
	return loaders.UserPostsLoader.Load(ctx, userID)
}

func GetUsersPosts(ctx context.Context, userIDs []string) ([][]*model.Post, error) {
	loaders := ForContext(ctx)
	return loaders.UserPostsLoader.LoadAll(ctx, userIDs)
}
