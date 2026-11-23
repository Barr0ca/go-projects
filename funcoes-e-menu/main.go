package main

import (
	"fmt"
	"os"
)

func main() {
	apresentacao()
	exibeMenu()

	opcao := escolheOpcao()

	switch opcao {
	case 1:
		fmt.Println("O dia está ótimo.")
	case 2:
		fmt.Println("Café da tarde é muito bom.")
	case 3:
		fmt.Println("A lua está muito bonita.")
	case 0:
		fmt.Println("Saindo...")
		os.Exit(0)
	default:
		fmt.Println("Não entendi.")
		os.Exit(1)
	}
}

func apresentacao() {
	nome := "Fulano"
	var versao float32 = 1.1
	fmt.Println("Olá,", nome)
	fmt.Println("Este código está na versão", versao, "\n")
}

func exibeMenu() {
	fmt.Println("1 - Bom Dia!")
	fmt.Println("2 - Boa Tarde!")
	fmt.Println("3 - Boa Noite!")
	fmt.Println("0 - Sair.")
}

func escolheOpcao() int {
	var opcao int
	fmt.Scan(&opcao)

	return opcao
}
