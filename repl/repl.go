package repl

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/sandeshsitaula/monkeyinter/evaluator"
	"github.com/sandeshsitaula/monkeyinter/lexer"
	"github.com/sandeshsitaula/monkeyinter/object"
	"github.com/sandeshsitaula/monkeyinter/parser"
)

const PROMPT = ">> "

var k int = len(os.Args)

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()
	for {
		if k <= 1 {
			fmt.Println()
			fmt.Printf(PROMPT)
		}

		scanned := scanner.Scan()

		if !scanned {
			return
		}
		line := scanner.Text()

		l := lexer.New(line)
		p := parser.New(l)
		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		evaluated := evaluator.Eval(program, env)

		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}
	}
}

func printParserErrors(out io.Writer, errors []string) {
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
