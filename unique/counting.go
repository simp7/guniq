package unique

import (
	"bufio"
	"fmt"
	"io"
)

type counting struct {
	prev string
}

func Counting() *counting {
	c := new(counting)
	return c
}

func (c *counting) Execute(input io.Reader, output io.Writer) {
	scanner := bufio.NewScanner(input)
	scanner.Scan()

	c.prev = scanner.Text()
	count := 1

	for scanner.Scan() {
		line := scanner.Text()
		if c.prev == line {
			count++
			continue
		}

		formatted := fmt.Sprintf("%4d %s\n", count, c.prev)
		output.Write([]byte(formatted))
		c.prev = line
		count = 1
	}

	formatted := fmt.Sprintf("%4d %s\n", count, c.prev)
	output.Write([]byte(formatted))
}
