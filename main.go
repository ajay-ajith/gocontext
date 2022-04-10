package main

import (
	"context"
	"net/http"
	"sync"
)

func main() {
	var crawlableUrls []string = []string{
		"https://www.google.co.in",
		"https://www.google.co.in",
		"https://www.google.co.in",
		"https://www.google.co.in",
		"https://www.google.co.in",
		"https://www.google.co.in",
		"https://www.google.co.in",
		"https://www.google.co.in",
		"https://www.google.co.in",
		"https://www.google.co.in",
		"https://www.google.co.in",
		"https://www.google.co.in",
		"https://www.google.co.in",
		"https://www.google.co.in",
		"https://www.google.co.in",
		"https://www.google.co.in",
		"https://www.google.co.in",
		"https://www.google.co.in",
		"https://www.google.co.in",
	}

	const MAXWORKERS = 3

	url_server := make(chan string)
	client := http.Client{}
	wg := sync.WaitGroup{}

	// starting workers
	for i := 0; i < MAXWORKERS; i++ {
		go Worker(i, &client, url_server, context.Background(), &wg)
		wg.Add(1)
	}

	// pushing URLs through the channel for workers to consume
	for _, url := range crawlableUrls {
		url_server <- url
	}
	close(url_server)

	// waiting for goroutines to finish
	wg.Wait()

}
