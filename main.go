package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Coin struct {
	Bitcoin struct {
		USD float64 `json:"usd"`
	} `json:"bitcoin"`
}

func main() {
	resp, err := http.Get("https://api.coingecko.com/api/v3/simple/price?ids=bitcoin&vs_currencies=usd")
	if err != nil {
		fmt.Println("Error getting data")
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading data")
	}

	var humanData Coin
	if err := json.Unmarshal(data, &humanData); err != nil {
		fmt.Println("Error unmarshaling json data")
	}
	fmt.Println(humanData)
	fmt.Printf("Bitcoin cost: %.2f\n", humanData.Bitcoin.USD)
}

// Web-Server Will Add Here
