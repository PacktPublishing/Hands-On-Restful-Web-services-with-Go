package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/Hands-On-Restful-Web-services-with-Go/chapter1/mirrors"
)

func findFastest(urls []string) []byte {
	var fastestURL string
	var fastestTime time.Duration

	for _, url := range urls {
		start := time.Now()
		_, err := http.Get(url + "/README")
		if err == nil {
			now := time.Now()
			latency := now.Sub(start) / time.Millisecond

			if fastestTime == 0 {
				fastestURL = url
				fastestTime = latency
				continue
			}

			if latency < fastestTime {
				fastestURL = url
				fastestTime = latency
			}

			fmt.Printf("Mirror: %s, Latency: %dms\n", url, latency)
		}
	}
	data := make(map[string]string)
	data["fastest_mirror"] = fastestURL
	data["latency"] = strconv.Itoa(int(fastestTime))
	resp, _ := json.Marshal(data)
	return resp
}

func main() {
	http.HandleFunc("/fastest-mirror", func(w http.ResponseWriter, r *http.Request) {
		respJSON := findFastest(mirrors.MirrorList)
		// d := make(map[string]int)
		// d["data"] = 3
		// respJSON, _ := json.Marshal(d)
		fmt.Printf(string(respJSON))
		w.Header().Set("Content-Type", "application/json")
		w.Write(respJSON)
	})
	port := ":8000"
	server := &http.Server{
		Addr:           port,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	fmt.Printf("Starting server on port %s\n", port)
	log.Fatal(server.ListenAndServe())
}
