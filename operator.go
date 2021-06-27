package expr

import (
	"fmt"
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

var ops = map[string]Operator{
	"**": &PlusOperator{
		priority: power,
	},
	"*": &MultiplyOperator{
		priority: multiply,
	},
	"/": &DivideOperator{
		priority: divide,
	},
	"%": &RemainderOperator{
		priority: remainder,
	},
	"+": &PlusOperator{
		priority: plus,
	},
	"-": &MinusOperator{
		priority: minus,
	},
	"<<": &ShiftLeftOperator{
		priority: shl,
	},
	">>": &ShiftRightOperator{
		priority: shr,
	},
	"<": &LessOperator{
		priority: lessThan,
	},
	"<=": &LessThanOperator{
		priority: lessOrEquals,
	},
	">": &MoreOperator{
		priority: greaterThan,
	},
	">=": &MoreThanOperator{
		priority: greaterOrEquals,
	},
	"==": &EqualOperator{
		priority: equals,
	},
	"!=": &NotEqualOperator{
		priority: notEquals,
	},
	//"&&": logicalAnd, "||": logicalOr,
}

var (
	LackParamError = fmt.Errorf("Lack of params")
	WrongParamType = fmt.Errorf("The type of param is not correct")
	DivideZero     = fmt.Errorf("Divide zero")
)

type Operator interface {
	GetPriority() int
	Cal(value ...string) (string, error)
}

type PlusOperator struct {
	priority int
}

func (p PlusOperator) GetPriority() int {
	return p.priority
}

func (p PlusOperator) Cal(value ...string) (string, error) {
	if len(value) != 2 {
		return "", LackParamError
	}

	param, err := strconv.ParseFloat(value[0], 64)
	if err != nil {
		return "", err
	}

	param1, err := strconv.ParseFloat(value[1], 64)
	if err != nil {
		return "", err
	}

	return strconv.FormatFloat(param+param1, 'g', -1, 64), nil
}

type MinusOperator struct {
	priority int
}

func (p MinusOperator) GetPriority() int {
	return p.priority
}

func (p MinusOperator) Cal(value ...string) (string, error) {
	if len(value) != 2 {
		return "", LackParamError
	}

	param, err := strconv.ParseFloat(value[0], 64)
	if err != nil {
		return "", err
	}

	param1, err := strconv.ParseFloat(value[1], 64)
	if err != nil {
		return "", err
	}

	return strconv.FormatFloat(param1-param, 'g', -1, 64), nil
}

type MultiplyOperator struct {
	priority int
}

func (p MultiplyOperator) GetPriority() int {
	return p.priority
}

func (p MultiplyOperator) Cal(value ...string) (string, error) {
	if len(value) != 2 {
		return "", LackParamError
	}

	param, err := strconv.ParseFloat(value[0], 64)
	if err != nil {
		return "", err
	}

	param1, err := strconv.ParseFloat(value[1], 64)
	if err != nil {
		return "", err
	}

	return strconv.FormatFloat(param*param1, 'g', -1, 64), nil
}

type DivideOperator struct {
	priority int
}

func (p DivideOperator) GetPriority() int {
	return p.priority
}

func (p DivideOperator) Cal(value ...string) (string, error) {
	if len(value) != 2 {
		return "", LackParamError
	}

	if value[0] == "0" {
		return "", DivideZero
	}

	param, err := strconv.ParseFloat(value[0], 64)
	if err != nil {
		return "", err
	}

	param1, err := strconv.ParseFloat(value[1], 64)
	if err != nil {
		return "", err
	}

	return strconv.FormatFloat(param1/param, 'g', -1, 64), nil
}

type RemainderOperator struct {
	priority int
}

func (p RemainderOperator) GetPriority() int {
	return p.priority
}

func (p RemainderOperator) Cal(value ...string) (string, error) {
	if len(value) != 2 {
		return "", LackParamError
	}

	if value[0] == "0" {
		return "", DivideZero
	}

	param, err := strconv.ParseUint(value[0], 10, 64)
	if err != nil {
		return "", err
	}

	param1, err := strconv.ParseUint(value[1], 10, 64)
	if err != nil {
		return "", err
	}

	ret := param1 % param
	return strconv.FormatUint(ret, 10), nil
}

type PowerOperator struct {
	priority int
}

func (p PowerOperator) GetPriority() int {
	return p.priority
}

func (p PowerOperator) Cal(value ...string) (string, error) {
	if len(value) != 2 {
		return "", LackParamError
	}

	param1, err := strconv.ParseFloat(value[0], 64)
	if err != nil {
		return "", err
	}

	param, err := strconv.ParseUint(value[1], 10, 64)
	if err != nil {
		return "", err
	}

	var ret float64 = 1
	for i := 0; i < int(param); i++ {
		ret = param1 * ret
	}

	return strconv.FormatFloat(ret, 'g', -1, 64), nil
}

type ShiftLeftOperator struct {
	priority int
}

func (s ShiftLeftOperator) GetPriority() int {
	return s.priority
}

func (s ShiftLeftOperator) Cal(value ...string) (string, error) {
	if len(value) != 2 {
		return "", LackParamError
	}

	param1, err := strconv.ParseFloat(value[0], 64)
	if err != nil {
		return "", err
	}

	param, err := strconv.ParseUint(value[1], 10, 64)
	if err != nil {
		return "", err
	}

	for i := 0; i < int(param); i++ {
		param1 = param1 * 2
	}

	return strconv.FormatFloat(param1, 'g', -1, 64), nil
}

type ShiftRightOperator struct {
	priority int
}

func (s ShiftRightOperator) GetPriority() int {
	return s.priority
}

func (s ShiftRightOperator) Cal(value ...string) (string, error) {
	if len(value) != 2 {
		return "", LackParamError
	}

	param1, err := strconv.ParseFloat(value[0], 64)
	if err != nil {
		return "", err
	}

	param, err := strconv.ParseUint(value[1], 10, 64)
	if err != nil {
		return "", err
	}

	for i := 0; i < int(param); i++ {
		param1 = param1 / 2
	}

	return strconv.FormatFloat(param1, 'g', -1, 64), nil
}

type LessOperator struct {
	priority int
}

func (s LessOperator) GetPriority() int {
	return s.priority
}

func (s LessOperator) Cal(value ...string) (string, error) {
	if len(value) != 2 {
		return "", LackParamError
	}

	param1, err := strconv.ParseFloat(value[0], 64)
	if err != nil {
		return "", err
	}

	param, err := strconv.ParseFloat(value[1], 64)
	if err != nil {
		return "", err
	}

	if param > param1 {
		return "1", nil
	} else {
		return "0", nil
	}
}

type LessThanOperator struct {
	priority int
}

func (s LessThanOperator) GetPriority() int {
	return s.priority
}

func (s LessThanOperator) Cal(value ...string) (string, error) {
	if len(value) != 2 {
		return "", LackParamError
	}

	param1, err := strconv.ParseFloat(value[0], 64)
	if err != nil {
		return "", err
	}

	param, err := strconv.ParseFloat(value[1], 64)
	if err != nil {
		return "", err
	}

	if param >= param1 {
		return "1", nil
	} else {
		return "0", nil
	}
}

type MoreOperator struct {
	priority int
}

func (s MoreOperator) GetPriority() int {
	return s.priority
}

func (s MoreOperator) Cal(value ...string) (string, error) {
	if len(value) != 2 {
		return "", LackParamError
	}

	param1, err := strconv.ParseFloat(value[0], 64)
	if err != nil {
		return "", err
	}

	param, err := strconv.ParseFloat(value[1], 64)
	if err != nil {
		return "", err
	}

	if param < param1 {
		return "1", nil
	} else {
		return "0", nil
	}
}

type MoreThanOperator struct {
	priority int
}

func (s MoreThanOperator) GetPriority() int {
	return s.priority
}

func (s MoreThanOperator) Cal(value ...string) (string, error) {
	if len(value) != 2 {
		return "", LackParamError
	}

	param1, err := strconv.ParseFloat(value[0], 64)
	if err != nil {
		return "", err
	}

	param, err := strconv.ParseFloat(value[1], 64)
	if err != nil {
		return "", err
	}

	if param <= param1 {
		return "1", nil
	} else {
		return "0", nil
	}
}

type EqualOperator struct {
	priority int
}

func (s EqualOperator) GetPriority() int {
	return s.priority
}

func (s EqualOperator) Cal(value ...string) (string, error) {
	if len(value) != 2 {
		return "", LackParamError
	}

	param1, err := strconv.ParseFloat(value[0], 64)
	if err != nil {
		return "", err
	}

	param, err := strconv.ParseFloat(value[1], 64)
	if err != nil {
		return "", err
	}

	if param == param1 {
		return "1", nil
	} else {
		return "0", nil
	}
}

type NotEqualOperator struct {
	priority int
}

func (s NotEqualOperator) GetPriority() int {
	return s.priority
}

func (s NotEqualOperator) Cal(value ...string) (string, error) {
	if len(value) != 2 {
		return "", LackParamError
	}

	param1, err := strconv.ParseFloat(value[0], 64)
	if err != nil {
		return "", err
	}

	param, err := strconv.ParseFloat(value[1], 64)
	if err != nil {
		return "", err
	}

	if param != param1 {
		return "1", nil
	} else {
		return "0", nil
	}
}
