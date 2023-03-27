package AutoPartsParser

import (
	"github.com/RB-PRO/AutoPartsParser/pkg/zzap"
)

func ZZapParse(ReqXlsx []Request) error {
	// Получить конфигурационынй файл
	ZZapUser, errorZZapNew := zzap.New("zzap.json")
	if errorZZapNew != nil {
		return errorZZapNew
	}

	// Получить регионы
	ErrorRegions := ZZapUser.GetRegionsV2()
	if ErrorRegions != nil {
		return ErrorRegions
	}

	return nil
}
