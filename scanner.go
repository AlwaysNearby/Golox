package golox

import (
	"strconv"
)

type Scanner struct {
	source  string
	tokens  []Token
	start   int
	current int
	line    int
}

var keywords map[string]TokenType

func NewScanner(source string) *Scanner {
	keywords = map[string]TokenType{
		"and":    AND,
		"class":  CLASS,
		"else":   ELSE,
		"false":  FALSE,
		"for":    FOR,
		"fun":    FUN,
		"if":     IF,
		"nil":    NIL,
		"or":     OR,
		"print":  PRINT,
		"return": RETURN,
		"super":  SUPER,
		"this":   THIS,
		"true":   TRUE,
		"var":    VAR,
		"while":  WHILE,
	}

	return &Scanner{
		source: source,
		line:   1,
	}
}

func (s *Scanner) ScanTokens() []Token {
	for !s.isAtEnd() {
		s.start = s.current
		s.scanToken()
	}

	s.tokens = append(s.tokens, Token{Type: EOF, Lexeme: "", Literal: nil, Line: s.line})
	return s.tokens
}

func (s *Scanner) isAtEnd() bool {
	return s.current >= len(s.source)
}

func (s *Scanner) scanToken() {
	switch ch := s.advance(); ch {
	case '(':
		s.addSimpleToken(LEFT_PAREN)
		break
	case ')':
		s.addSimpleToken(RIGHT_PAREN)
		break
	case '{':
		s.addSimpleToken(LEFT_BRACE)
		break
	case '}':
		s.addSimpleToken(RIGHT_BRACE)
		break
	case ',':
		s.addSimpleToken(COMMA)
		break
	case '.':
		s.addSimpleToken(DOT)
		break
	case '-':
		s.addSimpleToken(MINUS)
		break
	case '+':
		s.addSimpleToken(PLUS)
		break
	case ';':
		s.addSimpleToken(SEMICOLON)
		break
	case '*':
		s.addSimpleToken(STAR)
		break
	case '!':
		s.addSimpleToken(getExpectedToken(s.match('='), BANG_EQUAL, BANG))
		break
	case '=':
		s.addSimpleToken(getExpectedToken(s.match('='), EQUAL_EQUAL, EQUAL))
		break
	case '<':
		s.addSimpleToken(getExpectedToken(s.match('='), LESS_EQUAL, LESS))
		break
	case '>':
		s.addSimpleToken(getExpectedToken(s.match('='), GREATER_EQUAL, GREATER))
		break
	case '/':
		if s.match('/') {
			for s.peek() != '\n' && !s.isAtEnd() {
				s.advance()
			}
		} else if s.match('*') {
			openComments := []int{s.line}

			for len(openComments) > 0 && !s.isAtEnd() {
				if s.match('/') && s.match('*') {
					openComments = append(openComments, s.line)
				} else if s.match('*') && s.match('/') {
					openComments = openComments[:len(openComments)-1]
					continue
				}

				if s.peek() == '\n' {
					s.line++
				}

				s.advance()
			}

			if len(openComments) > 0 {
				for _, line := range openComments {
					sendError(line, "Blockcomment was never closed, no \"*/\" found")
				}
			}

		} else {
			s.addSimpleToken(SLASH)
		}
		break
	case ' ':
		break
	case '\r':
		break
	case '\t':
		break
	case '\n':
		s.line++
		break
	case '"':
		s.scanString()
		break
	default:
		if isDigit(ch) {
			s.scanNumber()
		} else if isAlpha(ch) {
			s.scanIdentifier()
		} else {
			sendError(s.line, "Unexpected characher.")
		}
		break
	}
}

func isAlpha(ch byte) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') && (ch == '_')
}

func isDigit(ch byte) bool {
	return ch >= '0' && ch <= '9'
}

func isAlphaNumeric(ch byte) bool {
	return isAlpha(ch) || isDigit(ch)
}

func getExpectedToken(isMatch bool, matchedToken, unmatchedToken TokenType) TokenType {
	if isMatch {
		return matchedToken
	}

	return unmatchedToken
}

func (s *Scanner) scanIdentifier() {
	for isAlphaNumeric(s.peek()) {
		s.advance()
	}

	text := s.source[s.start:s.current]
	if val, ok := keywords[text]; ok {
		s.addSimpleToken(val)
	} else {
		s.addSimpleToken(IDENTIFIER)
	}
}

func (s *Scanner) scanNumber() {
	for isDigit(s.peek()) {
		s.advance()
	}

	if s.peek() == '.' && isDigit(s.peekNext()) {
		s.advance()

		for isDigit(s.peek()) {
			s.advance()
		}
	}

	value, _ := strconv.ParseFloat(s.source[s.start:s.current], 64)

	s.addToken(NUMBER, value)
}

func (s *Scanner) scanString() {
	for s.peek() != '"' && !s.isAtEnd() {
		if s.peek() == '\n' {
			s.line++
		}

		s.advance()
	}

	if s.isAtEnd() {
		sendError(s.line, "Unterminated string.")
		return
	}

	s.advance()

	value := s.source[s.start+1 : s.current-1]
	s.addToken(STRING, value)
}

func (s *Scanner) peek() byte {
	if s.isAtEnd() {
		return 0
	}

	return s.source[s.current]
}

func (s *Scanner) peekNext() byte {
	if s.current+1 >= len(s.source) {
		return 0
	}

	return s.source[s.current+1]
}

func (s *Scanner) match(expected byte) bool {
	if s.isAtEnd() {
		return false
	}

	if s.source[s.current] != expected {
		return false
	}

	s.current++
	return true
}

func (s *Scanner) advance() byte {
	ch := s.source[s.current]
	s.current++
	return ch
}

func (s *Scanner) addSimpleToken(tokenType TokenType) {
	s.addToken(tokenType, nil)
}

func (s *Scanner) addToken(tokenType TokenType, literal interface{}) {
	text := s.source[s.start:s.current]
	s.tokens = append(s.tokens, Token{Type: tokenType, Lexeme: text, Literal: literal, Line: s.line})
}
