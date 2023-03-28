package zzap_test

import (
	"fmt"
	"testing"

	"github.com/RB-PRO/AutoPartsParser/pkg/zzap"
)

func TestGetSearchResultV3(t *testing.T) {
	// Создаём экземпляр zzap
	lap, lapError := zzap.New("zzap.json")
	if lapError != nil {
		t.Error(lapError)
	}

	// https://www.zzap.ru/public/search.aspx#rawdata=5QF413032B&codes_man=3261.1900952959&delivery_days=0;1;0.5
	request := zzap.GetSearchResultV3_request{
		Code_region:  "1",
		Search_text:  "",
		Partnumber:   "JNB000060", // JNB000060
		Class_man:    "LAND ROVER",
		Row_count:    "500",
		Type_request: "5",
	}

	GetSearchResultV3Resp, ErrorGetSearchResultV3 := lap.GetSearchResultV3(request)
	if ErrorGetSearchResultV3 != nil {
		t.Error(ErrorGetSearchResultV3)
	}

	for _, val := range GetSearchResultV3Resp.Table {
		fmt.Println(val.DeliveryDays, val.DescrDelivery)
	}

	fmt.Println("Товаров в наличии:", GetSearchResultV3Resp.PriceCountInstock)
	fmt.Println("Минимум цены в наличии:", GetSearchResultV3Resp.PriceMinInstock)
	fmt.Println("Средняя цена среди товаров в наличии", GetSearchResultV3Resp.PriceAvgInstock) // Средняя цена среди товаров в наличии
	fmt.Println("Максимум цены в наличии:", GetSearchResultV3Resp.PriceMaxInstock)

	fmt.Println("GetSearchResultV3: Всего найдено", len(GetSearchResultV3Resp.Table), "записей.")
}
