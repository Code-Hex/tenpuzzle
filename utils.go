package tenpuzzle

import (
	"crypto/rand"
	"log"
	"math/big"
	"unicode"
)

func pop(stack []int) (int, []int) {
	return stack[len(stack)-1], stack[:len(stack)-1]
}

// ランダムに演算子を生成
func opgen() rune {
	return op[random(4)]
}

// ランダムに数値を生成
func numgen() rune {
	return num[random(9)]
}

func isOp(c rune) bool {
	return !unicode.IsDigit(c)
}

func choiseTree(i int) (t *Tree) {
	switch i {
	case 0:
		t = buildTreeA()
	case 1:
		t = buildTreeB()
	case 2:
		t = buildTreeC()
	case 3:
		t = buildTreeD()
	case 4:
		t = buildTreeE()
	}

	return
}

func random(max int) int {
	n, err := rand.Int(rand.Reader, big.NewInt(int64(max)))
	if err != nil {
		log.Println(err)
	}
	return int(n.Int64())
}

// for sort package
func (t Trees) Len() int {
	return len(t)
}

func (t Trees) Less(i, j int) bool {
	return Fitness(t[i]) > Fitness(t[j])
}

func (t Trees) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}
