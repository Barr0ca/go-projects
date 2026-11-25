# Arrays e slices

Programa de terminal que repete um menu, utilizando "for",  mostra um array de tamanho fixo e um slice que cresce com `append`.

## Executar

Na pasta do projeto:

```bash
cd arrays-slices
go run main.go
```

Na raiz do repositório:

```bash
go run ./arrays-slices
```

O programa espera um número e volta ao menu até a opção `0`. Exemplo com a opção `1`:

```
Olá, Fulano
Este código mostra arrays e slices

1 - Mostrar array.
2 - Mostrar slice.
0 - Sair

1
Mostrando array...
0 - Bruno
1 - Carol
2 - João
3 - Simão

1 - Mostrar array.
2 - Mostrar slice.
0 - Sair
```

A opção `2` imprime o slice, acrescenta `Acari` e mostra o tamanho e a capacidade antes e depois:

```
2
Mostrando slice...
Cidades: [Natal Caicó Currais Novos] um total de 3 cidades.
A capacidade desse slice: 3

Adicionando Acari em cidades...

Cidades: [Natal Caicó Currais Novos Acari] um total de 4 cidades.
A capacidade desse slice: 6
```

A opção `0` imprime `Saindo...` e encerra com código de saída `0`. Qualquer outro valor imprime `Não entendi.` e mostra o menu de novo.

## Código

O fluxo está em [`main.go`](./main.go). O `for` sem condição repete o menu:

```go
for {
	exibeMenu()

	opcao := escolheOpcao()

	switch opcao {
	case 1:
		mostraArray()
	case 2:
		mostraSlice()
	// ...
	}
}
```

O array tem tamanho fixo. O `range` percorre o índice e o valor:

```go
var nomes [4]string

for i, nome := range nomes {
	fmt.Println(i, "-", nome)
}
```

O slice começa com três cidades. `append` devolve um slice novo. Quando o item extra não cabe, a capacidade dobra de `3` para `6`:

```go
cidades := []string{"Natal", "Caicó", "Currais Novos"}
cidades = append(cidades, "Acari")
```

| Função | Papel |
|--------|--------|
| `apresentacao` | Imprime o nome e o propósito do programa |
| `exibeMenu` | Lista as opções `1`, `2` e `0` |
| `escolheOpcao` | Lê um `int` com `fmt.Scan` e devolve esse valor |
| `mostraArray` | Preenche um `[4]string` e imprime índice e nome com `range` |
| `mostraSlice` | Imprime `len` e `cap`, acrescenta um item com `append` e imprime de novo |
