# Funções e menu

Programa de terminal que se apresenta, mostra um menu, lê um número e escolhe uma mensagem com `switch`.

## Executar

Na pasta do projeto:

```bash
cd funcoes-e-menu
go run main.go
```

Na raiz do repositório:

```bash
go run ./funcoes-e-menu
```

O programa espera um número. Exemplo com a opção `1`:

```
Olá, Fulano
Este código está na versão 1.1

1 - Bom Dia!
2 - Boa Tarde!
3 - Boa Noite!
0 - Sair.
1
O dia está ótimo.
```

A opção `2` imprime `Café da tarde é muito bom.` A opção `3` imprime `A lua está muito bonita.` As opções `1`, `2` e `3` encerram o programa em seguida. A opção `0` imprime `Saindo...` e encerra com código de saída `0`. Qualquer outro valor imprime `Não entendi.` e encerra com código `1`.

## Código

O fluxo está em [`main.go`](./main.go). `main` chama as funções e faz o `switch`:

```go
func main() {
	apresentacao()
	exibeMenu()

	opcao := escolheOpcao()

	switch opcao {
	case 1:
		fmt.Println("O dia está ótimo.")
	// ...
	}
}
```

| Função | Papel |
|--------|--------|
| `apresentacao` | Imprime nome e versão. `versao` é um `float32` |
| `exibeMenu` | Lista as opções `1`, `2`, `3` e `0` |
| `escolheOpcao` | Lê um `int` com `fmt.Scan` e devolve esse valor. O `int` depois do nome da função é o tipo do retorno |
