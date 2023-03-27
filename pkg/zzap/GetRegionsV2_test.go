package zzap_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/RB-PRO/AutoPartsParser/pkg/zzap"
)

func TestGetRegionsV2(t *testing.T) {
	// Создаём экземпляр zzap
	lap, lapError := zzap.New("zzap.json")
	if lapError != nil {
		t.Error(lapError)
	}

	errorGetRegionsV2 := lap.GetRegionsV2()
	if errorGetRegionsV2 != nil {
		t.Error(errorGetRegionsV2)
	}

	if len(lap.Regions) == 0 {
		t.Error(errors.New("Найдено НОЛЬ записей регионов"))
	}

	fmt.Println("GetRegionsV2: Всего найдено", len(lap.Regions), "записей.")
}
