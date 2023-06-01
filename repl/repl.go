package repl

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/sandeshsitaula/monkeyinter/evaluator"
	"github.com/sandeshsitaula/monkeyinter/lexer"
	"github.com/sandeshsitaula/monkeyinter/object"
	"github.com/sandeshsitaula/monkeyinter/parser"
)

const PROMPT = ">> "

var k int = len(os.Args)

var env *object.Environment = object.NewEnvironment()

func Start(in io.Reader, out io.Writer) {
	//	scanner := bufio.NewScanner(in)

	if k > 1 {
		var data []byte
		data, _ = ioutil.ReadFile(os.Args[1])
		runner(string(data), out)

	} else {
		for {
			scanner := bufio.NewScanner(in)
			fmt.Printf(PROMPT)

			scanned := scanner.Scan()

			if !scanned {
				return
			}
			line := scanner.Text()
			runner(line, out)
		}
	}
}

func runner(data string, out io.Writer) {

	l := lexer.New(data)
	p := parser.New(l)
	program := p.ParseProgram()
	if len(p.Errors()) != 0 {
		printParserErrors(out, p.Errors())
	}

	evaluated := evaluator.Eval(program, env)

	if evaluated != nil {
		io.WriteString(out, evaluated.Inspect())
		io.WriteString(out, "\n")
	}

}

func printParserErrors(out io.Writer, errors []string) {
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
