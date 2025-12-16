package main

import "fmt"

const usd_eur float64 = 0.87
const usd_rub float64 = 81.25

type MapSF64 = map[string]float64

func main() {
	val, first_val, second_val := valut()
	valut := map[string]*MapSF64{
		"RUB": &MapSF64{
			"USD": val / usd_rub,
			"EUR": (val / usd_rub) * usd_eur,
		},
		"USD": &MapSF64{
			"RUB": val * usd_rub,
			"EUR": val * usd_eur,
		},
		"EUR": &MapSF64{
			"RUB": (val / usd_eur) * usd_rub,
			"USD": val / usd_eur,
		},
	}
	res := convent(&valut, first_val, second_val)
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

func convent(valut *map[string]*MapSF64, first_val, second_val string) float64 {
	return (*(*valut)[first_val])[second_val]
}
