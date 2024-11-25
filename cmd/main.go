package main

import (
	database "TesisMVC/database/sqlc"
	"TesisMVC/internal/server/orchestrator"
	"TesisMVC/pkg/util"
	"database/sql"
	"log"
	"net/http"
)

func main() {

	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Failed to load config")
	}
	db, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Failed to load the database")
	}
	store := database.NewStore(db)
	orchestatrator, err := orchestrator.NewOrchestrator(store, config)
	if err != nil {
		log.Fatal("faild to load all servers", err.Error())
	}
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./pkg/static"))))
	http.HandleFunc("/login", orchestatrator.Auth.Login)
}
