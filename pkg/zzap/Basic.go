// # Документация API
//
// ###[БАЗОВАЯ ССЫЛКА ВЕБСЕРВИСА]
//
// ПОЛУЧИТЬ API_KEY
// Для того, чтобы воспользоваться ниже описанными методами выгрузки данных, необходимо получить api_key
// - Если вы покупатель, напишите нам на почту support@zzap.ru или создайте заявку в разделе "Заявки на тех.поддержку", и вам пришлют api_key
// - Если вы продавец и знаете, кто ваш менеджер, обратитесь к нему - менеджер вышлет вам api_key
// - Если вы продавец, но по каким-то причинам не знаете, кто ваш менеджер, или на данный момент к вам не прикреплён ни один из наших менеджеров, поступайте, как покупатель
//
// ### Правила API
// - API создан для интеграции с продавцами запчастей
// - недопустимо использовать API для выкачивания актуальных предложений или мониторинга через выкачивание
// - мы имеет право отозвать API ключ в любой момент без объяснения причин
// - частота запросов по API должна быть не выше 1 запроса в 3 секунды (1 запрос в 6-50 секунд, для методов которые начинаются с GetSearchSuggest* и GetSearchResult*, так же возможно появление каптчи в случае длительного превышения частоты запросов)
// - при использовании данных нашего сайта необходимо указывать в верхней части страницы (или у кнопки "Поиск") заметное упоминание о том, что данные предоставлены сайтом ZZap, а также размещать ссылку с переходом на ZZap (пример надписи: "Информация о запчастях предоставлена системой [ZZap]")
// - запросы через API обслуживаются группой серверов с балансированием нагрузки, и данные на них синхронизируются через репликацию. Обычно требуется несколько секунд, чтобы изменение проявилось всюду
//
// ### Важно:
// все даты выдаются по московскому времени!
//
// [БАЗОВАЯ ССЫЛКА ВЕБСЕРВИСА]: https://api.zzap.pro/webservice/datasharing.asmx
// [ZZap]: https://www.zzap.ru/
package zzap

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
)

const (
	URL      string = "https://api.zzap.pro/webservice/datasharing.asmx"
	Login    string = "login"
	Password string = "password"
	ApiKey   string = "api_key"
)

type Lap struct {
	Login    string `json:"login"`    // e-mail, указанный при регистрации (может быть пустым)
	Password string `json:"password"` // ваш пароль от аккаунта на сайте ZZap (может быть пустым)
	ApiKey   string `json:"api_key"`  // нужно попросить у нас

	Regions []GetRegionItem `json:"-"` // Список регионов
}

// Создать экземпляр логина/пароля/api для работы с сервисом zzap
func New(filename string) (Lap, error) {
	if lap, LapError := DataFile(filename); LapError != nil {
		return Lap{}, LapError
	} else {
		return lap, nil
	}
}

// Структура запросов данных
type MethodData struct {
	Key   string // Ключ запроса
	Param string // Параметр запроса
}

// Ядро запроса универсальное.
// Передаём метод и Data в формате []byte, нп выходе получаем ответ
// https://mailazy.com/blog/http-request-golang-with-best-practices/
func MakeRequest(MethodURL, method string, data []MethodData) ([]byte, error) {
	Request, ErrNewRequest := http.NewRequest(MethodURL, URL+"/"+method, nil)
	if ErrNewRequest != nil {
		return nil, ErrNewRequest
	}

	// appending to existing query args
	q := Request.URL.Query()
	for _, val := range data {
		q.Add(val.Key, val.Param)
	}

	// assign encoded query string to http request
	Request.URL.RawQuery = q.Encode()

	client := &http.Client{}
	resp, ErrClient := client.Do(Request)
	if ErrClient != nil {
		return nil, ErrClient
	}

	defer resp.Body.Close()

	bytesArray, ErrorReadAll := io.ReadAll(resp.Body)
	if ErrorReadAll != nil {
		return nil, ErrorReadAll
	}

	return bytesArray, nil
}

// Получение значение из файла
func DataFile(filename string) (Lap, error) {
	// Прочитать файл
	fileBytes, ErrorReadFile := os.ReadFile(filename)
	if ErrorReadFile != nil {
		return Lap{}, ErrorReadFile
	}
	// Распарсить
	var ReturnObject Lap
	ErrorUnmarshal := json.Unmarshal(fileBytes, &ReturnObject)
	if ErrorUnmarshal != nil {
		return Lap{}, ErrorUnmarshal
	}
	return ReturnObject, nil
}
