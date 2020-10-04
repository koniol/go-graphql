package main

import (
	"encoding/json"
	"fmt"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"graphqltest/auth"
	"graphqltest/database"
	"graphqltest/graph"
	"graphqltest/graph/generated"
	"log"
	"net/http"
	"os"
	"strconv"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	pg, _ := database.Connect()
	defer pg.Close()

	c := generated.Config{Resolvers: &graph.Resolver{
		UserRepo:     database.UserRepo{DB: pg},
		DocumentRepo: database.DocumentRepo{DB: pg},
	}}
	data := auth.NewAuthData(pg)

	router := chi.NewRouter()

	router.Use(auth.Middleware(pg))

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(c))
	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", graph.DataloaderMiddleware(pg, srv))
	router.Post("/auth", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("YEA")

		var l LoginPost
		err := json.NewDecoder(r.Body).Decode(&l)

		if err != nil {
			fmt.Println("err", err)
		}

		if l.Id != "" && l.Email != "" {
			err := data.UserExists(l.Id, l.Email)
			if err == nil {
				fmt.Println("CREATE TOKEN")
				id, _ := strconv.Atoi(l.Id)
				user := auth.UserToken{
					Id:      id,
					Email:   l.Email,
					IsAdmin: false,
				}
				token, err := auth.CreateToken(&user)

				if err != nil{
					w.WriteHeader(http.StatusUnauthorized)
					return
				}
				w.Header().Set("Authorization", token)
				w.WriteHeader(http.StatusOK)
				return
			}
		}
		w.Write([]byte("id "))
	})

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

type LoginPost struct {
	Id string `json:"id"`
	Email string `json:"email"`
}

