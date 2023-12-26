package main

import "fmt"

func main() {
	var variavel1 string = "Variável 1"

	varialve2 := "Variável 2"

	variavel3 := 1  

	fmt.Println(variavel1)
	fmt.Println(varialve2)
	fmt.Println(variavel3)

	var (
		varialvel4 string = "lalala"
		variavel5 string = "alalalal"
	)
    fmt.Println(varialvel4, variavel5)


	varialvel6, variavel7 := "Variável 6" , "Variável 7"
	fmt.Println(varialvel6, variavel7)

	const constante1 string = "Constante 1"

	fmt.Println(constante1)
	

	varialvel6 , variavel7 = variavel7, varialvel6
	fmt.Print(varialvel6, variavel7)
	

}
