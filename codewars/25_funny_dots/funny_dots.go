package funny_dots

import (
	"fmt"
	"strings"
)

const (
	IntersectionPoint = "+"
	HorizontalLine    = "---"
	VerticalLine      = "|"
	Space             = " "
	Content           = "o"
	Newline           = "\n"
)

// Average time complexity: O(n)
// Worst time complexity:   O(n)
// Space complexity:        O(1)
func Dot(n, m int) string {
	cellLine := fmt.Sprint(IntersectionPoint, HorizontalLine)
	rowLine := fmt.Sprint(strings.Repeat(cellLine, n), IntersectionPoint)
	cell := fmt.Sprint(VerticalLine, Space, Content, Space)
	row := fmt.Sprint(strings.Repeat(cell, n), VerticalLine)
	var canvas strings.Builder

	for ctr := m; ctr > 0; ctr -= 1 {
		canvas.WriteString(rowLine)
		canvas.WriteString(Newline)
		canvas.WriteString(row)
		canvas.WriteString(Newline)
	}

	canvas.WriteString(rowLine)

	return canvas.String()
}
