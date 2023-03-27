package AutoPartsParser

func Parse(AvtotoIsParse, ZZapIsParse bool) error {

	filename := "article.xlsx"

	ReqXlsx, errorXlsx := Xlsx(filename)
	if errorXlsx != nil {
		return errorXlsx
	}

	if ZZapIsParse {
		errorZzap := ZZapParse(ReqXlsx)
		if errorZzap != nil {
			return errorZzap
		}
	}

	return nil
}
