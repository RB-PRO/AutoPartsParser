// Обработка входного файла, который содержит название детали и производитель.

package AutoPartsParser

import (
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
