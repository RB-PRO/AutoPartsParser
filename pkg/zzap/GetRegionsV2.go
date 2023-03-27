package zzap

import (
	"encoding/json"
	"net/http"
)

// ## МЕТОД "РЕГИОНЫ ПОИСКА" (GETREGIONSV2)
// Аналогично GetRegions, но с дополнительным параметром login и password
//
// [GETREGIONSV2]: https://wiki.zzap.ru/api2-информация/#_Метод_регионы_поиска_GetRegionsVC__1

// Ответ результата выполнения метода GETREGIONSV2
type GetRegionsV2_response struct {
	Error    string          `json:"error"`     // если пусто, ошибок нет
	RowCount int             `json:"row_count"` // сколько строк вернулось
	Table    []GetRegionItem `json:"table"`     // Массив данных
}

// Структура ID региона
type GetRegionItem struct {
	CodeRegion  int    `json:"code_region"`  // Код региона
	ClassRegion string `json:"class_region"` // Название региона
}

func (lap *Lap) GetRegionsV2() error {

	// Собираем запрос
	data := make([]MethodData, 0)
	data = append(data, MethodData{Login, lap.Login})
	data = append(data, MethodData{Password, lap.Password})
	data = append(data, MethodData{ApiKey, lap.ApiKey})

	// Создаём запрос для метода GetRegionsV2
	responseByte, errorReq := MakeRequest(http.MethodGet, "GetRegionsV2", data)
	if errorReq != nil {
		return errorReq
	}

	//fmt.Println(string(responseByte[76 : len(responseByte)-9]))

	// Распасить ответ в структуру
	var GetRegionsV2Resp GetRegionsV2_response
	ErrorUnmarshal := json.Unmarshal(responseByte[76:len(responseByte)-9], &GetRegionsV2Resp)
	if ErrorUnmarshal != nil {
		return ErrorUnmarshal
	}

	lap.Regions = GetRegionsV2Resp.Table
	return nil
}
