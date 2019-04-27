package adt

import (
	"bytes"
	"math"
)

type Comparable interface{ Compare(c Comparable) int }
type rbcolor int

const (
	black	rbcolor	= iota
	red
)

type Interval struct {
	Begin	Comparable
	End	Comparable
}

func (ivl *Interval) Compare(c Comparable) int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	ivl2 := c.(*Interval)
	ivbCmpBegin := ivl.Begin.Compare(ivl2.Begin)
	ivbCmpEnd := ivl.Begin.Compare(ivl2.End)
	iveCmpBegin := ivl.End.Compare(ivl2.Begin)
	if ivbCmpBegin < 0 && iveCmpBegin <= 0 {
		return -1
	}
	if ivbCmpEnd >= 0 {
		return 1
	}
	return 0
}

type intervalNode struct {
	iv		IntervalValue
	max		Comparable
	left, right	*intervalNode
	parent		*intervalNode
	c		rbcolor
}

func (x *intervalNode) color() rbcolor {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if x == nil {
		return black
	}
	return x.c
}
func (n *intervalNode) height() int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if n == nil {
		return 0
	}
	ld := n.left.height()
	rd := n.right.height()
	if ld < rd {
		return rd + 1
	}
	return ld + 1
}
func (x *intervalNode) min() *intervalNode {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	for x.left != nil {
		x = x.left
	}
	return x
}
func (x *intervalNode) successor() *intervalNode {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if x.right != nil {
		return x.right.min()
	}
	y := x.parent
	for y != nil && x == y.right {
		x = y
		y = y.parent
	}
	return y
}
func (x *intervalNode) updateMax() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	for x != nil {
		oldmax := x.max
		max := x.iv.Ivl.End
		if x.left != nil && x.left.max.Compare(max) > 0 {
			max = x.left.max
		}
		if x.right != nil && x.right.max.Compare(max) > 0 {
			max = x.right.max
		}
		if oldmax.Compare(max) == 0 {
			break
		}
		x.max = max
		x = x.parent
	}
}

type nodeVisitor func(n *intervalNode) bool

func (x *intervalNode) visit(iv *Interval, nv nodeVisitor) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if x == nil {
		return true
	}
	v := iv.Compare(&x.iv.Ivl)
	switch {
	case v < 0:
		if !x.left.visit(iv, nv) {
			return false
		}
	case v > 0:
		maxiv := Interval{x.iv.Ivl.Begin, x.max}
		if maxiv.Compare(iv) == 0 {
			if !x.left.visit(iv, nv) || !x.right.visit(iv, nv) {
				return false
			}
		}
	default:
		if !x.left.visit(iv, nv) || !nv(x) || !x.right.visit(iv, nv) {
			return false
		}
	}
	return true
}

type IntervalValue struct {
	Ivl	Interval
	Val	interface{}
}
type IntervalTree struct {
	root	*intervalNode
	count	int
}

func (ivt *IntervalTree) Delete(ivl Interval) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	z := ivt.find(ivl)
	if z == nil {
		return false
	}
	y := z
	if z.left != nil && z.right != nil {
		y = z.successor()
	}
	x := y.left
	if x == nil {
		x = y.right
	}
	if x != nil {
		x.parent = y.parent
	}
	if y.parent == nil {
		ivt.root = x
	} else {
		if y == y.parent.left {
			y.parent.left = x
		} else {
			y.parent.right = x
		}
		y.parent.updateMax()
	}
	if y != z {
		z.iv = y.iv
		z.updateMax()
	}
	if y.color() == black && x != nil {
		ivt.deleteFixup(x)
	}
	ivt.count--
	return true
}
func (ivt *IntervalTree) deleteFixup(x *intervalNode) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	for x != ivt.root && x.color() == black && x.parent != nil {
		if x == x.parent.left {
			w := x.parent.right
			if w.color() == red {
				w.c = black
				x.parent.c = red
				ivt.rotateLeft(x.parent)
				w = x.parent.right
			}
			if w == nil {
				break
			}
			if w.left.color() == black && w.right.color() == black {
				w.c = red
				x = x.parent
			} else {
				if w.right.color() == black {
					w.left.c = black
					w.c = red
					ivt.rotateRight(w)
					w = x.parent.right
				}
				w.c = x.parent.color()
				x.parent.c = black
				w.right.c = black
				ivt.rotateLeft(x.parent)
				x = ivt.root
			}
		} else {
			w := x.parent.left
			if w.color() == red {
				w.c = black
				x.parent.c = red
				ivt.rotateRight(x.parent)
				w = x.parent.left
			}
			if w == nil {
				break
			}
			if w.left.color() == black && w.right.color() == black {
				w.c = red
				x = x.parent
			} else {
				if w.left.color() == black {
					w.right.c = black
					w.c = red
					ivt.rotateLeft(w)
					w = x.parent.left
				}
				w.c = x.parent.color()
				x.parent.c = black
				w.left.c = black
				ivt.rotateRight(x.parent)
				x = ivt.root
			}
		}
	}
	if x != nil {
		x.c = black
	}
}
func (ivt *IntervalTree) Insert(ivl Interval, val interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var y *intervalNode
	z := &intervalNode{iv: IntervalValue{ivl, val}, max: ivl.End, c: red}
	x := ivt.root
	for x != nil {
		y = x
		if z.iv.Ivl.Begin.Compare(x.iv.Ivl.Begin) < 0 {
			x = x.left
		} else {
			x = x.right
		}
	}
	z.parent = y
	if y == nil {
		ivt.root = z
	} else {
		if z.iv.Ivl.Begin.Compare(y.iv.Ivl.Begin) < 0 {
			y.left = z
		} else {
			y.right = z
		}
		y.updateMax()
	}
	z.c = red
	ivt.insertFixup(z)
	ivt.count++
}
func (ivt *IntervalTree) insertFixup(z *intervalNode) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	for z.parent != nil && z.parent.parent != nil && z.parent.color() == red {
		if z.parent == z.parent.parent.left {
			y := z.parent.parent.right
			if y.color() == red {
				y.c = black
				z.parent.c = black
				z.parent.parent.c = red
				z = z.parent.parent
			} else {
				if z == z.parent.right {
					z = z.parent
					ivt.rotateLeft(z)
				}
				z.parent.c = black
				z.parent.parent.c = red
				ivt.rotateRight(z.parent.parent)
			}
		} else {
			y := z.parent.parent.left
			if y.color() == red {
				y.c = black
				z.parent.c = black
				z.parent.parent.c = red
				z = z.parent.parent
			} else {
				if z == z.parent.left {
					z = z.parent
					ivt.rotateRight(z)
				}
				z.parent.c = black
				z.parent.parent.c = red
				ivt.rotateLeft(z.parent.parent)
			}
		}
	}
	ivt.root.c = black
}
func (ivt *IntervalTree) rotateLeft(x *intervalNode) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	y := x.right
	x.right = y.left
	if y.left != nil {
		y.left.parent = x
	}
	x.updateMax()
	ivt.replaceParent(x, y)
	y.left = x
	y.updateMax()
}
func (ivt *IntervalTree) rotateRight(x *intervalNode) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if x == nil {
		return
	}
	y := x.left
	x.left = y.right
	if y.right != nil {
		y.right.parent = x
	}
	x.updateMax()
	ivt.replaceParent(x, y)
	y.right = x
	y.updateMax()
}
func (ivt *IntervalTree) replaceParent(x *intervalNode, y *intervalNode) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	y.parent = x.parent
	if x.parent == nil {
		ivt.root = y
	} else {
		if x == x.parent.left {
			x.parent.left = y
		} else {
			x.parent.right = y
		}
		x.parent.updateMax()
	}
	x.parent = y
}
func (ivt *IntervalTree) Len() int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return ivt.count
}
func (ivt *IntervalTree) Height() int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return ivt.root.height()
}
func (ivt *IntervalTree) MaxHeight() int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return int((2 * math.Log2(float64(ivt.Len()+1))) + 0.5)
}

type IntervalVisitor func(n *IntervalValue) bool

func (ivt *IntervalTree) Visit(ivl Interval, ivv IntervalVisitor) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	ivt.root.visit(&ivl, func(n *intervalNode) bool {
		return ivv(&n.iv)
	})
}
func (ivt *IntervalTree) find(ivl Interval) (ret *intervalNode) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	f := func(n *intervalNode) bool {
		if n.iv.Ivl != ivl {
			return true
		}
		ret = n
		return false
	}
	ivt.root.visit(&ivl, f)
	return ret
}
func (ivt *IntervalTree) Find(ivl Interval) (ret *IntervalValue) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	n := ivt.find(ivl)
	if n == nil {
		return nil
	}
	return &n.iv
}
func (ivt *IntervalTree) Intersects(iv Interval) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	x := ivt.root
	for x != nil && iv.Compare(&x.iv.Ivl) != 0 {
		if x.left != nil && x.left.max.Compare(iv.Begin) > 0 {
			x = x.left
		} else {
			x = x.right
		}
	}
	return x != nil
}
func (ivt *IntervalTree) Contains(ivl Interval) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var maxEnd, minBegin Comparable
	isContiguous := true
	ivt.Visit(ivl, func(n *IntervalValue) bool {
		if minBegin == nil {
			minBegin = n.Ivl.Begin
			maxEnd = n.Ivl.End
			return true
		}
		if maxEnd.Compare(n.Ivl.Begin) < 0 {
			isContiguous = false
			return false
		}
		if n.Ivl.End.Compare(maxEnd) > 0 {
			maxEnd = n.Ivl.End
		}
		return true
	})
	return isContiguous && minBegin != nil && maxEnd.Compare(ivl.End) >= 0 && minBegin.Compare(ivl.Begin) <= 0
}
func (ivt *IntervalTree) Stab(iv Interval) (ivs []*IntervalValue) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if ivt.count == 0 {
		return nil
	}
	f := func(n *IntervalValue) bool {
		ivs = append(ivs, n)
		return true
	}
	ivt.Visit(iv, f)
	return ivs
}
func (ivt *IntervalTree) Union(inIvt IntervalTree, ivl Interval) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	f := func(n *IntervalValue) bool {
		ivt.Insert(n.Ivl, n.Val)
		return true
	}
	inIvt.Visit(ivl, f)
}

type StringComparable string

func (s StringComparable) Compare(c Comparable) int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	sc := c.(StringComparable)
	if s < sc {
		return -1
	}
	if s > sc {
		return 1
	}
	return 0
}
func NewStringInterval(begin, end string) Interval {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return Interval{StringComparable(begin), StringComparable(end)}
}
func NewStringPoint(s string) Interval {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return Interval{StringComparable(s), StringComparable(s + "\x00")}
}

type StringAffineComparable string

func (s StringAffineComparable) Compare(c Comparable) int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	sc := c.(StringAffineComparable)
	if len(s) == 0 {
		if len(sc) == 0 {
			return 0
		}
		return 1
	}
	if len(sc) == 0 {
		return -1
	}
	if s < sc {
		return -1
	}
	if s > sc {
		return 1
	}
	return 0
}
func NewStringAffineInterval(begin, end string) Interval {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return Interval{StringAffineComparable(begin), StringAffineComparable(end)}
}
func NewStringAffinePoint(s string) Interval {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return NewStringAffineInterval(s, s+"\x00")
}
func NewInt64Interval(a int64, b int64) Interval {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return Interval{Int64Comparable(a), Int64Comparable(b)}
}
func NewInt64Point(a int64) Interval {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return Interval{Int64Comparable(a), Int64Comparable(a + 1)}
}

type Int64Comparable int64

func (v Int64Comparable) Compare(c Comparable) int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	vc := c.(Int64Comparable)
	cmp := v - vc
	if cmp < 0 {
		return -1
	}
	if cmp > 0 {
		return 1
	}
	return 0
}

type BytesAffineComparable []byte

func (b BytesAffineComparable) Compare(c Comparable) int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	bc := c.(BytesAffineComparable)
	if len(b) == 0 {
		if len(bc) == 0 {
			return 0
		}
		return 1
	}
	if len(bc) == 0 {
		return -1
	}
	return bytes.Compare(b, bc)
}
func NewBytesAffineInterval(begin, end []byte) Interval {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return Interval{BytesAffineComparable(begin), BytesAffineComparable(end)}
}
func NewBytesAffinePoint(b []byte) Interval {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	be := make([]byte, len(b)+1)
	copy(be, b)
	be[len(b)] = 0
	return NewBytesAffineInterval(b, be)
}
