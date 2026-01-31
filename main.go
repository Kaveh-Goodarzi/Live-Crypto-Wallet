package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Coin interface {
	GetPrice() Coins
}

type Bitcoin struct {
	USD float64 `json:"usd"`
}

type Coins struct {
	BTC Bitcoint
	ETH Ethereum
}

func (btc *Bitcoin) GetPrice() Coins {
	resp, err := http.Get("https://api.coingecko.com/api/v3/simple/price?ids=bitcoin&vs_currencies=usd")
	if err != nil {
		fmt.Println("Error getting BTC data")
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading BTC data")
	}

	var finallyData Coins
	if err := json.Unmarshal(data, &finallyData); err != nil {
		fmt.Println("Error unmarshaling BTC data")
	}
	
	return finallyData

}

type Ethereum struct {
	USD float64 `json:"usd"`
}

func (eth *Ethereum) GetPrice() Coins {
	resp, err := http.Get("https://api.coingeko.com/api/v3/simple/price?ids=ethereum&vs_currencies=usd")
	if err != nil {
		fmt.Println("Error getting ETH data")
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading ETH data")
	}

	var finallyData Coins
	
	if err := json.Unmarshal(data, &finallyData); err != nil {
		fmt.Println("Error unmarshaling ETH data")
	}

	return finallyData

func main() {

}

// Web-Server Will Add Here
