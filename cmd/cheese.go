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
	cheese.StartServer(config)
}
