package main

import (
	"fmt"

	"github.com/ifvictr/cheese/pkg/cheese"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
}

func main() {
	fmt.Println("Starting Cheese…")
	config := cheese.NewConfig()

	// Start receiving messages
	fmt.Println(fmt.Sprintf("Listening on port %d", config.Port))
	cheese.StartServer(config)
}
