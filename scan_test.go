package scan_test

import (
	"fmt"
	"github.com/spellgen/scan"
	"strings"
	"testing"
)

type singleInt struct {
	v int
}

func (o singleInt) Parse(in string) (scan.LineScanner, bool) {
	n, err := fmt.Sscanf(in, "%d\n", &o.v)
	return o, n == 1 && err == nil
}

type empty struct{}

func (o empty) Parse(in string) (scan.LineScanner, bool) {
	n, err := fmt.Sscanf(in, "\n")
	return o, n == 0 && err == nil
}

type pair struct {
	a int
	b int
}

func (o pair) parse(in string) (scan.LineScanner, bool) {
	n, err := fmt.Sscanf(in, "%d,%d\n", &o.a, &o.b)
	return o, n == 2 && err == nil
}

func TestScanner1(t *testing.T) {
	input := "1\n\n2\n3,4"
	var si singleInt
	var e empty
	var p pair
	r := strings.NewReader(input)
	data, err := scan.ReadAll(r, si, e, p)
	if err != nil {
		t.Log(err)
		t.Fail()
	} else {
		for k, d := range data {
			t.Logf("k=%d, d=%#v", k, d)
		}
	}
}
