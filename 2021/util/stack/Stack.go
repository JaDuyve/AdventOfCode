package stack

type Stack []byte

func (s *Stack) Push(v byte) {
	*s = append([]byte{v}, *s...)
}

func (s *Stack) Pop() (byte, bool) {
	if len(*s) == 0 {
		var v byte
		return v, false
	}

	v := (*s)[0]
	*s = (*s)[1:]
	return v, true
}
