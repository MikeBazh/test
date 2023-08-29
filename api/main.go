package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type loggingRoundTripper struct {
	logger io.Writer
	next   http.RoundTripper
}

func (l loggingRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
	fmt.Fprintf(l.logger, "[%s] %s %s\n", time.Now().Format(time.ANSIC), r.Method, r.URL)
	return l.next.RoundTrip(r)
}

func main() {
	client := &http.Client{
		Transport: &loggingRoundTripper{
			logger: os.Stdout,
			next:   http.DefaultTransport,
		},
	}
	resp, err := client.Get("https://api.coincap.io/v2/assets/bitcoin")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	fmt.Println("Response status:", resp.StatusCode)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(body))
}

/* client := &http.Client{}
Transport: &loggingRoundTripper{
	logger: os.Stdout
} */
