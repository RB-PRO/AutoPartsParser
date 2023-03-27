package zzap

import (
	"encoding/json"
	"net/http"
)

// ## МЕТОД "РЕЗУЛЬТАТ ПОИСКА" (GETSEARCHRESULTV3)
// Аналогично GetSearchResult, но с дополнительными параметрами search_text и type_request [GetSearchResultV3]
// https://api.zzap.pro/webservice/datasharing.asmx/GetSearchResultV3
// Аналогично GetRegions, но с дополнительным параметром login и password
//
// [GetSearchResultV3]: https://api.zzap.pro/webservice/datasharing.asmx?op=GetSearchResultV3

// Запрос метода GetSearchResultV3
type GetSearchResultV3_request struct {
	Code_region  string // code_region из метода GetRegions
	Search_text  string // произвольная строка поиска
	Partnumber   string // номер запчасти
	Class_man    string // производитель запчасти
	Row_count    string // ограничение по кол-ву строк. по умолчанию 100, максимум 500
	Type_request string // тип поискового запроса: 0 - поиск любых запчастей по номеру, 1 - поиск только новых запчастей по номеру, 2 - поиск по б/у и уценке (по введённым в поисковую строку словам), 4 - любые предложения только по запрошенному номеру, 5 - новые только по запрошенному номеру
}

// Ответ результата выполнения метода GetSearchResultV3
type GetSearchResultV3_response struct {
	Error             string   `json:"error"`     // если пусто, ошибок нет	содержит текст ошибки, если таковая возникла при выполнении запроса.
	RowCount          int      `json:"row_count"` // сколько строк вернулось
	Terms             string   `json:"terms"`     // "колодок;колодку;колодкою; колодкой; колодке;колодках;колодками; колодкам; колодка;колодки;колодки" поле terms нужно для того, чтобы выделять жёлтым цветом, что найдено, т.к. сервер может искать с учетом склонений.
	ClassMan          string   `json:"class_man"`
	Logopath          string   `json:"logopath"`
	Partnumber        string   `json:"partnumber"`
	ClassCat          string   `json:"class_cat"`
	Imagepath         string   `json:"imagepath"`
	CodeCat           int      `json:"code_cat"`
	ClassCur          string   `json:"class_cur"`
	PriceCountInstock int      `json:"price_count_instock"`
	PriceMinInstock   float64  `json:"price_min_instock"`
	PriceAvgInstock   float64  `json:"price_avg_instock"`
	PriceMaxInstock   float64  `json:"price_max_instock"`
	PriceCountOrder   int      `json:"price_count_order"`
	PriceMinOrder     float64  `json:"price_min_order"`
	PriceAvgOrder     float64  `json:"price_avg_order"`
	PriceMaxOrder     float64  `json:"price_max_order"`
	ImagepathV2       []string `json:"imagepathV2"`
	CodeMan           int      `json:"code_man"`
	Table             []struct {
		CodeDocB         int64    `json:"code_doc_b"`
		CodeCat          int      `json:"code_cat"`
		DescrTypeSearch  string   `json:"descr_type_search"`
		TypeSearch       int      `json:"type_search"`
		ClassMan         string   `json:"class_man"`
		Logopath         string   `json:"logopath"`
		Partnumber       string   `json:"partnumber"`
		ClassCat         string   `json:"class_cat"`
		Imagepath        string   `json:"imagepath"`
		Qty              string   `json:"qty"`
		Instock          int      `json:"instock"`
		Wholesale        int      `json:"wholesale"`
		Local            int      `json:"local"`
		Price            string   `json:"price"`
		PriceDate        string   `json:"price_date"`
		DescrPrice       string   `json:"descr_price"`
		DescrQty         string   `json:"descr_qty"`
		ClassUser        string   `json:"class_user"`
		DescrRatingCount string   `json:"descr_rating_count"`
		Rating           int      `json:"rating"`
		DescrAddress     string   `json:"descr_address"`
		Phone1           string   `json:"phone1"`
		OrderText        string   `json:"order_text"`
		UserKey          string   `json:"user_key"`
		AddrMapGeo1      float64  `json:"addr_map_geo1"`
		AddrMapGeo2      float64  `json:"addr_map_geo2"`
		Used             int      `json:"used"`
		Apply            string   `json:"apply"`
		MinSumOrder      float64  `json:"min_sum_order"`
		DescrMinSumOrder string   `json:"descr_min_sum_order"`
		Shipment         string   `json:"shipment"`
		Courier          bool     `json:"courier"`
		InstockV2        bool     `json:"instockV2"`
		WholesaleV2      bool     `json:"wholesaleV2"`
		LocalV2          bool     `json:"localV2"`
		UsedV2           bool     `json:"usedV2"`
		PriceV2          float64  `json:"priceV2"`
		DescrPriceV2     string   `json:"descr_priceV2"`
		PriceOrig        float64  `json:"price_orig"`
		DescrPriceOrig   string   `json:"descr_price_orig"`
		DescrTypePrice   string   `json:"descr_type_price"`
		QtyV2            int      `json:"qtyV2"`
		QtyMax           int      `json:"qty_max"`
		DescrQtyV2       string   `json:"descr_qtyV2"`
		DeliveryDays     int      `json:"delivery_days"`
		DescrDelivery    string   `json:"descr_delivery"`
		TypeUser         string   `json:"type_user"`
		TypeUser2        string   `json:"type_user2"`
		TypePrice        string   `json:"type_price"`
		ImagepathV2      []string `json:"imagepathV2"`
		DescrPriceDate   string   `json:"descr_price_date"`
		Pack             int      `json:"pack"`
		DescrPack        string   `json:"descr_pack"`
		TypeChainSearch  int      `json:"type_chain_search"`
		Noorig           bool     `json:"noorig"`
		CodeMan          int      `json:"code_man"`
		Location         string   `json:"location"`
	} `json:"table"`
}

func (lap *Lap) GetSearchResultV3(request GetSearchResultV3_request) error {

	// Собираем запрос
	data := make([]MethodData, 0)
	data = append(data, MethodData{Login, lap.Login})
	data = append(data, MethodData{Password, lap.Password})
	data = append(data, MethodData{ApiKey, lap.ApiKey})

	// Собираем дополнительные параметры для запроса
	data = append(data, []MethodData{
		{"code_region", request.Code_region},
		{"search_text", request.Search_text},
		{"partnumber", request.Partnumber},
		{"class_man", request.Class_man},
		{"row_count", request.Row_count},
		{"type_request", request.Type_request},
	}...)

	// Создаём запрос для метода GetRegionsV2
	responseByte, errorReq := MakeRequest(http.MethodGet, "GetSearchResultV3", data)
	if errorReq != nil {
		return errorReq
	}

	// fmt.Println(string(responseByte[76 : len(responseByte)-9]))

	// Распасить ответ в структуру
	var GetRegionsV2Resp GetRegionsV2_response
	ErrorUnmarshal := json.Unmarshal(responseByte[76:len(responseByte)-9], &GetRegionsV2Resp)
	if ErrorUnmarshal != nil {
		return ErrorUnmarshal
	}

	lap.Regions = GetRegionsV2Resp.Table
	return nil
}
