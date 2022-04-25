package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/Rindrics/monkey/lexer"
	"github.com/Rindrics/monkey/token"
)

func Start(in io.Reader) {
	scanner := bufio.NewScanner(in)

	for {
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()
		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
