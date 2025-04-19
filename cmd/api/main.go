package main

import (
	"context"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	_ "github.com/lib/pq"
	"net/http"
	"os"
	"social-comments/api/graphql/generated"
	"social-comments/api/graphql/resolvers"
	"social-comments/internal/app"
	"social-comments/internal/app/db"
	"social-comments/internal/core/domain"
	"social-comments/internal/core/ports"
	"social-comments/internal/core/repository"
	"social-comments/internal/infrastructure/persistence/memory"
	"social-comments/internal/infrastructure/persistence/postgres"
	"social-comments/internal/infrastructure/pubsub"
	lg "social-comments/pkg/logging"
)

func main() {
	logger := lg.Default()
	logger.Info("Executing InitLogger.")

	envFile := ".env"
	if len(os.Args) >= 2 {
		envFile = os.Args[1]
	}

	logger.Info("Executing InitConfig.")
	logger.Info("Reading %s \n", envFile)
	if err := app.InitConfig(envFile); err != nil {
		logger.Error(err.Error())
	}

	logger.Info("Connecting to Postgres.")

	options := db.PostgresOptions{
		Name:     os.Getenv("POSTGRES_DBNAME"),
		Port:     os.Getenv("POSTGRES_PORT"),
		Host:     os.Getenv("POSTGRES_HOST"),
		User:     os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
	}

	postgresDb, err := db.NewPostgresDB(options)

	if err != nil {
		logger.Error(err.Error())
	}

	broker := pubsub.NewBroker()

	var gateways *repository.Gateways

	logger.Info("Creating Gateways.")
	logger.Info("USE_IN_MEMORY = ", os.Getenv("USE_IN_MEMORY"))

	if os.Getenv("USE_IN_MEMORY") == "true" {
		posts := memory.NewPostRepository(100)
		posts.Create(context.Background(), domain.Post{Title: "Test Post 1", Content: "Content 1", Author: "Admin"})
		posts.Create(context.Background(), domain.Post{Title: "Test Post 2", Content: "Content 2", Author: "Admin"})
		posts.DebugDumpPosts()
		comments := memory.NewCommentRepository(10)
		gateways = repository.NewGateways(posts, comments)
	} else {
		posts := postgres.NewPostRepository(postgresDb)
		posts.Create(context.Background(), domain.Post{Title: "Test Post 1", Content: "Content 1", Author: "Admin"})
		posts.Create(context.Background(), domain.Post{Title: "Test Post 2", Content: "Content 2", Author: "Admin"})
		comments := postgres.NewCommentRepository(postgresDb)
		gateways = repository.NewGateways(posts, comments)
	}

	logger.Info("Creating Services.")
	services := ports.NewServices(gateways, logger, broker)

	logger.Info("Creating graphql server.")
	port := os.Getenv("PORT")
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolvers.Resolver{
		Broker:         broker,
		PostService:    services.PostService,
		CommentService: services.CommentService,
		Logger:         *logger,
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	logger.Info("Connect to http://localhost:%s/ for GraphQL playground", port)
	http.ListenAndServe(":"+port, nil)
}
