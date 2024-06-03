package app

import (
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/eeQuillibrium/posts/graph"
	"github.com/eeQuillibrium/posts/graph/model"
	"github.com/eeQuillibrium/posts/graph/storage"
	"github.com/eeQuillibrium/posts/internal/service"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const defaultPort = "8080"

func (a *app) runHttpServer(service *service.Service, db *sqlx.DB) error {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	notifyChan := make(chan *model.Notification)
	defer close(notifyChan)

	st := storage.NewStorage()

	var srv http.Handler = handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{
		Resolvers: graph.NewResolver(service, a.log, notifyChan, st)}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	a.log.Infof("connect to http://localhost:%s/ for GraphQL playground", port)
	return http.ListenAndServe(":"+port, nil)
}
