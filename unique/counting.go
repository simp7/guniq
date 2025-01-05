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
	reader := bufio.NewReader(input)
	count := 0
	line, err := reader.ReadString('\n')
	if err != nil {
		return
	}
	c.prev = line

	for {
		count++

		line, err = reader.ReadString('\n')
		if err != nil {
			formatted := fmt.Sprintf("%d %s", count, c.prev)
			output.Write([]byte(formatted))
			return
		}

		if c.prev == line {
			continue
		}

		formatted := fmt.Sprintf("%d %s", count, c.prev)
		output.Write([]byte(formatted))
		c.prev = line
		count = 0
	}
}
