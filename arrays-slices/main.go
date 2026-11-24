package main

import (
	"fmt" // Bib para formatacao
	"os"  // Bib para comandos de sistema
)

func main() {
	apresentacao()

	// Nao existe "While" em Go, "for" sem parametros se torna um "While True"
	for {
		exibeMenu()

		opcao := escolheOpcao()

		switch opcao {
		case 1:
			mostraArray()
		case 2:
			mostraSlice()
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
	fmt.Println("Este código mostra arrays e slices")
}

func exibeMenu() {
	fmt.Println("\n1 - Mostrar array.")
	fmt.Println("2 - Mostrar slice.")
	fmt.Println("0 - Sair\n")
}

func escolheOpcao() int {
	var opcao int
	fmt.Scan(&opcao)

	return opcao
}

func mostraArray() {
	fmt.Println("Mostrando array...")
	var nomes [4]string // Em Go o array tem tamanho fixo, nesse caso 4

	nomes[0] = "Bruno"
	nomes[1] = "Carol"
	nomes[2] = "João"
	nomes[3] = "Simão"

	for i, nome := range nomes { // For clássico em Go, permite acessar o index "i" e o valor "nome" utilizando "range"
		fmt.Println(i, "-", nome)
	}
}

func mostraSlice() {
	fmt.Println("Mostrando slice...")
	cidades := []string{"Natal", "Caicó", "Currais Novos"} // Isso e um slice em Go (abstracao de um array)

	fmt.Println("Cidades:", cidades, "um total de", len(cidades), "cidades. \nA capacidade desse slice:", cap(cidades))

	fmt.Println("\nAdicionando Acari em cidades...\n")
	cidades = append(cidades, "Acari")

	fmt.Println("Cidades:", cidades, "um total de", len(cidades), "cidades. \nA capacidade desse slice:", cap(cidades))
}

// A capacidade de um slice é igual á quantidade de itens, no momento de sua criação,
// mas no momento em que a capacidade do slice estoura, o Go duplica a capacidade do
// slice para armazenar mais itens. Nesse caso, inicialmente a quantidade de itens era
// de 3 e a capacidade também, ao adicionar mais um item ao slice a capacidade duplicou,
// assim podendo armazenar 6 itens.
