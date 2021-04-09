package expr

import "unicode"

const (
	tokNumber = 1 << iota
	tokWord
	tokOp
	tokOpen
	tokClose
)

func tokenize(input []rune) (tokens []string, err error) {
	pos := 0
	expected := tokOpen | tokNumber | tokWord
	for pos < len(input) {
		tok := []rune{}
		c := input[pos]
		if unicode.IsSpace(c) {
			pos++
			continue
		}

		if unicode.IsNumber(c) {
			if expected&tokNumber == 0 {
				return nil, ErrUnexpectedNumber
			}

			expected = tokOp | tokClose
			for (c == '.' || unicode.IsNumber(c)) && pos < len(input) {
				tok = append(tok, input[pos])
				pos++
				if pos < len(input) {
					c = input[pos]
				} else {
					c = 0
				}
			}
		} else if unicode.IsLetter(c) {
			if expected&tokWord == 0 {
				return nil, ErrUnexpectedIdentifier
			}
			expected = tokOp | tokOpen | tokClose
			for (unicode.IsLetter(c) || unicode.IsNumber(c) || c == '_') && pos < len(input) {
				tok = append(tok, input[pos])
				pos++
				if pos < len(input) {
					c = input[pos]
				} else {
					c = 0
				}
			}
		} else if c == '@' {
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
				tok = append(tok, input[pos])
				pos++
				if pos < len(input) {
					c = input[pos]
				} else {
					c = 0
				}
			}
		} else if c == '(' || c == ')' {
			tok = append(tok, c)
			pos++
			if c == '(' && (expected&tokOpen) != 0 {
				expected = tokNumber | tokWord | tokOpen | tokClose
			} else if c == ')' && (expected&tokClose) != 0 {
				expected = tokOp | tokClose
			} else {
				return nil, ErrParen
			}
		} else {
			if expected&tokOp == 0 {
				if c != '-' && c != '^' && c != '!' {
					return nil, ErrOperandMissing
				}
				tok = append(tok, c, 'u')
				pos++
			} else {
				var lastOp string
				for !unicode.IsLetter(c) && !unicode.IsNumber(c) && !unicode.IsSpace(c) &&
					c != '_' && c != '(' && c != ')' && pos < len(input) {
					if _, ok := ops[string(tok)+string(input[pos])]; ok {
						tok = append(tok, input[pos])
						lastOp = string(tok)
					} else if lastOp == "" {
						tok = append(tok, input[pos])
					} else {
						break
					}
					pos++
					if pos < len(input) {
						c = input[pos]
					} else {
						c = 0
					}
				}
				if lastOp == "" {
					return nil, ErrBadOp
				}
			}
			expected = tokNumber | tokWord | tokOpen
		}
		tokens = append(tokens, string(tok))
	}
	return tokens, nil
}
