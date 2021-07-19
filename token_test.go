package expr

import (
	"testing"
)

func TestTokenize(t *testing.T) {
	t.Log(Tokenize([]rune("1+2")))
}

func TestEval(t *testing.T) {
	expression := NewExpression("a+b*c")
	expression.Variables = map[string]float64{
		"a": 1,
		"b": 2,
		"c": 12,
	}
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
