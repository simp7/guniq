package unique

import (
	"bufio"
	"fmt"
	"io"
)

type repeated struct {
	prev       string
	isCounting bool
}

func Repeated(isCounting bool) *repeated {
	r := new(repeated)
	r.isCounting = isCounting
	return r
}

func (r *repeated) print(output io.Writer, count int) {
	formatted := r.prev + "\n"
	if r.isCounting {
		formatted = fmt.Sprintf("%4d %s", count, formatted)
	}
	output.Write([]byte(formatted))
}

func (r *repeated) Execute(input io.Reader, output io.Writer) {
	scanner := bufio.NewScanner(input)
	scanner.Scan()
	line := scanner.Text()

	r.prev = line
	count := 1

	for scanner.Scan() {
		line := scanner.Text()
		if r.prev != line {
			if count > 1 {
				r.print(output, count)
			}
			r.prev = line
			count = 1
			continue
		}
		count++
	}

	if count > 1 {
		r.print(output, count)
	}
}
