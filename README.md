# Nestpay
NestPay (EST) (Akbank, İş Bankası, Finansbank, Denizbank, Kuveytturk, Halkbank, Anadolubank, ING Bank, Citibank, Cardplus, Ziraat Bankası) Omnipay Sanalpos API with golang

# Security
If you discover any security related issues, please email ozguryalcin@outlook.com instead of using the issue tracker.

# License
The MIT License (MIT). Please see License File for more information.


# Akbank sanalpos satış işlemi
```go
package main

import (
	"fmt"
	"nestpay/config"
	"nestpay/src"
)

func init() {
	config.Bank = "akbank" // Banka
	config.Client = ""     // Müşteri numarası
	config.User = ""       // Kullanıcı adı
	config.Pass = ""       // Şifre
}

func main() {
	api := nestpay.API{}
	request := nestpay.Request{}
	request.Name = config.User
	request.Password = config.Pass
	request.ClientId = config.Client
	// Ödeme
	request.Type = "Auth"
	request.Mode = "P"                          // TEST : "T" - PRODUCTION "P"
	request.IPAddress = ""                      // Müşteri IP adresi
	request.Number = ""                         // Kart numarası
	request.Expires = "xx/xx"                   // Kart son kullanma tarihi
	request.Cvv2Val = "xxx"                     // Kart Cvv2 Kodu
	request.Total = "0.00"                      // Satış tutarı
	request.Currency = config.Currencies["TRY"] // Para birimi
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
```


# Akbank sanalpos iade işlemi
```go
package main

import (
	"fmt"
	"nestpay/config"
	"nestpay/src"
)

func init() {
	config.Bank = "akbank" // Banka
	config.Client = ""     // Müşteri numarası
	config.User = ""       // Kullanıcı adı
	config.Pass = ""       // Şifre
}

func main() {
	api := nestpay.API{}
	request := nestpay.Request{}
	request.Name = config.User
	request.Password = config.Pass
	request.ClientId = config.Client
	// İade
	request.Type = "Credit"
	request.Mode = "P"                          // TEST : "T" - PRODUCTION "P"
	request.IPAddress = ""                      // Müşteri IP adresi
	request.Number = ""                         // Kart numarası
	request.Expires = "xx/xx"                   // Kart son kullanma tarihi
	request.Cvv2Val = "xxx"                     // Kart Cvv2 Kodu
	request.Total = "0.00"                      // Satış tutarı
	request.Currency = config.Currencies["TRY"] // Para birimi
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
```
