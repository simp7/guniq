package unique

import (
	"bufio"
	"io"
)

type singular struct {
	prev string
}

func Singular() *singular {
	r := new(singular)
	return r
}

func (s *singular) Execute(input io.Reader, output io.Writer) {
	isDuplicate := false
	reader := bufio.NewReader(input)
	line, err := reader.ReadString('\n')
	if err != nil {
		return
	}

	s.prev = line
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			output.Write([]byte(s.prev))
			return
		}

		if s.prev != line {
			if !isDuplicate {
				output.Write([]byte(s.prev))
			}
			s.prev = line
			isDuplicate = false
			continue
		}
		isDuplicate = true
	}
}
