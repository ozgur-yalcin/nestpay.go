package nestpay

import (
	"encoding/xml"
	"log"
	"net/http"
	"strings"
)

var EndPoints map[string]string = map[string]string{
	"asseco":      "https://entegrasyon.asseco-see.com.tr/fim/api",
	"akbank":      "https://www.sanalakpos.com/fim/api",
	"isbank":      "https://spos.isbank.com.tr/fim/api",
	"finansbank":  "https://www.fbwebpos.com/fim/api",
	"denizbank":   "https://denizbank.est.com.tr/fim/api",
	"kuveytturk":  "https://kuveytturk.est.com.tr/fim/api",
	"halkbank":    "https://sanalpos.halkbank.com.tr/fim/api",
	"anadolubank": "https://anadolusanalpos.est.com.tr/fim/api",
	"hsbc":        "https://vpos.advantage.com.tr/fim/api",
	"ziraatbank":  "https://sanalpos2.ziraatbank.com.tr/fim/api",
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

type API struct {
	Bank string
}

type Request struct {
	XMLName    xml.Name `xml:"CC5Request,omitempty"`
	Username   string   `xml:"Name,omitempty"`
	Password   string   `xml:"Password,omitempty"`
	ClientId   string   `xml:"ClientId,omitempty"`
	OrderId    string   `xml:"OrderId,omitempty"`
	GroupId    string   `xml:"GroupId,omitempty"`
	TransId    string   `xml:"TransId,omitempty"`
	UserId     string   `xml:"UserId,omitempty"`
	IPAddress  string   `xml:"IPAddress,omitempty"`
	Email      string   `xml:"Email,omitempty"`
	Mode       string   `xml:"Mode,omitempty"`
	Type       string   `xml:"Type,omitempty"`
	Number     string   `xml:"Number,omitempty"`
	Expires    string   `xml:"Expires,omitempty"`
	Cvv2Val    string   `xml:"Cvv2Val,omitempty"`
	Total      string   `xml:"Total,omitempty"`
	Currency   string   `xml:"Currency,omitempty"`
	Instalment string   `xml:"Instalment,omitempty"`

	PayerTxnId              string `xml:"PayerTxnId,omitempty"`
	PayerSecurityLevel      string `xml:"PayerSecurityLevel,omitempty"`
	PayerAuthenticationCode string `xml:"PayerAuthenticationCode,omitempty"`
	CardholderPresentCode   string `xml:"CardholderPresentCode,omitempty"`

	BillTo struct {
		Name       string `xml:"Name,omitempty"`
		Company    string `xml:"Company,omitempty"`
		Street1    string `xml:"Street1,omitempty"`
		Street2    string `xml:"Street2,omitempty"`
		Street3    string `xml:"Street3,omitempty"`
		City       string `xml:"City,omitempty"`
		StateProv  string `xml:"StateProv,omitempty"`
		PostalCode string `xml:"PostalCode,omitempty"`
		Country    string `xml:"Country,omitempty"`
		TelVoice   string `xml:"TelVoice,omitempty"`
	} `xml:"BillTo,omitempty"`

	ShipTo struct {
		Name       string `xml:"Name,omitempty"`
		Company    string `xml:"Company,omitempty"`
		Street1    string `xml:"Street1,omitempty"`
		Street2    string `xml:"Street2,omitempty"`
		Street3    string `xml:"Street3,omitempty"`
		City       string `xml:"City,omitempty"`
		StateProv  string `xml:"StateProv,omitempty"`
		PostalCode string `xml:"PostalCode,omitempty"`
		Country    string `xml:"Country,omitempty"`
		TelVoice   string `xml:"TelVoice,omitempty"`
	} `xml:"ShipTo,omitempty"`

	OrderItemList struct {
		OrderItem []struct {
			Id          string `xml:"Id,omitempty"`
			ItemNumber  string `xml:"ItemNumber,omitempty"`
			ProductCode string `xml:"ProductCode,omitempty"`
			Qty         string `xml:"Qty,omitempty"`
			Desc        string `xml:"Desc,omitempty"`
			Price       string `xml:"Price,omitempty"`
			Total       string `xml:"Total,omitempty"`
		} `xml:"OrderItem,omitempty"`
	} `xml:"OrderItemList,omitempty"`

	PbOrder struct {
		OrderType              string `xml:"OrderType,omitempty"`
		TotalNumberPayments    string `xml:"TotalNumberPayments,omitempty"`
		OrderFrequencyCycle    string `xml:"OrderFrequencyCycle,omitempty"`
		OrderFrequencyInterval string `xml:"OrderFrequencyInterval,omitempty"`
		Desc                   string `xml:"Desc,omitempty"`
		Price                  string `xml:"Price,omitempty"`
		Total                  string `xml:"Total,omitempty"`
	} `xml:"PbOrder,omitempty"`

	VersionInfo string `xml:"VersionInfo,omitempty"`
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

func (api *API) Transaction(request Request) (response Response) {
	postdata, _ := xml.Marshal(request)
	res, err := http.Post(EndPoints[api.Bank], "text/xml; charset=utf-8", strings.NewReader(xml.Header+string(postdata)))
	if err != nil {
		log.Println(err)
		return response
	}
	defer res.Body.Close()
	decoder := xml.NewDecoder(res.Body)
	decoder.Decode(&response)
	return response
}
