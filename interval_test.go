package diet

import (
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

// Tests a valid interval
func TestIntervalValid(t *testing.T) {
	validate := validator.New()
	i := Interval{0, 5}
	err := validate.Struct(i)
	assert.NoError(t, err)
}

// Tests a valid interval with a single member
func TestIntervalValidPoint(t *testing.T) {
	validate := validator.New()
	i := Interval{5, 5}
	err := validate.Struct(i)
	assert.NoError(t, err)
}

// Tests an interval with invalid ordering
func TestIntervalInvalid(t *testing.T) {
	validate := validator.New()
	i := Interval{5, 0}
	err := validate.Struct(i)
	assert.Error(t, err)
}

// Tests stringfication
func TestIntervalString(t *testing.T) {
	i := Interval{0, 5}
	assert.Equal(t, "[0, 5]", i.String())
}

// Tests membership of a given value
func TestIntervalHas(t *testing.T) {
	i := Interval{0, 5}
	assert.True(t, i.Has(3))
	assert.False(t, i.Has(-2))
	assert.False(t, i.Has(7))
}

// Tests overlap of intervals
func TestIntervalOverlaps(t *testing.T) {
	i := Interval{3, 5}
	assert.True(t, i.Overlaps(Interval{0, 3}))
	assert.True(t, i.Overlaps(Interval{4, 5}))
	assert.True(t, i.Overlaps(Interval{5, 7}))
	assert.False(t, i.Overlaps(Interval{0, 2}))
	assert.False(t, i.Overlaps(Interval{6, 7}))
}

// Tests whether an interval is contained by another.
func TestIntervalContains(t *testing.T) {
	i := Interval{2, 5}
	assert.True(t, i.Contains(Interval{3, 4}))
	assert.True(t, i.Contains(Interval{2, 4}))
	assert.True(t, i.Contains(Interval{3, 5}))
	assert.True(t, i.Contains(Interval{2, 5}))
	assert.False(t, i.Contains(Interval{1, 3}))
	assert.False(t, i.Contains(Interval{4, 6}))
	assert.False(t, i.Contains(Interval{0, 1}))
}

// Tests adjacency.
func TestIntervalIsAdjacent(t *testing.T) {
	i := Interval{3, 5}
	assert.True(t, i.IsAdjacent(Interval{0, 2}))
	assert.True(t, i.IsAdjacent(Interval{6, 8}))
	assert.False(t, i.IsAdjacent(Interval{-1, 1}))
	assert.False(t, i.IsAdjacent(Interval{7, 9}))
	assert.False(t, i.IsAdjacent(Interval{2, 4}))
	assert.False(t, i.IsAdjacent(Interval{4, 6}))
}
