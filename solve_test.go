package arithmetic

import (
	"reflect"
	"testing"
)

func TestSolve(t *testing.T) {

	for i, c := range []struct {
		in  []interface{}
		out interface{}
		err bool
	}{
		{
			in:  []interface{}{},
			out: nil,
			err: false,
		},
		{
			in:  []interface{}{minus{}},
			out: nil,
			err: true,
		},
		{
			in:  []interface{}{3.0, 2.0, minus{}},
			out: 1.0,
			err: false,
		},
		{
			in:  []interface{}{3.0, 2.0, minus{}, minus{}},
			out: nil,
			err: true,
		},
	} {

		out, err := Solve(c.in)
		if (err != nil) != c.err {
			t.Log("case", i+1, "unexpected error")
			t.Log("want:", c.err)
			t.Log("got: ", err)
			t.Fail()
			continue
		}

		if !reflect.DeepEqual(out, c.out) {
			t.Log("case", i+1, "unexpected output")
			t.Log("want:", c.out)
			t.Log("got: ", out)
			t.Fail()
		}
	}
}
