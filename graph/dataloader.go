package graph

import (
	"context"
	"errors"
	"github.com/go-pg/pg/v10"
	"graphqltest/models"
	"net/http"
	"sync"
	"time"
)

const userLoaderKey = "userLoader"

func DataloaderMiddleware(db *pg.DB, next http.Handler) http.Handler  {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userLoader := UserLoader{
			fetch: func(keys []string) ([]*models.User, []error) {
				var users []*models.User

				u := make(map[string]*models.User, len(keys))

				err := db.Model(&users).Where("id in (?)", pg.In(keys)).Select()

				if err != nil {
					return users, []error{errors.New("Cannot get users")}
				}

				for _, user := range users {
					u[user.ID] = user
				}

				result := make([]*models.User, len(keys))

				for i, key := range keys {
					result[i] = u[key]
				}

				return result, nil

			},
			wait:     1 * time.Millisecond,
			maxBatch: 100,
			mu:       sync.Mutex{},
		}


		ctx := context.WithValue(r.Context(), userLoaderKey, &userLoader)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// Gest user loader data
func GetUserLoader(ctx context.Context) *UserLoader {
	return ctx.Value(userLoaderKey).(*UserLoader)
}