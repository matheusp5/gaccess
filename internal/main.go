package main

import (
	"gaccess/external/data"
	"gaccess/internal/http"
)

func main() {
	database := data.InitDatabase()
	http.InitHttpServer(database)
}
