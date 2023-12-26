package main

import (
	"fmt"
	"modulo/auxiliar"
	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Escrevendo arquivos main")
	auxiliar.Escrever()

	erro := checkmail.ValidateFormat("valter.junior@asapcred.com.br")
	fmt.Println(erro)
	
	
}
