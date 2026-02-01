package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Price struct {
	USD float64 `json:"usd"`
}

type Prices map[string]Price

func GetPrices() (Prices, error) {
	resp, err := http.Get("https://api.coingecko.com/api/v3/simple/price?ids=bitcoin,ethereum,solana&vs_currencies=usd")

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

func handler(w http.ResponseWriter, r *http.Request) {
	prices, err := GetPrices()
	if err != nil {
		http.Error(w, "API Error", 500)
		return
	}

	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	tmpl.Execute(w, prices)
}

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", handler)

	fmt.Println("Listennig on 8080 port")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
