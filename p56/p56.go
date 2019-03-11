package main

import (
	"fmt"
)

// Definition for an interval.
type Interval struct {
	Start int
	End   int
}

type overlappingType int

const (
	error             overlappingType = -1
	no_overlapping    overlappingType = 0
	contains          overlappingType = 1
	contained_by      overlappingType = 2
	start_overlapping overlappingType = 3
	end_overlapping   overlappingType = 4
)

func (I *Interval) overlapsWith(J Interval) overlappingType {
	a, b := I.Start, I.End
	x, y := J.Start, J.End

	if x < a {
		if y < a {
			return no_overlapping
		} else if a <= y && y < b {
			return start_overlapping
		} else { // b <= y
			return contained_by
		}
	} else if a <= x && x <= b {
		if y <= b {
			return contains
		} else { // b < y
			return end_overlapping
		}
	} else { // b < x
		return no_overlapping
	}

	return error
}

func (I *Interval) mergeWith(J Interval) (ret []Interval) {
	switch I.overlapsWith(J) {
	case no_overlapping:
		fmt.Println("ERROR: mergeTwo error.")
		ret = nil
	case contains:
		ret = append(ret, *I)
	case contained_by:
		ret = append(ret, J)
	case start_overlapping:
		ret = append(ret, Interval{Start: J.Start, End: I.End})
	case end_overlapping:
		ret = append(ret, Interval{Start: I.Start, End: J.End})
	default:
		fmt.Println("ERROR: overlap error.")
		ret = nil
	}
	return
}

func merge(intervals []Interval) []Interval {
	if len(intervals) <= 0 {
		fmt.Println("ERROR: 0")
		return nil
	}

	if len(intervals) == 1 {
		return intervals
	}

	if len(intervals) == 2 {
		I, J := intervals[0], intervals[1]
		if I.overlapsWith(J) != no_overlapping {
			return I.mergeWith(J)
		}
	}

	if len(intervals) == 3 {
	}
	return nil
}

func MergeTest(intervals []Interval) []Interval {
	return merge(intervals)
}

func main() {

}
