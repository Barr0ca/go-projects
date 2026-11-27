package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv" // Bib que converte tipos em string
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
			reqInternet("monitoramento-sites/sites.txt")
		case 2:
			imprimeLogs("monitoramento-sites/logs.txt")
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
	fmt.Println("\n1 - Iniciar monitoramento.")
	fmt.Println("2 - Exibir logs.")
	fmt.Println("0 - Sair\n")
}

func escolheOpcao() int {
	var opcao int
	fmt.Scan(&opcao)

	return opcao
}

func reqInternet(arquivoTxt string) {
	fmt.Println("Realizando requisições...")
	sites := lerArquivo(arquivoTxt)

	for i := 0; i < MONITORAMENTOS; i++ {
		fmt.Println("\n---------------------------------------------", i+1, "---------------------------------------------")
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
		registraLogs(site, true)
	} else {
		fmt.Println("Site:", site, "está com problemas.", "Status:", resp.StatusCode)
		registraLogs(site, false)
	}
}

func lerArquivo(arquivoTxt string) []string {
	var sites []string

	arquivo, err := os.Open(arquivoTxt)

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

func registraLogs(site string, status bool) {
	// O arquivo de log e criado caso nao exista (O_CREATE) podendo escrever e ler (O_RDWR)
	// alem disso, com permissao total (0666)
	arquivo, err := os.OpenFile("monitoramento-sites/logs.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	// A linha e escrita no fim do arquivo devido ao "O_APPEND", flag de "OpenFile"
	// Para formatar o tempo com dd/mm/aaaa, por exemplo, verifica-se as constantes do Go: https://go.dev/src/time/format.go
	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - Online: " + strconv.FormatBool(status) + "\n")
	arquivo.Close()
}

func imprimeLogs(arquivoTxt string) {
	arquivo, err := os.Open("monitoramento-sites/logs.txt")

	copiaArquivo := lerArquivo(arquivoTxt)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	leitor := bufio.NewReader(arquivo)

	for c := 0; c < len(copiaArquivo)-1; c++ {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)

		if c%10 == 0 {
			fmt.Println("")
		}

		fmt.Println(linha)

		if err == io.EOF {
			break
		}
	}

	arquivo.Close()
}
