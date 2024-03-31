package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.45

import (
	"context"
	"fmt"
	"github/rowmur/insta-clone/internal/auth"
	"github/rowmur/insta-clone/internal/database"
	"github/rowmur/insta-clone/internal/graph/model"
	"github/rowmur/insta-clone/internal/helpers"
	"github/rowmur/insta-clone/internal/loaders"
	"time"

	"github.com/google/uuid"
)

// CreateUser is the resolver for the CreateUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, name string) (*model.CurrentUser, error) {
	currentTime := time.Now()
	dbUser, err := r.DBQueries.CreateUser(ctx, database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: currentTime,
		Name:      name,
	})
	if err != nil {
		return nil, err
	}

	user := helpers.DBUserToGqlCurrentUser(dbUser)
	return &user, nil
}

// Follow is the resolver for the follow field.
func (r *mutationResolver) Follow(ctx context.Context, userID string) (*model.User, error) {
	currentUser := auth.ForContext(ctx)
	if currentUser == nil {
		return nil, fmt.Errorf("access denied")
	}

	userToFollowGuid, err := uuid.Parse(userID)
	if err != nil {
		return nil, fmt.Errorf("invalid user ID")
	}

	_, err = r.DBQueries.CreateFollow(ctx, database.CreateFollowParams{
		UserID:          currentUser.ID,
		UserFollowingID: userToFollowGuid,
		CreatedAt:       time.Now(),
	})
	if err != nil {
		return nil, fmt.Errorf("something went wrong")
	}

	user, err := loaders.GetUser(ctx, userToFollowGuid.String())
	if err != nil {
		return nil, fmt.Errorf("something went wrong")
	}

	return user, nil
}

// Unfollow is the resolver for the unfollow field.
func (r *mutationResolver) Unfollow(ctx context.Context, userID string) (*string, error) {
	currentUser := auth.ForContext(ctx)
	if currentUser == nil {
		return nil, fmt.Errorf("access denied")
	}

	userToUnfollowGuid, err := uuid.Parse(userID)
	if err != nil {
		return nil, fmt.Errorf("invalid user ID")
	}

	err = r.DBQueries.Unfollow(ctx, database.UnfollowParams{
		UserID:          currentUser.ID,
		UserFollowingID: userToUnfollowGuid,
	})
	if err != nil {
		return nil, fmt.Errorf("something went wrong")
	}

	return &userID, nil
}

// CreatePost is the resolver for the createPost field.
func (r *mutationResolver) CreatePost(ctx context.Context, text string) (*model.Post, error) {
	currentUser := auth.ForContext(ctx)
	if currentUser == nil {
		return nil, fmt.Errorf("access denied")
	}

	dbPost, err := r.DBQueries.CreatePost(ctx, database.CreatePostParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UserID:    currentUser.ID,
		PostText:  text,
	})
	if err != nil {
		return nil, fmt.Errorf("something went wrong")
	}

	post := helpers.DBPostToGqlPost(dbPost)
	return &post, nil
}

// User is the resolver for the user field.
func (r *postResolver) User(ctx context.Context, obj *model.Post) (*model.User, error) {
	return loaders.GetUser(ctx, obj.User.ID)
}

// CurrentUser is the resolver for the CurrentUser field.
func (r *queryResolver) CurrentUser(ctx context.Context) (*model.CurrentUser, error) {
	dbUser := auth.ForContext(ctx)
	if dbUser == nil {
		return nil, fmt.Errorf("access denied")
	}
	user := helpers.DBUserToGqlCurrentUser(*dbUser)
	return &user, nil
}

// Users is the resolver for the Users field.
func (r *queryResolver) Users(ctx context.Context, ids []string) ([]*model.User, error) {
	if ids != nil {
		users, err := loaders.GetUsers(ctx, ids)
		if err != nil {
			return nil, fmt.Errorf("something went wrong")
		}

		return users, nil
	}

	dbUsers, err := r.DBQueries.GetUsers(ctx)
	if err != nil {
		return nil, fmt.Errorf("something went wrong")
	}

	users := helpers.DBUsersToGqlUsers(dbUsers)
	return users, nil
}

// Posts is the resolver for the posts field.
func (r *queryResolver) Posts(ctx context.Context) ([]*model.Post, error) {
	currentUser := auth.ForContext(ctx)
	if currentUser == nil {
		return nil, fmt.Errorf("access denied")
	}

	dbPosts, err := r.DBQueries.GetPostsForUser(ctx, currentUser.ID)
	if err != nil {
		return nil, fmt.Errorf("something went wrong")
	}

	posts := []*model.Post{}
	for _, dbPost := range dbPosts {
		post := helpers.DBPostToGqlPost(dbPost)
		posts = append(posts, &post)
	}

	return posts, nil
}

// Following is the resolver for the following field.
func (r *userResolver) Following(ctx context.Context, obj *model.User) ([]*model.Follow, error) {
	return userFollows(r, ctx, obj, true)
}

// Followers is the resolver for the followers field.
func (r *userResolver) Followers(ctx context.Context, obj *model.User) ([]*model.Follow, error) {
	return userFollows(r, ctx, obj, false)
}

// Posts is the resolver for the posts field.
func (r *userResolver) Posts(ctx context.Context, obj *model.User) ([]*model.Post, error) {
	return loaders.GetUserPosts(ctx, obj.ID)
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Post returns PostResolver implementation.
func (r *Resolver) Post() PostResolver { return &postResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

// User returns UserResolver implementation.
func (r *Resolver) User() UserResolver { return &userResolver{r} }

type mutationResolver struct{ *Resolver }
type postResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type userResolver struct{ *Resolver }