package main

import (
	"bytes"
	"fmt"
	"strings"
)

const (
	eq  byte   = '='
	pl  byte   = '+'
	mn  byte   = '-'
	ml  byte   = '*'
	dv  byte   = '/'
	sp  byte   = ' '
	lb  byte   = 10
	fin byte   = ';'
	ops string = "=+-*/%"
)

func main() {
	input := `x
y
some = test
z
`

	src := []byte(input)
	fmt.Println(src)

	fields := bytes.Fields(src)
	fmt.Println(fields)

	var set []string
	var id []byte

	for _, t := range src {
		// Detect spaces and finalizers
		if t == sp || t == lb || t == fin {
			if string(id) != "" {
				set = append(set, strings.TrimSpace(string(id)))
				id = []byte{}
				continue
			}
		}

		// Detect operators as identifiers
		if strings.Contains(ops, string(t)) {
			set = append(set, strings.TrimSpace(string(t)))
			id = []byte{}
			continue
		}

		id = append(id, t)
	}

	fmt.Println(set)
}
