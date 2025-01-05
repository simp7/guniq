package unique

import (
	"bufio"
	"io"
)

type standard struct {
	prev string
}

func Standard() *standard {
	u := new(standard)
	return u
}

func (s *standard) Execute(input io.Reader, output io.Writer) {
	reader := bufio.NewReader(input)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			return
		}
		if s.prev != line {
			output.Write([]byte(line))
			s.prev = line
		}
	}
}
