# Go Projects

Repositório de projetos em Go. Cada pasta é um programa próprio, com o seu módulo e o seu README. Um projeto novo entra como outra pasta. Ele não precisa ter relação com os que já estão aqui.

O `go.work` só permite rodar qualquer projeto a partir da raiz. Nenhum módulo importa outro.

## Pré-requisitos

- [Go 1.26+](https://go.dev/dl/)

## Projetos

| Projeto | Descrição |
|---------|-----------|
| [funcoes-e-menu](./funcoes-e-menu/) | Menu no terminal: funções, leitura com `fmt.Scan` e `switch` |
| [hello](./hello/) | Programa que imprime `Hello, World!` |

## Estrutura

```
go-projects/
├── README.md            # Este arquivo — índice dos projetos
├── go.work              # Workspace Go (liga os módulos para a raiz)
├── .gitignore
├── funcoes-e-menu/
│   ├── go.mod
│   ├── main.go
│   └── README.md
└── hello/
    ├── go.mod
    ├── main.go
    └── README.md
```

## Como executar um projeto

```bash
git clone https://github.com/Barr0ca/go-projects.git
cd go-projects
```

Na raiz:

```bash
go run ./hello
go run ./funcoes-e-menu
```

Ou dentro da pasta do projeto:

```bash
cd hello
go run main.go
```

`funcoes-e-menu` lê um número do teclado. O README desse projeto mostra um exemplo.

## Adicionar um projeto

Na raiz do repositório:

```bash
mkdir meu-projeto
cd meu-projeto
go mod init github.com/Barr0ca/go-projects/meu-projeto
cd ..
go work use ./meu-projeto
```

O projeto ganha um `README.md` com o que ele faz e como executar. A tabela desta página ganha uma linha para ele.
