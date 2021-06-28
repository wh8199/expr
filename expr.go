package expr

import (
	"errors"
	"fmt"
	"strconv"
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

func Compile(input string) (*Expression, error) {
	operationStack := NewStack(0)
	ret := NewStack(0)

	tokens, functionParams, err := tokenize([]rune(input))
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

			if o2 == nil || o2 == "(" || op.GetPriority() > ops[o2.(string)].GetPriority() {
				operationStack.Push(token)
				continue
			}

			for o2 != nil && o2 != "(" && op.GetPriority() <= ops[o2.(string)].GetPriority() {
				ret.Push(o2)
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

	return &Expression{
		stack:          ret,
		functionParams: functionParams,
		params:         []interface{}{},
	}, nil
}

type Expression struct {
	stack          *Stack
	functionParams map[string][]interface{}
	params         []interface{}
}

func (e *Expression) Run(vars map[string]interface{}) (interface{}, error) {
	tokenStack := e.stack
	functionParams := e.functionParams

	tokenStack.ResetIndex()
	e.params = e.params[:0]
	if tokenStack == nil || functionParams == nil {
		return "", fmt.Errorf("invalid expression")
	}

	for {
		token := tokenStack.IndexPopLeft()
		if token == nil {
			break
		}

		switch token.(type) {
		case string:
			token := token.(string)
			if operation, ok := ops[token]; ok {
				if len(e.params) < 2 {
					return "", fmt.Errorf("not enough params to exec '%s'", token)
				}

				var err error
				p1 := e.params[0]
				p2 := e.params[1]

				e.params = e.params[2:]

				switch p1.(type) {
				case string:
					p1, err = strconv.ParseFloat(p1.(string), 64)
					if err != nil {
						return nil, err
					}
				}

				switch p2.(type) {
				case string:
					p2, err = strconv.ParseFloat(p2.(string), 64)
					if err != nil {
						return nil, err
					}
				}

				ret, err := operation.Cal(1, 1)
				if err != nil {
					return nil, err
				}

				e.params = append(e.params, ret)
			} else if value, ok := vars[token]; ok {
				e.params = append(e.params, value)
			} else if funcs, ok := functions[token]; ok {
				rawParams := functionParams[token]
				newParams := make([]interface{}, len(rawParams))
				for index, rawParam := range rawParams {
					switch rawParam.(type) {
					case string:
						p, err := strconv.ParseFloat(rawParam.(string), 64)
						if err != nil {
							return nil, err
						}

						newParams[index] = p
					default:
						newParams[index] = rawParam
					}
				}

				result, err := funcs(newParams...)
				if err != nil {
					return nil, err
				}
				e.params = append(e.params, result)
			} else {
				e.params = append(e.params, token)
			}
		default:
			e.params = append(e.params, token)
		}
	}

	//log.Info(e.params)
	return 1, nil
}
