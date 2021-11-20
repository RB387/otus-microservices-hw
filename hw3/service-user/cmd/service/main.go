package main

import (
	"fmt"
	"github.com/RB387/otus-microservices-hw/hw2/service-user/internal/app"
	"github.com/RB387/otus-microservices-hw/hw2/service-user/internal/clients/psql"
	"github.com/RB387/otus-microservices-hw/hw2/service-user/internal/handlers/create_user"
	"github.com/RB387/otus-microservices-hw/hw2/service-user/internal/handlers/delete_user"
	"github.com/RB387/otus-microservices-hw/hw2/service-user/internal/handlers/get_user"
	"github.com/RB387/otus-microservices-hw/hw2/service-user/internal/handlers/health"
	"github.com/RB387/otus-microservices-hw/hw2/service-user/internal/handlers/update_user"
	"github.com/RB387/otus-microservices-hw/hw2/service-user/internal/storages/user"
	_ "github.com/lib/pq"
	"log"
)

const userNameParam = "username"

func main() {
	application := app.NewFromEnv()
	postgres, err := psql.NewFromEnv()
	if err != nil {
		log.Fatal(err)
	}

	storage := user.NewPostgresStorage(postgres)

	createHandler := create_user.New(storage)
	getHandler := get_user.New(storage, userNameParam)
	updateHandler := update_user.New(storage, userNameParam)
	deleteHandler := delete_user.New(storage, userNameParam)
	healthHandler := health.New()

	userPath := fmt.Sprintf("/user/:%s", userNameParam)

	application.POST("/user", createHandler.Handle)
	application.GET(userPath, getHandler.Handle)
	application.PATCH(userPath, updateHandler.Handle)
	application.DELETE(userPath, deleteHandler.Handle)

	application.GET("/health", healthHandler.Handle)

	log.Fatal(application.Run())
}
