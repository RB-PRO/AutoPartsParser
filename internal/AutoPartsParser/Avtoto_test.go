package AutoPartsParser_test

import (
	"fmt"
	"testing"

	"github.com/RB-PRO/AutoPartsParser/internal/AutoPartsParser"
	"github.com/RB-PRO/avtoto"
)

func TestGetBrandsByCode(t *testing.T) {

	// Получить данные из json-файла
	AvtotoDataAuf, ErrorDataFile := AutoPartsParser.DataFile("avtoto.json")
	if ErrorDataFile != nil {
		t.Error(ErrorDataFile)
	}

	// Создать экземпляр пользователя api avtoto
	user := avtoto.User{
		UserId:       AvtotoDataAuf.UserId,
		UserLogin:    AvtotoDataAuf.UserLogin,
		UserPassword: AvtotoDataAuf.UserPassword,
	}

	myBrand := avtoto.GetBrandsByCodeRequest{SearchCode: "N007603010406"}
	GetBrandsByCodeResp, ErrGetBrandsByCodeResp := user.GetBrandsByCode(myBrand)
	if ErrGetBrandsByCodeResp != nil {
		t.Error(ErrGetBrandsByCodeResp)
	}

	fmt.Printf("%+#v", GetBrandsByCodeResp)
}
