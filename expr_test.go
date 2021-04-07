package expr

import "testing"

func TestParse(t *testing.T) {
	tokens, err := tokenize([]rune("1 + (1 + 1)"))
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(tokens)

	expr, err := Parse("1 + (1 + a)", map[string]Var{
		"a": NewVar(Num(10.1)),
	}, nil)
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(expr.Eval())
}
