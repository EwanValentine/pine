package main

import (
	"os"

	"github.com/ewanvalentine/pine/repl"
)

func main() {
	repl.Start(os.Stdin, os.Stdout)
}
