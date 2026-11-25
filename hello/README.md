# Hello

Programa que imprime um nome e a versão e encerra.

## Executar

Na pasta do projeto:

```bash
cd hello
go run main.go
```

Na raiz do repositório:

```bash
go run ./hello
```

**Saída:**

```
Olá, Fulano
Este código está na versão 1.1

```

## Código

O programa está em [`main.go`](./main.go):

```go
func main() {
	nome := "Fulano"
	var versao float32 = 1.1
	fmt.Println("Olá,", nome)
	fmt.Println("Este código está na versão", versao, "\n")
}
```

| Elemento | Descrição |
|----------|-----------|
| `package main` | Pacote de um programa executável |
| `import "fmt"` | Pacote padrão de formatação e impressão |
| `func main()` | Função chamada ao iniciar o programa |
| `nome` | Texto impresso na saudação |
| `versao` | `float32` impresso na segunda linha |
| `fmt.Println` | Imprime os valores e pula uma linha |
