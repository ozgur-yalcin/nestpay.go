[![license](https://img.shields.io/:license-mit-blue.svg)](https://github.com/ozgur-yalcin/nestpay.go/blob/master/LICENSE.md)
[![documentation](https://pkg.go.dev/badge/github.com/ozgur-yalcin/nestpay.go)](https://pkg.go.dev/github.com/ozgur-yalcin/nestpay.go/src)

# Nestpay.go
NestPay POS API with golang

# Installation
```bash
go get github.com/ozgur-yalcin/nestpay.go
```

# Satış
```go
package main

import (
	"context"
	"encoding/xml"
	"fmt"

	nestpay "github.com/ozgur-yalcin/nestpay.go/src"
)

// Pos bilgileri
const (
	bankname = "asseco"    // Banka adı : "asseco","akbank","isbank","ziraatbank","halkbank","finansbank","teb"
	envmode  = "TEST"      // Çalışma ortamı (Production : "PROD" - Test : "TEST")
	clientid = "100100000" // Müşteri numarası
	username = "AKTESTAPI" // Kullanıcı adı
	password = "AKBANK01"  // Şifre
)

func main() {
	api, req := nestpay.Api(bankname, clientid, username, password)
	req.SetMode(envmode)

	req.SetIPAddress("1.2.3.4")           // Müşteri IPv4 adresi (zorunlu)
	req.SetCardNumber("4355084355084358") // Kart numarası (zorunlu)
	req.SetCardExpiry("12", "30")         // Son kullanma tarihi - AA,YY (zorunlu)
	req.SetCardCode("000")                // Kart arkasındaki 3 haneli numara (zorunlu)
	req.SetAmount("1.00", "TRY")          // Satış tutarı ve para birimi (zorunlu)
	req.SetInstallment("")                // Taksit sayısı (varsa)

	// Kişisel bilgiler (zorunlu)
	req.BillTo = new(nestpay.To)
	req.BillTo.Name = ""  // Kart sahibi
	req.BillTo.Phone = "" // Telefon numarası

	// Satış
	ctx := context.Background()
	if res, err := api.Auth(ctx, req); err == nil {
		pretty, _ := xml.MarshalIndent(res, " ", " ")
		fmt.Println(string(pretty))
	} else {
		fmt.Println(err)
	}
}
```

# İade
```go
package main

import (
	"context"
	"encoding/xml"
	"fmt"

	nestpay "github.com/ozgur-yalcin/nestpay.go/src"
)

// Pos bilgileri
const (
	bankname = "asseco"    // Banka adı : "asseco","akbank","isbank","ziraatbank","halkbank","finansbank","teb"
	envmode  = "TEST"      // Çalışma ortamı (Production : "PROD" - Test : "TEST")
	clientid = "100100000" // Müşteri numarası
	username = "AKTESTAPI" // Kullanıcı adı
	password = "AKBANK01"  // Şifre
)

func main() {
	api, req := nestpay.Api(bankname, clientid, username, password)
	req.SetMode(envmode)

	req.SetAmount("1.00", "TRY") // İade tutarı ve para birimi (zorunlu)
	req.SetOrderId("ORDER-")     // Sipariş numarası (zorunlu)

	// İade
	ctx := context.Background()
	if res, err := api.Refund(ctx, req); err == nil {
		pretty, _ := xml.MarshalIndent(res, " ", " ")
		fmt.Println(string(pretty))
	} else {
		fmt.Println(err)
	}
}
```

# İptal
```go
package main

import (
	"context"
	"encoding/xml"
	"fmt"

	nestpay "github.com/ozgur-yalcin/nestpay.go/src"
)

// Pos bilgileri
const (
	bankname = "asseco"    // Banka adı : "asseco","akbank","isbank","ziraatbank","halkbank","finansbank","teb"
	envmode  = "TEST"      // Çalışma ortamı (Production : "PROD" - Test : "TEST")
	clientid = "100100000" // Müşteri numarası
	username = "AKTESTAPI" // Kullanıcı adı
	password = "AKBANK01"  // Şifre
)

func main() {
	api, req := nestpay.Api(bankname, clientid, username, password)
	req.SetMode(envmode)

	req.SetOrderId("ORDER-") // Sipariş numarası (zorunlu)

	// İptal
	ctx := context.Background()
	if res, err := api.Cancel(ctx, req); err == nil {
		pretty, _ := xml.MarshalIndent(res, " ", " ")
		fmt.Println(string(pretty))
	} else {
		fmt.Println(err)
	}
}
```
