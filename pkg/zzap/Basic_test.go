package zzap_test

import (
	"fmt"
	"testing"

	"github.com/RB-PRO/AutoPartsParser/pkg/zzap"
)

func TestNew(t *testing.T) {
	lap, lapError := zzap.New("zzap.json")
	if lapError != nil {
		t.Error(lapError)
	}
	fmt.Printf("New: Получил данные: %+v", lap)
}
