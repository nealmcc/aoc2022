package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func() { file.Close() }()

	s := bufio.NewScanner(file)
	lines := make([]string, 0, 102)
	for s.Scan() {
		lines = append(lines, s.Text())
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}

	start := time.Now()
	sum := 0
	for _, s := range lines {
		sum += SNAFUToInt(s)
	}
	p1 := IntToSNAFU(sum)

	end := time.Now()

	fmt.Printf("part 1: %s in %s\n", p1, end.Sub(start))
}

func IntToSNAFU(n int) string {
	if n == 0 {
		return "0"
	}

	symbols := []byte{}
	var rem int
	for n != 0 {
		n, rem = n/5, n%5
		if rem > 2 {
			rem -= 5
			n += 1
		}
		switch rem {
		case -2:
			symbols = append(symbols, '=')
		case -1:
			symbols = append(symbols, '-')
		default:
			symbols = append(symbols, '0'+byte(rem))
		}
	}

	for i, j := 0, len(symbols)-1; i < j; i, j = i+1, j-1 {
		symbols[i], symbols[j] = symbols[j], symbols[i]
	}
	return string(symbols)
}

func SNAFUToInt(s string) int {
	var n int
	for i, x := len(s)-1, 1; i >= 0; i-- {
		switch s[i] {
		case '2', '1', '0':
			n += int('2'-s[i]) * x
		case '-':
			n += -1 * x
		case '=':
			n += -2 * x
		default:
			// ignore anything else
			continue
		}
		x *= 5
	}
	return n
}
