package main

import (
	"os"

	"github.com/slonegd-otus-go/10_envdir/internal"
)

func main() {
	internal.Execute(os.Stdin, os.Stdout, os.Stderr, os.Args)
}
