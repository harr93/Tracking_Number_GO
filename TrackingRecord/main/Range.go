package main

import "fmt"

type Range struct {

	lo, hi int
}


func (obj Range) containsRange(r Range) bool {
	return obj.lo <= r.lo && r.hi <= obj.hi
}

func (obj Range) equals(r Range) bool {
	return obj.lo == r.lo && obj.hi == r.hi
}
func (obj Range) isDisjoint (r Range) bool {
	return obj.lo > r.hi || obj.hi < r.lo
}

func (obj Range) isOverlapping (r Range) bool {
	return !(obj.isDisjoint(r))
}

func (obj Range) lessThan(r Range) bool {
	return obj.lo < r.lo;
}

type RELATION int
const (
SUBSET RELATION = 1 + iota
SUPERSET
LESSOVERLAP
MOREOVERLAP
LESSDISJOINT
MOREDISJOINT
SAME
)

func (obj Range) classify(r Range) RELATION {

	if (obj.equals(r)) {
		return SAME
	}

	if (obj.containsRange(r)) {
		return SUPERSET
	}

	if (obj.isDisjoint(r)) {

		if (obj.lo > r.hi) {
			return MOREDISJOINT
		}
		if (obj.lo < r.hi) {
			return LESSDISJOINT
		}
	}

	if (obj.lessThan(r)) {
		return LESSOVERLAP
	}

	return MOREOVERLAP
}

func (obj Range) getLo() int {
	return obj.lo
}

func (obj Range) getHi() int {
	return obj.hi
}

func (obj Range) setLo(lo int) {
	obj.lo = lo
}

func (obj Range) setHi(hi int) {
	obj.hi = hi
}

func main() {
	r1 := Range{1, 100000}
    fmt.Print(r1.getHi())
}