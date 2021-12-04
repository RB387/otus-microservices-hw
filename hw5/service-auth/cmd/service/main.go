package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"

	"github.com/RB387/otus-microservices-hw/hw5/service-auth/internal/clients/user"
	"github.com/RB387/otus-microservices-hw/hw5/service-auth/internal/handlers/auth"
	"github.com/RB387/otus-microservices-hw/hw5/service-auth/internal/handlers/login"
	"github.com/RB387/otus-microservices-hw/hw5/service-auth/internal/handlers/register"
	"github.com/RB387/otus-microservices-hw/hw5/service-auth/internal/token"
)

const defaultAddr = ":8080"

func getEnv(key, defaultValue string) string {
	val := os.Getenv(key)
	if val == "" {
		return defaultValue
	}

	return val
}

func main() {
	tokenManager, err := token.NewManager()
	if err != nil {
		log.Fatal(err)
	}

	usersClient := user.NewFromEnv()

	authHandler := auth.New(tokenManager)
	loginHandler := login.New(usersClient, tokenManager)
	registerHandler := register.New(usersClient, tokenManager)

	srv := gin.Default()
	srv.GET("/auth", authHandler.Handle)
	srv.POST("/login", loginHandler.Handle)
	srv.POST("/register", registerHandler.Handle)

	addr := getEnv("SERVICE_URL", defaultAddr)
	log.Fatal(srv.Run(addr))
}
