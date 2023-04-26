package internal

import (
	"log"
	"strconv"
	"strings"
	"time"

	bank "second_project/bank"
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
func FindMax(courses []bank.CurrencyRates) (string, string, float64) {
	maxString := strings.Replace(courses[0].Record[0].Value, ",", ".", 1)
	max, err := strconv.ParseFloat(maxString, 64)
	if err != nil {
		log.Fatal("невозможно перевести string в float в функции findMax:", err)
	}

	// здесь сохранится ответ
	id := ""
	date := ""
	value := 0.0

	for _, v := range courses {
		if len(v.Record) == 0 {
			log.Fatal("нет данных")
		}
		for i := 1; i < len(v.Record); i++ { // начинаем со второго элемента, потому что первый уже записан в max
			curValString := strings.Replace(v.Record[i].Value, ",", ".", 1)
			curValue, err := strconv.ParseFloat(curValString, 64)

			if err != nil {
				log.Fatal("невозможно перевести string в float в функции findMax:", err)
			}

			if curValue > max {
				max = curValue
				id = v.Record[i].Id
				date = v.Record[i].Date
				value = curValue
			}
		}
	}

	return id, date, value
}

// FindMin находит минимальное значение и возвращает дату, значение и ID валюты
func FindMin(courses []bank.CurrencyRates) (string, string, float64) {
	minString := strings.Replace(courses[0].Record[0].Value, ",", ".", 1)
	min, err := strconv.ParseFloat(minString, 64)
	if err != nil {
		log.Fatal("невозможно перевести string в float в функции findMin:", err)
	}
	ID := ""
	date := ""
	value := 0.0

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

			if curValue < min {
				min = curValue
				ID = v.Record[i].Id
				date = v.Record[i].Date
				value = curValue
			}
		}
	}

	return ID, date, value
}

// FindAverage находит среднее значение рубля
func FindAverage(courses []bank.CurrencyRates) float64 {
	sum := 0.0
	count := 0.0

	for _, v := range courses {
		if len(v.Record) == 0 {
			log.Fatal("нет данных")
		}
		for _, val := range v.Record {			
			curValString := strings.Replace(val.Value, ",", ".", 1)
			curValue, err := strconv.ParseFloat(curValString, 64)

			if err != nil {
				log.Fatal("невозможно перевести string в float в функции findAverage:", err)
			}

			sum += curValue
			count++
		}
	}

	return sum / count
}
