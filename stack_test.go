package arithmetic

import(
	"reflect"
	"testing"
)

func TestStack_push(t *testing.T) {
	st := &stack{}
	out := []interface{}{4, 5, 6}
	st.push(4)
	st.push(5)
	st.push(6)
	if !reflect.DeepEqual(st.values, out) {
		t.Fail()
	}
}

func TestStack_pop(t *testing.T) {
	st := &stack{[]interface{}{4, 5, 6}}
	out := []interface{}{6, 5, 4}
	for _, o := range out {
		v, ok := st.pop()
		if !ok {
			t.Log("stack should not be empty")
			t.Fail()
		}
		if v != o {
			t.Logf("%v != %v\n", v, o)
			t.Fail()
		}
	}
	_, ok := st.pop()
	if ok {
		t.Log("stack should be empty")
		t.Fail()
	}
}

func TestStack_inc(t *testing.T) {
	st := &stack{[]interface{}{4, 5, 6}}
	out := []interface{}{4, 5, 7}
	st.inc()
	if !reflect.DeepEqual(st.values, out) {
		t.Fail()
	}
}

func TestStack_slice(t *testing.T) {
	st := &stack{[]interface{}{4, 5, 6, 7, 8}}
	out, err := st.slice(2)
	if err != nil {
		t.Log("unexpected error, having", err)
		t.Fail()
	}
	
	if !reflect.DeepEqual(st.values, []interface{}{4, 5, 6}) {
		t.Logf("invalid stack values: %v\n", st.values)
		t.Fail()
	}
	if !reflect.DeepEqual(out, []interface{}{7, 8}) {
		t.Logf("invalid output values: %v\n", out)
		t.Fail()
	}
	
	_, err = st.slice(10)
	if err == nil {
		t.Log("expecting error, having none")
		t.Fail()
	}
}

func TestStack_popFloat(t *testing.T) {
	st := &stack{[]interface{}{4.0}}
	out, err := st.popFloat()
	if err != nil {
		t.Log("unexpected error:", err)
		t.Fail()
	}
	if out != 4.0 {
		t.Log("unexpected output:", out)
		t.Fail()
	}
	
	out, err = st.popFloat()
	if err == nil {
		t.Log("expected error, having none")
		t.Fail()
	}
	
	st.values = []interface{}{1}
	out, err = st.popFloat()
	if err == nil {
		t.Log("expected error, having none")
		t.Fail()
	}
}

func TestStack_popInt(t *testing.T) {
	st := &stack{[]interface{}{4}}
	out, err := st.popInt()
	if err != nil {
		t.Log("unexpected error:", err)
		t.Fail()
	}
	if out != 4 {
		t.Log("unexpected output:", out)
		t.Fail()
	}
	
	out, err = st.popInt()
	if err == nil {
		t.Log("expected error, having none")
		t.Fail()
	}
	
	st.values = []interface{}{1.0}
	out, err = st.popInt()
	if err == nil {
		t.Log("expected error, having none")
		t.Fail()
	}
}
