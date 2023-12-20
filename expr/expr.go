package expr

type ValueType int

const (
	INT ValueType = iota
	BOOL
	FLOAT
	STRING
)

var zeroValues = map[ValueType]Value{
	INT:    IntValue(0),
	BOOL:   BoolValue(false),
	FLOAT:  FloatValue(0.0),
	STRING: StringValue(""),
}

func ZeroFor(vt ValueType) Value {
	return zeroValues[vt]
}

type Value struct {
	v interface{}
	t ValueType
}

func (v Value) Type() ValueType {
	return v.t
}

func IntValue(v int) Value {
	return Value{v, INT}
}

func FloatValue(v float64) Value {
	return Value{v, FLOAT}
}

func BoolValue(v bool) Value {
	return Value{v, BOOL}
}

func StringValue(s string) Value {
	return Value{s, STRING}
}

func (dv Value) AsInt() int {
	if dv.t == INT {
		v, _ := dv.v.(int)
		return v
	}
	return 0
}

func (dv Value) AsFloat() float64 {
	if dv.t == FLOAT {
		v, _ := dv.v.(float64)
		return v
	}
	return 0.0
}

func (dv Value) AsBool() bool {
	if dv.t == BOOL {
		v, _ := dv.v.(bool)
		return v

	}
	return false
}

func (dv Value) AsString() string {
	if dv.t == STRING {
		v, _ := dv.v.(string)
		return v

	}
	return ""
}

type Env map[string]Value

func (v Value) Eval(e Env) Value {
	return v
}

type Id struct {
	name string
	t    ValueType
}

func (i Id) Type() ValueType {
	return i.t
}

func IntId(name string) Id {
	return Id{name, INT}
}

func (i Id) Eval(e Env) Value {
	return e[i.name]
}

type Bop struct {
	e1 Expr
	e2 Expr
	t  ValueType
}

func (b Bop) Type() ValueType {
	return b.t
}

type Plus Bop

func IntPlus(e1 Expr, e2 Expr) Plus {
	return Plus{e1, e2, INT}
}

func (b Plus) Eval(e Env) Value {
	if b.t == INT {
		return IntValue(b.e1.Eval(e).AsInt() + b.e2.Eval(e).AsInt())
	}
	if b.t == FLOAT {
		return FloatValue(b.e1.Eval(e).AsFloat() + b.e2.Eval(e).AsFloat())
	}
	return ZeroFor(b.t)
}

type Minus Bop

func (b Minus) Eval(e Env) Value {
	if b.t == INT {
		return IntValue(b.e1.Eval(e).AsInt() - b.e2.Eval(e).AsInt())
	}

	if b.t == FLOAT {
		return FloatValue(b.e1.Eval(e).AsFloat() - b.e2.Eval(e).AsFloat())
	}
	return ZeroFor(b.t)
}

type Times Bop

func (b Times) Eval(e Env) Value {
	if b.t == INT {
		return IntValue(b.e1.Eval(e).AsInt() * b.e2.Eval(e).AsInt())
	}

	if b.t == FLOAT {
		return FloatValue(b.e1.Eval(e).AsFloat() * b.e2.Eval(e).AsFloat())
	}
	return ZeroFor(b.t)
}

type DividedBy Bop

func (b DividedBy) Eval(e Env) Value {
	if b.t == INT {
		return IntValue(b.e1.Eval(e).AsInt() / b.e2.Eval(e).AsInt())
	}

	if b.t == FLOAT {
		return FloatValue(b.e1.Eval(e).AsFloat() / b.e2.Eval(e).AsFloat())
	}
	return ZeroFor(b.t)
}

type Modulo Bop

func (b Modulo) Eval(e Env) Value {
	if b.t == INT {
		return IntValue(b.e1.Eval(e).AsInt() % b.e2.Eval(e).AsInt())
	}

	return ZeroFor(b.t)
}

type And Bop

func (b And) Eval(e Env) Value {
	return BoolValue(b.e1.Eval(e).AsBool() && b.e2.Eval(e).AsBool())
}

type Or Bop

func (b Or) Eval(e Env) Value {
	return BoolValue(b.e1.Eval(e).AsBool() || b.e2.Eval(e).AsBool())
}

type Equal Bop

func (b Equal) Eval(e Env) Value {
	return BoolValue(b.e1.Eval(e) == b.e2.Eval(e))
}

type NotEqual Bop

func (b NotEqual) Eval(e Env) Value {
	return BoolValue(b.e1.Eval(e) != b.e2.Eval(e))
}

type Greater Bop

func (b Greater) Eval(e Env) Value {
	if b.t == INT {
		return BoolValue(b.e1.Eval(e).AsInt() > b.e2.Eval(e).AsInt())
	}

	if b.t == FLOAT {
		return BoolValue(b.e1.Eval(e).AsFloat() > b.e2.Eval(e).AsFloat())
	}
	return ZeroFor(BOOL)
}

type GreaterOrEqual Bop

func (b GreaterOrEqual) Eval(e Env) Value {
	if b.t == INT {
		return BoolValue(b.e1.Eval(e).AsInt() >= b.e2.Eval(e).AsInt())
	}

	if b.t == FLOAT {
		return BoolValue(b.e1.Eval(e).AsFloat() >= b.e2.Eval(e).AsFloat())
	}
	return ZeroFor(BOOL)
}

type Lesser Bop

func (b Lesser) Eval(e Env) Value {
	if b.t == INT {
		return BoolValue(b.e1.Eval(e).AsInt() < b.e2.Eval(e).AsInt())
	}

	if b.t == FLOAT {
		return BoolValue(b.e1.Eval(e).AsFloat() < b.e2.Eval(e).AsFloat())
	}
	return ZeroFor(BOOL)
}

type LesserOrEqual Bop

func (b LesserOrEqual) Eval(e Env) Value {
	if b.t == INT {
		return BoolValue(b.e1.Eval(e).AsInt() <= b.e2.Eval(e).AsInt())
	}

	if b.t == FLOAT {
		return BoolValue(b.e1.Eval(e).AsFloat() <= b.e2.Eval(e).AsFloat())
	}
	return ZeroFor(BOOL)
}

type Expr interface {
	Eval(e Env) Value
	Type() ValueType
}
