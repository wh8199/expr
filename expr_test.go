package expr

import "testing"

func TestParse(t *testing.T) {
	expr, err := Parse("1+@.a <= @.a", map[string]Var{
		"a": NewVar(Num(1)),
	}, nil)
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(expr.Eval())
}
