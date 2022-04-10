package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

func GetBody(client *http.Client, url string) ([]byte, error) {
	request, _ := http.NewRequest("GET", url, nil)
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	body, _ := ioutil.ReadAll(response.Body)
	return body, nil
}

// A worker accepts a url to visit and a context
// It visits the website, and pushes the data to a channel
func Worker(id int, client *http.Client, urls chan string, ctx context.Context, wg *sync.WaitGroup) {
	fmt.Printf("Worker %d is starting\n", id)
	for url := range urls {
		fmt.Printf("worker %d got url : %s\n", id, url)
		response, err := GetBody(client, url)
		if err != nil {
			fmt.Printf("Error visiting url : %s\n", url)
			continue
		}
		fmt.Printf("%d\n", len(string(response)))
	}
	wg.Done()
}
