package main

import (
	"log"

	"github.com/joho/godotenv"
)

func main() {

	var err error

	err = godotenv.Load()
	if err != nil {
		log.Panic(err)
	}

}