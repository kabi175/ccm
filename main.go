package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println(err)
		return
	}
	args := os.Args[1:]
	handler := handler{}
	switch args[0] {
	case "verify":
		handler.Verify()
	case "init":
		handler.Init()
	case "track":
		handler.Track()
	case "extract":
		handler.Extract(args[1])
	default:
		handler.Default()
	}
}
