package tenpuzzle

type Tree struct {
	Left  *Tree
	Right *Tree
	Val   rune // [+-*/] | [1-9]
}

type Trees []*Tree

var (
	op  = []rune{'*', '+', '-', '/'}
	num = []rune{'1', '2', '3', '4', '5', '6', '7', '8', '9'}
)

// 完成

// NewTree('/')
func NewTree(c rune) *Tree {
	return &Tree{nil, nil, c}
}

func findSameTree(t_base, t *Tree) **Tree {
	if !isOp(t.Val) {
		if isSameTree(t_base, t) {
			return &t
		}
		return nil
	}

	if left := findSameTree(t_base, t.Left); left != nil {
		return left
	}

	if right := findSameTree(t_base, t.Right); right != nil {
		return right
	}

	if isSameTree(t_base, t) {
		return &t
	}

	return nil
}

func candidateTree(t *Tree, depth int) **Tree {
	directions := make([]int, depth)
	for i := 0; i < depth; i++ {
		// 0: left, 1: right
		directions[i] = random(2)
	}

	return pickTree(directions, t)
}

func pickTree(directions []int, t *Tree) **Tree {
	if len(directions) == 0 || !isOp(t.Val) {
		return &t
	}
	direction := 0
	direction, directions = pop(directions)
	// 0: left, 1: right
	if direction == 0 {
		return pickTree(directions, t.Left)
	}
	return pickTree(directions, t.Right)
}

func swap(t1, t2 *Tree) {
	*t1, *t2 = *t2, *t1
}

// どの木にどんな木を挿入するか
// insert(&t.Left, *Tree{})
func insert(t1 **Tree, t2 *Tree) {
	*t1 = t2
}

// 形だけ同じか調べる
func isSameTree(t1, t2 *Tree) bool {
	if t1 == nil || t2 == nil {
		return false
	}
	//fmt.Println(string(t1.Val), string(t2.Val))
	if isOp(t1.Val) && isOp(t2.Val) {
		if !isSameTree(t1.Left, t2.Left) {
			return false
		}

		if !isSameTree(t1.Right, t2.Right) {
			return false
		}

		return true
	}

	if !isOp(t1.Val) && !isOp(t2.Val) {
		return true
	}

	return false
}

/* Like
       *
    +     -
   1 2   3 4
*/
func buildTreeA() *Tree {
	t := NewTree(opgen())
	// edge
	insert(&t.Left, NewTree(opgen()))
	insert(&t.Right, NewTree(opgen()))

	// leaf
	insert(&t.Left.Left, NewTree(numgen()))
	insert(&t.Left.Right, NewTree(numgen()))
	insert(&t.Right.Left, NewTree(numgen()))
	insert(&t.Right.Right, NewTree(numgen()))

	return t
}

/* Like
        *
     +     4
    - 3
   1 2
*/
func buildTreeB() *Tree {
	t := NewTree(opgen())

	// edge
	insert(&t.Left, NewTree(opgen()))
	insert(&t.Left.Left, NewTree(opgen()))

	// leaf
	insert(&t.Right, NewTree(numgen()))
	insert(&t.Left.Right, NewTree(numgen()))
	insert(&t.Left.Left.Left, NewTree(numgen()))
	insert(&t.Left.Left.Right, NewTree(numgen()))

	return t
}

/* Like
      *
   1     +
        2 -
         3 4
*/
func buildTreeC() *Tree {
	t := NewTree(opgen())

	// edge
	insert(&t.Right, NewTree(opgen()))
	insert(&t.Right.Right, NewTree(opgen()))

	// leaf
	insert(&t.Left, NewTree(numgen()))
	insert(&t.Right.Left, NewTree(numgen()))
	insert(&t.Right.Right.Left, NewTree(numgen()))
	insert(&t.Right.Right.Right, NewTree(numgen()))

	return t
}

/* Like
       *
    +     4
   1 -
    2 3
*/
func buildTreeD() *Tree {
	t := NewTree(opgen())
	// edge
	insert(&t.Left, NewTree(opgen()))
	insert(&t.Left.Right, NewTree(opgen()))

	// leaf
	insert(&t.Right, NewTree(numgen()))
	insert(&t.Left.Left, NewTree(numgen()))
	insert(&t.Left.Right.Left, NewTree(numgen()))
	insert(&t.Left.Right.Right, NewTree(numgen()))

	return t
}

/* Like
      *
   1     +
        - 4
       2 3
*/
func buildTreeE() *Tree {
	t := NewTree(opgen())
	// edge
	insert(&t.Right, NewTree(opgen()))
	insert(&t.Right.Left, NewTree(opgen()))

	// leaf
	insert(&t.Left, NewTree(numgen()))
	insert(&t.Right.Right, NewTree(numgen()))
	insert(&t.Right.Left.Left, NewTree(numgen()))
	insert(&t.Right.Left.Right, NewTree(numgen()))

	return t
}
