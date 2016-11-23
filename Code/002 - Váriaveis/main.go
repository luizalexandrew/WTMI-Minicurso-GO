package main

import "fmt"

func main() {

	// bool,string,int,int8,int16,int32,int64
	// uint,uint8,uint16,uint32,uint64,uintptr
	// byte,rune,float32,float64,complex64,complex128

	var numero int
	numero = 25
	fmt.Println(numero + 15)

	idade := 22

	fmt.Println("Minha idade é: ", idade)
	// idade = "Mandioca"
	// fmt.Println("Minha idade é: ", idade)

	nome, codigo := "Google", 1

	fmt.Println(nome, codigo)

	var final = "luiz"
	fmt.Println(final)

	var (
		personagen         = "Goku"
		inimigo1, inimigo2 = "Cell", "Majin BOO"
		vida               = 23
	)

	fmt.Println(personagen, inimigo1, inimigo2, "vida:", vida)

}
