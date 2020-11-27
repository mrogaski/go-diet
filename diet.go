package diet

import "github.com/go-playground/validator/v10"

type DiscreteIntervalEncodingTree struct {
	Interval *Interval                     // The interval contained by the node
	Left     *DiscreteIntervalEncodingTree // Intervals preceding the node interval
	Right    *DiscreteIntervalEncodingTree // Intervals succeeding the node interval
}

// IsEmpty tests whether the tree contains any intervals.
func (diet *DiscreteIntervalEncodingTree) IsEmpty() bool {
	return diet.Interval == nil
}

// single instance of Validate to cache struct information
var validate *validator.Validate

func init() {
	validate = validator.New()
}
