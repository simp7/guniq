package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
)

func main() {
	var reader io.Reader = os.Stdin
	var writer io.Writer = os.Stdout
	var err error

	args := os.Args

	switch len(args) {
	case 1:
	case 2:
		if args[1] == "-" {
			break
		}
		reader, err = os.Open(args[1])
		if err != nil {
			fmt.Println("Error: ", err)
			os.Exit(1)
		}
	default:
		readers := make([]io.Reader, 0)
		for i := 1; i < len(args)-2; i++ {
			readers[i-1], err = os.Open(args[i])
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

	u := NewUnique()
	u.Execute(reader, writer)
	os.Exit(0)
}

type unique struct {
	prev string
}

func NewUnique() *unique {
	u := new(unique)
	return u
}

func (u *unique) Execute(input io.Reader, output io.Writer) {
	reader := bufio.NewReader(input)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			return
		}
		if u.prev != line {
			output.Write([]byte(line))
			u.prev = line
		}
	}
}
