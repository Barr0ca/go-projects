# Usando arquivos

Programa de terminal que manipula arquivos, abrindo e lendo linhas de um .txt.

## Executar

Na raiz do repositório:

```bash
go run ./usando-arquivos
```

`lerArquivo` abre `usando-arquivos/nomes.txt`. Esse caminho parte da pasta em que o comando foi iniciado, então o comando acima precisa rodar na raiz do repositório.

O programa espera um número. Exemplo com a opção `1`:

```
Olá, Fulano
Este código manipula arquivos.

1 - Ler arquivo de nomes.
0 - Sair.

1
Abrindo arquivo...

```

A opção `0` imprime `Saindo...` e encerra com código de saída `0`. Qualquer outro valor imprime `Não entendi.` e encerra com código `1`.

O arquivo [`nomes.txt`](./nomes.txt) contém `Simão`, `Carol`, `João` e `Mateus`, um nome por linha.

## Código

O fluxo está em [`main.go`](./main.go). `main` chama as funções e faz o `switch`:

```go
switch opcao {
case 1:
	fmt.Println("Abrindo arquivo...\n")
	lerArquivo()
// ...
}
```

`os.Open` abre o arquivo. `bufio.NewReader` lê uma linha por vez até `io.EOF`. `strings.TrimSpace` tira o espaço e a quebra de linha de cada item. As linhas vão para um slice, que a função devolve:

```go
linha, err := leitor.ReadString('\n')
linha = strings.TrimSpace(linha)
nomes = append(nomes, linha)

if err == io.EOF {
	break
}
```

`main` chama `lerArquivo` e não imprime o slice devolvido. Se `os.Open` falha, o programa imprime `Ocorreu um erro:`.

| Função | Papel |
|--------|--------|
| `apresentacao` | Imprime o nome e o propósito do programa |
| `exibeMenu` | Lista as opções `1` e `0` |
| `escolheOpcao` | Lê um `int` com `fmt.Scan` e devolve esse valor |
| `lerArquivo` | Lê `nomes.txt` e devolve um `[]string` com as linhas |
