package main

import (
	"fmt"

	nestpay "github.com/OzqurYalcin/nestpay/src"
)

func main() {
	api := &nestpay.API{"akbank"} // "akbank","asseco","isbank","finansbank","denizbank","kuveytturk","halkbank","anadolubank","hsbc","ziraatbank"
	request := new(nestpay.Request)
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
	request.Currency = nestpay.Currencies["TRY"] // Para birimi
	// Fatura
	request.BillTo.Name = ""    // Kart sahibi
	request.BillTo.Company = "" // Fatura unvanı
	// 3D (varsa)
	request.PayerTxnId = nil
	request.PayerSecurityLevel = nil
	request.PayerAuthenticationCode = nil
	request.CardholderPresentCode = nil
	response := api.Transaction(request)
	if response.ProcReturnCode != "00" {
		if response.ErrMsg == "" {
			response.ErrMsg = "Banka bağlantısında hata oluştu"
		}
		fmt.Println(response.ProcReturnCode, response.ErrMsg)
	} else {
		fmt.Println(response.Response)
	}
}
