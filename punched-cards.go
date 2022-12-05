# https://codingcompetitions.withgoogle.com/codejam/round/0000000000876ff1/0000000000a4621b
// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	row, col := 2, 3
	m := makeMatrix(count(row), count(col), '.')
	for r, row := range m {
		for c := range row {
			if r == 0 || r%2 == 0 {
				if c == 0 || c%2 == 0 {
					m[r][c] = '+'
				} else {
					m[r][c] = '-'
				}
			} else {
				if c == 0 || c%2 == 0 {
					m[r][c] = '|'
				}
			}
		}
	}
	m[0][0] = '.'
	m[0][1] = '.'
	m[1][0] = '.'

	for _, row := range m {
		fmt.Println(string(row))
	}
}

func makeMatrix(row, col int, val rune) [][]rune {
	m := make([][]rune, row)
	for r := 0; r < row; r++ {
		m[r] = make([]rune, col)
		for c := 0; c < col; c++ {
			m[r][c] = val
		}
	}
	return m
}

func count(n int) int {
	return (n * 3) - (n - 1)
}
