package main

import (
	"log"

	"github.com/levigross/grequests"
)

func main() {
	resp, err := grequests.Get("http://httpbin.org/get", nil)
	// You can modify the request by passing an optional RequestOptions struct
	if err != nil {
		log.Fatalln("Unable to make request: ", err)
	}
	var returnData map[string]interface{}
	resp.JSON(&returnData)
	log.Println(returnData)

}
