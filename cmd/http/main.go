package main

import (
	"context"
	"fmt"
	"log"
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
		log.Fatalf("%s: %v", msg, err)
		os.Exit(1)
	}
}

func main() {
	// load .env configs
	conf, err := config.New()
	handleError(err, "unable to load .env configs")
	fmt.Println(".env configs loaded successfully")

	// init context
	ctx := context.Background()

	// init db connection
	db, err := postgres.New(ctx, conf.DB)
	handleError(err, "unable to connect with postgres db")
	fmt.Println("DB connection established successfully")

	// migrate dbs
	err = db.Migrate(&domain.User{}, &domain.Job{})
	handleError(err, "db migration failed")
	fmt.Println("DB migrated successfully")

	// dependency injections
	userRepo := repository.NewUserRepository(db)
	userSvc := service.NewUserService(userRepo)
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
	if err := r.Serve(conf.HTTP); err != nil {
		log.Fatal(err)
	}
}
