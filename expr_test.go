package expr

import (
	"testing"

	"gopkg.in/Knetic/govaluate.v3"
)

func TestParse(t *testing.T) {
	expr, err := Compile("1+2+2 * 10+(1+2) * 3")
	if err != nil {
		t.Error(err)
		return
	}

	ret, err := expr.Run(nil)
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(ret)
}

func TestGoevel(t *testing.T) {
	expString := `1 << 2.1+2`
	expression, err := govaluate.NewEvaluableExpression(expString)
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(expression.Evaluate(nil))
}

func BenchmarkRun(b *testing.B) {
	expr, err := Compile("1+2+2 * 10+(1+2) * 3")
	if err != nil {
		b.Error(err)
		return
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		expr.Run(nil)
	}
}

func BenchmarkGoevel(b *testing.B) {
	expString := `1+2+2 * 10+(1+2) * 3`
	expression, err := govaluate.NewEvaluableExpression(expString)
	if err != nil {
		b.Error(err)
		return
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		expression.Evaluate(nil)
	}
}
