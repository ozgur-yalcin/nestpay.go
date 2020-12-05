package main

import (
	"encoding/xml"
	"fmt"

	nestpay "github.com/OzqurYalcin/nestpay/src"
)

func main() {
	api := nestpay.API{"akbank"} // "akbank","asseco","isbank","finansbank","denizbank","kuveytturk","halkbank","anadolubank","hsbc","ziraatbank"
	request := nestpay.Request{}
	request.ClientId = "" // Müşteri No
	request.Username = "" // Kullanıcı adı
	request.Password = "" // Şifre
	// Ödeme
	request.Type = "Auth"
	request.Mode = "P"                           // TEST : "T" - PRODUCTION "P"
	request.IPAddress = ""                       // Müşteri IP adresi
	request.Number = ""                          // Kart numarası
	request.Expires = "xx/xx"                    // Kart son kullanma tarihi
	request.Cvv2Val = "xxx"                      // Kart Cvv2 Kodu
	request.Total = "0.00"                       // Satış tutarı
	request.Instalment = "0"                     // Taksit sayısı
	request.Currency = nestpay.Currencies["TRY"] // Para birimi
	// Fatura
	request.BillTo.Name = ""    // Kart sahibi
	request.BillTo.Company = "" // Fatura unvanı
	// 3D (varsa)
	//request.PayerTxnId = ""
	//request.PayerSecurityLevel = ""
	//request.PayerAuthenticationCode = ""
	//request.CardholderPresentCode = ""
	response := api.Transaction(request)
	pretty, _ := xml.MarshalIndent(response, " ", " ")
	fmt.Println(string(pretty))
}
