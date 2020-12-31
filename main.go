package main

import (
	"fmt"
	"github.com/wrporter/monkey/evaluator"
	"github.com/wrporter/monkey/lexer"
	"github.com/wrporter/monkey/object"
	"github.com/wrporter/monkey/parser"
	"github.com/wrporter/monkey/repl"
	"io/ioutil"
	"os"
	"os/user"
)

func main() {
	if len(os.Args) > 1 {
		executeFile()
	} else {
		executeREPL()
	}
}

func executeFile() {
	filename := os.Args[1]
	code, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("Failed to execute file %s due to error: %v\n", filename, err)
		os.Exit(1)
	}

	l := lexer.New(string(code))
	p := parser.New(l)

	program := p.ParseProgram()
	if len(p.Errors()) != 0 {
		repl.PrintParserErrors(os.Stderr, p.Errors())
		os.Exit(1)
	}

	env := object.NewEnvironment()
	evaluated := evaluator.Eval(program, env)

	fmt.Println(evaluated.Inspect())
}

func executeREPL() {
	u, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the Monkey programming language!\n",
		u.Username)
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
