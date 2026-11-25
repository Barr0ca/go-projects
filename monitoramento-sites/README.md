# Monitoramento de sites

Projeto que acompanha a disponibilidade de uma lista de sites. A lista fica em [`sites.txt`](./sites.txt), um endereço por linha. Cada verificação faz um `GET` e classifica a resposta pelo status HTTP.

Hoje a lista reúne Google, Cloudflare, GitHub, Wikipedia, Amazon, Microsoft, Apple, UOL, Globo e o portal do IFRN. Trocar um endereço no arquivo muda o que a próxima verificação consulta, sem alterar o código.

## Como o monitoramento funciona

Uma verificação percorre a lista três vezes. `MONITORAMENTOS` define essa quantidade. Entre uma rodada e a seguinte o projeto espera 1 minuto. `DELAY` vale `1` e a pausa é `DELAY * time.Minute`.

Em cada site, `http.Get` busca a página. O resultado cai em um de três casos:

| Situação | O que aparece |
|----------|----------------|
| Status de `200` a `299` | `foi carregado com sucesso!`, junto com o status |
| Qualquer outro status | `está com problemas.`, junto com o status |
| Falha na requisição | `Ocorreu um erro:`, seguido da causa |

A rodada começa com um separador numerado (`0`, `1` e `2`) e, antes de cada consulta, imprime `Verificando` e o endereço.

```
Realizando requisições...

--------------------------------------------- 0 ---------------------------------------------

Verificando https://www.google.com
Site: https://www.google.com foi carregado com sucesso! Status: 200
```

As rodadas `1` e `2` repetem a mesma lista. Como cada pausa dura 1 minuto, a verificação completa leva pelo menos 3 minutos.

Para iniciar, escolha `1`. A opção `0` encerra. Outro valor informa que a escolha não foi entendida e a lista de opções volta.

## Executar

Na raiz do repositório:

```bash
go run ./monitoramento-sites
```

A leitura abre `monitoramento-sites/sites.txt`. Esse caminho é resolvido a partir da pasta em que o comando foi iniciado.

## Código

A lógica está em [`main.go`](./main.go).

`reqInternet` carrega os endereços e repete o ciclo de verificação e pausa:

```go
sites := lerArquivo()

for i := 0; i < MONITORAMENTOS; i++ {
	for _, site := range sites {
		verificaSite(site)
	}
	time.Sleep(DELAY * time.Minute)
}
```

`verificaSite` faz a requisição e separa falha de rede e status HTTP:

```go
resp, err := http.Get(site)

if err != nil {
	fmt.Println("Ocorreu um erro:", err)
}
```

`lerArquivo` lê o arquivo linha a linha com `bufio.NewReader` até `io.EOF`. `strings.TrimSpace` remove espaços e a quebra de linha antes de guardar o endereço no slice.

| Função | Papel |
|--------|--------|
| `reqInternet` | Dispara as `MONITORAMENTOS` rodadas, com a pausa de `DELAY` minuto |
| `verificaSite` | Consulta um site e registra sucesso, problema ou erro |
| `lerArquivo` | Monta o slice de endereços a partir de `sites.txt` |
