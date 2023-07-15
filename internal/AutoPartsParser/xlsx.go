// Обработка входного файла, который содержит название детали и производитель.

package AutoPartsParser

import (
	"strconv"
	"time"

	"github.com/xuri/excelize/v2"
)

// Стурктура с запросом, которое будет отправлено серверу
type Request struct {
	Name        string
	Manufacture string
}

func Xlsx(filename string) (ReqXlsx []Request, ErrorFile error) {
	// Открываем файл
	FileXlsx, ErrorFile := excelize.OpenFile(filename)
	if ErrorFile != nil {
		return nil, ErrorFile
	}

	// Получить все строки на первом листе
	rows, err := FileXlsx.GetRows(FileXlsx.GetSheetName(0))
	if err != nil {
		return nil, ErrorFile
	}

	// Добавление в массив запроса
	for _, row := range rows {
		ReqXlsx = append(ReqXlsx, Request{
			Name:        row[0],
			Manufacture: row[1],
		})
	}

	// Закрываем файл
	ErrorFile = FileXlsx.Close()
	if ErrorFile != nil {
		return []Request{}, ErrorFile
	}
	return ReqXlsx, nil
}

// Пакет создан для оперирования с Excel файлами.
// А именно:
// - Создание файла
// - Добавление в него информации по источникам ZZap и avtoto
// Сохранить в файл
func SaveXlsx(AvtotoIsParse, ZZapIsParse bool, Avtoto []Avtoto_Output, ZZap []ZZap_Output) error {
	// Создаём файл
	f := excelize.NewFile()
	defer f.Close()

	if ZZapIsParse {
		f.NewSheet("ZZap")
		f.DeleteSheet("Sheet1")
		saveZZap(f, ZZap)
	}
	if AvtotoIsParse {
		f.NewSheet("Avtoto")
		f.DeleteSheet("Sheet1")
		saveAvtoto(f, Avtoto)
	}

	// Сохраняем
	dt := time.Now()
	if ErrSave := f.SaveAs("Цены от " + dt.Format("2006-01-02_15h04m") + ".xlsx"); ErrSave != nil {
		return ErrSave
	}
	return nil
}

// Сохранить результаты по ZZap
func saveZZap(f *excelize.File, ZZap []ZZap_Output) {
	f.SetCellValue("ZZap", "A1", "Название товара")
	f.SetCellValue("ZZap", "B1", "Артикул")
	f.SetCellValue("ZZap", "C1", "Производитель")
	f.SetCellValue("ZZap", "D1", "количество предложений в наличии запрашиваемой запчасти")
	f.SetCellValue("ZZap", "E1", "минимальная цена среди предложений в наличии запрашиваемой запчасти")
	f.SetCellValue("ZZap", "F1", "средняя цена среди предложений в наличии запрашиваемой запчасти")
	f.SetCellValue("ZZap", "G1", "Максимальная цена среди предложений в наличии запрашиваемой запчасти")
	for index, value := range ZZap {
		f.SetCellValue("ZZap", "A"+strconv.Itoa(index+2), value.RealName)          // Название детали
		f.SetCellValue("ZZap", "B"+strconv.Itoa(index+2), value.Name)              // Название детали
		f.SetCellValue("ZZap", "C"+strconv.Itoa(index+2), value.Manuf)             // Производитель
		f.SetCellValue("ZZap", "D"+strconv.Itoa(index+2), value.PriceCountInstock) // количество предложений в наличии запрашиваемой запчасти
		f.SetCellValue("ZZap", "E"+strconv.Itoa(index+2), value.PriceMinInstock)   // минимальная цена среди предложений в наличии запрашиваемой запчасти
		f.SetCellValue("ZZap", "F"+strconv.Itoa(index+2), value.PriceAvgInstock)   // средняя цена среди предложений в наличии запрашиваемой запчасти
		f.SetCellValue("ZZap", "G"+strconv.Itoa(index+2), value.PriceMaxInstock)   // Максимальная цена среди предложений в наличии запрашиваемой запчасти
	}
}

// Сохранить результаты по Avtoto
func saveAvtoto(f *excelize.File, Avtoto []Avtoto_Output) {
	ssheet := "Avtoto"
	f.SetCellValue(ssheet, "A1", "Артикул")
	f.SetCellValue(ssheet, "B1", "Производитель")
	f.SetCellValue(ssheet, "C1", "Название")

	f.SetCellValue(ssheet, "E1", "Цена1")
	f.SetCellValue(ssheet, "F1", "Цена1*0.9")
	f.SetCellValue(ssheet, "G1", "Склад1")
	f.SetCellValue(ssheet, "H1", "Доставка1")
	f.SetCellValue(ssheet, "I1", "Количество1")

	f.SetCellValue(ssheet, "K1", "Цена2")
	f.SetCellValue(ssheet, "L1", "Цена2*0.9")
	f.SetCellValue(ssheet, "M1", "Склад2")
	f.SetCellValue(ssheet, "N1", "Доставка2")
	f.SetCellValue(ssheet, "O1", "Количество2")

	f.SetCellValue(ssheet, "Q1", "Цена3")
	f.SetCellValue(ssheet, "R1", "Цена3*0.9")
	f.SetCellValue(ssheet, "S1", "Склад3")
	f.SetCellValue(ssheet, "T1", "Доставка3")
	f.SetCellValue(ssheet, "U1", "Количество3")

	for index, value := range Avtoto {
		f.SetCellValue(ssheet, "A"+strconv.Itoa(index+2), value.SKU)
		f.SetCellValue(ssheet, "B"+strconv.Itoa(index+2), value.Manufacture)
		//f.SetCellValue(ssheet, "C"+strconv.Itoa(index+2), value.Name)

		if len(value.Parts.Parts) >= 1 {
			f.SetCellValue(ssheet, "C"+strconv.Itoa(index+2), value.Parts.Parts[0].Name)

			f.SetCellValue(ssheet, "E"+strconv.Itoa(index+2), value.Parts.Parts[0].Price)
			f.SetCellValue(ssheet, "F"+strconv.Itoa(index+2), float64(value.Parts.Parts[0].Price)*0.9)
			f.SetCellValue(ssheet, "G"+strconv.Itoa(index+2), value.Parts.Parts[0].Storage)
			f.SetCellValue(ssheet, "H"+strconv.Itoa(index+2), value.Parts.Parts[0].Delivery)
			f.SetCellValue(ssheet, "I"+strconv.Itoa(index+2), value.Parts.Parts[0].MaxCount)
		}
		if len(value.Parts.Parts) >= 2 {
			f.SetCellValue(ssheet, "K"+strconv.Itoa(index+2), value.Parts.Parts[1].Price)
			f.SetCellValue(ssheet, "L"+strconv.Itoa(index+2), float64(value.Parts.Parts[1].Price)*0.9)
			f.SetCellValue(ssheet, "M"+strconv.Itoa(index+2), value.Parts.Parts[1].Storage)
			f.SetCellValue(ssheet, "N"+strconv.Itoa(index+2), value.Parts.Parts[1].Delivery)
			f.SetCellValue(ssheet, "O"+strconv.Itoa(index+2), value.Parts.Parts[1].MaxCount)
		}
		if len(value.Parts.Parts) >= 3 {
			f.SetCellValue(ssheet, "Q"+strconv.Itoa(index+2), value.Parts.Parts[2].Price)
			f.SetCellValue(ssheet, "R"+strconv.Itoa(index+2), float64(value.Parts.Parts[2].Price)*0.9)
			f.SetCellValue(ssheet, "S"+strconv.Itoa(index+2), value.Parts.Parts[2].Storage)
			f.SetCellValue(ssheet, "T"+strconv.Itoa(index+2), value.Parts.Parts[2].Delivery)
			f.SetCellValue(ssheet, "U"+strconv.Itoa(index+2), value.Parts.Parts[2].MaxCount)
		}
	}
}
