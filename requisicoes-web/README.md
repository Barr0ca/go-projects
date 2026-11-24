# Requisições web

Programa de terminal que repete um menu e monitora o status HTTP de três sites com `http.Get`.

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

O programa espera um número e volta ao menu até a opção `0`. A opção `1` consulta UOL, G1 e o portal do IFRN, três vezes, e espera 5 segundos depois de cada rodada. Exemplo:

```
Olá, Fulano
Este código fará requisições na internet

1 - Realizar requisição.
0 - Sair

1
Realizando requisições...

--------------------------------------------- 0 ---------------------------------------------

Verificando https://www.uol.com.br/
Site: https://www.uol.com.br/ foi carregado com sucesso! Status: 200

Verificando https://g1.globo.com/
Site: https://g1.globo.com/ foi carregado com sucesso! Status: 200

Verificando https://portal.ifrn.edu.br/
Site: https://portal.ifrn.edu.br/ foi carregado com sucesso! Status: 200

--------------------------------------------- 1 ---------------------------------------------
```

As rodadas `1` e `2` repetem os mesmos três sites. Um status de `200` a `299` imprime `foi carregado com sucesso!`. Os demais imprimem `está com problemas.`

A opção `0` imprime `Saindo...` e encerra com código de saída `0`. Qualquer outro valor imprime `Não entendi.` e mostra o menu de novo.

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

`MONITORAMENTOS` vale `3` e `DELAY` vale `5`. `reqInternet` percorre o slice de sites em cada rodada e pausa com `time.Sleep`:

```go
for i := 0; i < MONITORAMENTOS; i++ {
	for _, site := range sites {
		verificaSite(site)
	}
	time.Sleep(DELAY * time.Second)
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
| `reqInternet` | Repete o monitoramento `MONITORAMENTOS` vezes, com pausa de `DELAY` segundos |
| `verificaSite` | Faz o `GET` de um site e imprime o resultado a partir de `resp.StatusCode` |
