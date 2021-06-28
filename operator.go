package expr

import (
	"errors"
	"fmt"
	"math"
	"strconv"
)

const (
	power     = 8
	multiply  = 8
	divide    = 8
	remainder = 8

	plus  = 7
	minus = 7

	shl = 6
	shr = 6

	lessThan        = 5
	lessOrEquals    = 5
	greaterThan     = 5
	greaterOrEquals = 5
	equals          = 5
	notEquals       = 5
)

var (
	NotFloat64Error = errors.New("param is not number")
	LackParamError  = fmt.Errorf("lack of params")
	WrongParamType  = fmt.Errorf("the type of param is not correct")
	DivideZero      = fmt.Errorf("divide zero")
)

type Operator interface {
	GetPriority() int
	Cal(p1, p2 float64) (float64, error)
}

type baseOperator struct {
	priority int
}

func (b baseOperator) GetPriority() int {
	return b.priority
}

var ops = map[string]Operator{
	/*
		"**": &PowerOperator{
			baseOperator{
				priority: power,
			},
		},*/
	"*": &MultiplyOperator{
		baseOperator{
			priority: multiply,
		},
	},
	/*
		"/": &DivideOperator{
			baseOperator{
				priority: divide,
			},
		},
		"%": &RemainderOperator{
			baseOperator{
				priority: remainder,
			},
		},*/

	"+": &PlusOperator{
		baseOperator{
			priority: plus,
		},
	},
	/*
		"-": &MinusOperator{
			baseOperator{
				priority: minus,
			},
		},
		"<<": &ShiftLeftOperator{
			baseOperator{
				priority: shl,
			},
		},
		">>": &ShiftRightOperator{
			baseOperator{
				priority: shr,
			},
		},
		"<": &LessOperator{
			baseOperator{
				priority: lessThan,
			},
		},
		"<=": &LessThanOperator{
			baseOperator{
				priority: lessOrEquals,
			},
		},
		">": &MoreOperator{
			baseOperator{
				priority: greaterThan,
			},
		},
		">=": &MoreThanOperator{
			baseOperator{
				priority: greaterOrEquals,
			},
		},
		"==": &EqualOperator{
			baseOperator{
				priority: equals,
			},
		},
		"!=": &NotEqualOperator{
			baseOperator{
				priority: notEquals,
			},
		},*/
}

type PlusOperator struct {
	baseOperator
}

func (p *PlusOperator) Cal(p1, p2 float64) (float64, error) {
	return p1 + p2, nil
}

type MinusOperator struct {
	baseOperator
}

func (p MinusOperator) Cal(value ...interface{}) (interface{}, error) {
	if len(value) != 2 {
		return "", LackParamError
	}

	param, ok := value[0].(float64)
	if !ok {
		return "", NotFloat64Error
	}

	param1, ok := value[1].(float64)
	if !ok {
		return "", NotFloat64Error
	}

	return strconv.FormatFloat(param1-param, 'g', -1, 64), nil
}

type MultiplyOperator struct {
	baseOperator
}

func (p *MultiplyOperator) Cal(p1, p2 float64) (float64, error) {
	return p1 * p2, nil
}

type DivideOperator struct {
	baseOperator
}

func (p DivideOperator) Cal(value ...interface{}) (interface{}, error) {
	if len(value) != 2 {
		return "", LackParamError
	}

	if value[0] == "0" {
		return "", DivideZero
	}

	param, ok := value[0].(float64)
	if !ok {
		return "", NotFloat64Error
	}

	param1, ok := value[1].(float64)
	if !ok {
		return "", NotFloat64Error
	}

	return strconv.FormatFloat(param1/param, 'g', -1, 64), nil
}

type RemainderOperator struct {
	baseOperator
}

func (p RemainderOperator) Cal(value ...interface{}) (interface{}, error) {
	if len(value) != 2 {
		return "", LackParamError
	}

	if value[0] == "0" {
		return "", DivideZero
	}

	param, ok := value[0].(int64)
	if !ok {
		return "", NotFloat64Error
	}

	param1, ok := value[1].(int64)
	if !ok {
		return "", NotFloat64Error
	}

	ret := param1 % param
	return strconv.FormatInt(ret, 10), nil
}

type PowerOperator struct {
	baseOperator
}

func (p PowerOperator) Cal(value ...interface{}) (interface{}, error) {
	if len(value) != 2 {
		return "", LackParamError
	}

	param, ok := value[0].(float64)
	if !ok {
		return "", NotFloat64Error
	}

	param1, ok := value[1].(float64)
	if !ok {
		return "", NotFloat64Error
	}

	return math.Pow(param, param1), nil
}

type ShiftLeftOperator struct {
	baseOperator
}

func (s ShiftLeftOperator) Cal(value ...interface{}) (interface{}, error) {
	if len(value) != 2 {
		return "", LackParamError
	}

	param, ok := value[0].(float64)
	if !ok {
		return "", NotFloat64Error
	}

	param1, ok := value[1].(float64)
	if !ok {
		return "", NotFloat64Error
	}

	for i := 0; i < int(param1); i++ {
		param = param * 2
	}

	return strconv.FormatFloat(param, 'g', -1, 64), nil
}

type ShiftRightOperator struct {
	baseOperator
}

func (s ShiftRightOperator) Cal(value ...interface{}) (interface{}, error) {
	if len(value) != 2 {
		return "", LackParamError
	}

	param, ok := value[0].(float64)
	if !ok {
		return "", NotFloat64Error
	}

	param1, ok := value[1].(uint64)
	if !ok {
		return "", NotFloat64Error
	}
	for i := 0; i < int(param1); i++ {
		param = param / 2
	}

	return strconv.FormatFloat(param, 'g', -1, 64), nil
}

type LessOperator struct {
	baseOperator
}

func (s LessOperator) Cal(value ...interface{}) (interface{}, error) {
	if len(value) != 2 {
		return "", LackParamError
	}

	param, ok := value[0].(float64)
	if !ok {
		return "", NotFloat64Error
	}

	param1, ok := value[1].(float64)
	if !ok {
		return "", NotFloat64Error
	}

	if param > param1 {
		return 1, nil
	} else {
		return 0, nil
	}
}

type LessThanOperator struct {
	baseOperator
}

func (s LessThanOperator) Cal(value ...interface{}) (interface{}, error) {
	if len(value) != 2 {
		return "", LackParamError
	}

	param, ok := value[0].(float64)
	if !ok {
		return "", NotFloat64Error
	}

	param1, ok := value[1].(float64)
	if !ok {
		return "", NotFloat64Error
	}

	if param >= param1 {
		return 1, nil
	} else {
		return 0, nil
	}
}

type MoreOperator struct {
	baseOperator
}

func (s MoreOperator) Cal(value ...interface{}) (interface{}, error) {
	if len(value) != 2 {
		return "", LackParamError
	}

	param, ok := value[0].(float64)
	if !ok {
		return "", NotFloat64Error
	}

	param1, ok := value[1].(float64)
	if !ok {
		return "", NotFloat64Error
	}

	if param < param1 {
		return 1, nil
	} else {
		return 0, nil
	}
}

type MoreThanOperator struct {
	baseOperator
}

func (s MoreThanOperator) Cal(value ...interface{}) (interface{}, error) {
	if len(value) != 2 {
		return "", LackParamError
	}

	param, ok := value[0].(float64)
	if !ok {
		return "", NotFloat64Error
	}

	param1, ok := value[1].(float64)
	if !ok {
		return "", NotFloat64Error
	}

	if param < param1 {
		return 1, nil
	} else {
		return 0, nil
	}
}

type EqualOperator struct {
	baseOperator
}

func (s EqualOperator) Cal(value ...interface{}) (interface{}, error) {
	if len(value) != 2 {
		return "", LackParamError
	}

	param, ok := value[0].(float64)
	if !ok {
		return "", NotFloat64Error
	}

	param1, ok := value[1].(float64)
	if !ok {
		return "", NotFloat64Error
	}

	if param == param1 {
		return 1, nil
	} else {
		return 0, nil
	}
}

type NotEqualOperator struct {
	baseOperator
}

func (s NotEqualOperator) Cal(value ...interface{}) (interface{}, error) {
	if len(value) != 2 {
		return "", LackParamError
	}

	param, ok := value[0].(float64)
	if !ok {
		return "", NotFloat64Error
	}

	param1, ok := value[1].(float64)
	if !ok {
		return "", NotFloat64Error
	}

	if param == param1 {
		return 1, nil
	} else {
		return 0, nil
	}
}
