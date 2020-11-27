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
func (it Interval) Overlaps(that Interval) bool {
	if it.Begin <= that.End && that.Begin <= it.End {
		return true
	}
	return false
}

// Contains tests whether the interval contains another interval.
func (it Interval) Contains(that Interval) bool {
	if it.Begin <= that.Begin && it.End >= that.End {
		return true
	}
	return false
}
