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
