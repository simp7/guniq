package unique

import (
	"bufio"
	"io"
)

type repeated struct {
	prev string
}

func Repeated() *repeated {
	r := new(repeated)
	return r
}

func (r *repeated) Execute(input io.Reader, output io.Writer) {
	printed := false
	reader := bufio.NewReader(input)
	line, err := reader.ReadString('\n')
	if err != nil {
		return
	}
	r.prev = line

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			return
		}
		if r.prev != line {
			r.prev = line
			printed = false
			continue
		}
		if !printed {
			output.Write([]byte(line))
			printed = true
		}
	}
}
