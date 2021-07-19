package expr

import (
	"log"
	"strconv"
	"unicode"
)

/*
	exp -> term { addop term }
	term -> factor { mulop factor }
	factor -> number | (+|-)number | (exp)
	addop -> + | -
	mulop -> * | / | %
*/

const (
	NumberToken = iota + 1
	OperatorToken
	EofToken
	LeftBracket
	RightBracket
	Variable
)

type Token struct {
	TokenType  int
	StringData string
	DoubleData float64
}

func (t *Token) IsAddOp() bool {
	return t.TokenType == OperatorToken && (t.StringData == "+" || t.StringData == "-")
}

func (t *Token) IsMulOp() bool {
	return t.TokenType == OperatorToken && (t.StringData == "*" || t.StringData == "/" || t.StringData == "%")
}

func Tokenize(input []rune) []Token {
	tokens := []Token{}
	index := 0
	for {
		switch {
		case index >= len(input):
			tokens = append(tokens, Token{
				TokenType: EofToken,
			})
			index++
			return tokens
		case unicode.IsSpace(input[index]):
			index++
			continue
		case unicode.IsNumber(input[index]):
			begin := index

			for ; index < len(input); index++ {
				if unicode.IsNumber(input[index]) || input[index] == '.' {
				} else {
					break
				}
			}

			number, err := strconv.ParseFloat(string(input[begin:index]), 64)
			if err != nil {
				log.Panic(err)
			}

			tokens = append(tokens, Token{
				TokenType:  NumberToken,
				DoubleData: number,
			})

		case input[index] == '+' || input[index] == '-' || input[index] == '*' || input[index] == '/' || input[index] == '%':
			tokens = append(tokens, Token{
				TokenType:  OperatorToken,
				StringData: string(input[index]),
			})
			index++
		case input[index] == '(':
			tokens = append(tokens, Token{
				TokenType: LeftBracket,
			})
			index++
		case input[index] == ')':
			tokens = append(tokens, Token{
				TokenType: RightBracket,
			})
			index++
		case unicode.IsLetter(input[index]):
			begin := index

			for ; index < len(input); index++ {
				if unicode.IsLetter(input[index]) {
				} else {
					break
				}
			}

			tokens = append(tokens, Token{
				TokenType:  Variable,
				StringData: string(input[begin:index]),
			})
		default:
			index++
		}
	}
}

type Expression struct {
	Input     string
	Tokens    []Token
	Index     int
	Variables map[string]float64
}

func NewExpression(input string) *Expression {
	return &Expression{
		Input: input,
		Index: 0,
	}
}

func (e *Expression) Eval() float64 {
	if len(e.Input) == 0 {
		return 0
	}

	e.Index = 0
	return e.expr()
}

func (e *Expression) Tokenize() {
	e.Tokens = Tokenize([]rune(e.Input))
}

func (e *Expression) expr() float64 {
	result := e.term()

	for e.Tokens[e.Index].IsAddOp() {
		oper := e.Tokens[e.Index].StringData
		if oper == "+" {
			e.Index++
			result += e.term()
		} else if oper == "-" {
			e.Index++
			result -= e.term()
		} else {
			break
		}
	}

	return result
}

func (e *Expression) factory() float64 {
	var result float64

	if e.Tokens[e.Index].TokenType == LeftBracket {
		e.Index++
		result = e.expr()
		if e.Tokens[e.Index].TokenType != RightBracket {
			log.Fatal("expect right bracket")
		}
	} else if e.Tokens[e.Index].StringData == "-" {
		e.Index++
		result = -(e.Tokens[e.Index].DoubleData)
	} else if e.Tokens[e.Index].StringData == "+" {
		e.Index++
		result = e.Tokens[e.Index].DoubleData
	} else if e.Tokens[e.Index].TokenType == Variable {
		result = e.Variables[e.Tokens[e.Index].StringData]
	} else {
		result = e.Tokens[e.Index].DoubleData
	}
	e.Index++

	return result
}

func (e *Expression) term() float64 {
	result := e.factory()

	for e.Tokens[e.Index].IsMulOp() {
		oper := e.Tokens[e.Index].StringData
		if oper == "*" {
			e.Index++
			result *= e.factory()
		} else if oper == "/" {
			e.Index++
			result = result / e.factory()
		} else if oper == "%" {
			e.Index++
			result = float64(int64(result) % int64(e.factory()))
		} else {
			break
		}
	}

	return result
}
