package config

const (
	// parser names
	ParseCity     = "ParseCity"
	ParseCityList = "ParseCityList"
	ParseProfile  = "ParseProfile"
	NilParser     = "NilParser"

	// Service ports
	//ItemSaverPost = 1234
	//WorkerPort0 = 9000
	// ElasticSearch
	ElasticIndex = "dating_profile"
	// RPC Endpoints
	ItemSaverRpc = "ItemSaverService.Save"
	CrawlServiceRpc = "CrawlService.Process"

	// Rate Limiting
	Qps = 20

)
