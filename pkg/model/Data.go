package model

type Data struct {
	DcNumber int `json:"dcNumber"`
	LoadingTime string `json:"loadingTime"`
	RamConsumptions string `json:"ramConsumption"`
	RequestTime string `json:"requestTime"`
	MostUsedDc string `json:"mostUsedDc"`
	AverageDcProperties string `json:"averageDcProperties"`
	DcList []DcInfo            `json:"dcList"`
}

type DcInfo struct {
	DcNumber int `json:"dcNumber"`
	Features string `json:"features"`
	AverageRequestTime string `json:"averageRequestTime"`
	LastRequest string `json:"lastRequest"`
	LastUpdated string `json:"lastUpdated"`
}

