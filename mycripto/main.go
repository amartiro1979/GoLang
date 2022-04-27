package main

import (
    "flag"
    "fmt"
    "log"
	"mycripto/client"
)



func main() {

    nameOfCrypto := flag.String(
      "crypto", "BTC", "Input the name of the CryptoCurrency you would like to know the price of",
    )
    flag.Parse()

    crypto, err := client.FetchCrypto(*nameOfCrypto)
    if err != nil {
        log.Println(err)
      }

  fmt.Println(crypto)
}