package main

import "fmt"

func main() {
	var tabuadaCinco = tabuada(5)
	tabuadaCinco()
}

func tabuada(valor int) func() {
	return func() {
		for index := 0; index <= 10; index++ {
			fmt.Println(index, " * ", valor, " = ", valor*index)
		}
	}
}
