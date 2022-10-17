package main

import (
	"flag"
	"log"
	"os"

	"github.com/acbrown/plug-lang/ast"
	"github.com/acbrown/plug-lang/lexer/lexer"
	"github.com/acbrown/plug-lang/parser"
)

var (
	src = flag.String("src", "", "source file to compile")
)

func compile(data string) {
	l := lexer.NewLexer(data)
	p := parser.NewParser(l)
	log.Print(ast.ParseProgram(p))
}

func main() {
	flag.Parse()

	if *src == "" {
		log.Fatal("No src file specified")
	}

	data, err := os.ReadFile(*src)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	compile(string(data))
}
