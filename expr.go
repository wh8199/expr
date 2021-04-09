package expr

import (
	"errors"
	"fmt"
)

var (
	ErrParen                = errors.New("parenthesis mismatch")
	ErrUnexpectedNumber     = errors.New("unexpected number")
	ErrUnexpectedIdentifier = errors.New("unexpected identifier")

	ErrBadCall        = errors.New("function call expected")
	ErrBadVar         = errors.New("variable expected in assignment")
	ErrBadOp          = errors.New("unknown operator or function")
	ErrOperandMissing = errors.New("missing operand")
)

func Compile(input string) (*Stack, error) {
	operationStack := NewStack(0)
	ret := NewStack(0)

	tokens, err := tokenize([]rune(input))
	if err != nil {
		return nil, err
	}

	for _, token := range tokens {
		if token == "(" {
			operationStack.Push("(")
		} else if token == ")" {
			for !operationStack.IsEmpty() && operationStack.Peek() != "(" {
				ret.Push(operationStack.Pop())
			}

			operationStack.Pop()
		} else if op, ok := ops[token]; ok {
			o2 := operationStack.Peek()
			for ops[o2] != nil && op.GetPriority() > ops[o2].GetPriority() {
				ret.Push(token)
				operationStack.Pop()
				o2 = operationStack.Peek()
			}
			operationStack.Push(token)
		} else {
			ret.Push(token)
		}
	}

	for !operationStack.IsEmpty() {
		op := operationStack.Pop()
		ret.Push(op)
	}

	return ret, nil
}

func Run(stack *Stack, vars map[string]string) (string, error) {
	doubleStack := NewStack(stack.Length())

	for !stack.IsEmpty() {
		token := stack.PopLeft()

		if operation, ok := ops[token]; ok {
			if doubleStack.Length() < 2 {
				return "", fmt.Errorf("Not enough params to exec '%s'", token)
			}

			param1 := doubleStack.Pop()
			param2 := doubleStack.Pop()

			ret, err := operation.Cal(param1, param2)
			if err != nil {
				return "", err
			}

			doubleStack.Push(ret)
		} else if value, ok := vars[token]; ok {
			doubleStack.Push(value)
		} else {
			doubleStack.Push(token)
		}
	}

	return doubleStack.Pop(), nil
}
