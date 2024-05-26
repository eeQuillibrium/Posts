package app

import (
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/eeQuillibrium/posts/graph"
	loaders "github.com/eeQuillibrium/posts/graph/loader"
	"github.com/eeQuillibrium/posts/graph/model"
	"github.com/eeQuillibrium/posts/internal/service"
	"github.com/jmoiron/sqlx"
)

const defaultPort = "8080"

func (a *app) runHttpServer(service *service.Service, db *sqlx.DB) error {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	notifyChan := make(chan *model.Notification)
	defer close(notifyChan)

	var srv http.Handler = handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{
		Resolvers: graph.NewResolver(service, a.log, notifyChan)}))

	srvLoader := loaders.Middleware(db, srv)
	srvLoader.SetErrorPresenter(func(ctx context.Context, e error) *gqlerror.Error {
		err := graphql.DefaultErrorPresenter(ctx, e)

		var myErr *MyError
		if errors.As(e, &myErr) {
			err.Message = "Eeek!"
		}

		return err
	})
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srvLoader)

	a.log.Infof("connect to http://localhost:%s/ for GraphQL playground", port)
	a.log.Infof("asdsa")
	return http.ListenAndServe(":"+port, nil)
}
