package main

import (
	"fmt"
	"github.com/emicklei/go-restful"
	"io"
	"net/http"
	"time"
)

func main() {
	// Create a web service
	webservice := new(restful.WebService)
	// Create a route and attach it to handler in the service
	webservice.Route(webservice.GET("/ping").To(pingTime))
	// Add the service to application
	restful.Add(webservice)
	http.ListenAndServe(":8000", nil)
}

func pingTime(req *restful.Request, resp *restful.Response) {
	// Write to the response
	io.WriteString(resp, fmt.Sprintf("%s", time.Now()))
}
