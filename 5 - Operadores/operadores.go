package main

import "fmt"

func main() {

	// ARITMÉTICOS
	soma := 1 + 2
	subtracao := 1 - 2
	divisao := 10 / 4
	multiplicacao := 10 * 5
	restoDaDivisao := 10 % 2

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDaDivisao)

	var numero1 int16 = 10
	var numero2 int16 = 5
	soma2 := numero1 + numero2
	fmt.Println(soma2)

	// FIM DOS ARITMÉTICOS

	//ATRIBUIÇAO
	var variavel1 string = "String"
	variavel2 := "String2"
	fmt.Println(variavel1, variavel2)

	//OPERADORES RELACIONAIS
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 > 2)
	fmt.Println(1 < 2)
	fmt.Println(1 != 2)
	// FIM DOS RELACIONAIS

	//OPERADORES LÓGICOS
	fmt.Println("/////==========/////")
	verdadeiro, falso := true, false

	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro)
	fmt.Println(!falso)
	//FIM OPERADORES LÓGICOS

	//OPERADORES UNÁRIOS

	numero := 10
	numero++
	numero += 15
	fmt.Println(numero)

	numero--
	numero -=20
	numero *=3
	numero /=10
	numero %=3

	fmt.Print(numero)

	//FIM OPERADORES UNÁRIOS


}
