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
	scanner := bufio.NewScanner(input)
	scanner.Scan()
	line := scanner.Text()
	output.Write([]byte(line + "\n"))
	s.prev = line

	for scanner.Scan() {
		line = scanner.Text()
		if s.prev != line {
			output.Write([]byte(line + "\n"))
			s.prev = line
		}
	}
}
