package nestpay

import (
	"encoding/xml"
	"fmt"
	"nestpay/config"
	"net/http"
	"strings"

	"golang.org/x/net/html/charset"
)

type API struct {
}

type Request struct {
	XMLName    xml.Name    `xml:"CC5Request,omitempty"`
	Name       interface{} `xml:"Name,omitempty"`
	Password   interface{} `xml:"Password,omitempty"`
	ClientId   interface{} `xml:"ClientId,omitempty"`
	OrderId    interface{} `xml:"OrderId,omitempty"`
	GroupId    interface{} `xml:"GroupId,omitempty"`
	TransId    interface{} `xml:"TransId,omitempty"`
	UserId     interface{} `xml:"UserId,omitempty"`
	IPAddress  interface{} `xml:"IPAddress,omitempty"`
	Email      interface{} `xml:"Email,omitempty"`
	Mode       interface{} `xml:"Mode,omitempty"`
	Type       interface{} `xml:"Type,omitempty"`
	Number     interface{} `xml:"Number,omitempty"`
	Expires    interface{} `xml:"Expires,omitempty"`
	Cvv2Val    interface{} `xml:"Cvv2Val,omitempty"`
	Total      interface{} `xml:"Total,omitempty"`
	Currency   interface{} `xml:"Currency,omitempty"`
	Instalment interface{} `xml:"Instalment,omitempty"`

	PayerTxnId              interface{} `xml:"PayerTxnId,omitempty"`
	PayerSecurityLevel      interface{} `xml:"PayerSecurityLevel,omitempty"`
	PayerAuthenticationCode interface{} `xml:"PayerAuthenticationCode,omitempty"`
	CardholderPresentCode   interface{} `xml:"CardholderPresentCode,omitempty"`

	BillTo struct {
		Name       interface{} `xml:"Name,omitempty"`
		Company    interface{} `xml:"Company,omitempty"`
		Street1    interface{} `xml:"Street1,omitempty"`
		Street2    interface{} `xml:"Street2,omitempty"`
		Street3    interface{} `xml:"Street3,omitempty"`
		City       interface{} `xml:"City,omitempty"`
		StateProv  interface{} `xml:"StateProv,omitempty"`
		PostalCode interface{} `xml:"PostalCode,omitempty"`
		Country    interface{} `xml:"Country,omitempty"`
		TelVoice   interface{} `xml:"TelVoice,omitempty"`
	} `xml:"BillTo,omitempty"`

	ShipTo struct {
		Name       interface{} `xml:"Name,omitempty"`
		Company    interface{} `xml:"Company,omitempty"`
		Street1    interface{} `xml:"Street1,omitempty"`
		Street2    interface{} `xml:"Street2,omitempty"`
		Street3    interface{} `xml:"Street3,omitempty"`
		City       interface{} `xml:"City,omitempty"`
		StateProv  interface{} `xml:"StateProv,omitempty"`
		PostalCode interface{} `xml:"PostalCode,omitempty"`
		Country    interface{} `xml:"Country,omitempty"`
		TelVoice   interface{} `xml:"TelVoice,omitempty"`
	} `xml:"ShipTo,omitempty"`

	OrderItemList struct {
		OrderItem []struct {
			Id          interface{} `xml:"Id,omitempty"`
			ItemNumber  interface{} `xml:"ItemNumber,omitempty"`
			ProductCode interface{} `xml:"ProductCode,omitempty"`
			Qty         interface{} `xml:"Qty,omitempty"`
			Desc        interface{} `xml:"Desc,omitempty"`
			Price       interface{} `xml:"Price,omitempty"`
			Total       interface{} `xml:"Total,omitempty"`
		} `xml:"OrderItem,omitempty"`
	} `xml:"OrderItemList,omitempty"`

	PbOrder struct {
		OrderType              interface{} `xml:"OrderType,omitempty"`
		TotalNumberPayments    interface{} `xml:"TotalNumberPayments,omitempty"`
		OrderFrequencyCycle    interface{} `xml:"OrderFrequencyCycle,omitempty"`
		OrderFrequencyInterval interface{} `xml:"OrderFrequencyInterval,omitempty"`
		Desc                   interface{} `xml:"Desc,omitempty"`
		Price                  interface{} `xml:"Price,omitempty"`
		Total                  interface{} `xml:"Total,omitempty"`
	} `xml:"PbOrder,omitempty"`

	Extra struct {
		RECURRINGID           interface{} `xml:"RECURRINGID,omitempty"`
		RECURRINGCOUNT        interface{} `xml:"RECURRINGCOUNT,omitempty"`
		AUTH_DTTM             interface{} `xml:"AUTH_DTTM,omitempty"`
		CAPTURE_DTTM          interface{} `xml:"CAPTURE_DTTM,omitempty"`
		PLANNED_START_DTTM    interface{} `xml:"PLANNED_START_DTTM,omitempty"`
		ORIG_TRANS_AMT        interface{} `xml:"ORIG_TRANS_AMT,omitempty"`
		CAPTURE_AMT           interface{} `xml:"CAPTURE_AMT,omitempty"`
		PROC_RET_CD           interface{} `xml:"PROC_RET_CD,omitempty"`
		CHARGE_TYPE_CD        interface{} `xml:"CHARGE_TYPE_CD,omitempty"`
		TRANS_STAT            interface{} `xml:"TRANS_STAT,omitempty"`
		ORDERSTATUS           interface{} `xml:"ORDERSTATUS,omitempty"`
		TRXDATE               interface{} `xml:"TRXDATE,omitempty"`
		TRXCOUNT              interface{} `xml:"TRXCOUNT,omitempty"`
		HOSTDATE              interface{} `xml:"HOSTDATE,omitempty"`
		HOSTMSG               interface{} `xml:"HOSTMSG,omitempty"`
		HOST_REF_NUM          interface{} `xml:"HOST_REF_NUM,omitempty"`
		PAN                   interface{} `xml:"PAN,omitempty"`
		MDSTATUS              interface{} `xml:"MDSTATUS,omitempty"`
		ECI_3D                interface{} `xml:"ECI_3D,omitempty"`
		CAVV_3D               interface{} `xml:"CAVV_3D,omitempty"`
		XID_3D                interface{} `xml:"XID_3D,omitempty"`
		ORD_ID                interface{} `xml:"ORD_ID,omitempty"`
		SETTLEID              interface{} `xml:"SETTLEID,omitempty"`
		TRANS_ID              interface{} `xml:"TRANS_ID,omitempty"`
		ERRORCODE             interface{} `xml:"ERRORCODE,omitempty"`
		AUTH_CODE             interface{} `xml:"AUTH_CODE,omitempty"`
		NUMCODE               interface{} `xml:"NUMCODE,omitempty"`
		ORDERHISTORY          interface{} `xml:"ORDERHISTORY,omitempty"`
		MAILORDER             interface{} `xml:"MAILORDER,omitempty"`
		SUBMERCHANTNAME       interface{} `xml:"SUBMERCHANTNAME,omitempty"`
		SUBMERCHANTID         interface{} `xml:"SUBMERCHANTID,omitempty"`
		SUBMERCHANTPOSTALCODE interface{} `xml:"SUBMERCHANTPOSTALCODE,omitempty"`
		SUBMERCHANTCITY       interface{} `xml:"SUBMERCHANTCITY,omitempty"`
		SUBMERCHANTCOUNTRY    interface{} `xml:"SUBMERCHANTCOUNTRY,omitempty"`
	} `xml:"Extra,omitempty"`

	VersionInfo interface{} `xml:"VersionInfo,omitempty"`
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
	Extra          struct {
		RECURRINGID           string `xml:"RECURRINGID,omitempty"`
		RECURRINGCOUNT        string `xml:"RECURRINGCOUNT,omitempty"`
		AUTH_DTTM             string `xml:"AUTH_DTTM,omitempty"`
		CAPTURE_DTTM          string `xml:"CAPTURE_DTTM,omitempty"`
		PLANNED_START_DTTM    string `xml:"PLANNED_START_DTTM,omitempty"`
		ORIG_TRANS_AMT        string `xml:"ORIG_TRANS_AMT,omitempty"`
		CAPTURE_AMT           string `xml:"CAPTURE_AMT,omitempty"`
		PROC_RET_CD           string `xml:"PROC_RET_CD,omitempty"`
		CHARGE_TYPE_CD        string `xml:"CHARGE_TYPE_CD,omitempty"`
		TRANS_STAT            string `xml:"TRANS_STAT,omitempty"`
		ORDERSTATUS           string `xml:"ORDERSTATUS,omitempty"`
		TRXDATE               string `xml:"TRXDATE,omitempty"`
		TRXCOUNT              string `xml:"TRXCOUNT,omitempty"`
		HOSTDATE              string `xml:"HOSTDATE,omitempty"`
		HOSTMSG               string `xml:"HOSTMSG,omitempty"`
		HOST_REF_NUM          string `xml:"HOST_REF_NUM,omitempty"`
		PAN                   string `xml:"PAN,omitempty"`
		MDSTATUS              string `xml:"MDSTATUS,omitempty"`
		ECI_3D                string `xml:"ECI_3D,omitempty"`
		CAVV_3D               string `xml:"CAVV_3D,omitempty"`
		XID_3D                string `xml:"XID_3D,omitempty"`
		ORD_ID                string `xml:"ORD_ID,omitempty"`
		SETTLEID              string `xml:"SETTLEID,omitempty"`
		TRANS_ID              string `xml:"TRANS_ID,omitempty"`
		ERRORCODE             string `xml:"ERRORCODE,omitempty"`
		AUTH_CODE             string `xml:"AUTH_CODE,omitempty"`
		NUMCODE               string `xml:"NUMCODE,omitempty"`
		ORDERHISTORY          string `xml:"ORDERHISTORY,omitempty"`
		MAILORDER             string `xml:"MAILORDER,omitempty"`
		SUBMERCHANTNAME       string `xml:"SUBMERCHANTNAME,omitempty"`
		SUBMERCHANTID         string `xml:"SUBMERCHANTID,omitempty"`
		SUBMERCHANTPOSTALCODE string `xml:"SUBMERCHANTPOSTALCODE,omitempty"`
		SUBMERCHANTCITY       string `xml:"SUBMERCHANTCITY,omitempty"`
		SUBMERCHANTCOUNTRY    string `xml:"SUBMERCHANTCOUNTRY,omitempty"`
		CARDBRAND             string `xml:"CARDBRAND,omitempty"`
		CARDHOLDERNAME        string `xml:"CARDHOLDERNAME,omitempty"`
		MRK1ISLEMSAYAC        string `xml:"MRK1ISLEMSAYAC,omitempty"`
		MRK2ISLEMSAYAC        string `xml:"MRK2ISLEMSAYAC,omitempty"`
		MRK1CIROSAYAC         string `xml:"MRK1CIROSAYAC,omitempty"`
		MRK2CIROSAYAC         string `xml:"MRK2CIROSAYAC,omitempty"`
		RFM1SAYAC             string `xml:"RFM1SAYAC,omitempty"`
		RFM2SAYAC             string `xml:"RFM2SAYAC,omitempty"`
		RFM3SAYAC             string `xml:"RFM3SAYAC,omitempty"`
		CCBCHIPPARA           string `xml:"CCBCHIPPARA,omitempty"`
		CCBCHIPPARABAKIYE     string `xml:"CCBCHIPPARABAKIYE,omitempty"`
		CCBCHIPPARAACIKLAMA   string `xml:"CCBCHIPPARAACIKLAMA,omitempty"`
		PCBCHIPPARA           string `xml:"PCBCHIPPARA,omitempty"`
		PCBCHIPPARABAKIYE     string `xml:"PCBCHIPPARABAKIYE,omitempty"`
		PCBCHIPPARAACIKLAMA   string `xml:"PCBCHIPPARAACIKLAMA,omitempty"`
		XCBCHIPPARA           string `xml:"XCBCHIPPARA,omitempty"`
		XCBCHIPPARABAKIYE     string `xml:"XCBCHIPPARABAKIYE,omitempty"`
		XCBCHIPPARAACIKLAMA   string `xml:"XCBCHIPPARAACIKLAMA,omitempty"`
		ARTITAKSIT            string `xml:"ARTITAKSIT,omitempty"`
		ERTELEMESABITTARIH    string `xml:"ERTELEMESABITTARIH,omitempty"`
		ERTELEMETAKSIT        string `xml:"ERTELEMETAKSIT,omitempty"`
		KAZANILANCEKILISADEDI string `xml:"KAZANILANCEKILISADEDI,omitempty"`
		TOPLAMCEKILISBAKIYESI string `xml:"TOPLAMCEKILISBAKIYESI,omitempty"`
		MESAJBASILACAK        string `xml:"MESAJBASILACAK,omitempty"`
	} `xml:"Extra,omitempty"`
}

func (api *API) Transaction(request Request) (response Response) {
	postdata, _ := xml.Marshal(request)
	res, err := http.Post(config.APIURL, "text/xml; charset=utf-8", strings.NewReader(strings.ToLower(xml.Header)+string(postdata)))
	if err != nil {
		fmt.Println(err)
		return response
	}
	defer res.Body.Close()
	decoder := xml.NewDecoder(res.Body)
	decoder.CharsetReader = charset.NewReaderLabel
	decoder.Decode(&response)
	return response
}
