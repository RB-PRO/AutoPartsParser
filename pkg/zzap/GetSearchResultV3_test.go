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
		Partnumber:   "5QF413032B",
		Class_man:    "vag",
		Row_count:    "100",
		Type_request: "0",
	}

	errorGetRegionsV2 := lap.GetSearchResultV3(request)
	if errorGetRegionsV2 != nil {
		t.Error(errorGetRegionsV2)
	}

	fmt.Println("GetRegionsV2: Всего найдено", len(lap.Regions), "записей.")
}
