package scan

import (
	"bufio"
	"fmt"
	"io"
)

// LineScanner parses a line (string) and returns a copy of the result along with an ok/no good signal
type LineScanner interface {
	Parse(string) (LineScanner, bool)
}

// ReadAll reads all lines provided by 'r' and attempts to parse each line with the passed 'scanners'
func ReadAll(r io.Reader, scanners ...LineScanner) ([]LineScanner, error) {
	s := bufio.NewScanner(r)
	ret := make([]LineScanner, 0)
	line := 0
	for s.Scan() {
		line++
		for _, sr := range scanners {
			txt := s.Text()
			obj, ok := sr.Parse(txt)
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

