package main

import (
	"fmt"
	"nestpay/config"
	"nestpay/src"
)

func init() {
	config.Client = "" // Müşteri numarası
	config.User = ""   // Kullanıcı adı
	config.Pass = ""   // Şifre
}

func main() {
	api := nestpay.API{}
	request := nestpay.Request{}
	request.Name = config.User
	request.Password = config.Pass
	request.ClientId = config.Client
	// Ödeme
	request.Type = "Auth"
	request.Mode = "P"
	request.IPAddress = ""    // Müşteri IP adresi
	request.Number = ""       // Kart numarası
	request.Expires = "xx/xx" // Kart son kullanma tarihi
	request.Cvv2Val = ""      // Kart Cvv2 Kodu
	request.Total = "0.00"    // Satış tutarı
	request.Currency = "949"  // Para birimi (949 : TRY)
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
