package main

import "fmt"

func main() {

	var valorConverter float32

	fmt.Print("Digite a quantidade de graus Fahrenheit: ")
	fmt.Scanln(&valorConverter)

	temperaturaConvertida := fahrenheitToCelsius(valorConverter)

	fmt.Println(valorConverter, "° Fahrenheit são ", temperaturaConvertida, "° Celsius")
	fmt.Printf("Celsius %.2f\n", temperaturaConvertida)
}

func fahrenheitToCelsius(origem float32) float32 {

	return (origem - 32) / 1.8
}

// func celsiusToFahrenheit(origem float32) float32 {
//
// 	return origem*1.8 + 32
//
// }
