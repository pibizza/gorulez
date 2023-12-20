package expr

func (e Value) Validate(env Env) bool {
	if e.t == INT {
		_, ok := e.v.(int)
		return ok
	}

	if e.t == FLOAT {
		_, ok := e.v.(float64)
		return ok
	}

	return false

}

func (e Plus) Validate(env Env) bool {
	if e.t == FLOAT && e.e1.Type() == FLOAT && e.e2.Type() == FLOAT {
		return true

	}
	if e.t == INT && e.e1.Type() == INT && e.e2.Type() == INT {
		return true

	}
	return false
}
