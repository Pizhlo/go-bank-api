package internal

import (
	"log"
	"strconv"
	"strings"
	"time"

	bank "github.com/Pizhlo/go-bank-api/bank"
)

const (
	format = "02/01/2006"
)

// MakeDates создает список дат, по которым будет произведен запрос. days - количество дней
func MakeDates(days int) []string {
	if days < 0 {
		log.Fatal("количество дней должно быть положительным")
	}
	now := time.Now()
	res := []string{}

	date := now.AddDate(0, 0, -days).Format(format)
	res = append(res, date)
	res = append(res, now.Format(format))

	return res
}

// FindMax находит максимальное значение и возвращает дату, значение и ID валюты
// func FindMax(courses []bank.CurrencyRates) (string, string, float64) {
// 	maxString := strings.Replace(courses[0].Record[0].Value, ",", ".", 1)
// 	max, err := strconv.ParseFloat(maxString, 64)
// 	if err != nil {
// 		log.Fatal("невозможно перевести string в float в функции findMax:", err)
// 	}

// 	// здесь сохранится ответ
// 	id := ""
// 	date := ""
// 	value := 0.0

// 	for _, v := range courses {
// 		if len(v.Record) == 0 {
// 			log.Fatal("нет данных")
// 		}
// 		for i := 1; i < len(v.Record); i++ { // начинаем со второго элемента, потому что первый уже записан в max
// 			curValString := strings.Replace(v.Record[i].Value, ",", ".", 1)
// 			curValue, err := strconv.ParseFloat(curValString, 64)

// 			if err != nil {
// 				log.Fatal("невозможно перевести string в float в функции findMax:", err)
// 			}

// 			if curValue > max {
// 				max = curValue
// 				id = v.Record[i].Id
// 				date = v.Record[i].Date
// 				value = curValue
// 			}
// 		}
// 	}

// 	return id, date, value
// }

// FindMinMaxAverage находит минимальное, максимальное и среднее значение и возвращает дату, значение и ID валюты
func FindMinMaxAverage(courses []bank.CurrencyRates) (string, string, float64, string, string, float64, float64) {
	minString := strings.Replace(courses[0].Record[0].Value, ",", ".", 1)
	min, err := strconv.ParseFloat(minString, 64)
	max := min
	if err != nil {
		log.Fatal("невозможно перевести string в float в функции findMin:", err)
	}
	minID := ""
	minDate := ""
	minValue := 0.0

	maxID := ""
	maxDate := ""
	maxValue := 0.0

	sum := 0.0
	count := 0.0

	for _, v := range courses {
		if len(v.Record) == 0 {
			log.Fatal("нет данных")
		}
		for i := 1; i < len(v.Record); i++ {
			curValString := strings.Replace(v.Record[i].Value, ",", ".", 1)
			curValue, err := strconv.ParseFloat(curValString, 64)

			if err != nil {
				log.Fatal("невозможно перевести string в float в функции findMin:", err)
			}

			if curValue / float64(v.Record[i].Nominal) < min {
				min = curValue / float64(v.Record[i].Nominal)
				minID = v.Record[i].Id
				minDate = v.Record[i].Date
				minValue = curValue
			}

			if curValue / float64(v.Record[i].Nominal) > max {
				max = curValue / float64(v.Record[i].Nominal)
				maxID = v.Record[i].Id
				maxDate = v.Record[i].Date
				maxValue = curValue
			}

			sum += curValue / float64(v.Record[i].Nominal)
			count++
		}
	}

	return minID, minDate, minValue, maxID, maxDate, maxValue, sum / count
}

// FindAverage находит среднее значение рубля
// func FindAverage(courses []bank.CurrencyRates) float64 {
// 	sum := 0.0
// 	count := 0.0

// 	for _, v := range courses {
// 		if len(v.Record) == 0 {
// 			log.Fatal("нет данных")
// 		}
// 		for _, val := range v.Record {			
// 			curValString := strings.Replace(val.Value, ",", ".", 1)
// 			curValue, err := strconv.ParseFloat(curValString, 64)

// 			if err != nil {
// 				log.Fatal("невозможно перевести string в float в функции findAverage:", err)
// 			}

// 			sum += curValue
// 			count++
// 		}
// 	}

// 	return sum / count
// }
