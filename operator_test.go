package expr

import "testing"

func TestAdd(t *testing.T) {
	p := PlusOperator{}

	t.Log(p.Cal("1", "2"))
}

func TestRemainder(t *testing.T) {
	r := RemainderOperator{}

	t.Log(r.Cal("1", "2"))
}

func TestPower(t *testing.T) {
	p := PowerOperator{}

	t.Log(p.Cal("2", "2"))
}

func TestShiftLeft(t *testing.T) {
	s := ShiftLeftOperator{}

	t.Log(s.Cal("2", "4"))
}

func TestLess(t *testing.T) {
	l := LessOperator{}

	t.Log(l.Cal("1", "2"))
}

func TestLessThan(t *testing.T) {
	lt := LessThanOperator{}

	t.Log(lt.Cal("1", "0"))
}

func TestMoreThan(t *testing.T) {
	mt := MoreThanOperator{}

	t.Log(mt.Cal("1", "1"))
}

func BenchmarkAdd(b *testing.B) {
	p := PlusOperator{}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		p.Cal()
	}
}

func BenchmarkNormalAdd(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = 1 + 2
	}
}
