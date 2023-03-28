package AutoPartsParser

import "fmt"

// Пропарсить по данным из файла filename и сохранить результат в Xlsx
func Parse(AvtotoIsParse, ZZapIsParse bool) (ErrorParse error) {
	var ZZapOutput []ZZap_Output
	var AvtotoOutput []Avtoto_Output
	filename := "article.xlsx"

	ReqXlsx, errorXlsx := Xlsx(filename)
	if errorXlsx != nil {
		return errorXlsx
	}

	fmt.Println("[RB_PRO]: AVTOTO")
	if AvtotoIsParse {
		AvtotoOutput, ErrorParse = AvtotoParse(ReqXlsx)
		if ErrorParse != nil {
			return ErrorParse
		}
	}

	fmt.Println("[RB_PRO]: ZZAP")
	if ZZapIsParse {
		ZZapOutput, ErrorParse = ZZapParse(ReqXlsx)
		if ErrorParse != nil {
			return ErrorParse
		}
	}

	SaveXlsx(AvtotoIsParse, ZZapIsParse, AvtotoOutput, ZZapOutput)
	return nil
}
