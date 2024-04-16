package main

import (
	"log"
	"mail-service/cmd"
)

func main() {
	log.Default().Println("Starting server...")
	cmd.Start()
}
