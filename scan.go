package scan

import (
	"bufio"
	"fmt"
	"io"
)

// LineScanner parses a line (string) and returns a copy of the result along with an ok/no good signal
type LineScanner interface {
	parse(string) (LineScanner, bool)
}

func ReadAll(r io.Reader, scanners ...LineScanner) ([]LineScanner, error) {
	s := bufio.NewScanner(r)
	ret := make([]LineScanner, 0)
	line := 0
	for s.Scan() {
		line++
		for _, sr := range scanners {
			txt := s.Text()
			obj, ok := sr.parse(txt)
			if ok {
				ret = append(ret, obj)
				break
			}
		}
		if len(ret)!=line {
			return nil, fmt.Errorf("no match for supplied scanners on line %d", line)
		}
	}

	return ret, nil
}

