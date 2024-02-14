package main

import (
	"context"
	"log"
	"net/http"
	"os"

	authmodule "github.com/EzequielBPullolil/auth_service/src/auth_module"
	usermodule "github.com/EzequielBPullolil/auth_service/src/user_module"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	HOST, PORT := os.Getenv("HOST"), os.Getenv("PORT")
	server := http.NewServeMux()
	connectionPool, _ := pgxpool.New(context.Background(), os.Getenv("DB_URI"))
	user_repo := usermodule.NewUserRepository(connectionPool)
	user_repo.CreateTables()
	usermodule.HandleUserRoute(server, user_repo)
	authmodule.HandleAuthRoutes(server, user_repo)
	log.Printf(`Server listen on "%s:%s`, HOST, PORT)
	log.Fatal(http.ListenAndServe(HOST+":"+PORT, server))
}
