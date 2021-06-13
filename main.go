package main

import (
	"fmt"
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
	if len(args) > 1 {
		fmt.Println("args length exceded")
		return
	}
	switch args[0] {
	case "init":
		handleInit("./user-data/valid_users.json")
	case "verify":
		handleVerify()
	case "track":
		handleTrack("./user-data/valid_users.json")
	default:
		fmt.Printf("%s cmd not found", args[0])
	}
}
