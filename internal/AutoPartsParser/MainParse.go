package AutoPartsParser

func Parse(AvtotoIsParse, ZZapIsParse bool) (ErrorParse error) {
	var ZZapOutput []ZZap_Output
	var AvtotoOutput []Avtoto_Output
	filename := "article.xlsx"

	ReqXlsx, errorXlsx := Xlsx(filename)
	if errorXlsx != nil {
		return errorXlsx
	}

	if AvtotoIsParse {
		AvtotoOutput, ErrorParse = AvtotoParse(ReqXlsx)
		if ErrorParse != nil {
			return ErrorParse
		}
	}
	if ZZapIsParse {
		ZZapOutput, ErrorParse = ZZapParse(ReqXlsx)
		if ErrorParse != nil {
			return ErrorParse
		}
	}

	SaveXlsx(AvtotoIsParse, ZZapIsParse, AvtotoOutput, ZZapOutput)
	return nil
}
