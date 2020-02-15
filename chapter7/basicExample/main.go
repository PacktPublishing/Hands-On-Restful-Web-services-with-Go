package main

import (
	"log"

	"github.com/Hands-On-Restful-Web-services-with-Go/chapter7/basicExample/helper"
)

func main() {
	_, err := helper.InitDB()
	if err != nil {
		log.Println(err)
	}

	log.Println("Database tables are successfully initialized.")
}
