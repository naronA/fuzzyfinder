package main

import (
	"fmt"
	"net/http"
	"os"
	"sort"
	"strings"

	"github.com/gocarina/gocsv"
	"github.com/labstack/echo"
	"github.com/naronA/fuzzyfinder/score"
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

func load() []string {
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
func findAddress(addresses []string) echo.HandlerFunc {
	return func(c echo.Context) error {
		term := c.QueryParams()["term"]
		finders := make([]score.Finder, 0)
		for _, add := range addresses {
			f := score.Finder{Source: add, Inputs: term}
			finders = append(finders, f)
		}
		output := ""
		for _, f := range finders {
			output += f.String() + "\n"
		}
		return c.String(http.StatusOK, output)
	}
}

func main() {
	addresses := load()
	finders := make(score.Finders, 0)
	for _, add := range addresses {
		f := score.Finder{Source: add, Inputs: []string{"本郷"}}
		if f.Score() >= 1 {
			finders = append(finders, f)
		}
	}
	sort.Sort(sort.Reverse(finders))
	for _, f := range finders[:30] {
		if f.String() != "" {
			fmt.Printf("%v : %v\n", f.Highlight(), f.Score())
		}
	}

}
