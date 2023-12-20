package expr

import "testing"

func TestEvalPlus(t *testing.T) {
	e1 := IntValue(3)
	e2 := IntValue(4)

	e3 := IntPlus(e1, e2)

	env := make(map[string]Value)

	if !(e3.Eval(env).AsInt() == 7) {
		t.Error("3 +4 != 7")
	}
}
