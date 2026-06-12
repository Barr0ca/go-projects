# Go Projects

Repositório pessoal de projetos em Go. Cada pasta é um projeto independente com seu próprio módulo.

## Pré-requisitos

- [Go 1.26+](https://go.dev/dl/)

## Projetos

| Projeto | Descrição | Status |
|---------|-----------|--------|
| [hello](./hello/) | Hello World — primeiro programa | ✅ Concluído |

## Estrutura

```
go-projects/
├── README.md          # Este arquivo — índice dos projetos
├── go.work            # Workspace Go (liga todos os módulos)
├── .gitignore
└── hello/             
    ├── go.mod
    ├── main.go
    └── README.md
```

## Como executar um projeto

```bash
# Clonar o repositório
git clone https://github.com/Barr0ca/go-projects.git
cd go-projects

# Executar um projeto específico
cd hello
go run main.go
```

Ou, na raiz do repositório (com `go.work`):

```bash
go run ./hello
```
