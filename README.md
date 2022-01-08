[![license](https://img.shields.io/:license-mit-blue.svg)](https://github.com/ozgur-soft/nestpay.go/blob/master/LICENSE.md)
[![documentation](https://pkg.go.dev/badge/github.com/ozgur-soft/nestpay.go)](https://pkg.go.dev/github.com/ozgur-soft/nestpay.go/src)

# Nestpay.go
NestPay (EST) (Asseco, Akbank, İş Bankası, Ziraat Bankası, Halkbank, Finansbank, TEB) Virtual POS API with golang

# Installation
```bash
go get github.com/ozgur-soft/nestpay.go
```

# Sanalpos satış işlemi
```go
package main

import (
	"context"
	"encoding/xml"
	"fmt"

	nestpay "github.com/ozgur-soft/nestpay.go/src"
)

func main() {
	// Banka adı : "akbank","isbank","ziraatbank","halkbank","finansbank","teb"
	api, req := nestpay.Api("banka adı", "müşteri no", "kullanıcı adı", "şifre")
	// Test : "T" - Production "P" (zorunlu)
	req.SetMode("P")
	// Müşteri IPv4 adresi (zorunlu)
	req.SetIPAddress("1.2.3.4")
	// Kart numarası (zorunlu)
	req.SetCardNumber("4242424242424242")
	// Son kullanma tarihi (Ay ve yılın son 2 hanesi) AA,YY (zorunlu)
	req.SetCardExpiry("02", "20")
	// Cvv2 kodu (kartın arka yüzündeki 3 haneli numara) (zorunlu)
	req.SetCardCode("000")
	// Satış tutarı (zorunlu)
	req.SetAmount("1.00")
	// Para birimi (zorunlu)
	req.SetCurrency("TRY")
	// Taksit sayısı (varsa)
	req.SetInstalment("")

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

	nestpay "github.com/ozgur-soft/nestpay.go/src"
)

func main() {
	// Banka adı : "akbank","isbank","ziraatbank","halkbank","finansbank","teb"
	api, req := nestpay.Api("banka adı", "müşteri no", "kullanıcı adı", "şifre")
	// Test : "T" - Production "P" (zorunlu)
	req.SetMode("P")
	// Sipariş numarası (zorunlu)
	req.SetOrderId("ORDER-")
	// İade tutarı (zorunlu)
	req.SetAmount("1.00")
	// Para birimi (zorunlu)
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

	nestpay "github.com/ozgur-soft/nestpay.go/src"
)

func main() {
	// Banka adı : "akbank","isbank","ziraatbank","halkbank","finansbank","teb"
	api, req := nestpay.Api("banka adı", "müşteri no", "kullanıcı adı", "şifre")
	// Test : "T" - Production "P" (zorunlu)
	req.SetMode("P")
	// Sipariş numarası (zorunlu)
	req.SetOrderId("ORDER-")
	// İptal tutarı (zorunlu)
	req.SetAmount("1.00")
	// Para birimi (zorunlu)
	req.SetCurrency("TRY")

	// İptal
	ctx := context.Background()
	res := api.Cancel(ctx, req)
	pretty, _ := xml.MarshalIndent(res, " ", " ")
	fmt.Println(string(pretty))
}
```
