package unique

import (
	"bufio"
	"io"
)

type singular struct {
	prev       string
	isCounting bool
}

func Singular(isCounting bool) *singular {
	r := new(singular)
	r.isCounting = isCounting
	return r
}

func (s *singular) print(output io.Writer) {
	formatted := s.prev + "\n"
	if s.isCounting {
		formatted = "   1 " + formatted
	}
	output.Write([]byte(formatted))
}

func (s *singular) Execute(input io.Reader, output io.Writer) {
	isDuplicate := false
	scanner := bufio.NewScanner(input)
	scanner.Scan()
	s.prev = scanner.Text()

	for scanner.Scan() {
		line := scanner.Text()
		if s.prev != line {
			if !isDuplicate {
				s.print(output)
			}
			s.prev = line
			isDuplicate = false
		} else {
			isDuplicate = true
		}
		s.prev = line
	}

	if !isDuplicate {
		s.print(output)
	}
}
