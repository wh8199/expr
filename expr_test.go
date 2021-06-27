package expr

import "testing"

func TestParse(t *testing.T) {
	expr, err := Compile("2 ** 2 > 3")
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(expr)

	ret, err := Run(expr, map[string]string{
		"a": "2",
		"b": "1",
		"c": "2",
	})
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(ret)
}

func BenchmarkRun(b *testing.B) {

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		expr, err := Compile("1+2")
		if err != nil {
			b.Error(err)
			return
		}

		Run(expr, map[string]string{
			"aa": "1",
		})
	}
}
