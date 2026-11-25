package main

import (
	"bufio" // Bib que manipula arquivos
	"fmt"
	"io" // Bib de suporte para entrada e saida de dados
	"os"
	"strings"
)

func main() {
	apresentacao()
	exibeMenu()

	opcao := escolheOpcao()

	switch opcao {
	case 1:
		fmt.Println("Abrindo arquivo...\n")
		lerArquivo()
	case 0:
		fmt.Println("Saindo...")
		os.Exit(0)
	default:
		fmt.Println("Não entendi.\n")
		os.Exit(1)
	}
}

func apresentacao() {
	nome := "Fulano"
	fmt.Println("Olá,", nome)
	fmt.Println("Este código manipula arquivos.\n")
}

func exibeMenu() {
	fmt.Println("1 - Ler arquivo de nomes.")
	fmt.Println("0 - Sair.")
	fmt.Println("")
}

func escolheOpcao() int {
	var opcao int
	fmt.Scan(&opcao)

	return opcao
}

func lerArquivo() []string {
	var nomes []string

	arquivo, err := os.Open("usando-arquivos/nomes.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	leitor := bufio.NewReader(arquivo)
	for {
		// Aspas simples informa que e o valor em bytes de "\n",
		// alem disso, ReadString recebe como parametro o valor
		// em que ele deve parar de ler, nesse caso "\n" (quebra de linha)
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)      // Retira os espacos entre as linhas

		nomes = append(nomes, linha)

		if err == io.EOF { // EOF significa que chegou ao fim do arquivo, assim saimos do "for"
			break
		}
	}

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	arquivo.Close()

	return nomes
}
