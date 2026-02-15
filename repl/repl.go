// Package repl provides an interactive REPL environment
package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/suasuasuasuasua/building-an-interpreter-go/lexer"
	"github.com/suasuasuasuasua/building-an-interpreter-go/token"
)

const PROMPT = ">> "

// scan and lex an input stream
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprintf(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		// createa a new lexer with the current line
		line := scanner.Text()
		l := lexer.New(line)

		// loop over the tokens until we encounter an EOF
		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Fprintf(out, "%+v\n", tok)
		}
	}
}
