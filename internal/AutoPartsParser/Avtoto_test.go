package AutoPartsParser_test

import (
	"fmt"
	"strings"
	"testing"
	"time"

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

	myBrand := avtoto.GetBrandsByCodeRequest{SearchCode: "FSG500030"}
	GetBrandsByCodeResp, ErrGetBrandsByCodeResp := user.GetBrandsByCode(myBrand)
	if ErrGetBrandsByCodeResp != nil {
		t.Error(ErrGetBrandsByCodeResp)
	}

	fmt.Printf("len - %v\n%+#v\n", len(GetBrandsByCodeResp.Brands), GetBrandsByCodeResp)
	for i, val := range GetBrandsByCodeResp.Brands {
		fmt.Println(i, val.Name)
	}

}

func TestSearch(t *testing.T) {
	// Получить данные из json-файла
	AvtotoDataAuf, _ := AutoPartsParser.DataFile("avtoto.json")

	// Создать экземпляр пользователя api avtoto
	user := avtoto.User{
		UserId:       AvtotoDataAuf.UserId,
		UserLogin:    AvtotoDataAuf.UserLogin,
		UserPassword: AvtotoDataAuf.UserPassword,
	}
	////////////////////
	ValueInputName := "XGB500020"
	ValueInputManufacture := "LAND ROVER"
	GetBrandsByCodeReq := avtoto.GetBrandsByCodeRequest{
		SearchCode: ValueInputName,
	}
	GetBrandsByCodeResp, ErrGetBrandsByCode := user.GetBrandsByCode(GetBrandsByCodeReq)
	if ErrGetBrandsByCode != nil {
		t.Error(ErrGetBrandsByCode)
	}

	fmt.Printf(">>%+v\n", GetBrandsByCodeResp)
	var BrandID string
	for _, val := range GetBrandsByCodeResp.Brands {
		// fmt.Println("---", val.Name, val.Manuf, ValueInput.Manufacture)
		if strings.EqualFold(val.Manuf, ValueInputManufacture) {
			BrandID = val.Manuf
		}
	}

	fmt.Println("BrandID", BrandID)
	//
	// Создать запрос старта поиска
	SearchStartReq := avtoto.SearchStartRequest{
		SearchCode: ValueInputName,
		// Brand:       BrandID,
		SearchCross: "off",
	}
	SearchStartResp, ErrSearchStartRequest := user.SearchStartRequest(SearchStartReq)
	if ErrSearchStartRequest != nil {
		t.Error(ErrSearchStartRequest)
	}

	//
	//
	fmt.Println("[RB_PRO]: Ждём 15 секунд")
	time.Sleep(15 * time.Second)
	// Формируем запрос
	SearchGetParts2Req := avtoto.SearchGetParts2Request{
		ProcessSearchId: SearchStartResp.ProcessSearchID,
		Limit:           500,
	}

	// Выполняем запрос
	SearchResp, ErrSearchGetParts2 := SearchGetParts2Req.SearchGetParts2()
	if ErrSearchGetParts2 != nil {
		t.Error(ErrSearchGetParts2)
	}

	fmt.Println(SearchResp)

	for i, val := range SearchResp.Parts {
		fmt.Printf("%d.\t%d\t%s\t%s\t%s\n", i, val.Price, val.Storage, val.Manuf, val.Name)
	}

	//

	SearchResp = AutoPartsParser.Avtoto_Filter(SearchResp, ValueInputManufacture)

	for i, val := range SearchResp.Parts {
		fmt.Printf("%d.\t%d\t%s\t%s\t%s\n", i, val.Price, val.Storage, val.Manuf, val.Name)
	}
}
