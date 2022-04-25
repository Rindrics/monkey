package repl

import (
	"bufio"
	"fmt"
	"io"
)

func Start(in io.Reader) {
	scanner := bufio.NewScanner(in)
	scanned := scanner.Scan()
	if !scanned {
		return
	}
	fmt.Println("from repl:", scanner.Text())
}
