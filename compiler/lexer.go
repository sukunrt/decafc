package main

import (
	"bufio"
	"io"
)

type Type int

const (
	TypeUnknown Type = iota
	TypeOp
)

type token struct {
	Type         Type
	Line, Column int
	Value        string
	Error        error
}

type lexer struct {
	// s is the scanner for the source
	// s is at the beginning of the next token
	s *bufio.Scanner
	// line and column of the next token in s
	line, column int
	next         []token
}

func NewLexer(reader io.Reader) *lexer {
	s := bufio.NewScanner(reader)
	s.Split(bufio.ScanRunes)
	s.Scan()
	return &lexer{s: s, line: 1, column: 1}
}

func (l *lexer) Peek() token {
	if len(l.next) == 0 {
		l.advance()
	}
	if len(l.next) == 0 {
		return token{}
	}
	return l.next[len(l.next)-1]
}

func (l *lexer) Pop() token {
	if len(l.next) == 0 {
		l.advance()
	}
	if len(l.next) == 0 {
		return token{}
	}
	t := l.next[len(l.next)-1]
	l.next = l.next[:len(l.next)-1]
	return t
}

func (l *lexer) advance() {
LOOP:
	for {
		switch l.s.Text() {
		case "\n":
			// treat new line as the zeroth character of the next line
			// next Scan call will read the first character
			// in to column 1
			l.line++
			l.column = 0
		case " ", "\t":
		default:
			break LOOP
		}
		l.s.Scan()
		l.column++
	}

	advanceScanner := true
	defer func() {
		if advanceScanner {
			l.s.Scan()
			l.column++
		}
	}()

	txt := l.s.Text()
	var tp Type
	var value string
	switch txt {
	case "+":
		value = txt
		tp = TypeOp
	default:
	}
	l.next = append(l.next, token{Type: tp, Line: l.line, Column: l.column, Value: value})
}
