package main

import (
	"github.com/Fudjin322/todoApp/pkg/handler"
	"github.com/Fudjin322/todoApp/pkg/http"
	"log"
)

func main() {

	handlers := new(handler.Handler)

	srv := new(http.Server)
	if err := srv.Run("8080", handlers.InitRouter()); err != nil {
		log.Fatalf("error occured while running http server: %s", err)
	}
}
