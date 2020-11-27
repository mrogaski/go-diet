package diet

import (
	"fmt"
)

// Interval is a closed interval over integers.
type Interval struct {
	Begin int // member of the interval with the minimum value
	End   int `validate:"gtefield=Begin"` // member of the interval with the maximum value
}

// String represents the interval using a standard [x, y] notation.
func (it Interval) String() string {
	return fmt.Sprintf("[%d, %d]", it.Begin, it.End)
}

// Has tests membership of a given integer.
func (it Interval) Has(x int) bool {
	if it.Begin <= x && x <= it.End {
		return true
	}
	return false
}

// Overlaps tests whether two intervals have common members.
func (it Interval) Overlaps(other Interval) bool {
	if it.Begin <= other.End && other.Begin <= it.End {
		return true
	}
	return false
}

// Contains tests whether the interval contains another interval.
func (it Interval) Contains(other Interval) bool {
	if it.Begin <= other.Begin && it.End >= other.End {
		return true
	}
	return false
}

// IsAdjacent tests whether intervals are adjacent.
func (it Interval) IsAdjacent(other Interval) bool {
	if it.Begin-1 == other.End || it.End+1 == other.Begin {
		return true
	}
	return false
}
