package main

import (
	"encoding/json"
	"fmt"
	"github.com/sahil-9/chapter1/mirrors"
	"log"
	"net/http"
	"time"
)

type response struct {
	FastestURL string `json:"fastest_url"`
	Latency time.Duration `json:"latency"`
}

func main() {
	http.HandleFunc("/fastest-mirror", func(w http.ResponseWriter,
		request *http.Request) {
		response := findFastest(mirrors.MirrorList)
		respJSON, _ := json.Marshal(response)
		w.Header().Set("Content-type", "application/json")
		w.Write(respJSON)
	})
	port := ":9000"
	server := &http.Server{
		Addr: port,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	fmt.Printf("Starting server on port %s\n", port)
	log.Fatal(server.ListenAndServe())
}

func findFastest(urls []string) response {
	urlChan := make(chan string)
	latencyChan := make(chan time.Duration)
	for _, url := range urls {
		mirrorURL := url
		go func() {
			start := time.Now()
			_, err := http.Get(mirrorURL + "/README")
			latency := time.Now().Sub(start) / time.Millisecond
			if err == nil {
				urlChan <- mirrorURL
				latencyChan <- latency
			}
		}()
	}
	return response{FastestURL: <-urlChan, Latency: <-latencyChan}
}