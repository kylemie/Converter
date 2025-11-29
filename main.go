package main

import "fmt"

const usd_eur float64 = 0.87
const usd_rub float64 = 81.25

func main() {
	val, first_val, second_val := valut()
	res := convent(val, first_val, second_val)
	fmt.Printf("Итоговая сумма: %.1f", res)
}
func valut() (float64, string, string) {
	for {
		var val float64
		var first_val, second_val string
		fmt.Print("Введите исходную валюту RUB/USD/EUR: ")
		fmt.Scan(&first_val)
		if first_val != "RUB" && first_val != "USD" && first_val != "EUR" {
			fmt.Println("Неправильно введена валюта, попробуйте снова")
			continue
		}
		fmt.Print("Введите число: ")
		fmt.Scan(&val)
		if val <= 0 {
			fmt.Println("Число должно быть положительным, попробуйте снова")
			continue
		}
		switch first_val {
		case "RUB":
			fmt.Print("Введите целевую валюту USD/EUR: ")
		case "USD":
			fmt.Print("Введите целевую валюту RUB/EUR: ")
		case "EUR":
			fmt.Print("Введите целевую валюту RUB/USD: ")
		}
		fmt.Scan(&second_val)
		switch first_val {
		case "RUB":
			if second_val != "USD" && second_val != "EUR" {
				fmt.Println("Неправильно введена валюта, попробуйте снова")
				continue
			}
		case "USD":
			if second_val != "RUB" && second_val != "EUR" {
				fmt.Println("Неправильно введена валюта, попробуйте снова")
				continue
			}
		case "EUR":
			if second_val != "RUB" && second_val != "USD" {
				fmt.Println("Неправильно введена валюта, попробуйте снова")
				continue
			}
		}
		return val, first_val, second_val
	}
}

func convent(val float64, first_val, second_val string) float64 {
	var res float64
	switch {
	case first_val == "RUB":
		switch {
		case second_val == "EUR":
			res = (val / usd_rub) * usd_eur
		case second_val == "USD":
			res = val / usd_rub
		}
	case first_val == "EUR":
		switch {
		case second_val == "RUB":
			res = (val / usd_eur) * usd_rub
		case second_val == "USD":
			res = val / usd_eur
		}
	case first_val == "USD":
		switch {
		case second_val == "RUB":
			res = val * usd_rub
		case second_val == "EUR":
			res = val * usd_eur
		}
	}
	return res
}
