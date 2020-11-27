package diet

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Empty trees
func TestDiscreteIntervalEncodingTreeIsEmpty(t *testing.T) {
	tree := &DiscreteIntervalEncodingTree{}
	assert.True(t, tree.IsEmpty())

	tree.Interval = &Interval{0, 2}
	assert.False(t, tree.IsEmpty())
}

// Tree validity
func TestDiscreteIntervalEncodingTreeValidateEmpty(t *testing.T) {
	tree := DiscreteIntervalEncodingTree{}
	assert.NoError(t, tree.Validate())
}

func TestDiscreteIntervalEncodingTreeValidate(t *testing.T) {
	left := &DiscreteIntervalEncodingTree{Interval: &Interval{0, 2}}
	right := &DiscreteIntervalEncodingTree{Interval: &Interval{8, 10}}
	tree := &DiscreteIntervalEncodingTree{Interval: &Interval{4, 6}, Left: left, Right: right}
	assert.NoError(t, tree.Validate())
}

func TestDiscreteIntervalEncodingTreeValidateInvalid(t *testing.T) {
	sub := &DiscreteIntervalEncodingTree{Interval: &Interval{0, 2}}
	tree := &DiscreteIntervalEncodingTree{Left: sub}
	assert.Error(t, tree.Validate())

	tree = &DiscreteIntervalEncodingTree{Right: sub}
	assert.Error(t, tree.Validate())
}

func TestDiscreteIntervalEncodingTreeValidateInvalidSubtree(t *testing.T) {
	sub := &DiscreteIntervalEncodingTree{Left: &DiscreteIntervalEncodingTree{}}
	tree := &DiscreteIntervalEncodingTree{Interval: &Interval{0, 2}, Left: sub}
	assert.Error(t, tree.Validate())

	tree = &DiscreteIntervalEncodingTree{Interval: &Interval{0, 2}, Right: sub}
	assert.Error(t, tree.Validate())
}
