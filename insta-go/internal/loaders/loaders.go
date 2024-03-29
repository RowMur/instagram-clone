package loaders

import (
	"context"
	"github/rowmur/insta-clone/internal/database"
	"github/rowmur/insta-clone/internal/graph/model"
	"net/http"
	"time"

	"github.com/vikstrous/dataloadgen"
)

type ctxKey string

const loadersKey = ctxKey("dataloaders")

type Loaders struct {
	UserLoader      *dataloadgen.Loader[string, *model.User]
	UserPostsLoader *dataloadgen.Loader[string, []*model.Post]
}

func NewLoaders(dbQueries *database.Queries) *Loaders {
	ur := &userReader{dbQueries: dbQueries}
	upr := &userPostsReader{dbQueries: dbQueries}
	return &Loaders{
		UserLoader:      dataloadgen.NewLoader(ur.getUsers, dataloadgen.WithWait(time.Millisecond)),
		UserPostsLoader: dataloadgen.NewLoader(upr.getPostsByUserIDs, dataloadgen.WithWait(time.Millisecond)),
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
