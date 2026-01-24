package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/yehezkiel1086/go-gin-nextjs-auth/internal/adapter/config"
	"github.com/yehezkiel1086/go-gin-nextjs-auth/internal/adapter/handler"
	"github.com/yehezkiel1086/go-gin-nextjs-auth/internal/adapter/storage/postgres"
	"github.com/yehezkiel1086/go-gin-nextjs-auth/internal/adapter/storage/postgres/repository"
	"github.com/yehezkiel1086/go-gin-nextjs-auth/internal/core/domain"
	"github.com/yehezkiel1086/go-gin-nextjs-auth/internal/core/service"
)

func handleError(err error, msg string) {
	if err != nil {
		slog.Error(msg, "error", err)
		os.Exit(1)
	}
}

func main() {
	// load .env configs
	conf, err := config.New()
	handleError(err, "unable to load .env configs")
	slog.Info(".env configs loaded successfully", "app", conf.App.Name, "env", conf.App.Env)

	// init context
	ctx := context.Background()

	// init db connection
	db, err := postgres.New(ctx, conf.DB)
	handleError(err, "unable to connect with postgres db")
	slog.Info("postgres db connected successfully", "db", conf.DB.Host + ":" + conf.DB.Port)

	// migrate dbs
	err = db.Migrate(&domain.User{}, &domain.Job{})
	handleError(err, "migration failed")
	slog.Info("dbs migrated successfully")

	// dependency injections
	userRepo := repository.NewUserRepository(db)
	userSvc := service.NewUserService(conf.HTTP, conf.JWT, userRepo)
	userHandler := handler.NewUserHandler(userSvc)

	authSvc := service.NewAuthService(conf.JWT, userRepo)
	authHandler := handler.NewAuthHandler(conf.JWT, authSvc)

	jobRepo := repository.NewJobRepository(db)
	jobSvc := service.NewJobService(jobRepo)
	jobHandler := handler.NewJobHandler(jobSvc)

	// init router
	r := handler.NewRouter(
		conf.HTTP,
		conf.JWT,
		userHandler,
		authHandler,
		jobHandler,
	)

	// start server
	err = r.Serve(conf.HTTP)
	handleError(err, "failed to serve backend api")
}
