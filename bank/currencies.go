package bank

import (
	"encoding/xml"
	"fmt"
	"log"
	"strings"

	internet "github.com/Pizhlo/go-bank-api/internet"

	"golang.org/x/net/html/charset"
)

// Currency хранит в себе информацию о валютах
type Currency struct {
	Item []struct {
		ID         string `xml:"ID,attr"`
		Name       string `xml:"Name"`
		EngName    string `xml:"EngName"`
		Nominal    uint   `xml:"Nominal"`
		ParentCode string `xml:"ParentCode"`
	} `xml:"Item"`
}

// CurrencyRates хранит в себе информацию о курсе валют
type CurrencyRates struct {
	ID     string `xml:"ID,attr"`
	Record []struct {
		Date    string `xml:"Date,attr"`
		Id      string `xml:"Id,attr"`
		Nominal uint   `xml:"Nominal"`
		Value   string `xml:"Value"`
	} `xml:"Record"`
}

var (
	cur   Currency
	rates CurrencyRates
)

// GetCurRates запрашивает с сервера курс всех валют за указанный период
func GetCurRates(days []string, curs Currency) []CurrencyRates {
	url := "http://www.cbr.ru/scripts/XML_dynamic.asp?date_req1=%s&date_req2=%s&VAL_NM_RQ=%s"
	res := []CurrencyRates{}
	for _, item := range curs.Item {

		resp := internet.MakeRequest(fmt.Sprintf(url, days[0], days[1], item.ParentCode))
		decoder := xml.NewDecoder(resp)
		decoder.CharsetReader = charset.NewReaderLabel
		defer resp.Close()

		err := decoder.Decode(&rates)
		if err != nil {
			log.Fatal("не удалось кодировать ответ c сервера:", err)
		}

		res = append(res, rates)

	}

	return res

}

// GetCurrencies создает список валют, по которым будет произведен запрос
func GetCurrencies() (Currency, map[string]string) {
	curNames := map[string]string{} // для хранения названий и кодов валют
	url := "http://www.cbr.ru/scripts/XML_val.asp?d=0"

	resp := internet.MakeRequest(url)

	defer resp.Close()

	decoder := xml.NewDecoder(resp)
	decoder.CharsetReader = charset.NewReaderLabel
	err := decoder.Decode(&cur)
	if err != nil {
		fmt.Println(err)
	}

	for _, val := range cur.Item {
		curNames[strings.TrimSpace(val.ParentCode)] = val.Name
	}

	return cur, curNames
}
