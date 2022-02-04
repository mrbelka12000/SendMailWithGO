package main

import (
	"github.com/joho/godotenv"
	"log"
	"mailsender/handler"
	"mailsender/repo"
	"mailsender/server"
)

func main() {
	log.SetFlags(log.LstdFlags|log.Lshortfile)
	if err := godotenv.Load(); err != nil {
		log.Fatalln(err)
	}
	confirm := repo.InitConfirm()
	handler := handler.InitHandler(confirm)
	if err := server.Run(handler); err != nil {
		log.Fatalln(err)
	}
}
