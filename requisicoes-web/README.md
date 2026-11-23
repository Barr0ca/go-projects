# Requisições web

Programa de terminal que repete um menu e consulta o status HTTP de um site com `http.Get`.

## Executar

Na pasta do projeto:

```bash
cd requisicoes-web
go run main.go
```

Na raiz do repositório:

```bash
go run ./requisicoes-web
```

O programa espera um número e volta ao menu até a opção `0`. Exemplo com a opção `1`:

```
Olá, Fulano
Este código fará requisições na internet

1 - Realizar requisição.
0 - Sair

1
Realizando requisição...
Site: https://httpbin.org/status/404 está com problemas. Status: 404
1 - Realizar requisição.
0 - Sair

0
Saindo...
```

A opção `0` imprime `Saindo...` e encerra com código de saída `0`. Qualquer outro valor imprime `Não entendi.` e mostra o menu de novo.

O endereço usado é `https://httpbin.org/status/404`. Um status de `200` a `299` imprime `foi carregado com sucesso!`. Os demais imprimem `está com problemas.`

## Código

O fluxo está em [`main.go`](./main.go). O `for` sem condição repete o menu:

```go
for {
	exibeMenu()

	opcao := escolheOpcao()

	switch opcao {
	case 1:
		reqInternet()
	// ...
	}
}
```

`http.Get` devolve a resposta e um `error`. O `_` descarta o `error` e fica só a resposta:

```go
resp, _ := http.Get(site)
```

| Função | Papel |
|--------|--------|
| `apresentacao` | Imprime o nome e o propósito do programa |
| `exibeMenu` | Lista as opções `1` e `0` |
| `escolheOpcao` | Lê um `int` com `fmt.Scan` e devolve esse valor |
| `reqInternet` | Faz o `GET` e imprime o resultado a partir de `resp.StatusCode` |
