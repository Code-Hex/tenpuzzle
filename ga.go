package tenpuzzle

import (
	"fmt"
	"math"
	"sort"
	"time"
)

var formula map[string]bool = map[string]bool{}

func Run(num int) {
	// 個体を作る (num > 100 以上 && 偶数)
	trees := GAInit(num)
	for {
		// 優秀な順でソート(降順)
		sort.Sort(trees)
		Enumerate(trees)
		trees = Kill(trees)
		trees = Crossing(trees)
		trees = Mutation(trees, num)
		time.Sleep(100 * time.Millisecond)
	}
}

// 適当に生成
func GAInit(init_num int) (trees Trees) {
	for i := 0; i < init_num; i++ {
		trees = append(trees, choiseTree(i%5))
	}
	return
}

func Enumerate(trees Trees) {
	for i := 0; Fitness(trees[i]) == 1 && i < len(trees)-1; i++ {
		f := Formula(trees[i], true)
		if _, ok := formula[f]; !ok {
			fmt.Println(f)
		}
		formula[f] = true
	}
}

// 個体の半分を死滅させる
func Kill(trees Trees) Trees {
	return trees[len(trees)/2:]
}

// 交叉
// Trees == []*Tree
// KilledTree / 2 の数作成
// if KilledTree == 50 then 25
func Crossing(trees Trees) Trees {
	// 木の深さ最大3
	max := 3
	cpln := len(trees) / 2
	_ts := make(Trees, cpln)
	copy(_ts, trees)

	for i := 0; i+1 < cpln; i += 2 {
		// 0 が一番の親なので, 含ませてはいけないから rand.Int()%(max-1)+1
		t_base := candidateTree(_ts[i], random(max-1)+1)
		t := findSameTree(*t_base, _ts[i+1])
		if t != nil {
			swap(*t_base, *t)
			trees = append(trees, _ts[i], _ts[i+1])
		}
	}

	return trees
}

// 突然変異
func Mutation(trees Trees, num int) Trees {
	// 木の深さ最大3
	max := 3
	size := num - len(trees)
	half := size / 2

	// 演算子の突然変異
	mtrees1 := make(Trees, half)
	copy(mtrees1, trees)
	for i := 0; i < half; i++ {
		t := candidateTree(mtrees1[i], random(max))
		renewal(*t)
	}

	// 数値の突然変異
	mtrees2 := make(Trees, size-half)
	copy(mtrees2, trees)
	for i := 0; i < size-half; i++ {
		t := candidateTree(mtrees2[i], max)
		renewal(*t)
	}

	trees = append(trees, mtrees1...)
	return append(trees, mtrees2...)
}

func renewal(t *Tree) {
	if isOp(t.Val) {
		// 演算子の突然変異
		_op := t.Val
		for _op == t.Val {
			t.Val = op[random(4)]
		}
	} else {
		// 数値の突然変異
		_num := t.Val
		for _num == t.Val {
			t.Val = num[random(9)]
		}
	}
}

// 適応度関数
// 10 から遠ければ遠いほど優秀
// つまりこの値が 0 に近ければ近いほど良い
func Fitness(t *Tree) float64 {
	return 1.0 / (1.0 + math.Abs(Calc(t)-10.0))
}
