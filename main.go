package main

import (
	"fmt"
	"monkey-interpreter-go/repl"
	"os"
)

func main() {
	fmt.Printf("monke\n")
	repl.Start(os.Stdin, os.Stdout)
}
