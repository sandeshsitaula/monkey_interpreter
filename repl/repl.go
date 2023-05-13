package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/sandeshsitaula/monkeyinter/lexer"
	"github.com/sandeshsitaula/monkeyinter/token"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()
		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v", tok)
		}
	}
}
