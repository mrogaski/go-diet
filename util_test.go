package diet

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMin(t *testing.T) {
	assert.Equal(t, min(2, 5), 2)
	assert.Equal(t, min(5, 2), 2)
	assert.Equal(t, min(5, 5), 5)
}

func TestMax(t *testing.T) {
	assert.Equal(t, max(2, 5), 5)
	assert.Equal(t, max(5, 2), 5)
	assert.Equal(t, min(5, 5), 5)
}
