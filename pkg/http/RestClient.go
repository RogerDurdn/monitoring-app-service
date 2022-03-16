package http

import (
	"github.com/RogerDurdn/MonitoringApp/pkg/model"
	"math/rand"
	"strconv"
	"time"
)

type DataSource interface {
	FetchData() model.Data
}

type RestSource struct {
}

type MockSource struct {
}

func (s *MockSource) FetchData() model.Data {
	return model.Data{
		DcNumber: rN(10, 1000),
		LoadingTime: rS() +" -s",
		RamConsumptions: rS() +"%",
		RequestTime: rS() +" -s",
		MostUsedDc: "#" + rS(),
		AverageDcProperties:"("+rS() +")",
		DcList: generateList(10),
	}
}

func generateList(size int) []model.DcInfo {
	dcList := make([]model.DcInfo,0)
	for i := 1; i <= size; i++ {
		dcList = append(dcList, model.DcInfo{
			DcNumber: i,
			Features: "#"+rS(),
			AverageRequestTime: rSs() + "%",
			LastRequest: rSs() + "-s",
			LastUpdated: rSs()+ "-hr",
		})
	}
	return dcList
}

func rSs()string  {
	rn := rN(2, 510)
	return strconv.Itoa(rn)
}

func rS()string  {
	rn := rN(5, 100)
	return strconv.Itoa(rn)
}

func rN(min, max int) int  {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max - min) + max
}