# Hello

Programa que imprime uma linha e encerra.

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
Hello, World!
```

## Código

O programa está em [`main.go`](./main.go):

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}
```

| Elemento | Descrição |
|----------|-----------|
| `package main` | Pacote de um programa executável |
| `import "fmt"` | Pacote padrão de formatação e impressão |
| `func main()` | Função chamada ao iniciar o programa |
| `fmt.Println` | Imprime o texto e pula uma linha |
