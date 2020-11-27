package diet

import (
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

func TestIntervalValid(t *testing.T) {
	validate := validator.New()
	i := Interval{0, 5}
	err := validate.Struct(i)
	assert.NoError(t, err)
}

func TestIntervalValidPoint(t *testing.T) {
	validate := validator.New()
	i := Interval{5, 5}
	err := validate.Struct(i)
	assert.NoError(t, err)
}

func TestIntervalInvalid(t *testing.T) {
	validate := validator.New()
	i := Interval{5, 0}
	err := validate.Struct(i)
	assert.Error(t, err)
}


func TestIntervalString(t *testing.T) {
	i := Interval{0, 5}
	assert.Equal(t, "[0, 5]", i.String())
}
