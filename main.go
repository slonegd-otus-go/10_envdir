package main

import (
	"log"
	"os"

	"github.com/slonegd-otus-go/10_envdir/internal"
)

func main() {
	err := internal.Execute(os.Stdin, os.Stdout, os.Stderr, os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
}
