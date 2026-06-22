package main

import "fmt"

func main() {
	nome := "Ian"
	var versao float32 = 1.1

	fmt.Println("Olá,", nome)
	fmt.Println("Este código está na versão", versao, "\n")

	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - Ver logs")
	fmt.Println("0 - Sair do programa")

	var opcao int
	fmt.Scan(&opcao)

	// if opcao == 1 {
	// 	fmt.Println("Iniciando monitoramento...")
	// } else if opcao == 2 {
	// 	fmt.Println("Exibindo logs...")
	// } else if opcao == 0 {
	// 	fmt.Println("Saindo...")
	// } else {
	// 	fmt.Println("Opção inválida!")
	// }

	switch opcao {
	case 1:
		fmt.Println("Iniciando monitoramento...")
	case 2:
		fmt.Println("Exibindo logs...")
	case 0:
		fmt.Println("Saindo...")
	default:
		fmt.Println("Opção inválida!")
	}
}
