package graph

import (
	"context"
	"fmt"
	"github/rowmur/insta-clone/internal/database"
	"github/rowmur/insta-clone/internal/graph/model"
	"github/rowmur/insta-clone/internal/loaders"

	"github.com/google/uuid"
)

func userFollows(r *userResolver, ctx context.Context, obj *model.User, following bool) ([]*model.Follow, error) {
	currentUserId := uuid.MustParse(obj.ID)

	var dbFollows []database.Follow
	var err error
	if following {
		dbFollows, err = r.DBQueries.GetFollowsByUser(ctx, currentUserId)
	} else {
		dbFollows, err = r.DBQueries.GetFollowersByUser(ctx, currentUserId)
	}
	if err != nil {
		return nil, fmt.Errorf("something went wrong")
	}

	follows := []*model.Follow{}
	for _, dbFollow := range dbFollows {
		var otherUserId uuid.UUID
		if following {
			otherUserId = dbFollow.UserFollowingID
		} else {
			otherUserId = dbFollow.UserID
		}

		user, err := loaders.GetUser(ctx, otherUserId.String())
		if err != nil {
			return nil, fmt.Errorf("something went wrong")
		}

		follow := model.Follow{
			FollowingSince: dbFollow.CreatedAt.String(),
			User:           user,
		}

		follows = append(follows, &follow)
	}

	return follows, nil
}
