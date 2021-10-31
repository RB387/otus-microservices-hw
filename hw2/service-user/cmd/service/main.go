package main

import (
	"fmt"
	"github.com/RB387/otus-microservices-hw/hw2/service-user/internal/app"
	"github.com/RB387/otus-microservices-hw/hw2/service-user/internal/clients/psql"
	"github.com/RB387/otus-microservices-hw/hw2/service-user/internal/handlers/create_user"
	"github.com/RB387/otus-microservices-hw/hw2/service-user/internal/handlers/delete_user"
	"github.com/RB387/otus-microservices-hw/hw2/service-user/internal/handlers/get_user"
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

	getHandler := get_user.New(storage, userNameParam)
	deleteHandler := delete_user.New(storage, userNameParam)
	updateHandler := update_user.New(storage, userNameParam)
	createHandler := create_user.New(storage)

	userPath := fmt.Sprintf("/user/:%s", userNameParam)

	application.GET(userPath, getHandler.Handle)
	application.DELETE(userPath, deleteHandler.Handle)
	application.PATCH(userPath, updateHandler.Handle)
	application.POST("/user", createHandler.Handle)

	log.Fatal(application.Run())
}
