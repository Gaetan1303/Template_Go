package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Gaetan1303/Template_Go/config"
	"github.com/Gaetan1303/Template_Go/internal/routes"
)

func main() {
	config.LoadEnv()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r := routes.SetupRouter()
	log.Printf("Serveur démarré sur :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
