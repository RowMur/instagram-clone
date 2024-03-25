package loaders

import (
	"context"
	"github/rowmur/insta-clone/internal/database"
	"github/rowmur/insta-clone/internal/graph/model"
	"github/rowmur/insta-clone/internal/helpers"
	"net/http"
	"time"

	"github.com/vikstrous/dataloadgen"
)

type ctxKey string

const loadersKey = ctxKey("dataloaders")

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

type Loaders struct {
	UserLoader *dataloadgen.Loader[string, *model.User]
}

func NewLoaders(dbQueries *database.Queries) *Loaders {
	ur := &userReader{dbQueries: dbQueries}
	return &Loaders{
		UserLoader: dataloadgen.NewLoader(ur.getUsers, dataloadgen.WithWait(time.Millisecond)),
	}
}

func Middleware(dbQueries *database.Queries) func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			loader := NewLoaders(dbQueries)
			r = r.WithContext(context.WithValue(r.Context(), loadersKey, loader))
			h.ServeHTTP(w, r)
		})
	}
}

func ForContext(ctx context.Context) *Loaders {
	return ctx.Value(loadersKey).(*Loaders)
}

func GetUser(ctx context.Context, userID string) (*model.User, error) {
	loaders := ForContext(ctx)
	return loaders.UserLoader.Load(ctx, userID)
}

func GetUsers(ctx context.Context, userIDs []string) ([]*model.User, error) {
	loaders := ForContext(ctx)
	return loaders.UserLoader.LoadAll(ctx, userIDs)
}
