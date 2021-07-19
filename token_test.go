package expr

import (
	"testing"
)

func TestTokenize(t *testing.T) {
	t.Log(Tokenize([]rune("1+2")))
}

func TestEval(t *testing.T) {
	expression := NewExpression("1++2--3")

	expression.Tokenize()
	t.Log(expression.Eval())
}
