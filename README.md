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
go build -o coin cmd/main.go
```

- If you don't have go installed in you computer, you can [build it by Docker](<[##build-by-docker](https://github.com/WandersonLontra/coin-cli?tab=readme-ov-file#%EF%B8%8F-build-by-docker)>)

### Adding to PATH

#### macOS/Linux

```bash
mv coin /usr/local/bin/
```

#### Windows

1. Move the executable to a permanent location (e.g., `C:\Program Files\coin\`)

```bash
move coin.exe "C:\coin-cli\"
```

2. Add to PATH through System Properties:

```bash
setx PATH "%PATH%;C:\coin-cli"
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
MVR = 15.40
FKP = 0.79
GTQ = 7.70
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

- The CLI uses the API of [https://fixer.io](fixer.io).
- The CLI needs an env file called `configs.env` at the same directory of the executable file built.
- The CLI uses a cache system to store exchange rates locally. Default cache duration is set to 12 hours but you can change it through the env var <TTL_CACHE_IN_HOURS>.

## 🔑 API Key

This tool requires an API key from a currency exchange rate provider as an environment variable. Checkout the file `configs.env.example`:

```bash
BASE_URL = "https://data.fixer.io/api/latest"

ACCESS_KEY = "<YOUR_API_KEY>"

TTL_CACHE_IN_HOURS = 12

DEV_MODE = false #set true only in development mode
```

## 🛠️ Build By Docker

First of all, you need change the GOOS parameter in dockerfile according to you Operational System

Use `darwin` for MacOS, `linux` or `windows`

```docker
RUN CGO_ENABLED=0 GOOS=darwin go build -o coin cmd/main.go
```

To build and extract the binary, run these commands:

1 - Build the Docker image

```bash
docker build -t coin-cli .
```

2 - Create a temporary container

```bash
docker create --name temp coin-cli
```

3 - Copy the binary from the container to your local directory

```bash
docker cp temp:/coin ./coin
```

4 - Clean up the temporary container

```bash
docker rm temp
```

## 📝 License

MIT License
