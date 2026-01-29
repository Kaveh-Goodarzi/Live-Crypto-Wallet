# Crypto Price Tracker 🪙

A simple Go application that fetches real-time cryptocurrency prices from CoinGecko API.

---

## ✨ Features
- Real-time Bitcoin price fetching
- Clean and minimal Go code
- Easy to extend and modify
- No external dependencies (only standard library)

---

## 🚀 Quick Start

### Prerequisites
- Go 1.16 or higher

### Installation
1. Clone the repository:
```bash
git clone https://github.com/Kaveh-Goodarzi/Live-Crypto-Wallet.git
cd Live-Crypto-Wallet
```

2. Run the application:

```bash

go run main.go

```

---

## 📁 Project Structure
```
crypto-tracker/
├── main.go          # Main application code
├── README.md        # This file
└── go.mod           # Go module file
```

---

## 🛠️ How It Works

1. Sends HTTP GET request to CoinGecko API

2. Parses JSON response

3. Displays Bitcoin price in USD

---

## 🔧 Code Example
```go

// Simple struct to match API response
type CoinData struct {
    Bitcoin struct {
        USD float64 `json:"usd"`
    } `json:"bitcoin"`
}
```

---

## 🌟 Why This Project?

- **Educational**: Perfect for beginners learning Go and API integration

- **Minimal**: Only uses Go standard library

- **Foundation**: Can be extended to full-featured crypto tracker

---

## 📈 Future Improvements

* Add more cryptocurrencies

* Create web interface

* Add price history tracking

* Implement WebSocket for real-time updates

---

## 🤝 Contributing

Feel free to fork this project and submit pull requests!

---

## 📄 License

MIT License - see LICENSE file for details

---

## 👨💻 Author

**Kaveh Goodarzi**

- GitHub: @Kaveh-Goodarzi

