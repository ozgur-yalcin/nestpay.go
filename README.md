[![license](https://img.shields.io/:license-mit-blue.svg)](https://github.com/ozgur-soft/nestpay/blob/master/LICENSE.md)
[![documentation](https://pkg.go.dev/badge/github.com/ozgur-soft/nestpay)](https://pkg.go.dev/github.com/ozgur-soft/nestpay/src)

# Nestpay
NestPay (EST) (Asseco, Akbank, İş Bankası, Ziraat Bankası, Halkbank, Finansbank, TEB) Sanal POS API with golang

# Installation
```bash
go get github.com/ozgur-soft/nestpay
```

# Sanalpos satış işlemi
```go
package main

import (
	"context"
	"encoding/xml"
	"fmt"

	nestpay "github.com/ozgur-soft/nestpay/src"
)

func main() {
	// banka adı : "akbank","isbank","ziraatbank","halkbank","finansbank","teb"
	api, req := nestpay.Api("banka adı", "müşteri no", "kullanıcı adı", "şifre")
	// TEST : "T" - PRODUCTION "P"
	req.SetMode("P")
	// Müşteri IP adresi (zorunlu)
	req.SetIPAddress("1.2.3.4")
	// Kart numarası
	req.SetCardNumber("4242424242424242")
	// Son kullanma tarihi (Ay ve Yılın son 2 hanesi) AA,YY
	req.SetExpires("02", "20")
	// Cvv2 Kodu (kartın arka yüzündeki 3 haneli numara)
	req.SetCvv2("000")
	// Satış tutarı
	req.SetAmount("1.00")
	// Taksit sayısı
	req.SetInstalment("")
	// Para birimi
	req.SetCurrency("TRY")

	// Fatura
	req.BillTo = new(nestpay.To)
	req.BillTo.Name = ""     // Kart sahibi
	req.BillTo.TelVoice = "" // Telefon numarası

	// 3D (varsa)
	//req.PayerTxnId = ""
	//req.PayerSecurityLevel = ""
	//req.PayerAuthenticationCode = ""
	//req.CardholderPresentCode = ""

	// Satış
	ctx := context.Background()
	res := api.Pay(ctx, req)
	pretty, _ := xml.MarshalIndent(res, " ", " ")
	fmt.Println(string(pretty))
}
```

# Sanalpos iade işlemi
```go
package main

import (
	"context"
	"encoding/xml"
	"fmt"

	nestpay "github.com/ozgur-soft/nestpay/src"
)

func main() {
	// banka adı : "akbank","isbank","ziraatbank","halkbank","finansbank","teb"
	api, req := nestpay.Api("banka adı", "müşteri no", "kullanıcı adı", "şifre")
	// TEST : "T" - PRODUCTION "P"
	req.SetMode("P")
	// Sipariş numarası
	req.SetOrderId("ORDER-")
	// Satış tutarı
	req.SetAmount("1.00")
	// Para birimi
	req.SetCurrency("TRY")

	// İade
	ctx := context.Background()
	res := api.Refund(ctx, req)
	pretty, _ := xml.MarshalIndent(res, " ", " ")
	fmt.Println(string(pretty))
}
```

# Sanalpos iptal işlemi
```go
package main

import (
	"context"
	"encoding/xml"
	"fmt"

	nestpay "github.com/ozgur-soft/nestpay/src"
)

func main() {
	// banka adı : "akbank","isbank","ziraatbank","halkbank","finansbank","teb"
	api, req := nestpay.Api("banka adı", "müşteri no", "kullanıcı adı", "şifre")
	// TEST : "T" - PRODUCTION "P"
	req.SetMode("P")
	// Sipariş numarası
	req.SetOrderId("ORDER-")
	// Satış tutarı
	req.SetAmount("1.00")
	// Para birimi
	req.SetCurrency("TRY")

	// İade
	ctx := context.Background()
	res := api.Cancel(ctx, req)
	pretty, _ := xml.MarshalIndent(res, " ", " ")
	fmt.Println(string(pretty))
}
```
