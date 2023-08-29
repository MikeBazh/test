package coincap

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"
	"time"
)

type Client struct {
	client *http.Client
}

func NewClient(timeout time.Duration) (*Client, error) {
	if timeout == 0 {
		return nil, errors.New("timeout can't be zero")
	}
	return &Client{
		client: &http.Client{
			Timeout: timeout,
			Transport: &loggingRoundTripper{
				logger: os.Stdout,
				next:   http.DefaultTransport,
			},
		},
	}, nil
}

func (c Client) GetAssets() ([]Asset, error) {
	resp, err := c.client.Get("https://api.coincap.io/v2/assets")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	//	fmt.Println(string(body))

	var r assetsResponse
	if err = json.Unmarshal(body, &r); err != nil {
		return nil, err
	}

	return r.Assets, nil
}
