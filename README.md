# netatmo 





</br>

  - [📖 describe()](#-describe)
  - [🧑‍💻 use()](#-use)
    - [🌡 temp](#-temp)
    - [🎧 noise](#-noise)
    - [🌫 co2](#-co2)
    - [💦 humidity](#-humidity)
    - [⚙️ firmware](#️-firmware)
    - [📶 wifi](#-wifi)
    - [🕐 lastupgrade](#-lastupgrade)
    - [📈 pressure](#-pressure)
  - [📜 develop()](#-develop)
  - [💾 install()](#-install)

## 📖 describe()
`netatmo` is a tiny CLI based on the [cobra](https://github.com/spf13/cobra)
 framework written in [go-lang](https://golang.org/). Its mostly just for fun, but the purpose is retrieving and displaying data in the command line from netatmo weather api.

 </br>

## 🧑‍💻 use()
`netatmo` has several subcommands which can be used to get different data from your Netatmo Weather Station. There is still some work remaining to get all commands working.

### 🌡 temp
 ```shell
$ netatmo temp -o, --outdoor | -i , --indoor
 ```

### 🎧 noise
 ```shell
$ netatmo noise
 ```

### 🌫 co2
 ```shell
$ netatmo co2
 ```

### 💦 humidity
 ```shell
$ netatmo humidity
 ```

### ⚙️ firmware
 ```shell
$ netatmo firmware

 ```
### 📶 wifi
  ```shell
$ netatmo wifi
 ```

### 🕐 lastupgrade
  ```shell
$ netatmo lastupgrade
 ```

 ### 📈 pressure
  ```shell
$ netatmo pressure
 ```


</br>

## 📜 develop()
 * You need to have your own [Netatmo Weather Station](https://www.netatmo.com/en-eu/weather/weatherstation) in order to use this CLI
 * Sign up at [netatmo](https://dev.netatmo.com/apps/) and create an app to get `clientId` and `clientSecret` in order to retrieve data from your Netatmo Weateher Station through the API.
  
</br>

 ## 💾 install()
  * Make sure [go](https://golang.org/) is installed
  * Clone this repo
  * install dependencies and build:
```shell
$ go install && go build
```
* create a config file called `$HOME/.netatmo.yaml` with this content:
  
```yaml
netatmo:
  clientID: YOUR_CLIENT_ID
  clientSecret: YOUR_CLIENT_SECRET
  username: YOUR_NETATMO_USERNAME
  password: YOUR_NETATMO_PASSWORD
```
* If everything is correct you should be able to run:
```
$ netatmo

Uses the Netatmo Weatherstation API to get your indoor/outdoor
temperature, co2 level, nois level, humidity, firmware data, wifi signal strength,
and more

Usage:
  netatmo [flags]
  netatmo [command]

Examples:
netatmo temp --indoor

Available Commands:
  co2         read co2 data from netatmo station
  firmware    read firmware data from netatmo station
  help        Help about any command
  humidity    read humidity data from netatmo station
  noise       read noise data from netatmo station
  temp        read temperature data from netatmo station
  wifi        read wifi data from netatmo station
  pressure    read pressure data from netatmo station

Flags:
      --config string   config file (default is $HOME/.netatmo.yaml)
  -d, --debug           debug logging
  -h, --help            help for netatmo

Use "netatmo [command] --help" for more information about a command.
```

</br>
