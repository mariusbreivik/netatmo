# 🌤️ netatmo

> Your weather station in your terminal. Because sometimes you just need to know the CO2 level without opening an app.

[![Build](https://github.com/mariusbreivik/netatmo/actions/workflows/build.yml/badge.svg)](https://github.com/mariusbreivik/netatmo/actions/workflows/build.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/mariusbreivik/netatmo)](https://goreportcard.com/report/github.com/mariusbreivik/netatmo)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](LICENSE)

---

## 📖 What is this?

`netatmo` is a lightweight CLI tool built with [Cobra](https://github.com/spf13/cobra) and [Go](https://golang.org/) that lets you fetch data from your [Netatmo Weather Station](https://www.netatmo.com/en-eu/weather/weatherstation) right in your terminal.


---

## ✨ Features

- 🌡️ **Temperature** — Indoor and outdoor readings
- 💧 **Humidity** — Indoor and outdoor levels
- 🌫️ **CO2** — Carbon dioxide concentration
- 🔊 **Noise** — Sound level in decibels
- 🌀 **Pressure** — Atmospheric pressure
- 📶 **WiFi** — Signal strength monitoring
- ⚙️ **Firmware** — Device firmware info
- 🔐 **Secure Auth** — OAuth2 with automatic token refresh
- 🐚 **Shell Completion** — Bash, Zsh, and Fish support

---

## 🚀 Quick Start

### Prerequisites

- [Go](https://golang.org/) 1.24 or later
- A [Netatmo Weather Station](https://www.netatmo.com/en-eu/weather/weatherstation)
- A Netatmo developer account at [dev.netatmo.com](https://dev.netatmo.com/apps/)

### Installation

```shell
# Clone the repo
git clone https://github.com/mariusbreivik/netatmo.git
cd netatmo

# Build it
go build -o netatmo .

# (Optional) Install globally
go install
```

### Configuration

All configuration lives in one file: `~/.netatmo-config.json`

**Step 1: Set up your API credentials**

1. Head to [dev.netatmo.com/apps](https://dev.netatmo.com/apps/)
2. Create an app (or use an existing one)
3. Grab your **Client ID** and **Client Secret**

```shell
netatmo configure
# Follow the prompts, or use flags:
netatmo configure --client-id YOUR_ID --client-secret YOUR_SECRET
```

**Step 2: Authenticate**

1. In your app on [dev.netatmo.com](https://dev.netatmo.com/apps/), scroll to **Token generator**
2. Select scope `read_station` and click **Generate Token**
3. Copy both the **access token** and **refresh token**

```shell
netatmo login
# Follow the prompts, or use flags:
netatmo login --access-token YOUR_TOKEN --refresh-token YOUR_REFRESH
```

> 💡 **Pro tip:** Tokens auto-refresh when they expire. Set it and forget it!

---

## 🧑‍💻 Usage

### 🌡️Temperature
```shell
# Indoor temperature
netatmo temp --indoor
netatmo temp -i

# Outdoor temperature
netatmo temp --outdoor
netatmo temp -o
```

### 💧 Humidity
```shell
netatmo humidity --indoor   # or -i
netatmo humidity --outdoor  # or -o
```

### 🌫 CO2 Level
```shell
netatmo co2
```
> 🌿 Keep it under 1000 ppm for a happy brain!

### 🔊 Noise Level
```shell
netatmo noise
```
> 🔇 Measured in decibels. Library quiet? Or rock concert?

### 🌀 Pressure
```shell
netatmo pressure
```
> 📊 Atmospheric pressure for weather nerds.

### Shell Completion
```shell
# Bash
netatmo completion bash > /etc/bash_completion.d/netatmo

# Zsh
netatmo completion zsh > "${fpath[1]}/_netatmo"

# Fish
netatmo completion fish > ~/.config/fish/completions/netatmo.fish
```

---

## 📋 All Commands

```
Usage:
  netatmo [command]

Available Commands:
  co2         Read CO2 data from netatmo station
  completion  Generate shell autocompletion scripts
  configure   Configure Netatmo API credentials
  firmware    Read firmware data from netatmo station
  help        Help about any command
  humidity    Read humidity data from netatmo station
  login       Store Netatmo API tokens for authentication
  noise       Read noise data from netatmo station
  pressure    Read pressure data from netatmo station
  temp        Read temperature data from netatmo station
  wifi        Read wifi data from netatmo station

Flags:
  -d, --debug   Enable debug logging
  -h, --help    Show help
```

---

## 🛠️ Development

Want to contribute? Awesome! 🎉

```shell
# Clone
git clone https://github.com/mariusbreivik/netatmo.git
cd netatmo

# Run tests
go test ./...

# Build
go build -o netatmo .

# Run
./netatmo --help
```

---

## 📄 License

[Apache License 2.0](LICENSE)

---

<p align="center">
  Made with ☕ and curiosity
</p>
