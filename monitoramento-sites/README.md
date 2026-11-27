# Monitoramento de sites

Projeto que acompanha a disponibilidade de uma lista de sites e guarda o resultado de cada consulta. A lista fica em [`sites.txt`](./sites.txt), um endereço por linha. O histórico fica em [`logs.txt`](./logs.txt).

Cada verificação faz um `GET` e classifica a resposta pelo status HTTP. Em seguida grava data, hora, endereço e se o site respondeu como online.

A lista reúne Google, Cloudflare, GitHub, League of Legends, Amazon, Microsoft, Apple, UOL, Globo e o portal do IFRN. Trocar um endereço em `sites.txt` muda o que a próxima verificação consulta.

## Como o monitoramento funciona

A opção `1` percorre a lista três vezes. `MONITORAMENTOS` define essa quantidade. As rodadas aparecem numeradas de `1` a `3`. Entre uma rodada e a seguinte o projeto espera 1 minuto: `DELAY` vale `1` e a pausa é `DELAY * time.Minute`. A verificação completa leva pelo menos 3 minutos.

Em cada site, `http.Get` busca a página. O resultado cai em um de três casos:

| Situação | O que aparece | Registro |
|----------|----------------|----------|
| Status de `200` a `299` | `foi carregado com sucesso!`, junto com o status | `Online: true` |
| Qualquer outro status | `está com problemas.`, junto com o status | `Online: false` |
| Falha na requisição | `Ocorreu um erro:`, seguido da causa | — |

Antes de cada consulta, a saída mostra `Verificando` e o endereço.

```
Realizando requisições...

--------------------------------------------- 1 ---------------------------------------------

Verificando https://www.google.com
Site: https://www.google.com foi carregado com sucesso! Status: 200
```

## Logs

`registraLogs` cria `logs.txt` se o arquivo ainda não existe e acrescenta uma linha no fim. O texto usa o layout `dia/mês/ano hora:minuto:segundo`:

```
27/11/2026 18:16:54 - https://www.github.com - Online: true
```

A opção `2` lê esse arquivo e imprime o histórico. A cada 10 registros, uma linha em branco separa o bloco seguinte.

A opção `0` encerra. Outro valor informa que a escolha não foi entendida e as opções voltam.

## Executar

Na raiz do repositório:

```bash
go run ./monitoramento-sites
```

A leitura de sites abre `monitoramento-sites/sites.txt` e a dos logs abre `monitoramento-sites/logs.txt`. Esses caminhos são resolvidos a partir da pasta em que o comando foi iniciado.

## Código

A lógica está em [`main.go`](./main.go).

`reqInternet` carrega os endereços e repete o ciclo de verificação e pausa:

```go
sites := lerArquivo(arquivoTxt)

for i := 0; i < MONITORAMENTOS; i++ {
	for _, site := range sites {
		verificaSite(site)
	}
	time.Sleep(DELAY * time.Minute)
}
```

`verificaSite` faz a requisição. No sucesso e no status fora da faixa, chama `registraLogs`:

```go
resp, err := http.Get(site)

if resp.StatusCode > 199 && resp.StatusCode < 300 {
	registraLogs(site, true)
} else {
	registraLogs(site, false)
}
```

O registro abre o arquivo com `os.OpenFile` e as flags `O_CREATE`, `O_RDWR` e `O_APPEND`:

```go
arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - Online: " + strconv.FormatBool(status) + "\n")
```

`lerArquivo` lê um arquivo linha a linha com `bufio.NewReader` até `io.EOF`. `strings.TrimSpace` remove espaços e a quebra de linha. A mesma função serve para a lista de sites e para contar o que `imprimeLogs` vai mostrar.

| Função | Papel |
|--------|--------|
| `reqInternet` | Dispara as `MONITORAMENTOS` rodadas, com a pausa de `DELAY` minuto |
| `verificaSite` | Consulta um site, classifica o status e pede o registro |
| `registraLogs` | Acrescenta data, hora, endereço e `Online: true` ou `Online: false` em `logs.txt` |
| `imprimeLogs` | Mostra o histórico, com uma linha em branco a cada 10 registros |
| `lerArquivo` | Devolve as linhas de um arquivo em um `[]string` |
