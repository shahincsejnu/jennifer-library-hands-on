package main

import (
	"fmt"
	. "github.com/dave/jennifer/jen"
	"log"
)

func main () {
	f := NewFile("main")

	f.Func().Id("main").Params().Block(
		Qual("fmt", "Println").Call(Lit("Hello, World")),
	)

	// for testing, a file or statement can be rendered with the fmt package using the %#v verb
	// but this only for testing, not recommended for use in production
	// cause any error will cause a panic
	fmt.Printf("%#v", f)

	// saving the generated code in given file
	err := f.Save("generated-by-first-test.go")
	if err != nil {
		log.Fatal(err)
		return
	}
}