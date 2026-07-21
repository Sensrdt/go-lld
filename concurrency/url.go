package main

import (
	"fmt"
	"net/http"
	"sync"
)

type UrlResponse struct {
	Url        string
	StatusCode int
	Error      error
}

func testUrl(url string, successChannel chan UrlResponse, failureChannel chan UrlResponse) {

	resp, err := http.Get(url)
	if err != nil {
		failureChannel <- UrlResponse{Url: url, Error: err}
		return
	}
	successChannel <- UrlResponse{Url: url, StatusCode: resp.StatusCode}
}

func ProcessUrls() {
	urls := []string{"https://www.google.com", "https://www.facebook.com", "https://www.twitter.com", "https://www.instagram.com", "https://www.no-work-it-is.com"}

	successChannel := make(chan UrlResponse, len(urls)-1)
	failureChannel := make(chan UrlResponse, len(urls)-1)

	var wg sync.WaitGroup

	for i := range urls {
		wg.Add(1)
		go func(urlIdx int) {
			defer wg.Done()
			testUrl(urls[urlIdx], successChannel, failureChannel)
		}(i)
	}

	wg.Wait()
	close(successChannel)
	close(failureChannel)

	for v := range successChannel {
		fmt.Println("Success: ", v)
	}
	for v := range failureChannel {
		fmt.Println("Failure: ", v)
	}

	fmt.Println("Done")

}
