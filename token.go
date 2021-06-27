package expr

import (
	"unicode"
)

const (
	tokNumber = 1 << iota
	tokWord
	tokOp
	tokOpen
	tokClose
)

func tokenize(input []rune) (tokens []string, err error) {
	pos := 0
	begin := 0

	expected := tokOpen | tokNumber | tokWord
	for pos < len(input) {
		c := input[pos]
		if unicode.IsSpace(c) {
			pos++
			continue
		}

		begin = pos

		switch {
		case unicode.IsNumber(c):
			if expected&tokNumber == 0 {
				return nil, ErrUnexpectedNumber
			}

			expected = tokOp | tokClose
			for (c == '.' || unicode.IsNumber(c)) && pos < len(input) {
				pos++
				if pos >= len(input) {
					break
				}

				c = input[pos]
			}
		case unicode.IsLetter(c):
			if expected&tokWord == 0 {
				return nil, ErrUnexpectedIdentifier
			}

			expected = tokOp | tokOpen | tokClose
			for (unicode.IsLetter(c) || unicode.IsNumber(c) || c == '_') && pos < len(input) {
				pos++

				if pos >= len(input) {
					break
				}
				c = input[pos]
			}
		case c == '@':
			if expected&tokWord == 0 {
				return nil, ErrUnexpectedIdentifier
			}

			if pos+2 >= len(input) {
				return nil, ErrUnexpectedIdentifier
			}

			if input[pos+1] != '.' {
				return nil, ErrUnexpectedIdentifier
			}

			c = input[pos+2]
			pos = pos + 2

			expected = tokOp | tokOpen | tokClose
			for (unicode.IsLetter(c) || unicode.IsNumber(c) || c == '_') && pos < len(input) {
				pos++
				if pos >= len(input) {
					break
				}
				c = input[pos]
			}
		case c == '(' || c == ')':
			pos++
			if c == '(' && (expected&tokOpen) != 0 {
				expected = tokNumber | tokWord | tokOpen | tokClose
			} else if c == ')' && (expected&tokClose) != 0 {
				expected = tokOp | tokClose
			} else {
				return nil, ErrParen
			}
		default:
			if expected&tokOp == 0 {
				if c != '-' && c != '^' && c != '!' {
					return nil, ErrOperandMissing
				}
				pos++
			} else {
				lastOp := string(input[begin:pos])
				for !unicode.IsLetter(c) && !unicode.IsNumber(c) && !unicode.IsSpace(c) &&
					c != '_' && c != '(' && c != ')' && pos < len(input) {

					if _, ok := ops[string(input[begin:pos])+string(input[pos])]; ok {
						lastOp = lastOp + string(input[pos])
					} else {
						break
					}

					pos++
					if pos >= len(input) {
						break
					}

					c = input[pos]
				}

				if lastOp == "" {
					return nil, ErrBadOp
				}
			}
			expected = tokNumber | tokWord | tokOpen
		}
		tokens = append(tokens, string(input[begin:pos]))
	}
	return tokens, nil
}
