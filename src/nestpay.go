package nestpay

import (
	"context"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"encoding/xml"
	"errors"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"

	"golang.org/x/net/html/charset"
)

var EndPoints = map[string]string{
	"asseco":   "https://entegrasyon.asseco-see.com.tr/fim/api",
	"asseco3D": "https://entegrasyon.asseco-see.com.tr/fim/est3Dgate",

	"anadolu":   "https://anadolusanalpos.est.com.tr/fim/api",
	"anadolu3D": "https://anadolusanalpos.est.com.tr/fim/est3Dgate",

	"akbank":   "https://www.sanalakpos.com/fim/api",
	"akbank3D": "https://www.sanalakpos.com/fim/est3Dgate",

	"isbank":   "https://spos.isbank.com.tr/fim/api",
	"isbank3D": "https://spos.isbank.com.tr/fim/est3Dgate",

	"ziraatbank":   "https://sanalpos2.ziraatbank.com.tr/fim/api",
	"ziraatbank3D": "https://sanalpos2.ziraatbank.com.tr/fim/est3Dgate",

	"halkbank":   "https://sanalpos.halkbank.com.tr/fim/api",
	"halkbank3D": "https://sanalpos.halkbank.com.tr/fim/est3Dgate",

	"finansbank":   "https://www.fbwebpos.com/fim/api",
	"finansbank3D": "https://www.fbwebpos.com/fim/est3Dgate",

	"teb":   "https://sanalpos.teb.com.tr/fim/api",
	"teb3D": "https://sanalpos.teb.com.tr/fim/est3Dgate",
}

var CurrencyCode = map[string]string{
	"TRY": "949",
	"YTL": "949",
	"TRL": "949",
	"TL":  "949",
	"USD": "840",
	"EUR": "978",
	"GBP": "826",
	"JPY": "392",
}

var CurrencyISO = map[string]string{
	"949": "TRY",
	"840": "USD",
	"978": "EUR",
	"826": "GBP",
	"392": "JPY",
}

type API struct {
	Bank string
	Key  string
}

type Request struct {
	XMLName       xml.Name  `xml:"CC5Request,omitempty"`
	Username      string    `xml:"Name,omitempty"`
	Password      string    `xml:"Password,omitempty"`
	ClientId      string    `xml:"ClientId,omitempty" form:"clientid,omitempty"`
	OrderId       string    `xml:"OrderId,omitempty" form:"oid,omitempty"`
	GroupId       string    `xml:"GroupId,omitempty"`
	TransId       string    `xml:"TransId,omitempty"`
	UserId        string    `xml:"UserId,omitempty"`
	IPAddress     string    `xml:"IPAddress,omitempty" form:"clientip,omitempty"`
	Email         string    `xml:"Email,omitempty"`
	Mode          string    `xml:"Mode,omitempty"`
	StoreType     string    `xml:",omitempty" form:"storetype,omitempty"`
	IslemTipi     string    `xml:"Type,omitempty" form:"islemtipi,omitempty"`
	TranType      string    `xml:"TranType,omitempty" form:"TranType,omitempty"`
	CardNumber    string    `xml:"Number,omitempty" form:"pan,omitempty"`
	CardMonth     string    `xml:",omitempty" form:"Ecom_Payment_Card_ExpDate_Month,omitempty"`
	CardYear      string    `xml:",omitempty" form:"Ecom_Payment_Card_ExpDate_Year,omitempty"`
	CardExpiry    string    `xml:"Expires,omitempty"`
	CardCode      string    `xml:"Cvv2Val,omitempty" form:"cv2,omitempty"`
	Total         string    `xml:"Total,omitempty" form:"amount,omitempty"`
	Currency      string    `xml:"Currency,omitempty" form:"currency,omitempty"`
	Installment   string    `xml:"Instalment,omitempty" form:"Instalment,omitempty"`
	Taksit        string    `xml:"Instalment,omitempty" form:"taksit,omitempty"`
	XID           string    `xml:"PayerTxnId,omitempty"`
	ECI           string    `xml:"PayerSecurityLevel,omitempty"`
	CAVV          string    `xml:"PayerAuthenticationCode,omitempty"`
	PresentCode   string    `xml:"CardholderPresentCode,omitempty"`
	BillTo        *To       `xml:"BillTo,omitempty"`
	ShipTo        *To       `xml:"ShipTo,omitempty"`
	PbOrder       *Pb       `xml:"PbOrder,omitempty"`
	OrderItemList *ItemList `xml:"OrderItemList,omitempty"`
	Random        string    `xml:",omitempty" form:"rnd,omitempty"`
	Hash          string    `xml:",omitempty" form:"hash,omitempty"`
	HashAlgorithm string    `xml:",omitempty" form:"hashAlgorithm,omitempty"`
	OkUrl         string    `xml:",omitempty" form:"okUrl,omitempty"`
	FailUrl       string    `xml:",omitempty" form:"failUrl,omitempty"`
	Lang          string    `xml:"lang,omitempty" form:"lang,omitempty"`
	VersionInfo   string    `xml:"VersionInfo,omitempty"`
}

type To struct {
	Name       string `xml:"Name,omitempty" form:"BillToName,omitempty"`
	Company    string `xml:"Company,omitempty" form:"BillToCompany,omitempty"`
	Street1    string `xml:"Street1,omitempty"`
	Street2    string `xml:"Street2,omitempty"`
	Street3    string `xml:"Street3,omitempty"`
	City       string `xml:"City,omitempty"`
	StateProv  string `xml:"StateProv,omitempty"`
	PostalCode string `xml:"PostalCode,omitempty"`
	Country    string `xml:"Country,omitempty"`
	Phone      string `xml:"TelVoice,omitempty" form:"phone,omitempty"`
}

type Pb struct {
	OrderType              string `xml:"OrderType,omitempty"`
	TotalNumberPayments    string `xml:"TotalNumberPayments,omitempty"`
	OrderFrequencyCycle    string `xml:"OrderFrequencyCycle,omitempty"`
	OrderFrequencyInterval string `xml:"OrderFrequencyInterval,omitempty"`
	Desc                   string `xml:"Desc,omitempty"`
	Price                  string `xml:"Price,omitempty"`
	Total                  string `xml:"Total,omitempty"`
}

type Item struct {
	Id          string `xml:"Id,omitempty"`
	ItemNumber  string `xml:"ItemNumber,omitempty"`
	ProductCode string `xml:"ProductCode,omitempty"`
	Qty         string `xml:"Qty,omitempty"`
	Desc        string `xml:"Desc,omitempty"`
	Price       string `xml:"Price,omitempty"`
	Total       string `xml:"Total,omitempty"`
}

type ItemList struct {
	Items []*Item `xml:"OrderItem,omitempty"`
}

type Response struct {
	XMLName        xml.Name `xml:"CC5Response,omitempty"`
	OrderId        string   `xml:"OrderId,omitempty"`
	GroupId        string   `xml:"GroupId,omitempty"`
	TransId        string   `xml:"TransId,omitempty"`
	Response       string   `xml:"Response,omitempty"`
	AuthCode       string   `xml:"AuthCode,omitempty"`
	HostRefNum     string   `xml:"HostRefNum,omitempty"`
	ProcReturnCode string   `xml:"ProcReturnCode,omitempty"`
	ErrMsg         string   `xml:"ErrMsg,omitempty"`
}

func IPv4(r *http.Request) (ip string) {
	ipv4 := []string{
		r.Header.Get("X-Real-Ip"),
		r.Header.Get("X-Forwarded-For"),
		r.RemoteAddr,
	}
	for _, ipaddress := range ipv4 {
		if ipaddress != "" {
			ip = ipaddress
			break
		}
	}
	return strings.Split(ip, ":")[0]
}

func Random(n int) string {
	const alphanum = "123456789"
	var bytes = make([]byte, n)
	source := rand.NewSource(time.Now().UnixNano())
	rand := rand.New(source)
	rand.Read(bytes)
	for i, b := range bytes {
		bytes[i] = alphanum[b%byte(len(alphanum))]
	}
	return string(bytes)
}

func HEX(data string) (hash string) {
	b, err := hex.DecodeString(data)
	if err != nil {
		log.Println(err)
		return hash
	}
	hash = string(b)
	return hash
}

func SHA512(data string) (hash string) {
	h := sha512.New()
	h.Write([]byte(data))
	hash = hex.EncodeToString(h.Sum(nil))
	return hash
}

func B64(data string) (hash string) {
	hash = base64.StdEncoding.EncodeToString([]byte(data))
	return hash
}

func D64(data string) []byte {
	b, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		log.Println(err)
		return nil
	}
	return b
}

func Hash(form url.Values, keys []string, secret string) string {
	hash := []string{}
	for _, k := range keys {
		hash = append(hash, form.Get(k))
	}
	hash = append(hash, secret)
	return B64(HEX(SHA512(strings.Join(hash, "|"))))
}

func Api(bank, clientid, username, password string) (*API, *Request) {
	api := new(API)
	api.Bank = bank
	request := new(Request)
	request.ClientId = clientid
	request.Username = username
	request.Password = password
	request.BillTo = new(To)
	return api, request
}

func (api *API) SetStoreKey(key string) {
	api.Key = key
}

func (request *Request) SetMode(mode string) {
	switch mode {
	case "TEST":
		request.Mode = "T"
	case "PROD":
		request.Mode = "P"
	default:
		request.Mode = mode
	}
}

func (request *Request) SetIPAddress(ip string) {
	request.IPAddress = ip
}

func (request *Request) SetPhoneNumber(phone string) {
	request.BillTo.Phone = phone
}

func (request *Request) SetCardHolder(holder string) {
	request.BillTo.Name = holder
}

func (request *Request) SetCardNumber(number string) {
	request.CardNumber = number
}

func (request *Request) SetCardExpiry(month, year string) {
	request.CardExpiry = month + "/" + year
	request.CardMonth = month
	request.CardYear = year
}

func (request *Request) SetCardCode(code string) {
	request.CardCode = code
}

func (request *Request) SetAmount(total string, currency string) {
	request.Total = total
	request.Currency = CurrencyCode[currency]
}

func (request *Request) SetInstallment(ins string) {
	request.Installment = ins
}

func (request *Request) SetTaksit(ins string) {
	request.Taksit = ins
}

func (request *Request) SetOrderId(oid string) {
	if oid != "" {
		request.OrderId = oid
	}
}

func (api *API) PreAuth(ctx context.Context, req *Request) (Response, error) {
	switch api.Bank {
	case "halkbank":
		req.TranType = "PreAuth"
	default:
		req.IslemTipi = "PreAuth"
	}
	return api.Transaction(ctx, req)
}

func (api *API) Auth(ctx context.Context, req *Request) (Response, error) {
	switch api.Bank {
	case "halkbank":
		req.TranType = "Auth"
	default:
		req.IslemTipi = "Auth"
	}
	return api.Transaction(ctx, req)
}

func (api *API) PreAuth3D(ctx context.Context, req *Request) (Response, error) {
	switch api.Bank {
	case "halkbank":
		req.TranType = "PreAuth"
	default:
		req.IslemTipi = "PreAuth"
	}
	return api.Transaction(ctx, req)
}

func (api *API) Auth3D(ctx context.Context, req *Request) (Response, error) {
	switch api.Bank {
	case "halkbank":
		req.TranType = "Auth"
	default:
		req.IslemTipi = "Auth"
	}
	return api.Transaction(ctx, req)
}

func (api *API) PreAuth3Dhtml(ctx context.Context, req *Request) (string, error) {
	switch api.Bank {
	case "halkbank":
		req.TranType = "PreAuth"
	default:
		req.IslemTipi = "PreAuth"
	}
	req.HashAlgorithm = "ver3"
	req.StoreType = "3d"
	req.Random = Random(6)
	form, err := QueryString(req)
	if err == nil {
		keys := []string{}
		for k := range form {
			if strings.ToLower(k) != "hash" && strings.ToLower(k) != "encoding" {
				keys = append(keys, k)
			}
		}
		sort.Slice(keys, func(i, j int) bool { return strings.ToLower(keys[i]) < strings.ToLower(keys[j]) })
		req.Hash = Hash(form, keys, api.Key)
	}
	return api.Transaction3D(ctx, req)
}

func (api *API) Auth3Dhtml(ctx context.Context, req *Request) (string, error) {
	switch api.Bank {
	case "halkbank":
		req.TranType = "Auth"
	default:
		req.IslemTipi = "Auth"
	}
	req.HashAlgorithm = "ver3"
	req.StoreType = "3d"
	req.Random = Random(6)
	form, err := QueryString(req)
	if err == nil {
		keys := []string{}
		for k := range form {
			if strings.ToLower(k) != "hash" && strings.ToLower(k) != "encoding" {
				keys = append(keys, k)
			}
		}
		sort.Slice(keys, func(i, j int) bool { return strings.ToLower(keys[i]) < strings.ToLower(keys[j]) })
		req.Hash = Hash(form, keys, api.Key)
	}
	return api.Transaction3D(ctx, req)
}

func (api *API) PostAuth(ctx context.Context, req *Request) (Response, error) {
	switch api.Bank {
	case "halkbank":
		req.TranType = "PostAuth"
	default:
		req.IslemTipi = "PostAuth"
	}
	return api.Transaction(ctx, req)
}

func (api *API) Refund(ctx context.Context, req *Request) (Response, error) {
	switch api.Bank {
	case "halkbank":
		req.TranType = "Credit"
	default:
		req.IslemTipi = "Credit"
	}
	return api.Transaction(ctx, req)
}

func (api *API) Cancel(ctx context.Context, req *Request) (Response, error) {
	switch api.Bank {
	case "halkbank":
		req.TranType = "Void"
	default:
		req.IslemTipi = "Void"
	}
	return api.Transaction(ctx, req)
}

func (api *API) Transaction(ctx context.Context, req *Request) (res Response, err error) {
	payload, err := xml.Marshal(req)
	if err != nil {
		return res, err
	}
	request, err := http.NewRequestWithContext(ctx, "POST", EndPoints[api.Bank], strings.NewReader(xml.Header+string(payload)))
	if err != nil {
		return res, err
	}
	request.Header.Set("Content-Type", "text/xml; charset=utf-8")
	client := new(http.Client)
	response, err := client.Do(request)
	if err != nil {
		return res, err
	}
	defer response.Body.Close()
	decoder := xml.NewDecoder(response.Body)
	decoder.CharsetReader = charset.NewReaderLabel
	if err := decoder.Decode(&res); err != nil {
		return res, err
	}
	if code, err := strconv.Atoi(res.ProcReturnCode); err == nil {
		switch code {
		case 0:
			return res, nil
		default:
			return res, errors.New(res.ErrMsg)
		}
	} else {
		return res, errors.New(res.ErrMsg)
	}
}

func (api *API) Transaction3D(ctx context.Context, req *Request) (res string, err error) {
	payload, err := QueryString(req)
	if err != nil {
		return res, err
	}
	keys := []string{}
	for k := range payload {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool { return strings.ToLower(keys[i]) < strings.ToLower(keys[j]) })
	html := []string{}
	html = append(html, `<!DOCTYPE html>`)
	html = append(html, `<html>`)
	html = append(html, `<head>`)
	html = append(html, `<meta http-equiv="Content-Type" content="text/html; charset=utf-8">`)
	html = append(html, `<script type="text/javascript">function submitonload() {document.payment.submit();document.getElementById('button').remove();document.getElementById('body').insertAdjacentHTML("beforeend", "Lütfen bekleyiniz...");}</script>`)
	html = append(html, `</head>`)
	html = append(html, `<body onload="javascript:submitonload();" id="body" style="text-align:center;margin:10px;font-family:Arial;font-weight:bold;">`)
	html = append(html, `<form action="`+EndPoints[api.Bank+"3D"]+`" method="post" name="payment">`)
	for _, k := range keys {
		html = append(html, `<input type="hidden" name="`+k+`" value="`+payload.Get(k)+`">`)
	}
	html = append(html, `<input type="hidden" name="encoding" value="UTF-8">`)
	html = append(html, `<input type="submit" value="Gönder" id="button">`)
	html = append(html, `</form>`)
	html = append(html, `</body>`)
	html = append(html, `</html>`)
	res = B64(strings.Join(html, "\n"))
	return res, err
}
