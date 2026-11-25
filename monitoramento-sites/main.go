package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

const MONITORAMENTOS = 3
const DELAY = 1

func main() {
	apresentacao()

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

func reqInternet() {
	fmt.Println("Realizando requisições...")
	sites := lerArquivo()

	for i := 0; i < MONITORAMENTOS; i++ {
		fmt.Println("\n---------------------------------------------", i, "---------------------------------------------")
		for _, site := range sites {
			fmt.Println("\nVerificando", site)
			verificaSite(site)
		}
		time.Sleep(DELAY * time.Minute)
	}
	fmt.Println("")
}

func verificaSite(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	if resp.StatusCode > 199 && resp.StatusCode < 300 {
		fmt.Println("Site:", site, "foi carregado com sucesso!", "Status:", resp.StatusCode)
	} else {
		fmt.Println("Site:", site, "está com problemas.", "Status:", resp.StatusCode)
	}
}

func lerArquivo() []string {
	var sites []string

	arquivo, err := os.Open("monitoramento-sites/sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	leitor := bufio.NewReader(arquivo)

	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)

		sites = append(sites, linha)

		if err == io.EOF {
			break
		}
	}

	arquivo.Close()

	return sites
}
