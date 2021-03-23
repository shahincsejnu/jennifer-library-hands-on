# jennifer-library-hands-on

- Followed [jennifer repo](https://github.com/dave/jennifer)

## Intro

* Jennifer is a code generator for Go

* Example:
    - for this code:
        ```go
        package main

        import (
            "fmt"
        
            . "github.com/dave/jennifer/jen"
        )
        
        func main() {
            f := NewFile("main")
            f.Func().Id("main").Params().Block(
                Qual("fmt", "Println").Call(Lit("Hello, world")),
            )
            fmt.Printf("%#v", f)
        }
        ```
    - Output:
        ```go
        package main

        import "fmt"
        
        func main() {
            fmt.Println("Hello, world")
        }
        ```