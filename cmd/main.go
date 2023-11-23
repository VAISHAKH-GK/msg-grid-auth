package main

import (
	"os"

	"github.com/VAISHAKH-GK/msg-grid-auth/server"
	"github.com/joho/godotenv"
)

func main() {
	// Loading enviornment variables
	godotenv.Load(".env")
	var port = os.Getenv("PORT")

	// starting the server
	server.Run(port)
}
