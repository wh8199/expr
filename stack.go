package expr

// Simple string stack implementation
type Stack struct {
	data []string
}

func (s *Stack) Reset() {
	s.data = s.data[:0]
}

func (s *Stack) Push(str string) {
	s.data = append(s.data, str)
}

func (ss *Stack) Peek() string {
	if len(ss.data) == 0 {
		return ""
	}

	return ss.data[len(ss.data)-1]
}

func (ss *Stack) Pop() string {
	if len(ss.data) == 0 {
		return ""
	}

	ret := ss.data[len(ss.data)-1]
	ss.data = ss.data[:len(ss.data)-1]
	return ret
}

func (ss *Stack) PopLeft() string {
	if len(ss.data) == 0 {
		return ""
	}

	ret := ss.data[0]
	ss.data = ss.data[1:]
	return ret
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
			data: []string{},
		}
	}

	data := make([]string, 0, cap)
	return &Stack{
		data: data,
	}
}
