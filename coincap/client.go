package coincap

import (
	"encoding/json"
	"errors"
	"fmt"
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
		return nil, errors.New("timeout cant be zero")
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

func (c Client) GetAssets() ([]AssetData, error) {
	resp, err := c.client.Get("https://api.coincap.io/v2/assets")

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err

	}

	var r AssetsResp

	if err = json.Unmarshal(body, &r); err != nil {
		return nil, err

	}

	return r.Assets, nil
}

func (c Client) GetAsset(name string) (AssetData, error) {
	url := fmt.Sprintf("https://api.coincap.io/v2/assets/%s", name)
	resp, err := c.client.Get(url)

	if err != nil {
		return AssetData{}, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return AssetData{}, err

	}

	var r AssetResp

	if err = json.Unmarshal(body, &r); err != nil {
		return AssetData{}, err

	}

	return r.Asset, nil
}
