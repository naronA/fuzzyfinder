package address

import (
	"os"
	"strings"

	"github.com/gocarina/gocsv"
)

type zenkoku struct {
	AddressCD        string `csv:"住所CD"`
	TodofukenCD      string `csv:"都道府県CD"`
	ShikuchosonCD    string `csv:"市区町村CD"`
	ChoikiCD         string `csv:"町域CD"`
	Zip              string `csv:"郵便番号"`
	JigyosyoFlag     string `csv:"事業所フラ"`
	HaichiFlag       int    `csv:"廃止フラグ"`
	Todofuken        string `csv:"都道府県"`
	TodofukenKana    string `csv:"都道府県カナ"`
	Shikuchoson      string `csv:"市区町村"`
	ShikuchosonKana  string `csv:"市区町村カナ"`
	Choiki           string `csv:"町域"`
	ChoikiKana       string `csv:"町域カナ"`
	ChoikiHosoku     string `csv:"町域補足"`
	KyotoTorina      string `csv:"京都通り名"`
	Azachome         string `csv:"字丁目"`
	AzachomeKana     string `csv:"字丁目カナ"`
	Hosoku           string `csv:"補足"`
	Jigyosyomei      string `csv:"事業所名"`
	JigyosyomeiKana  string `csv:"事業所名カナ"`
	JigyosyomeiJusyo string `csv:"事業所住所"`
	ShinjyusyoCD     string `csv:"新住所CD"`
}

func Load() []string {
	file, err := os.Open("address/zenkoku.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	zenkokus := []*zenkoku{}

	if err := gocsv.UnmarshalFile(file, &zenkokus); err != nil {
		panic(err)
	}
	allAddress := make([]string, 0)
	for _, z := range zenkokus {
		zip := "〒" + z.Zip
		prefecture := z.Todofuken
		city := z.Shikuchoson
		town := z.Choiki
		kyoto := z.KyotoTorina
		aza := z.Azachome
		// jigyo := z.Jigyosyomei
		// jigyoAddress := z.JigyosyomeiJusyo
		add := strings.Join([]string{zip, prefecture, city, town, kyoto, aza}, "")
		allAddress = append(allAddress, add)
	}
	return allAddress
}
