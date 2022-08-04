package segmenttree

import (
	"errors"
	"fmt"
)

type TreeElem interface {
	//Operation(a, b TreeElem) TreeElem
} // should support operation Operation

type Operation func(TreeElem, TreeElem) TreeElem // should have neutralElem; should be distributive

type Segtree struct {
	n           int
	lazyN       int
	op          Operation
	tree        []TreeElem
	neutralElem TreeElem
}

func nextPowerOfTwo(x int) int {
	x--
	x |= x >> 1
	x |= x >> 2
	x |= x >> 4
	x |= x >> 8
	x |= x >> 16
	x++
	return x
}

func NewTree(s []TreeElem, op Operation, neutralElem TreeElem) (*Segtree, error) {
	n := len(s)
	if n == 0 {
		return nil, errors.New("Error: slice is empty")
	}
	if n >= 1<<30 {
		return nil, errors.New("Error: slice is too large")
	}
	lazyN := nextPowerOfTwo(n)
	t := &Segtree{
		n:           n,
		lazyN:       lazyN,
		op:          op,
		tree:        make([]TreeElem, 4*lazyN),
		neutralElem: neutralElem,
	}
	t.build(s)
	return t, nil
}

func (t *Segtree) build(s []TreeElem) { // should only be called in NewTree function
	for i := 0; i < t.n; i++ {
		t.tree[t.lazyN-1+i] = s[i]
	}
	for i := t.n; i < t.lazyN; i++ {
		t.tree[t.lazyN-1+i] = t.neutralElem
	}
	for i := t.lazyN - 2; i >= 0; i-- {
		t.tree[i] = t.op(t.tree[2*i+1], t.tree[2*i+2])
	}
}

func (t Segtree) implQuery(left, right int) TreeElem {
	leftRes := t.neutralElem
	rightRes := t.neutralElem
	for left < right {
		if left%2 == 0 {
			leftRes = t.op(leftRes, t.tree[left])
		}
		//left--
		left /= 2
		if right%2 == 1 {
			rightRes = t.op(t.tree[right], rightRes)
		}
		//right--
		right /= 2
		right--
	}
	if left == right {
		leftRes = t.op(leftRes, t.tree[left]) //left
	}
	return t.op(leftRes, rightRes)
}

func (t Segtree) Query(left, right int) (TreeElem, error) {
	if left < 0 || right >= t.n {
		return nil, errors.New(fmt.Sprintf("Incorrect query range; indexes should be in [0, %d]", t.n-1))
	}
	return t.implQuery(left+t.lazyN-1, right+t.lazyN-1), nil
}

func (t *Segtree) Set(index int, value TreeElem) error {
	if index < 0 || index >= t.n {
		return errors.New(fmt.Sprintf("Incorrect query range; indexes should be in [0, %d]", t.n-1))
	}
	i := t.lazyN - 1 + index
	t.tree[i] = value
	for par := (i - 1) >> 1; par >= 0; par = (i - 1) >> 1 {
		if i%2 == 0 {
			t.tree[par] = t.op(t.tree[i-1], t.tree[i])
		} else {
			t.tree[par] = t.op(t.tree[i], t.tree[i+1])
		}
		i = par
	}
	return nil
}

func (t *Segtree) Apply(index int, value TreeElem) error { // Apply Operation(elem[index], value)
	if index < 0 || index >= t.n {
		return errors.New(fmt.Sprintf("Incorrect query range; indexes should be in [0, %d]", t.n-1))
	}
	i := t.lazyN - 1 + index
	t.tree[i] = t.op(t.tree[i], value)
	for par := (i - 1) >> 1; par >= 0; par = (i - 1) >> 1 {
		if i%2 == 0 {
			t.tree[par] = t.op(t.tree[i-1], t.tree[i])
		} else {
			t.tree[par] = t.op(t.tree[i], t.tree[i+1])
		}
		i = par
	}
	return nil
}
