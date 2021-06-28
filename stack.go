package expr

// Simple string stack implementation
type Stack struct {
	data  []interface{}
	index int
}

func (s *Stack) Reset() {
	s.data = s.data[:0]
}

func (s *Stack) Push(str interface{}) {
	s.data = append(s.data, str)
}

func (ss *Stack) Peek() interface{} {
	if len(ss.data) == 0 {
		return nil
	}

	return ss.data[len(ss.data)-1]
}

func (ss *Stack) Pop() interface{} {
	if len(ss.data) == 0 {
		return nil
	}

	ret := ss.data[len(ss.data)-1]
	ss.data = ss.data[:len(ss.data)-1]
	return ret
}

func (ss *Stack) PopLeft() interface{} {
	if len(ss.data) == 0 {
		return nil
	}

	ret := ss.data[0]
	ss.data = ss.data[1:]
	return ret
}

func (ss *Stack) IndexPopLeft() interface{} {
	if len(ss.data) == 0 || ss.index >= len(ss.data) {
		return nil
	}

	ret := ss.data[ss.index]
	ss.index++
	return ret
}

func (ss *Stack) ResetIndex() {
	ss.index = 0
}

func (ss *Stack) IsEmpty() bool {
	return len(ss.data) == 0
}

func (ss *Stack) Length() int {
	return len(ss.data)
}

func NewStack(cap int) *Stack {
	if cap <= 0 {
		return &Stack{
			data: []interface{}{},
		}
	}

	data := make([]interface{}, 0, cap)
	return &Stack{
		data: data,
	}
}
