package AutoPartsParser

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/RB-PRO/avtoto"
	"github.com/cheggaaa/pb"
)

// Структура вывода, которая отправляется прямиком в файл
type Avtoto_Output struct {
	SKU         string // Артикул
	Manufacture string // Производитель

	ProcessSearchId string // идентификатор процесса поиска (тип: строка). Необходим для отслеживания результатов процесса поиска.
	IsOk            bool   // Если true - то информация уже собрана

	Parts avtoto.SearchGetParts2Response // Структура, в которй содержится вся полезная инфомрация

}

func AvtotoParse(ReqXlsx []Request) ([]Avtoto_Output, error) {

	// Получить данные из json-файла
	AvtotoDataAuf, ErrorDataFile := DataFile("avtoto.json")
	if ErrorDataFile != nil {
		return []Avtoto_Output{}, ErrorDataFile
	}

	// Создать экземпляр пользователя api avtoto
	user := avtoto.User{
		UserId:       AvtotoDataAuf.UserId,
		UserLogin:    AvtotoDataAuf.UserLogin,
		UserPassword: AvtotoDataAuf.UserPassword,
	}

	// Создаём массив, который будет подан на вывод.
	var output []Avtoto_Output

	// Цикл по всему входному файлу
	bar := pb.StartNew(len(ReqXlsx)).Prefix("[200 мс]: Запускаем поиск") // Включить показывание процесса
	for _, ValueInput := range ReqXlsx {

		// Поиск кода бренда
		//fmt.Println("ValueInput.Manufacture", ValueInput.Name)

		// Делаем запрос на получение кода
		GetBrandsByCodeReq := avtoto.GetBrandsByCodeRequest{
			SearchCode: ValueInput.Name,
		}
		GetBrandsByCodeResp, ErrGetBrandsByCodeResp := user.GetBrandsByCode(GetBrandsByCodeReq)
		if ErrGetBrandsByCodeResp != nil {
			return []Avtoto_Output{}, ErrGetBrandsByCodeResp
		}
		var BrandID string
		for _, val := range GetBrandsByCodeResp.Brands {
			//fmt.Println(val.Name, ValueInput.Manufacture)
			if val.Name == strings.ToUpper(ValueInput.Manufacture) {
				BrandID = val.Manuf
			}
		}

		// Создать запрос старта поиска
		SearchStartReq := avtoto.SearchStartRequest{
			SearchCode:  ValueInput.Name,
			Brand:       BrandID,
			SearchCross: "off",
		}

		// Выполнение самого запроса "до талого"
		var SearchStartResp avtoto.SearchStartResponse
		for {
			time.Sleep(200 * time.Millisecond)
			var ErrorStart error
			SearchStartResp, ErrorStart = user.SearchStartRequest(SearchStartReq)
			//fmt.Println(SearchStartResp.Error())
			if SearchStartResp.Error() == "" {
				break
			}
			if ErrorStart != nil {
				return []Avtoto_Output{}, ErrorStart // log.Println(ErrorStart)
			}
		}

		// Формирование выходного массива
		output = append(output, Avtoto_Output{
			SKU:             ValueInput.Name,
			Manufacture:     ValueInput.Manufacture,
			ProcessSearchId: SearchStartResp.ProcessSearchID,
		})
		bar.Increment() // Прибавляем 1 к отображению
	}
	bar.Finish() // Завершить

	fmt.Println("[RB_PRO]: Ждём 10 секунд")
	time.Sleep(10 * time.Second)
	bar2 := pb.StartNew(len(ReqXlsx)).Prefix("[300 мс]: Опрашиваем результаты") // Включить показывание процесса

	AllSize := len(output)
	TecalSize := 0
	for TecalSize != AllSize { // Цикл. Пока не найдено решение по каждому запросу
		for Index := range output { // Идём по всем запрашиваемым ID
			if !output[Index].IsOk { // Если не готов ответ на этот запрос

				// Формируем запрос
				SearchGetParts2Req := avtoto.SearchGetParts2Request{
					ProcessSearchId: output[Index].ProcessSearchId,
					Limit:           500,
				}

				// Выполняем запрос
				SearchResp, ErrorSearch := SearchGetParts2Req.SearchGetParts2()
				if ErrorSearch != nil {
					continue // Повторить поиск по заданному ID процесса
				}

				// Проверка на неготовность обработки поска
				if SearchResp.Status() == "Запрос в обработке" {
					continue // Повторить поиск по заданному ID процесса
				}

				// Профильтровать полученные результаты
				output[Index].Parts = Avtoto_Filter(SearchResp, output[Index].Manufacture)

				// Обработка готового события
				TecalSize++
				output[Index].IsOk = true
				bar2.Increment() // Прибавляем 1 к отображению
			}
		}
	}
	bar2.Finish() // Завершить

	return output, nil
}

// Фильтрация по бизнес-логике
func Avtoto_Filter(SearchResp avtoto.SearchGetParts2Response, manuf string) avtoto.SearchGetParts2Response {
	var NewParts avtoto.SearchGetParts2Response

	// Срок доставки меньше 7 и колличество больше 1
	for _, value := range SearchResp.Parts {
		delivery, _ := strconv.Atoi(value.Delivery)
		MaxCount, _ := strconv.Atoi(value.MaxCount)
		if delivery < 7 && (MaxCount > 1 || MaxCount == -1) {
			NewParts.Parts = append(NewParts.Parts, value)
		}
	}

	// if len(NewParts.Parts) > 1 {
	sort.Slice(NewParts.Parts, func(i, j int) (less bool) {
		return NewParts.Parts[i].Price < NewParts.Parts[j].Price
	})
	// }

	// Цикл по всем параметрам
	var NewPartsParts avtoto.SearchGetParts2Response
	for _, value := range NewParts.Parts {
		if strings.TrimSpace(strings.ToLower(value.Manuf)) == strings.TrimSpace(strings.ToLower(manuf)) {
			NewPartsParts.Parts = append(NewPartsParts.Parts, value)
			// OutputSearchGetParts2Response.Parts = append(OutputSearchGetParts2Response.Parts, value)
			// SearchResp.Parts = append(SearchResp.Parts[:i], SearchResp.Parts[i+1])
		}
	}

	return NewPartsParts
}
