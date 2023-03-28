package AutoPartsParser

import (
	"fmt"
	"log"
	"sort"
	"strconv"
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

	/*
		Product [3]struct {
			Price           int    // Цена
			Storage         string // [*] Склад
			Delivery        string // [*] Срок доставки
			MaxCount        string // [*] Максимальное количество для заказа, остаток по складу. Значение "-1" - означает "много" или "неизвестно"
			DeliveryPercent int    // [**] Процент успешных закупок из общего числа заказов
		}
	*/
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
		fmt.Println("ValueInput.Manufacture", ValueInput.Name)
		GetBrandsByCodeReq := avtoto.GetBrandsByCodeRequest{
			SearchCode: ValueInput.Name,
		}
		GetBrandsByCodeResp, ErrGetBrandsByCodeResp := user.GetBrandsByCode(GetBrandsByCodeReq)
		if ErrGetBrandsByCodeResp != nil {
			return []Avtoto_Output{}, ErrGetBrandsByCodeResp
		}
		var BrandID string
		for _, val := range GetBrandsByCodeResp.Brands {
			fmt.Println(val.Name, ValueInput.Manufacture)
			if val.Name == ValueInput.Manufacture {
				BrandID = val.Manuf
			}
		}

		// Если такой бренд не найден
		if BrandID == "" {
			//continue
		}

		fmt.Println("BrandID", BrandID)

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
			fmt.Println(SearchStartResp.Error())
			if SearchStartResp.Error() == "" {
				break
			}
			if ErrorStart != nil {
				log.Println(ErrorStart)
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
	defer bar2.Finish()                                                         // Завершить

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
				fmt.Println(SearchResp.Status())
				if ErrorSearch != nil {
					continue
				}
				if SearchResp.Status() == "Запрос в обработке" {
					continue
				}

				output[Index].Parts = Avtoto_Filter(SearchResp)

				// Обработка готового события
				TecalSize++
				output[Index].IsOk = true
				bar2.Increment() // Прибавляем 1 к отображению
				//fmt.Println("Найден результат", SearchResp.Status())
			}
		}
	}

	return output, nil
}

// Фильтрация по бизнес-логике
func Avtoto_Filter(SearchResp avtoto.SearchGetParts2Response) avtoto.SearchGetParts2Response {
	var NewParts avtoto.SearchGetParts2Response

	// Срок доставки меньше 7 и колличество больше 1
	for _, value := range SearchResp.Parts {
		fmt.Println("->", value.Delivery, value.MaxCount)
		delivery, _ := strconv.Atoi(value.Delivery)
		MaxCount, _ := strconv.Atoi(value.MaxCount)
		if delivery < 7 && (MaxCount > 1 || MaxCount == -1) {
			NewParts.Parts = append(NewParts.Parts, value)
		}
	}
	fmt.Println("len", len(NewParts.Parts))

	if len(NewParts.Parts) > 1 {
		sort.Slice(NewParts.Parts, func(i, j int) (less bool) {
			return NewParts.Parts[i].Price < NewParts.Parts[j].Price
		})
	}

	return SearchResp
}
