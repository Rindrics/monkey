package main

import (
	"os"

	"github.com/Rindrics/monkey/repl"
)

func main() {
	repl.Start(os.Stdin)
}
