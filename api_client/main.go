package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type assetsResponse struct {
	Data      []assetData `json:"data"`
	Timestamp int64       `json:"timestamp"`
}

type assetData struct {
	ID            string `json:"id"`
	Rank          string `json:"rank"`
	Symbol        string `json:"symbol"`
	Name          string `json:"name"`
	Supply        string `json:"supply"`
	MaxSupply     string `json:"maxSupply"`
	MarketCapUsd  string `json:"marketCapUsd"`
	VolumeUsd24Hr string `json:"volumeUsd24Hr"`
	PriceUsd      string `json:"priceUsd"`
}
type loggingRoundTripper struct {
	logger io.Writer
	next   http.RoundTripper
}

func (d assetData) Info() string {
	return fmt.Sprintf("[ID] %s [Rank] %s [Symbol] %s", d.ID, d.Rank, d.Symbol)
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
	resp, err := client.Get("https://api.coincap.io/v2/assets")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	fmt.Println("Response status:", resp.StatusCode)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	//	fmt.Println(string(body))

	var r assetsResponse
	if json.Unmarshal(body, &r); err != nil {
		log.Fatal(err)
	}
	for _, asset := range r.Data {
		fmt.Println(asset.Info())
	}
}

/* client := &http.Client{}
Transport: &loggingRoundTripper{
	logger: os.Stdout
} */
