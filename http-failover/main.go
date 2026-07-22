package main

/* implement logic if a http.Get resolves in 500ms then return the response if not, trigger next url get and if any resolve first then return that but if all fails then return failure */

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

type Result struct {
	body []byte
	err  error
}

func GetWithFailover(urls []string) ([]byte, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	resultCh := make(chan []byte, 1)
	faileChan := make(chan error, len(urls))

	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()

			req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
			if err != nil {
				faileChan <- err
				return
			}

			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				faileChan <- err
				return
			}
			defer resp.Body.Close()

			if resp.StatusCode != http.StatusOK {
				faileChan <- fmt.Errorf("status code %d", resp.StatusCode)
				return
			}

			body, err := io.ReadAll(resp.Body)
			if err != nil {
				faileChan <- err
				return
			}

			select {
			case resultCh <- body:
				cancel()
			default:
			}
		}(url)

		timer := time.NewTimer(500 * time.Millisecond)
		defer timer.Stop()

		select {
		case <-timer.C:
			continue
		case response := <-resultCh:
			return response, nil
		case <-faileChan:
			// continue to next URL
		}
	}

	wg.Wait()
	close(resultCh)

	select {
	case response := <-resultCh:
		return response, nil
	default:
	}

	return nil, fmt.Errorf("all requests failed")
}

func main() {
	urls := []string{
		"https://xasd.com",
		"https://economize.cloud",
		"https://facebook.com",
	}

	body, err := GetWithFailover(urls)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Success! Body length:", len(body))
}
