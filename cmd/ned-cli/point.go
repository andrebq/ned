package main

import "fmt"

func newPoint(width, height int) point {
	return point{line: height, col: width}
}

func (p point) String() string {
	return fmt.Sprintf("[line: %v / col: %v]", p.line, p.col)
}
