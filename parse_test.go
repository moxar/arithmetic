package arithmetic

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestParse(t *testing.T) {
	file, err := os.Open("tests.txt")
	if err != nil {
		t.Log(err)
		t.Fail()
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var i int
	for scanner.Scan() {
		i++
		line := scanner.Text()
		idx := strings.LastIndex(line, "=")
		exp, res := line[:idx], strings.TrimSpace(line[idx+1:])
		out, err := Parse(exp)
		if err != nil {
			out = err
		}
		out = fmt.Sprintf("%v", out)
		if out != res {
			t.Logf("line %d: [%s] returned [%s]", i, line, out)
			t.Fail()
		}
	}

	err = scanner.Err()
	if err != nil {
		t.Log(err)
		t.Fail()
	}
}
