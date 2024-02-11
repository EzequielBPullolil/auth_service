package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/EzequielBPullolil/auth_service/users"
	"github.com/EzequielBPullolil/auth_service/users/auth"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {

	server := http.NewServeMux()
	connectionPool, _ := pgxpool.New(context.Background(), os.Getenv("DB_URI"))
	user_repo := users.NewUserRepository(connectionPool)
	user_repo.CreateTables()
	users.HandleUserRoute(server, user_repo)
	auth.HandleAuthRoutes(server, user_repo)
	log.Println("Server start at port 8030")
	log.Fatal(http.ListenAndServe(":8030", server))
}
