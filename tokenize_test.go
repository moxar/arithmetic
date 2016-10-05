package arithmetic

import (
	"reflect"
	"testing"
)

func TestTokenize(t *testing.T) {

	for i, c := range []struct {
		in  string
		out []interface{}
		err bool
	}{
		{
			in:  " ",
			out: nil,
			err: false,
		},
		{
			in:  ":D",
			out: nil,
			err: true,
		},
		{
			in:  "5€",
			out: nil,
			err: true,
		},
		{
			in:  "5 €",
			out: nil,
			err: true,
		},
		{
			in:  "undef - 5",
			out: nil,
			err: true,
		},
		{
			in:  "3 - undef",
			out: nil,
			err: true,
		},
		{
			in:  "3 - 5",
			out: []interface{}{3.0, minus{}, 5.0},
			err: false,
		},
		{
			in:  "4 5 6",
			out: []interface{}{4.0, 5.0, 6.0},
			err: false,
		},
		{
			in:  "3 * 5",
			out: []interface{}{3.0, multiply{}, 5.0},
			err: false,
		},
	} {

		out, err := Tokenize(c.in)
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
