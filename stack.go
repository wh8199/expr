package expr

// Simple string stack implementation
type stringStack struct {
	data []string
}

func (s *stringStack) Reset() {
	s.data = s.data[:0]
}

func (s *stringStack) Push(str string) {
	s.data = append(s.data, str)
}

func (ss *stringStack) Peek() string {
	if len(ss.data) == 0 {
		return ""
	}

	return ss.data[len(ss.data)-1]
}

func (ss *stringStack) Pop() string {
	if len(ss.data) == 0 {
		return ""
	}

	ret := ss.data[len(ss.data)-1]
	ss.data = ss.data[:len(ss.data)-1]
	return ret
}

func (ss *stringStack) PopLeft() string {
	if len(ss.data) == 0 {
		return ""
	}

	ret := ss.data[0]
	ss.data = ss.data[1:]
	return ret
}

func (ss *stringStack) IsEmpty() bool {
	return len(ss.data) == 0
}

func (ss *stringStack) Length() int {
	return len(ss.data)
}

func NewStack(cap int) *stringStack {
	if cap <= 0 {
		return &stringStack{
			data: []string{},
		}
	}

	data := make([]string, 0, cap)
	return &stringStack{
		data: data,
	}
}
