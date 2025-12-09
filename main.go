package main

import (
	"context"
	"log"
	"session-14/cmd"
	"session-14/database"
	"session-14/handler"
	"session-14/repository"
	"session-14/service"
)

func main() {
	// Initialize database connection
	db, err := database.InitDB()
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	defer db.Close(context.Background())

	// Initialize objects for repository, service, and handler
	repoReport := repository.NewRepoReport(db)
	serviceReport := service.NewServiceReport(&repoReport)
	handlerReport := handler.NewReportHandler(&serviceReport)

	cmd.HomePage(handlerReport)
}
