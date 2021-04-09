package expr

import "testing"

func TestParse(t *testing.T) {
	expr, err := Parse("1+ @.a*@.a > 1")
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(expr)

	ret, err := Run(expr, map[string]string{
		"a": "2",
	})
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(ret)
}

func BenchmarkRun(b *testing.B) {
	expr, err := Parse("1+@.a*3")
	if err != nil {
		b.Error(err)
		return
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		Run(expr, map[string]string{
			"aa": "1",
		})
	}
}
