package main

import (
	"fmt"      // Bib para formatacao
	"net/http" // Bib para internet
	"os"       // Bib para comandos de sistema
	"time"     // Bib para comandos de tempo
)

const MONITORAMENTOS = 3
const DELAY = 5

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
	fmt.Println("Este código fará requisições na internet")
}

func exibeMenu() {
	fmt.Println("\n1 - Realizar requisição.")
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
	fmt.Println("Realizando requisições...")
	sites := []string{"https://www.uol.com.br/", "https://g1.globo.com/", "https://portal.ifrn.edu.br/"}

	for i := 0; i < MONITORAMENTOS; i++ {
		fmt.Println("\n---------------------------------------------", i, "---------------------------------------------")
		for _, site := range sites {
			fmt.Println("\nVerificando", site)
			verificaSite(site)
		}
		time.Sleep(DELAY * time.Second)
	}
	fmt.Println("")
}

func verificaSite(site string) {
	resp, _ := http.Get(site)

	if resp.StatusCode > 199 && resp.StatusCode < 300 {
		fmt.Println("Site:", site, "foi carregado com sucesso!", "Status:", resp.StatusCode)
	} else {
		fmt.Println("Site:", site, "está com problemas.", "Status:", resp.StatusCode)
	}
}
