package main

import (
	"echoApi/config"
	"echoApi/infastructure/persistence"
	"echoApi/interfaces"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("env file not found")
	}
}

func main() {
	serveMux := http.NewServeMux()
	logger := log.New(os.Stdout, "echo-api:", log.LstdFlags)

	appUrl := os.Getenv("APP_URL")
	postgresConfig := config.NewPostgresConfig()
	dbConfig := postgresConfig.GetConfig()
	logger.Printf("%#v", dbConfig)

	service, err := persistence.NewRepositories(
		dbConfig.Driver,
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Port,
		dbConfig.Host,
		dbConfig.Name)

	if err != nil {
		logger.Panic(err)
	}
	defer service.Close()
	service.AutoMigrate()

	userHandler := interfaces.NewUser(logger, service.User)
	echoHandler := interfaces.NewEcho(logger, service.Echo)
	serveMux.Handle("/user", userHandler)
	serveMux.Handle("/echo", echoHandler)

	server := http.Server{
		Addr:    appUrl,
		Handler: serveMux,
	}
	server.ListenAndServe()
}
