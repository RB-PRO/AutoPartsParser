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
	Error             string   `json:"error"`               // если пусто, ошибок нет	содержит текст ошибки, если таковая возникла при выполнении запроса.
	RowCount          int      `json:"row_count"`           // сколько строк вернулось
	Terms             string   `json:"terms"`               // "колодок;колодку;колодкою; колодкой; колодке;колодках;колодками; колодкам; колодка;колодки;колодки" поле terms нужно для того, чтобы выделять жёлтым цветом, что найдено, т.к. сервер может искать с учетом склонений.
	ClassMan          string   `json:"class_man"`           // производитель запрашиваемой запчасти
	Logopath          string   `json:"logopath"`            // ссылка на превью логотипа производителя 30px на 30px запрашиваемой запчасти
	Partnumber        string   `json:"partnumber"`          // номер запрашиваемой запчасти
	ClassCat          string   `json:"class_cat"`           // наименование запрашиваемой запчасти
	Imagepath         string   `json:"imagepath"`           // ссылка на превью фото запчасти 60px на 60px**** запрашиваемой запчасти
	CodeCat           int      `json:"code_cat"`            // внутренний уникальный код позиции запрашиваемой запчасти
	ClassCur          string   `json:"class_cur"`           // валюта (например, белорусские рубли) запрашиваемой запчасти
	PriceCountInstock int      `json:"price_count_instock"` // количество предложений в наличии запрашиваемой запчасти
	PriceMinInstock   float64  `json:"price_min_instock"`   // минимальная цена среди предложений в наличии запрашиваемой запчасти
	PriceAvgInstock   float64  `json:"price_avg_instock"`   // средняя цена среди предложений в наличии запрашиваемой запчасти
	PriceMaxInstock   float64  `json:"price_max_instock"`   // максимальная цена среди предложений в наличии запрашиваемой запчасти
	PriceCountOrder   int      `json:"price_count_order"`   // количество предложений под заказ запрашиваемой запчасти
	PriceMinOrder     float64  `json:"price_min_order"`     // минимальная цена среди предложений под заказ запрашиваемой запчасти
	PriceAvgOrder     float64  `json:"price_avg_order"`     // средняя цена среди предложений под заказ запрашиваемой запчасти
	PriceMaxOrder     float64  `json:"price_max_order"`     // максимальная цена среди предложений под заказ запрашиваемой запчасти
	ImagepathV2       []string `json:"imagepathV2"`         // массив ссылок на превью фото запчасти 60px на 60px**** запрашиваемой запчасти
	CodeMan           int      `json:"code_man"`            // внутренний код производителя позиции (соответствия кодов и названий всех производителей - в методе GetBrands) запрашиваемой запчасти
	Table             []struct {
		CodeDocB         int64    `json:"code_doc_b"`          // внутренний уникальный код предложения (нужно передавать в GetSearchResultOne)
		CodeCat          int      `json:"code_cat"`            // внутренний уникальный код позиции
		DescrTypeSearch  string   `json:"descr_type_search"`   // тип предложения
		TypeSearch       int      `json:"type_search"`         // 10 - Запрошенный номер (cпец. предложения), 13 - Запрошенный номер, 21- Замены (cпец. предложения), 31 - Замены, 50 - Запрошенный номер (недостоверные предложения), 34 - Деталь, как составляющие, 54 - Детали, как составляющие (недостоверные предложения), 14 - Запрошенный номер б/у и уценка, 15 - Результат поиска по б/у и уценка,
		ClassMan         string   `json:"class_man"`           // производитель
		Logopath         string   `json:"logopath"`            // ссылка на превью логотипа производителя 30px на 30px
		Partnumber       string   `json:"partnumber"`          // номер
		ClassCat         string   `json:"class_cat"`           // наименование
		Imagepath        string   `json:"imagepath"`           // ссылка на превью фото запчасти 60px на 60px****
		Qty              string   `json:"qty"`                 // кол-во, текст!
		Instock          int      `json:"instock"`             // V2 наличие на складе, если true, то надо зеленым подсвечивать квадратик с кол-вом
		Wholesale        int      `json:"wholesale"`           // V2 тип цены, опт или розница, если true, то надо желтым подсвечивать квадратик с ценой
		Local            int      `json:"local"`               // V2 если true, то предложение локальное, если false, то из другого региона
		Price            string   `json:"price"`               // V2 цена (в той валюте, в которой смотрел покупатель)
		PriceDate        string   `json:"price_date"`          // дата публикации
		DescrPrice       string   `json:"descr_price"`         // V2 цена текстом с коротким обозначением валюты, в которой смотрел покупатель
		DescrQty         string   `json:"descr_qty"`           // подпись под кол-вом
		ClassUser        string   `json:"class_user"`          // наименование продавца
		DescrRatingCount string   `json:"descr_rating_count"`  // сколько отзывов текстом
		Rating           int      `json:"rating"`              // рейтинг: если 0, ничего не показывать, если от 1 до 5, то показывать 5 звезд, заливая соотв. кол-во звезд
		DescrAddress     string   `json:"descr_address"`       // местоположение
		Phone1           string   `json:"phone1"`              // телефон
		OrderText        string   `json:"order_text"`          // три варианта: "Заказать" (можно оформить заказ у продавца через сайт ZZap), "Купить" (товар есть в наличии у данного продавца, можно совершить покупку в этот же день, оформить покупку можно на сайте ZZap) или "" (нет возможности заказать через сайт ZZap, необходимо связаться с продавцом)
		UserKey          string   `json:"user_key"`            // "ключ" продавца, с помощью которого вы сможете оставить отзыв по заказу, оформленному на сайте ZZap (метод MakeOrderRating), а также можете посмотреть информацию о продавце (метод GetUserInfo)
		AddrMapGeo1      float64  `json:"addr_map_geo1"`       // широта (координаты местонахождения продавца)
		AddrMapGeo2      float64  `json:"addr_map_geo2"`       // долгота (координаты местонахождения продавца)
		Used             int      `json:"used"`                // V2 тип предложения: false – обычное предложение, true – б/у и уценка
		Apply            string   `json:"apply"`               // условия продажи, если указаны продавцом
		MinSumOrder      float64  `json:"min_sum_order"`       // минимальная сумма заказа, если указана продавцом
		DescrMinSumOrder string   `json:"descr_min_sum_order"` // минимальная сумма заказа текстом, если указана продавцом
		Shipment         string   `json:"shipment"`            // условия доставки
		Courier          bool     `json:"courier"`             // доставка курьером: если true - есть, если false - нет
		InstockV2        bool     `json:"instockV2"`           // наличие на складе, если true, то надо зеленым подсвечивать квадратик с кол-вом
		WholesaleV2      bool     `json:"wholesaleV2"`         // тип цены, опт или розница, если true, то надо желтым подсвечивать квадратик с ценой
		LocalV2          bool     `json:"localV2"`             // если true, то предложение локальное, если false, то из другого региона
		UsedV2           bool     `json:"usedV2"`              // тип предложения: false – обычное предложение, true – б/у и уценка
		PriceV2          float64  `json:"priceV2"`             // цена (в той валюте, в которой смотрел покупатель)
		DescrPriceV2     string   `json:"descr_priceV2"`       // цена текстом с коротким обозначением валюты, в которой смотрел покупатель
		PriceOrig        float64  `json:"price_orig"`          // цена (в той валюте, в которой публиковал продавец)
		DescrPriceOrig   string   `json:"descr_price_orig"`    // цена текстом (в той валюте, в которой публиковал продавец)
		DescrTypePrice   string   `json:"descr_type_price"`    // отдельно тип цены – «Только для юр. лиц и ИП» или «» (пусто)
		QtyV2            int      `json:"qtyV2"`               // оличество, указанное продавцом в прайсе (числом). кроме обычных значений 10, 20, 4, могут быть спец. значения: -1 («На заказ»), -2 («В наличии»), 100012 («>12 шт.»)
		QtyMax           int      `json:"qty_max"`             // максимально допустимое количество для заказа по конкретному предложению продавца
		DescrQtyV2       string   `json:"descr_qtyV2"`         // количество, указанное продавцом в прайсе (текстом), соответственно количеству qtyV2 могут быть значения: «5 шт.», «15 шт.», «На заказ», «В наличии», «>20 шт.»
		DeliveryDays     int      `json:"delivery_days"`       // количество дней поставки (вычисленное нами в зависимости от указанного в прайсе)
		DescrDelivery    string   `json:"descr_delivery"`      // срок поставки текстом (примеры: «7-15 дней», «14-20 дней (плюс время на доставку из г. Санкт-Петербург)»)
		TypeUser         string   `json:"type_user"`           // тип пользователя*
		TypeUser2        string   `json:"type_user2"`          // тип пользователя 2**
		TypePrice        string   `json:"type_price"`          // тип прайс-листа***
		ImagepathV2      []string `json:"imagepathV2"`         // массив ссылок на превью фото запчасти 60px на 60px****
		DescrPriceDate   string   `json:"descr_price_date"`    // давности обновления прайс-листа
		Pack             int      `json:"pack"`                // кратность (упаковка) числом
		DescrPack        string   `json:"descr_pack"`          // кратность (упаковка) текстом
		TypeChainSearch  int      `json:"type_chain_search"`   // основные значения: 0 - запрашиваемые номера, 1 - замены, 3 - выбор производителя, 10 - номер не найден, 11 - нет предложений
		Noorig           bool     `json:"noorig"`              // если true, значит, предложение помечено продавцом в прайс-листе или нашей системой при проверке во время публикации как неоригинальное
		CodeMan          int      `json:"code_man"`            // внутренний код производителя позиции (соответствия кодов и названий всех производителей - в методе GetBrands)
		Location         string   `json:"location"`            // город и метро продавца
	} `json:"table"`
	// 	*type_user
	// 'K' - медаль "Официальный автодилер"
	// '1' - медаль "Больше года на сайте"
	// '2' - медаль "Больше 3х лет на сайте"
	// '3' - медаль "Есть уставные документы"
	// '4' - медаль "Есть фотографии магазина"
	// '5' - международная отправка
	// 'P' - Доставка товара курьером
	// 'Y' - Отправка товара в регионы
	// 'B' - Покупатель
	// 'A' - Продавец
	// 'I' - Интернет магазин
	// 'S' - Автосервис/служба установки
	// 'R' - Торговля в розницу
	// 'W' - Торговля оптом
	// 'F' - Частное лицо
	// 'U' - Юридическое лицо

	// **type_user2
	// 'Y' - Понизить до недостоверных
	// 'K' - Принимаем к оплате банковские карты
	// 'M' - Ручная проверка (снимает пометку Регистрация не подтверждена)
	// 'I' - Неликвидный товар (клиент продаёт неликвид)
	// 'W' - Шиномонтаж
	// 'R' - Не хотят переписываться с покупателями
	// 'Z' - Подтверждена гарантия наличия
	// 'F' - Подтверждение наличия
	// 'R' - Не переписываются с покупателями

	// ***type_price
	// 'R' - Прайс-лист для розницы
	// 'W' - Прайс-лист для юр.лиц и ИП
	// 'M' - Запретить загрузку прайса с сайта как файл
	// 'I' - Неликвидный товар
	// 'P' - Есть самовывоз
}

func (lap *Lap) GetSearchResultV3(request GetSearchResultV3_request) (GetSearchResultV3_response, error) {

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
		return GetSearchResultV3_response{}, errorReq
	}

	// fmt.Println(string(responseByte[76 : len(responseByte)-9]))

	// Распасить ответ в структуру
	var GetSearchResultV3Resp GetSearchResultV3_response
	ErrorUnmarshal := json.Unmarshal(responseByte[76:len(responseByte)-9], &GetSearchResultV3Resp)
	if ErrorUnmarshal != nil {
		return GetSearchResultV3_response{}, ErrorUnmarshal
	}

	return GetSearchResultV3Resp, nil
}
