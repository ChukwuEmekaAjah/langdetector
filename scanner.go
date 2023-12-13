// Package langdetector uses a scanner configuration to determine if the provided
// source code is of the specified type in the scanner used to parse it.
// It calculates the likelihood of a source code being same as a provided
// scanner language based on the number of errors generated while scanning
// the source code and other factors like number of
package langdetector

import (
	"fmt"
	"os"
	"regexp"
	"slices"
)

// LangSpec specifies unique characteristics of the language being evaluated
type LangSpec struct {
	language          string
	keywords          []string
	identifierPattern string
}

// Scanner holds the config for each programming language scanner
// It accepts the language's keywords, identifier regex etc.
type Scanner struct {
	keywords          []string
	identifierPattern string
	source            string
	line              uint
	current           uint
	start             uint
	errors            []error
	language          string
	tokens            []Token
	score             int
}

// Token captures the details of a token
type Token struct {
	tokenType string
	literal   string
	lexeme    string
}

// NewScanner creates a new instance of scanner
func NewScanner(spec LangSpec) *Scanner {
	s := Scanner{
		language: spec.language,
		errors:   make([]error, 0),
		keywords: spec.keywords,
	}
	if spec.identifierPattern != "" {
		s.identifierPattern = spec.language
	} else {
		s.identifierPattern = "^[a-zA-Z_$]+[0-9a-zA-Z_$]*"
	}
	return &s
}

// ScanTokens reads the characters in source text and converts them to tokens
func (s *Scanner) ScanTokens() []Token {
	for !s.isAtEnd() {
		s.start = s.current
		ch := s.advance()

		switch ch {
		case '(':
			s.addToken("LEFT_PAREN", string(ch), string(ch))
		case ')':
			s.addToken("RIGHT_PAREN", string(ch), string(ch))
		case '{':
			s.addToken("LEFT_BRACE", string(ch), string(ch))
		case '}':
			s.addToken("RIGHT_BRACE", string(ch), string(ch))
		case ',':
			s.addToken("COMMA", string(ch), string(ch))
		case ';':
			s.addToken("SEMI_COLON", string(ch), string(ch))
		case '.':
			s.addToken("DOT", string(ch), string(ch))
		case '!':
			if s.match('=') {
				s.addToken("BANG_EQUAL", string(s.source[s.start:s.current]), string(s.source[s.start:s.current]))
			} else {
				s.addToken("BANG", string(ch), string(ch))
			}
		case '=':
			if s.match('=') {
				s.addToken("EQUAL_EQUAL", string(s.source[s.start:s.current]), string(s.source[s.start:s.current]))
			} else {
				s.addToken("ASSIGN", string(ch), string(ch))
			}
		case '<':
			if s.match('=') {
				s.addToken("LESS_THAN_EQUAL", string(s.source[s.start:s.current]), string(s.source[s.start:s.current]))
			} else {
				s.addToken("LESS_THAN", string(ch), string(ch))
			}
		case '>':
			if s.match('=') {
				s.addToken("GREATER_THAN_EQUAL", string(s.source[s.start:s.current]), string(s.source[s.start:s.current]))
			} else {
				s.addToken("GREATER_THAN", string(ch), string(ch))
			}
		case '*':
			s.addToken("STAR", string(ch), string(ch))
		case '/':
			s.addToken("SLASH", string(ch), string(ch))
		case '+':
			if s.match('+') {
				s.addToken("PLUS_PLUS", string(s.source[s.start:s.current]), string(s.source[s.start:s.current]))
			} else if s.match('=') {
				s.addToken("PLUS_ASSIGN", string(s.source[s.start:s.current]), string(s.source[s.start:s.current]))
			} else {
				s.addToken("PLUS", string(ch), string(ch))
			}
		case '-':
			if s.match('-') {
				s.addToken("MINUS_MINUS", string(s.source[s.start:s.current]), string(s.source[s.start:s.current]))
			} else if s.match('=') {
				s.addToken("MINUS_ASSIGN", string(s.source[s.start:s.current]), string(s.source[s.start:s.current]))
			} else {
				s.addToken("MINUS", string(ch), string(ch))
			}
		case '%':
			if s.match('=') {
				s.addToken("MODULO_ASSIGN", string(s.source[s.start:s.current]), string(s.source[s.start:s.current]))
			} else {
				s.addToken("MODULO", string(ch), string(ch))
			}
		case '\t':
		case ' ':
		case '\n':
			s.line++
		case '\r':
		case '"':
			s.string()
		case '\'':
			s.string()
		default:
			if s.isNumber(ch) {
				s.number()
			} else if s.isAlpha(ch) || s.isSpecialChar(ch) || s.isNumber(ch) {
				s.identifier()
			} else {
				s.errors = append(s.errors, fmt.Errorf("Invalid character %s in source code", string(ch)))
			}
		}
	}
	return s.tokens
}

func (s *Scanner) addToken(tokenType, lexeme, literal string) {
	s.tokens = append(s.tokens, Token{tokenType: tokenType, literal: literal, lexeme: lexeme})
}

func (s Scanner) isSpecialChar(ch byte) bool {
	if (ch == '_') || (ch == '$') {
		return true
	}
	return false
}

func (s Scanner) isAlpha(ch byte) bool {
	if (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') {
		return true
	}
	return false
}

func (s *Scanner) identifier() {
	for s.isAlpha(s.peek(0)) || s.isSpecialChar(s.peek(0)) || s.isNumber(s.peek(0)) {
		s.advance()
	}

	identifier := s.source[s.start:s.current]

	if slices.Contains(s.keywords, identifier) {
		s.addToken("KEYWORD", identifier, identifier)
		s.Score()
	} else if ok, _ := regexp.MatchString(s.identifierPattern, identifier); ok {
		s.addToken("IDENTIFIER", identifier, identifier)
	} else {
		s.errors = append(s.errors, fmt.Errorf("Invalid identifier %s in source code", identifier))
	}
}

func (s Scanner) isNumber(ch byte) bool {
	return ch >= '0' && ch <= '9'
}

func (s *Scanner) number() {
	for s.isNumber(s.peek(0)) {
		s.advance()
	}

	// checking here so we don't catch multiple dots in a number
	if s.peek(0) == '.' && s.isNumber(s.peek(1)) && !s.isAtEnd() {
		s.advance()
		for s.isNumber(s.peek(0)) {
			s.advance()
		}
	}
	s.addToken("NUMBER", s.source[s.start:s.current], s.source[s.start:s.current])
}

func (s *Scanner) string() {
	// we use peek instead of match because we want to isolate the errors
	// for loop can fail because of it's at end or there is no closing quotes
	for s.peek(0) != '"' && !s.isAtEnd() {
		if s.peek(0) == '\n' {
			s.line++
		}
		s.advance()
	}

	if s.isAtEnd() {
		s.errors = append(s.errors, fmt.Errorf("Invalid string without closing quotes"))
		return
	}
	s.advance()

	s.addToken("STRING", s.source[s.start+1:s.current-1], s.source[s.start+1:s.current-1])
}

func (s *Scanner) advance() byte {

	if !s.isAtEnd() {
		s.current++
		return s.source[s.current-1]
	}
	return 0
}

func (s *Scanner) match(ch byte) bool {
	if s.isAtEnd() {
		return false
	}

	if ch != s.source[s.current] {
		return false
	}

	s.current++
	return true
}

func (s Scanner) peek(position uint) byte {
	if s.isAtEnd() {
		return byte(0)
	}
	return s.source[s.current+position]
}

func (s Scanner) isAtEnd() bool {
	if s.current >= uint(len(s.source)) {
		return true
	}
	return false
}

// WithKeywords adds keys that the language uses
func (s *Scanner) WithKeywords(keywords []string) *Scanner {
	s.keywords = append(s.keywords, keywords...)
	return s
}

// WithSource sets the source code it will scan through
func (s *Scanner) WithSource(source string) *Scanner {
	s.source = source
	return s
}

// WithSourcefile reads the source file and uses the source data as input
func (s *Scanner) WithSourcefile(sourceFile string) (*Scanner, error) {
	content, err := os.ReadFile(sourceFile)
	if err != nil {
		return nil, err
	}

	return s.WithSource(string(content)), err
}

// Score increments the score for the language
func (s *Scanner) Score() {
	s.score++
}

// GetScore returns the score of the scanner from the language spec
func (s Scanner) GetScore() int {
	return s.score
}
