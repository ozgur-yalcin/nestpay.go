package config

var (
	Bank   string = ""
	Client string = ""
	User   string = ""
	Pass   string = ""
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
