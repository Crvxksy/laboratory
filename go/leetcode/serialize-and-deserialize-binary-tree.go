package main

import (
	"fmt"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(n *TreeNode) string {
	res := []byte{}
	vs := [32]byte{}
	vi := 0
	nodes := []*TreeNode{nil}
	for len(nodes) > 0 {
		if n == nil {
			n = nodes[len(nodes)-1]
			nodes = nodes[:len(nodes)-1]
			res = append(res, ',')
		} else {
			if n.Val == 0 {
				res = append(res, '0')
			} else {
				if n.Val < 0 {
					res = append(res, '-')
					n.Val = -n.Val
				}
				for n.Val > 0 {
					vs[vi] = byte(n.Val%10) + '0'
					vi++
					n.Val /= 10
				}
				for vi > 0 {
					vi--
					res = append(res, vs[vi])
				}
			}
			if n.Right == nil {
				res = append(res, '.')
			} else {
				nodes = append(nodes, n.Right)
			}
			res = append(res, ',')
			n = n.Left
		}
	}
	return string(res)
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	val, sidx := 0, -1
	var root *TreeNode
	var node **TreeNode = &root
	var stack = []*TreeNode{{}}
	for i, b := range []byte(data) {
		if b == '.' || b == '-' {
			continue
		}
		if b != ',' {
			val = val*10 + int(b-'0')
			continue
		}
		if i-1 == sidx {
			node = &stack[len(stack)-1].Right
			stack = stack[:len(stack)-1]
		} else {
			if data[sidx+1:sidx+2] == "-" {
				val = -val
			}
			*node = &TreeNode{Val: val}
			if data[i-1:i] != "." {
				stack = append(stack, *node)
			}
			node = &(*node).Left
			val = 0
		}
		sidx = i
	}
	return root
}

// Serializes a tree to a single string.
func (this *Codec) serialize1(root *TreeNode) string {
	res := []byte{}
	vs := [32]byte{}
	vi := 0
	var addRes func(n *TreeNode)
	addRes = func(n *TreeNode) {
		if n == nil {
			res = append(res, ',')
		} else {
			if n.Val == 0 {
				res = append(res, '0')
			} else {
				if n.Val < 0 {
					res = append(res, '-')
					n.Val = -n.Val
				}
				for n.Val > 0 {
					vs[vi] = byte(n.Val%10) + '0'
					vi++
					n.Val /= 10
				}
				for vi > 0 {
					vi--
					res = append(res, vs[vi])
				}
			}
			res = append(res, ',')
			addRes(n.Left)
			addRes(n.Right)
		}
	}
	addRes(root)
	return string(res)
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize1(data string) *TreeNode {
	val, sidx := 0, -1
	sp := []byte{}
	vals := []int{}
	for i, b := range []byte(data) {
		if b != ',' {
			if b == '-' {
				continue
			}
			if b -= '0'; b < 10 {
				val = val*10 + int(b)
			}
			continue
		}
		if i-1 == sidx {
			sp = append(sp, 0)
		} else {
			if data[sidx+1:sidx+2] == "-" {
				val = -val
			}
			vals = append(vals, val)
			val = 0
			sp = append(sp, 1)
		}
		sidx = i
	}
	i, si := 0, -1
	var build func() *TreeNode
	build = func() *TreeNode {
		si++
		if sp[si] == 0 {
			return nil
		}
		v := vals[i]
		i++
		return &TreeNode{v, build(), build()}
	}
	return build()
}

func main() {
	ser := Constructor()
	deser := Constructor()
	root := &TreeNode{Val: -11,
		Left: &TreeNode{Val: -21},
		Right: &TreeNode{Val: 0,
			Left: &TreeNode{Val: -31,
				Right: &TreeNode{Val: 41, Right: &TreeNode{Val: -41},
					Left: &TreeNode{Val: -51,
						Left: &TreeNode{Val: -61,
							Left:  &TreeNode{Val: -71},
							Right: &TreeNode{Val: 71}},
						Right: &TreeNode{Val: -10}}}}},
	}
	data := ser.serialize(root)
	fmt.Println(data)
	ans := deser.deserialize(data)
	fmt.Println(deser.serialize(ans))
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
