package nestpay

import (
	"context"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"encoding/xml"
	"errors"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

var EndPoints map[string]string = map[string]string{
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

var Currencies map[string]string = map[string]string{
	"TRY": "949",
	"YTL": "949",
	"TRL": "949",
	"TL":  "949",
	"USD": "840",
	"EUR": "978",
	"GBP": "826",
	"JPY": "392",
}

var CurrencyISO map[string]string = map[string]string{
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
	XMLName         xml.Name  `xml:"CC5Request,omitempty"`
	Username        string    `xml:"Name,omitempty"`
	Password        string    `xml:"Password,omitempty"`
	ClientId        string    `xml:"ClientId,omitempty" form:"clientid,omitempty"`
	OrderId         string    `xml:"OrderId,omitempty" form:"oid,omitempty"`
	GroupId         string    `xml:"GroupId,omitempty"`
	TransId         string    `xml:"TransId,omitempty"`
	UserId          string    `xml:"UserId,omitempty"`
	IPAddress       string    `xml:"IPAddress,omitempty"`
	Email           string    `xml:"Email,omitempty"`
	Mode            string    `xml:"Mode,omitempty"`
	StoreType       string    `xml:",omitempty" form:"storetype,omitempty"`
	TransactionType string    `xml:"Type,omitempty" form:"islemtipi,omitempty"`
	CardNumber      string    `xml:"Number,omitempty" form:"pan,omitempty"`
	CardMonth       string    `xml:",omitempty" form:"Ecom_Payment_Card_ExpDate_Month,omitempty"`
	CardYear        string    `xml:",omitempty" form:"Ecom_Payment_Card_ExpDate_Year,omitempty"`
	CardExpiry      string    `xml:"Expires,omitempty"`
	CardCode        string    `xml:"Cvv2Val,omitempty" form:"cv2,omitempty"`
	Total           string    `xml:"Total,omitempty" form:"amount,omitempty"`
	Currency        string    `xml:"Currency,omitempty" form:"currency,omitempty"`
	Installment     string    `xml:"Instalment,omitempty" form:"taksit,omitempty"`
	XID             string    `xml:"PayerTxnId,omitempty"`
	ECI             string    `xml:"PayerSecurityLevel,omitempty"`
	CAVV            string    `xml:"PayerAuthenticationCode,omitempty"`
	PresentCode     string    `xml:"CardholderPresentCode,omitempty"`
	BillTo          *To       `xml:"BillTo,omitempty"`
	ShipTo          *To       `xml:"ShipTo,omitempty"`
	PbOrder         *Pb       `xml:"PbOrder,omitempty"`
	OrderItemList   *ItemList `xml:"OrderItemList,omitempty"`
	Random          string    `xml:",omitempty" form:"rnd,omitempty"`
	Hash            string    `xml:",omitempty" form:"hash,omitempty"`
	OkUrl           string    `xml:",omitempty" form:"okUrl,omitempty"`
	FailUrl         string    `xml:",omitempty" form:"failUrl,omitempty"`
	VersionInfo     string    `xml:"VersionInfo,omitempty"`
}

type To struct {
	Name       string `xml:"Name,omitempty" form:"cardholder,omitempty"`
	Company    string `xml:"Company,omitempty"`
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
	rand.Seed(time.Now().UnixNano())
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

func SHA1(data string) (hash string) {
	h := sha1.New()
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

func Hash(data string) string {
	return B64(HEX(SHA1(data)))
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
	request.Currency = Currencies[currency]
}

func (request *Request) SetInstallment(ins string) {
	request.Installment = ins
}

func (request *Request) SetOrderId(oid string) {
	if oid != "" {
		request.OrderId = oid
	}
}

func (api *API) PreAuth(ctx context.Context, req *Request) (Response, error) {
	req.TransactionType = "PreAuth"
	return api.Transaction(ctx, req)
}

func (api *API) Auth(ctx context.Context, req *Request) (Response, error) {
	req.TransactionType = "Auth"
	return api.Transaction(ctx, req)
}

func (api *API) PreAuth3D(ctx context.Context, req *Request) (Response, error) {
	req.TransactionType = "PreAuth"
	return api.Transaction(ctx, req)
}

func (api *API) Auth3D(ctx context.Context, req *Request) (Response, error) {
	req.TransactionType = "Auth"
	return api.Transaction(ctx, req)
}

func (api *API) PreAuth3Dhtml(ctx context.Context, req *Request) (string, error) {
	req.StoreType = "3d"
	req.TransactionType = "PreAuth"
	req.Random = Random(6)
	req.Hash = Hash(req.ClientId + req.OrderId + req.Total + req.OkUrl + req.FailUrl + req.TransactionType + req.Installment + req.Random + api.Key)
	return api.Transaction3D(ctx, req)
}

func (api *API) Auth3Dhtml(ctx context.Context, req *Request) (string, error) {
	req.StoreType = "3d"
	req.TransactionType = "Auth"
	req.Random = Random(6)
	req.Hash = Hash(req.ClientId + req.OrderId + req.Total + req.OkUrl + req.FailUrl + req.TransactionType + req.Installment + req.Random + api.Key)
	return api.Transaction3D(ctx, req)
}

func (api *API) PostAuth(ctx context.Context, req *Request) (Response, error) {
	req.TransactionType = "PostAuth"
	return api.Transaction(ctx, req)
}

func (api *API) Refund(ctx context.Context, req *Request) (Response, error) {
	req.TransactionType = "Credit"
	return api.Transaction(ctx, req)
}

func (api *API) Cancel(ctx context.Context, req *Request) (Response, error) {
	req.TransactionType = "Void"
	return api.Transaction(ctx, req)
}

func (api *API) Transaction(ctx context.Context, req *Request) (res Response, err error) {
	postdata, err := xml.Marshal(req)
	if err != nil {
		return res, err
	}
	request, err := http.NewRequestWithContext(ctx, "POST", EndPoints[api.Bank], strings.NewReader(xml.Header+string(postdata)))
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
	if err := decoder.Decode(&res); err != nil {
		return res, err
	}
	switch res.ProcReturnCode {
	case "00":
		return res, nil
	default:
		return res, errors.New(res.ErrMsg)
	}
}

func (api *API) Transaction3D(ctx context.Context, req *Request) (res string, err error) {
	postdata, err := QueryString(req)
	if err != nil {
		return res, err
	}
	html := []string{}
	html = append(html, `<!DOCTYPE html>`)
	html = append(html, `<html>`)
	html = append(html, `<head>`)
	html = append(html, `<script type="text/javascript">function submitonload() {document.payment.submit();document.getElementById('button').remove();document.getElementById('body').insertAdjacentHTML("beforeend", "Lütfen bekleyiniz...");}</script>`)
	html = append(html, `</head>`)
	html = append(html, `<body onload="javascript:submitonload();" id="body" style="text-align:center;margin:10px;font-family:Arial;font-weight:bold;">`)
	html = append(html, `<form action="`+EndPoints[api.Bank+"3D"]+`" method="post" name="payment">`)
	for k := range postdata {
		html = append(html, `<input type="hidden" name="`+k+`" value="`+postdata.Get(k)+`">`)
	}
	html = append(html, `<input type="submit" value="Gönder" id="button">`)
	html = append(html, `</form>`)
	html = append(html, `</body>`)
	html = append(html, `</html>`)
	res = B64(strings.Join(html, "\n"))
	return res, err
}
