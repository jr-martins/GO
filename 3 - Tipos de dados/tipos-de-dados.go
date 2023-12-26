package main

import (
	"errors"
	"fmt"
)

func main(){
	var numero int = 1000000000

	fmt.Println(numero)


	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)


	var numeroReal2 float64 = 123000000.45
	fmt.Println(numeroReal2)

	 numeroReal3 := 123000000.45
	fmt.Println(numeroReal3)


	char := 'B'
	fmt.Println(char)

	var boleano bool = true
	fmt.Println(boleano)

	var erro error
	fmt.Println(erro)


	var errors error = errors.New("Erro interno")
	fmt.Println(errors)
} 