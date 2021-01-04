package repl

import (
	"github.com/c-bata/go-prompt"
	"github.com/wrporter/monkey/evaluator"
	"github.com/wrporter/monkey/lexer"
	"github.com/wrporter/monkey/object"
	"github.com/wrporter/monkey/parser"
	"io"
	"os"
)

const PROMPT = ">> "

func completer(in prompt.Document) []prompt.Suggest {
	return prompt.FilterHasPrefix(nil, in.GetWordBeforeCursor(), true)
}

func Start(in io.Reader, out io.Writer) {
	env := object.NewEnvironment()

	p := prompt.New(
		func(input string) {
			if input == "exit" {
				os.Exit(0)
			}

			l := lexer.New(input)
			p := parser.New(l)

			program := p.ParseProgram()
			if len(p.Errors()) != 0 {
				PrintParserErrors(out, p.Errors())
				return
			}

			evaluated := evaluator.Eval(program, env)
			if evaluated != nil {
				io.WriteString(out, evaluated.Inspect())
				io.WriteString(out, "\n")
			}
		},
		completer,
		prompt.OptionPrefix(PROMPT),
	)

	p.Run()
}

func PrintParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, "Parser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
