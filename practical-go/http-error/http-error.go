package main

import (
	"fmt"
	"io"
	"net/http"
)

type HTTPError struct {
	StatusCode int
	URL        string
}

func (h *HTTPError) Error() string {
	return fmt.Sprintf("http status code = %d, url = %s", h.StatusCode, h.URL)
}

func ReadContents(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, &HTTPError{
			StatusCode: resp.StatusCode,
			URL:        url,
		}
	}
	return io.ReadAll(resp.Body)
}

func main() {
	rc, err := ReadContents("https://www.yahoo.co.jp/")
	if err != nil {
		panic(err)
	}

	fmt.Print(rc)
}
