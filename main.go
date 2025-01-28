package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/simp7/guniq/unique"
)

type Unique interface {
	Execute(input io.Reader, output io.Writer)
}

func main() {
	var reader io.Reader = os.Stdin
	var writer io.Writer = os.Stdout
	var err error
	var u Unique = unique.Standard()

	counting := flag.Bool("c", false, "add count of repeated lines")
	repeated := flag.Bool("d", false, "print only repeated lines")
	flag.Parse()

	if *counting {
		u = unique.Counting()
	}

	if *repeated {
		u = unique.Repeated()
	}

	args := flag.Args()

	switch len(args) {
	case 0:
	case 1:
		if args[0] == "-" {
			break
		}
		reader, err = os.Open(args[0])
		if err != nil {
			fmt.Println("Error: ", err)
			os.Exit(0)
		}
	default:
		readers := make([]io.Reader, 0)
		for i := 0; i < len(args)-1; i++ {
			readers[i], err = os.Open(args[i])
			if err != nil {
				fmt.Println("Error: ", err)
				os.Exit(1)
			}
		}
		reader = io.MultiReader(readers...)
		writer, err = os.Open(args[len(args)-1])
		if errors.Is(err, os.ErrNotExist) {
			err = nil
			writer, err = os.Create(args[len(args)-1])
		}

		if err != nil {
			fmt.Println("Error: ", err)
			os.Exit(1)
		}
	}

	u.Execute(reader, writer)
	os.Exit(0)
}
