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
| [requisicoes-web](./requisicoes-web/) | Menu que consulta o status HTTP de um site com `http.Get` |

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
├── hello/
│   ├── go.mod
│   ├── main.go
│   └── README.md
└── requisicoes-web/
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
go run ./requisicoes-web
```

Ou dentro da pasta do projeto:

```bash
cd hello
go run main.go
```

`funcoes-e-menu` e `requisicoes-web` leem um número do teclado. O README de cada um mostra um exemplo.

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
