# Hello

Primeiro projeto do repositório — um "Hello, World!" em Go.

## Executar

```bash
cd hello
go run main.go
```

**Saída:**
```
Hello, World! Boa noite.
```

## Código

```go
package main          // Pacote main = programa executável

import "fmt"         // Pacote padrão para entrada/saída

func main() {        // Ponto de entrada do programa
    fmt.Println("Hello, World!")
}
```

| Elemento | Descrição |
|----------|-----------|
| `package main` | Obrigatório em todo executável Go |
| `import "fmt"` | Importa o pacote de formatação |
| `func main()` | Função executada ao iniciar o programa |
| `fmt.Println` | Imprime texto e pula linha |
