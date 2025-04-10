package main

import (
	"abrxu.com/gin/api/router"
	_ "abrxu.com/gin/pkg/db"
	psql "abrxu.com/gin/pkg/db"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	psql.Connect()
	router.Initialize()
}
