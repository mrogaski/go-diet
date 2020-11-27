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
