package tenpuzzle

import "fmt"

func Formula(t *Tree, isfirst bool) string {
	if t.Val == '*' || t.Val == '/' {
		return fmt.Sprintf("%s %c %s", Formula(t.Left, false), t.Val, Formula(t.Right, false))
	}

	if t.Val == '+' || t.Val == '-' {
		if isfirst {
			return fmt.Sprintf("%s %c %s", Formula(t.Left, false), t.Val, Formula(t.Right, false))
		}
		return fmt.Sprintf("(%s %c %s)", Formula(t.Left, false), t.Val, Formula(t.Right, false))
	}

	return string(t.Val)
}

func Calc(t *Tree) float64 {
	if t.Val == '*' {
		return Calc(t.Left) * Calc(t.Right)
	}

	if t.Val == '+' {
		return Calc(t.Left) + Calc(t.Right)
	}

	if t.Val == '-' {
		return Calc(t.Left) - Calc(t.Right)
	}

	if t.Val == '/' {
		right := Calc(t.Right)
		// for division error
		if right == 0 {
			return 0
		}
		return Calc(t.Left) / right
	}

	return float64(t.Val - '0')
}
