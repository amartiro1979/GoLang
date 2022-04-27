package client

import (
   "encoding/json"
   //"fmt"
   "log"
   "net/http"
   "mycripto/model"
   //"io/ioutil"
   
)

//Fetch is exported ...
func FetchCrypto(crypto string) (string, error) {
   
   URL := "https://api.hitbtc.com/api/3/public/currency/" + crypto

   resp, err := http.Get(URL)
   if err != nil {
      log.Fatal("HTTP error")
   }
   defer resp.Body.Close()

   var cResp model.Cryptoresponse

   if err := json.NewDecoder(resp.Body).Decode(&cResp); err != nil {
      log.Fatal("Decode error")
   }
   return cResp.TextOutput(), nil
}