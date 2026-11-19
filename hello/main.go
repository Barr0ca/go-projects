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
		fmt.Println("Iniciando monitoramento...")
	case 2:
		fmt.Println("Exibindo logs...")
	case 0:
		fmt.Println("Saindo...")
		os.Exit(0)
	default:
		fmt.Println("Opção inválida!")
		os.Exit(-1)
	}
}

func apresentacao() {
	nome := "Ian"
	var versao float32 = 1.1
	fmt.Println("Olá,", nome)
	fmt.Println("Este código está na versão", versao, "\n")
}

func exibeMenu() {
	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - Ver logs")
	fmt.Println("0 - Sair do programa")
}

func escolheOpcao() int { // "int" após a função para informar o tipo da saída
	var opcao int
	fmt.Scan(&opcao)

	return opcao
}
