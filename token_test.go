package expr

import (
	"testing"
)

func TestTokenize(t *testing.T) {
	t.Log(Tokenize([]rune("1+2")))
}

func TestEval(t *testing.T) {
	expression := NewExpression("3%2+1")

	expression.Tokenize()
	t.Log(expression.Eval())
}
