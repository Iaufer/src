package coincap

import "fmt"

type AssetsResp struct {
	Assets    []AssetData `json:"data"`
	Timestamp int64       `json:"timestamp"`
}

type AssetResp struct {
	Asset     AssetData `json:"data"`
	Timestamp int64     `json:"timestamp"`
}

type AssetData struct {
	ID           string `json:"id"`
	Rank         string `json:"rank"`
	Symbol       string `json:"symbol"`
	Name         string `json:"name"`
	Supply       string `json:"supply"`
	MaxSupply    string `json:"maxSupply"`
	MarketCapUSD string `json:"marketCapUSD"`
	VolumeUSD24h string `json:"volumeUsd24Hr"`
	PriceUSD     string `json:"priceUsd"`
}

func (d AssetData) Info() string {
	return fmt.Sprintf("[ID] %s | [RANK] %s | [SYNBOL] %s | [NAME] %s [PRICE] %s", d.ID, d.Rank, d.Symbol, d.Name, d.PriceUSD)
}
