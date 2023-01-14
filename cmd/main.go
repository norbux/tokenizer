package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
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
	var inputFile string
	flag.StringVar(&inputFile, "f", "", "Input file")
	flag.Parse()

	input, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}

	defer input.Close()

	reader := bufio.NewReader(input)
	buf := make([]byte, 1)

	var set []string
	var id []byte

	for {
		_, err := reader.Read(buf)
		if err != nil {
			if err != io.EOF {
				log.Println(err)
			}
			break
		}
		// Detect spaces and finalizers
		if buf[0] == sp || buf[0] == lb || buf[0] == fin {
			if string(id) != "" {
				set = append(set, strings.TrimSpace(string(id)))
				id = []byte{}
				continue
			}
		}

		// Detect operators as identifiers
		if strings.Contains(ops, string(buf[0])) {
			set = append(set, strings.TrimSpace(string(buf[0])))
			id = []byte{}
			continue
		}

		id = append(id, buf[0])
	}

	fmt.Println(set)
}
