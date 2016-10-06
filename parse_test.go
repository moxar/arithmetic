package arithmetic

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"testing"
)

func Test_Parse(t *testing.T) {
	file, err := os.Open("tests.txt")
	if err != nil {
		t.Log(err)
		t.Fail()
		return
	}
	defer file.Close()

	type Case struct {
		line int
		in   string
		out  string
		err  bool
	}

	scanner := bufio.NewScanner(file)
	var i int
	var cases []Case
	for scanner.Scan() {
		i++
		line := scanner.Text()
		idx := strings.LastIndex(line, "=")
		if idx == -1 {
			continue
		}
		res := strings.TrimSpace(line[idx+1:])
		cases = append(cases, Case{
			line: i,
			in:   strings.TrimSpace(line[:idx]),
			out:  res,
			err:  res == "error",
		})
	}

	err = scanner.Err()
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	for _, c := range cases {
		out, err := Parse(c.in)
		if (err != nil) != c.err {
			if err == nil {
				t.Logf("line %d [%s = %s]: expected error\n", c.line, c.in, c.out)
				t.Fail()
			} else {
				t.Logf("line %d [%s = %s]: unexpected error %s\n", c.line, c.in, c.out, err)
				t.Fail()
			}
			continue
		}

		if err != nil {
			continue
		}

		if fmt.Sprintf("%v", out) != c.out {
			t.Logf("line %d [%s = %s]: unexpected output: %v\n", c.line, c.in, c.out, out)
			t.Fail()
			continue
		}
	}
}
