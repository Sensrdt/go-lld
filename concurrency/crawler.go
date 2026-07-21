package main

import (
	"fmt"
	"net/http"
	"sync"
)

type Response struct {
	Result bool
	Error  error
	Cached bool
}

type UrlResponseRegister struct {
	Response map[string]Response
	mu       sync.Mutex
}

func (urr *UrlResponseRegister) Crawl(url string, register chan<- Response) {
	urr.mu.Lock()
	response, present := urr.Response[url]
	urr.mu.Unlock()
	if present {
		response.Cached = true
		register <- response
		return
	}
	_, err := http.Get(url)
	if err != nil {
		urr.mu.Lock()
		response := Response{
			Result: false,
			Error:  err,
		}
		// urr.Response = make(map[url]Response)
		urr.Response[url] = response
		urr.mu.Unlock()
		register <- response

	} else {
		urr.mu.Lock()
		response := Response{
			Result: true,
			Error:  nil,
		}
		// urr.Response = make(map[url]Response)
		urr.Response[url] = response
		urr.mu.Unlock()
		register <- response
	}

}

func Handler() {
	urlResponseRegister := &UrlResponseRegister{
		Response: make(map[string]Response),
		mu:       sync.Mutex{},
	}

	register := make(chan Response)

	wg := sync.WaitGroup{}

	urls := []string{"https://www.google.com", "https://www.facebook.com", "https://www.twitter.com", "https://www.instagram.com", "https://www.no-work-it-is.com", "https://www.facebook.com"}
	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			urlResponseRegister.Crawl(url, register)
		}(url)
	}

	go func() {
		wg.Wait()
		close(register)
	}()

	for value := range register {
		fmt.Println("Cached: ", value.Cached)

		if value.Result {
			println("Success")
		} else {
			println("Failure: ", value.Error)
		}
	}

}
