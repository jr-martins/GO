package main

import "fmt"

func main() {

	numero := -1

	if numero > 15 {
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("Menor igual a 15")
	}

	if outroNumero := numero; outroNumero > 0 {
		fmt.Println("Número é maior que zero")
	} else {
		fmt.Println("Número é menor que zero")
	}

}
