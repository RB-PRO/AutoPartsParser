package AutoPartsParser

import (
	"github.com/RB-PRO/AutoPartsParser/pkg/zzap"
)

func ZZapParse(ReqXlsx []Request) error {
	ZZapUser, errorZZapNew := zzap.New("lap.json")
	if errorZZapNew != nil {
		return errorZZapNew
	}

	ErrorRegions := ZZapUser.GetRegionsV2()
	if ErrorRegions != nil {
		return ErrorRegions
	}

	return nil
}
