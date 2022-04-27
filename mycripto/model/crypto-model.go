package model

import (
   "fmt"
)

type Net struct {
    Network    				string   `json:"network"`
	Protocol      			string   `json:"protocol"` 
	Default    				bool     `json:"default"`
	Payin_enabled    		bool     `json:"payin_enabled"`
	Payout_enabled 			bool     `json:"payout_enabled"`
	Precision_payout    	string   `json:"precision_payout"`
	Payout_fee    			string   `json:"payout_fee"`
	Payout_is_payment_id    bool     `json:"payout_is_payment_id"`
	Payin_payment_id    	bool     `json:"payin_payment_id"`
	Payin_confirmations    	int      `json:"payin_confirmations"`
	Low_processing_time    	string   `json:"low_processing_time"`
	High_processing_time    string   `json:"high_processing_time"`
	Avg_processing_time    	string   `json:"avg_processing_time"`
}
type Cryptoresponse struct {
    Full_name       			string    `json:"full_name"`
    Crypto   					bool      `json:"crypto"`
	Payin_enabled 		    	bool   	  `json:"payin_enabled"`
	Payout_enabled 			    bool      `json:"payout_enabled"`
	Transfer_enabled 		    bool      `json:"transfer_enabled"`
	Precision_transfer 		    string    `json:"precision_transfer"`
	Networks 					[] Net
}

func (c Cryptoresponse) TextOutput() string {
   
  p := fmt.Sprintf("Full_name: %s\n", c.Full_name)

	p += fmt.Sprintf("Crypto : %t\n" +
	"Payin_enabled: %t\n" +
	"Payout_enabled: %t\n" +
	"Transfer_enabled %t\n" +
	"Precision_transfer: %s\n" +
	"Networks: %s\n" +
	"Protocol: %s\n" +
	"Default: %t\n" +
	"Payin_enabled: %t\n" +
	"Payout_enabled: %t\n" +
	"Precision_payout: %s\n" +
	"Payout_fee: %s\n" +
	"Payout_is_payment_id: %t\n" +
	"Payin_payment_id: %t\n" +
	"Payin_confirmations: %d\n" +
	"Low_processing_time: %s\n" +
	"High_processing_time: %s\n" +
	"Avg_processing_time: %s\n",
	c.Crypto, 
	c.Payin_enabled, 
	c.Payout_enabled, 
	c.Transfer_enabled, 
	c.Precision_transfer,	
	c.Networks[0].Network, 
	c.Networks[0].Protocol, 
	c.Networks[0].Default, 
	c.Networks[0].Payin_enabled, 
	c.Networks[0].Payout_enabled, 
	c.Networks[0].Precision_payout, 
	c.Networks[0].Payout_fee, 
	c.Networks[0].Payout_is_payment_id, 
	c.Networks[0].Payin_payment_id, 
	c.Networks[0].Payin_confirmations, 
	c.Networks[0].Low_processing_time, 
	c.Networks[0].High_processing_time, 
	c.Networks[0].Avg_processing_time)
  
  
	p += "ABC"
       return p
}  