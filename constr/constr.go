package constr

type Constraint interface {
	Match(o interface{}, e Env) bool
}

type Exprconstr struct {
	e    Expr
	name string
}

func (ec Exprconstr) Match(o interface{}) bool {
	env := make(Env)

	env[ec.name] = IntValue(o.(int))

	return ec.e.Eval(env).AsBool()
}
