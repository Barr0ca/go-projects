# Go Projects

Repositório de projetos em Go. Cada pasta é um programa próprio, com o seu módulo e o seu README. Um projeto novo entra como outra pasta. Ele não precisa ter relação com os que já estão aqui.

O `go.work` só permite rodar qualquer projeto a partir da raiz. Nenhum módulo importa outro.

## Pré-requisitos

- [Go 1.26+](https://go.dev/dl/)

## Projetos

| Projeto | Descrição |
|---------|-----------|
| [arrays-slices](./arrays-slices/) | Menu que mostra um array fixo e um slice que cresce com `append` |
| [funcoes-e-menu](./funcoes-e-menu/) | Menu no terminal: funções, leitura com `fmt.Scan` e `switch` |
| [hello](./hello/) | Programa que imprime um nome e a versão |
| [monitoramento-sites](./monitoramento-sites/) | Monitora sites listados em um arquivo e grava cada resultado em `logs.txt` |
| [requisicoes-web](./requisicoes-web/) | Menu que monitora o status HTTP de vários sites com `http.Get` |
| [usando-arquivos](./usando-arquivos/) | Menu que lê as linhas de um arquivo de texto |

## Estrutura

```
go-projects/
├── README.md            # Este arquivo — índice dos projetos
├── go.work              # Workspace Go (liga os módulos para a raiz)
├── .gitignore
├── arrays-slices/
│   ├── go.mod
│   ├── main.go
│   └── README.md
├── funcoes-e-menu/
│   ├── go.mod
│   ├── main.go
│   └── README.md
├── hello/
│   ├── go.mod
│   ├── main.go
│   └── README.md
├── monitoramento-sites/
│   ├── go.mod
│   ├── logs.txt
│   ├── main.go
│   ├── README.md
│   └── sites.txt
├── requisicoes-web/
│   ├── go.mod
│   ├── main.go
│   └── README.md
└── usando-arquivos/
    ├── go.mod
    ├── main.go
    ├── nomes.txt
    └── README.md
```

## Como executar um projeto

```bash
git clone https://github.com/Barr0ca/go-projects.git
cd go-projects
```

Na raiz:

```bash
go run ./arrays-slices
go run ./funcoes-e-menu
go run ./hello
go run ./monitoramento-sites
go run ./requisicoes-web
go run ./usando-arquivos
```

Ou dentro da pasta do projeto:

```bash
cd hello
go run main.go
```

`arrays-slices`, `funcoes-e-menu`, `monitoramento-sites`, `requisicoes-web` e `usando-arquivos` leem um número do teclado. O README de cada um mostra um exemplo. `monitoramento-sites` e `usando-arquivos` abrem um arquivo pelo caminho a partir da raiz do repositório.

## Adicionar um projeto

Na raiz do repositório:

```bash
mkdir meu-projeto
cd meu-projeto
go mod init github.com/Barr0ca/go-projects/meu-projeto
cd ..
go work use ./meu-projeto
```

Esses comandos criam a pasta, o `go.mod` e o registro no `go.work`. Em seguida, escreva o `README.md` do projeto, com o que ele faz e como executar, e acrescente uma linha na tabela **Projetos** desta página.
