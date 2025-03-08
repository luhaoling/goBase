package dataStructureAndAlgorithms

type StackInterface interface {
	push(int)
	poll() int
	getTop() int
}

type stack struct {
	arr []int
	len int
	cap int
}

func (s *stack) push(i int) {
	s.arr[s.len] = i
	return
}

func (s *stack) poll() int {
	val := s.arr[s.len]
	for i := 0; i < s.len; i++ {
		s.arr[i] = s.arr[i+1]
	}

	return val
}

func (s *stack) getTop() int {
	return s.arr[s.len]
}

func NewStack(cap int) StackInterface {
	return &stack{
		arr: []int{},
		len: 0,
		cap: cap,
	}
}
