package main

import "fmt"

func main() {

	transactions := make([]float64, 0, 30)
	for {
		fmt.Println("Здравствуйте, скики вы сегодня потратили ? Напишите число > ")
		coastUser := getScanTransaction()

		if coastUser <= 0 {
			break
		}

		transactions = append(transactions, coastUser)

		fmt.Println("Ваша общаяя сумма трат:", addingSum(transactions))
	}
}

func addingSum(transactions []float64) float64 {
	var sum float64
	for _, value := range transactions {
		sum += value
	}
	return sum
}

func getScanTransaction() float64 {
	var coastUser float64

	fmt.Scan(&coastUser)
	return coastUser
}
