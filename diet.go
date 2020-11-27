package diet

import "github.com/go-playground/validator/v10"

// DiscreteIntervalEncodingTree is a structure for storing subsets of integers as encoded intervals.
type DiscreteIntervalEncodingTree struct {
	Interval *Interval                     `validate:"required_with=Left Right"` // The interval contained by the node
	Left     *DiscreteIntervalEncodingTree // Intervals preceding the node interval
	Right    *DiscreteIntervalEncodingTree // Intervals succeeding the node interval
}

// IsEmpty tests whether the tree contains any intervals.
func (diet *DiscreteIntervalEncodingTree) IsEmpty() bool {
	return diet.Interval == nil
}

// Validate checks the validity of the DiscreteIntervalEncodingTree struct and returns errors if the struct is invalid.
func (diet *DiscreteIntervalEncodingTree) Validate() error {
	if diet == nil {
		return nil
	}

	err := diet.Left.Validate()
	if err != nil {
		return err
	}

	err = diet.Right.Validate()
	if err != nil {
		return err
	}

	return validate.Struct(diet)
}

// single instance of Validate to cache struct information
var validate *validator.Validate

func init() {
	validate = validator.New()
}
