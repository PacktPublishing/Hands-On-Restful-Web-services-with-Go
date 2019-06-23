package main

import (
	"log"

	helper "github.com/Hands-On-Restful-Web-services-with-Go/chapter7/basicExample/helper"
)

func main() {
	db, err := helper.InitDB()
	if err != nil {
		log.Println(db)
	}
}
