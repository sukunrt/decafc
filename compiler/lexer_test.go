package main

import (
	"fmt"
	"slices"
	"strings"
	"testing"
)

func TestLexer(t *testing.T) {
	newToken := func(typ Type, l, c int, v string) token {
		return token{
			Type:   typ,
			Line:   l,
			Column: c,
			Value:  v,
		}
	}
	tc := []struct {
		input  string
		output []token
	}{
		{
			input: "+ + +\n + + +",
			output: []token{
				newToken(TypeOp, 1, 1, "+"),
				newToken(TypeOp, 1, 3, "+"),
				newToken(TypeOp, 1, 5, "+"),
				newToken(TypeOp, 2, 2, "+"),
				newToken(TypeOp, 2, 4, "+"),
				newToken(TypeOp, 2, 6, "+"),
			},
		},
	}
	for i, c := range tc {
		t.Run(fmt.Sprintf("test-%d", i), func(t *testing.T) {
			r := strings.NewReader(c.input)
			l := NewLexer(r)
			var output []token
			for {
				t := l.Pop()
				if t.Type == TypeUnknown {
					break
				}
				output = append(output, t)
			}
			if !slices.Equal(output, c.output) {
				t.Fatalf("unequal outputs\n%v\n%v", output, c.output)
			}
		})
	}
}
