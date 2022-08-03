# Segment tree implementation in Go

Usage example:

```go
n := 10
slc := make([]segtree.TreeElem, n)
for i := 0; i < n; i++ {
	slc[i] = i
}
s, err := segtree.NewTree(slc, func(a, b segtree.TreeElem) segtree.TreeElem {return a.(int) + b.(int)}, 0)
if err != nil {
	fmt.Println(err)
}
fmt.Println(s.Query(0, n-1)) // [0,1,2,3,4,5,6,7,8,9] is initial slice; sum for 0 to n 45
s.Set(0, 5)                  // [5,1,2,3,4,5,6,7,8,9] now
fmt.Println(s.Query(0, n-1)) // sum = 50
s.Apply(0, 10)               // [15,1,2,3,4,5,6,7,8,9] now
fmt.Println(s.Query(0, n-1))
fmt.Println(s.Query(0, 3))
fmt.Println(s.Query(n-3, n-1))
fmt.Println(s.Query(0, 0))
```

1. Use segment.NewTree(slc, op, neutralElem) to create a new segment tree  
* slc should be []segtree.TreeElem  
* op should be the operation you want to calculate for ranges \[i,j\] (you can also apply it to tree elements \[currently one by one\])  
signature is  
func(segtree.TreeElem, segtree.TreeElem) segtree.TreeElem  
* neutralElem is a neutral element for operation op [for any element a: op(a,0) = op(0,a) = a] (that is, 0 for sum; 1 for product; false for XOR etc.)  
2. Currently tree has 3 methods:  
### Query(left, right int) (TreeElem, error)
Applies Operation for range [left, right] (inclusive); returns result and error (if given wrong range)
### Set(index int, value TreeElem) error
Sets element on position i to value; returns error if it's not possible to do that
### Apply(index int, value TreeElem) error
Sets element on position i to Operation(element, value); returns error if it's not possible to do that

TODO:  
Add apply for ranges  
Write examples, tests, benchmarks  
