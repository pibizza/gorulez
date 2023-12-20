package expr

import "testing"

func TestValidateInt(t *testing.T) {
	good := Value{3, INT}

	env := make(map[string]Value)
	if !good.Validate(env) {
		t.Error("3 is not a valid int")
	}
	bad := Value{3.0, INT}

	if bad.Validate(env) {
		t.Error("3.0 is a valid int")
	}

}

func TestValidateFloat(t *testing.T) {
	good := Value{3.0, FLOAT}

	env := make(map[string]Value)
	if !good.Validate(env) {
		t.Error("3.0 is not a valid FLOAT")
	}
	bad := Value{3, FLOAT}

	if bad.Validate(env) {
		t.Error("3 is a valid FLOAT")
	}

}

func TestValidatePlus(t *testing.T) {
	v1 := FloatValue(1.0)
	v2 := FloatValue(2.0)

	plus := Plus{v1, v2, FLOAT}
	env := make(map[string]Value)
	if !plus.Validate(env) {
		t.Error("1.0 + 2.0 should be valid and it is not")
	}

	v3 := IntValue(3)
	v4 := IntValue(4)

	intplus := Plus{v3, v4, INT}

	if !intplus.Validate(env) {
		t.Error("1 + 2 should be valid and it is not")
	}
}
