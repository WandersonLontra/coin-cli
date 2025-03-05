# 🪙 Coin CLI

A currency converter and exchange rates tool right in your terminal.

## ⚡️ Features

- Real-time currency exchange rates
- Local caching system
- Simple and intuitive commands
- Support for multiple currencies

## 🚀 Installation

### Building from source

```bash
go build -o coin main.go
```

### Adding to PATH

#### macOS/Linux

```bash
mv coin /usr/local/bin/
```

#### Windows

1. Move the executable to a permanent location (e.g., `C:\Program Files\coin\`)

```bash
move coin.exe "C:\Program Files\coin\"
```

2. Add to PATH through System Properties:

```bash
setx PATH "%PATH%;C:\Program Files\coin"
```

## 📖 Usage

### List available currencies

```bash
coin list
```

Example output:

```
Available currencies:
USD = 1.00
EUR = 0.93
USD = 1.00
EUR = 0.93
...
```

### Change base of currencies at list command

```bash
coin list -s BTC,USD,EUR,CAD -b BTC
```

Example output:

```
CAD = 127889.35
EUR = 82673.22
BTC = 1.00
USD = 89026.57
```

### Convert currencies

```bash
coin convert -f USD -t EUR -a 100
```

Example output:

```
100 USD = 85.23 EUR
Rate: 1.00 USD = 0.93 EUR
```

## ⚙️ Configuration

- The CLI uses the API of [https://fixer.io](fixer.io)
- The CLI needs an env file called `configs.env`
- The CLI uses a cache system to store exchange rates locally. Default cache duration is set to 12 hours.

## 🔑 API Key

This tool requires an API key from a currency exchange rate provider as an environment variable. Checkout the file `configs.env.example`:

```bash
BASE_URL = "https://data.fixer.io/api/latest"

ACCESS_KEY = "<YOUR_API_KEY>"

TTL_CACHE_IN_HOURS = 12
```

## 📝 License

MIT License
