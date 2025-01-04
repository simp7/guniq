package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("not implemented yet")
		os.Exit(1)
	}

	file, err := os.Open(args[1])
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	defer file.Close()

	u := NewUnique()
	result := u.Execute(file)
	for _, line := range result {
		fmt.Print(line)
	}
	os.Exit(0)
}

type unique struct {
	prev string
}

func NewUnique() *unique {
	u := new(unique)
	return u
}

func (u *unique) Execute(file *os.File) []string {
	result := make([]string, 0)
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			return result
		}
		if u.prev != line {
			result = append(result, line)
			u.prev = line
		}
	}
}
