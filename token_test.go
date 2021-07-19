package expr

import (
	"testing"
)

func TestTokenize(t *testing.T) {
	t.Log(Tokenize([]rune("1+2")))
}

func TestEval(t *testing.T) {
	expression := NewExpression("2 * ( 1 + 4) + 2")

	expression.Tokenize()
	t.Log(expression.Eval())
}

func BenchmarkEval(b *testing.B) {
	expression := NewExpression("1++2--3+4+10")
	expression.Tokenize()

	for i := 0; i < b.N; i++ {
		expression.Eval()
	}
}
