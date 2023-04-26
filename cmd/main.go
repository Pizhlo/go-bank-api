package main

import (
	"fmt"
	bank "github.com/Pizhlo/go-bank-api/bank"
	internal "github.com/Pizhlo/go-bank-api/internal"
)

func main() {
	days := 90                    // количество интересующих дней
	dates := internal.MakeDates(days) // создали список дат
	fmt.Printf("Рассматриваемый период: %s - %s\n\n", dates[0], dates[1])

	currencies, namesCodeMap := bank.GetCurrencies() // получаем список всех валют
	curRates := bank.GetCurRates(dates, currencies)  // получаем курс всех валют за указанный период

	average := internal.FindAverage(curRates)            // нашли среднее значение рубля
	maxID, maxDate, maxVal := internal.FindMax(curRates) // нашли максимум
	minID, minDate, minVal := internal.FindMin(curRates) // нашли минимум

	fmt.Printf("Максимальный курс валюты: %f P; название валюты: %s; дата: %s\n", maxVal, namesCodeMap[maxID], maxDate)
	fmt.Printf("Минимальный курс валюты: %f P; название валюты: %s; дата: %s\n", minVal, namesCodeMap[minID], minDate)
	fmt.Printf("Средний курс рубля: %f\n", average)

}
