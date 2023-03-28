package AutoPartsParser

import (
	"strconv"
	"time"

	"github.com/RB-PRO/AutoPartsParser/pkg/zzap"
	"github.com/cheggaaa/pb"
)

type ZZap_Output struct {
	Name              string  // Артикул
	RealName          string  // Название товара
	Manuf             string  // Производитель
	PriceCountInstock int     `json:"price_count_instock"` // количество предложений в наличии запрашиваемой запчасти
	PriceMinInstock   float64 `json:"price_min_instock"`   // минимальная цена среди предложений в наличии запрашиваемой запчасти
	PriceAvgInstock   float64 `json:"price_avg_instock"`   // средняя цена среди предложений в наличии запрашиваемой запчасти
	PriceMaxInstock   float64 `json:"price_max_instock"`   // максимальная цена среди предложений в наличии запрашиваемой запчасти
}

func ZZapParse(ReqXlsx []Request) ([]ZZap_Output, error) {
	// Получить конфигурационынй файл
	ZZapUser, errorZZapNew := zzap.New("zzap.json")
	if errorZZapNew != nil {
		return []ZZap_Output{}, errorZZapNew
	}

	// Получить регионы
	ErrorRegions := ZZapUser.GetRegionsV2()
	if ErrorRegions != nil {
		return []ZZap_Output{}, ErrorRegions
	}

	// Ищем Московский регион
	RegionMSK := ZZapUser.FindRegionMSK()

	// Создаём массив, который будет подан на вывод.
	var output []ZZap_Output

	// Цикл по всему входному файлу
	bar := pb.StartNew(len(ReqXlsx)).Prefix("Пауза между запросами - 4 сек.") // Включить показывание процесса
	defer bar.Finish()                                                        // Завершить
	for _, ValueInput := range ReqXlsx {
		time.Sleep(4 * time.Second)

		request := zzap.GetSearchResultV3_request{
			Code_region:  strconv.Itoa(RegionMSK),
			Search_text:  "",
			Partnumber:   ValueInput.Name,
			Class_man:    ValueInput.Manufacture,
			Row_count:    "500",
			Type_request: "5",
		}
		GetSearchResultV3Resp, ErrorGetSearchResultV3 := ZZapUser.GetSearchResultV3(request)
		if ErrorGetSearchResultV3 != nil {
			return []ZZap_Output{}, ErrorGetSearchResultV3
		}
		output = append(output, ZZap_Output{
			Name:              ValueInput.Name,
			Manuf:             ValueInput.Manufacture,
			RealName:          GetSearchResultV3Resp.ClassCat,
			PriceCountInstock: GetSearchResultV3Resp.PriceCountInstock,
			PriceMinInstock:   GetSearchResultV3Resp.PriceMinInstock,
			PriceAvgInstock:   GetSearchResultV3Resp.PriceAvgInstock,
			PriceMaxInstock:   GetSearchResultV3Resp.PriceMaxInstock,
		})
		bar.Increment() // Прибавляем 1 к отображению
	}

	return output, nil
}
