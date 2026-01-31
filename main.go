package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type Price struct {
	USD float64 `json:"usd"`
}

type Prices map[string]Price

func GetPrices(ids []string) (Prices, error) {
	url := fmt.Sprintf(
		"https://api.coingecko.com/api/v3/simple/price?ids=%s&vs_currencies=usd",
		strings.Join(ids, ","),
	)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var prices Prices
	if err := json.NewDecoder(resp.Body).Decode(&prices); err != nil {
		return nil, err
	}

	return prices, nil
}

func main() {
	coins := []string{
		"bitcoin",
		"ethereum",
		"solana",
	}

	prices, err := GetPrices(coins)
	if err != nil {
		panic(err)
	}

	for _, coin := range coins {
		fmt.Printf("%s price: %0.2f USD\n", coin, prices[coin].USD)
	}
}
