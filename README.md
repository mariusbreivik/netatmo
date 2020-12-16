# netatmo

</br>

  - [📖 describe()](#-describe)
  - [📜 prepare()](#-prepare)
  - [💾 install()](#-install)
  - [🧑‍💻 use()](#-use)
    - [🌡 temp](#-temp)
    - [🎧 noise](#-noise)
    - [🌫 co2](#-co2)
    - [💦 humidity](#-humidity)
    - [⚙️ firmware](#️-firmware)
    - [📶 wifi](#-wifi)
    - [🕐 lastupgrade](#-lastupgrade)

## 📖 describe()
`netatmo` is a tiny CLI based on the [cobra](https://github.com/spf13/cobra)
 framework written in [go-lang](https://golang.org/). Its mostly just for fun, but the purpose is retrieving and displaying data in the command line from netatmo weather api.

 </br>


## 📜 prepare()
 * Install go-lang
 * Sign up at [netatmo](https://dev.netatmo.com/apidocumentation/weather) to get credentials
  
</br>

 ## 💾 install()
  * install dependencies and build:
```shell
$ go install && go build
```
* create config file called ```$HOME/.netatmo.yaml``` with this content:
  
```yaml
netatmo:
  clientID: YOUR_CLIENT_ID
  clientSecret: YOUR_CLIENT_SECRET
  username: YOUR_NETATMO_USERNAME
  password: YOUR_NETATMO_PASSWORD
```

</br>

## 🧑‍💻 use()
```netatmo```CLI serves multiple usages. 

 ### 🌡 temp
 ```shell
$ netatmo --temp|-t outdoor|indoor
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
