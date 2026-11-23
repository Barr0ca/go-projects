package main

import (
	"fmt"      // Bib para formatacao
	"net/http" // Bib para internet
	"os"       // Bib para comandos de sistema
)

func main() {
	apresentacao()

	// Nao existe "While" em Go, "for" sem parametros se torna um "While True"
	for {
		exibeMenu()

		opcao := escolheOpcao()

		switch opcao {
		case 1:
			reqInternet()
		case 0:
			fmt.Println("Saindo...")
			os.Exit(0)
		default:
			fmt.Println("Não entendi.\n")
		}
	}
}

func apresentacao() {
	nome := "Fulano"
	fmt.Println("Olá,", nome)
	fmt.Println("Este código fará requisições na internet \n")
}

func exibeMenu() {
	fmt.Println("1 - Realizar requisição.")
	fmt.Println("0 - Sair\n")
}

func escolheOpcao() int {
	var opcao int
	fmt.Scan(&opcao)

	return opcao
}

// Funcao "Get" retorna 2 valores, para ignorar
// e utilizar apenas 1 desses valores utiliza-se
// "_" no espaco do valor a ser ignorado, nesse
// caso vamos ignorar o retorno de "error".

func reqInternet() {
	fmt.Println("Realizando requisição...")
	site := "https://httpbin.org/status/404" // ou 200
	resp, _ := http.Get(site)

	if resp.StatusCode > 199 && resp.StatusCode < 300 {
		fmt.Println("Site:", site, "foi carregado com sucesso!", "Status:", resp.StatusCode)
	} else {
		fmt.Println("Site:", site, "está com problemas.", "Status:", resp.StatusCode)
	}
}
