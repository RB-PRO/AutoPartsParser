package AutoPartsParser_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/RB-PRO/AutoPartsParser/internal/AutoPartsParser"
)

func TestXlsx(t *testing.T) {
	filename := "article_test.xlsx"
	ReqXlsx, errorXlsx := AutoPartsParser.Xlsx(filename)
	if errorXlsx != nil {
		t.Error(errorXlsx)
	}
	fmt.Println(ReqXlsx)
	if len(ReqXlsx) == 0 {
		t.Error(errors.New("Длина данных из файла " + filename + " равна нулю."))
	}
	fmt.Println("Файл", filename)
	for index, value := range ReqXlsx {
		fmt.Println(index, "-", value.Name, ",", value.Manufacture)
	}
}
