package stack

type Stack struct {
	storage []string
}

func (s *Stack) Push(element string) {
	s.storage = append(s.storage, element)
}

func (s *Stack) Pop() string {
	if len(s.storage) == 0 {
		return ""
	} else {
		value := s.storage[len(s.storage)-1]
		s.storage = s.storage[:len(s.storage)-1]
		return value
	}
}

func (s *Stack) Peek() string {
	if len(s.storage) > 0 {
		return ""
	} else {
		return s.storage[len(s.storage)-1]
	}
}

func (s *Stack) count() int {
	return len(s.storage)
}

func (s *Stack) isEmpty() bool {
	return s.count() == 0
}

func buildStack() *Stack {
	return &Stack{storage: []string{}}
}

func (s Stack) String() string {
	result := "---- top --- \n"
	for _, n := range s.storage {
		result = result + n + "\n"
	}
	result = result + "-------"

	return result
}

func CreateFrom(array []string) *Stack {
	stack := buildStack()
	for _, n := range array {
		stack.Push(n)
	}
	return stack
}

func CheckParentheses(expression string) bool {
	stack := buildStack()
	for _, rune := range expression {
		char := string(rune)
		switch char {
		case "(":
			stack.Push(char)
		case ")":
			if stack.isEmpty() {
				return false
			} else {
				stack.Pop()
			}
		default:
		}
	}
	return stack.isEmpty()
}
