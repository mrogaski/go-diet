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

// Insert adds an interval to the tree.
func (diet *DiscreteIntervalEncodingTree) Insert(it Interval) error {
	if diet.Interval == nil {
		diet.Interval = &it
		return nil
	}
	if *diet.Interval == it {
		return nil
	}
	if diet.Interval.Overlaps(it) || diet.Interval.IsAdjacent(it) {
		merged, err := diet.Interval.Merge(it)
		if err != nil {
			return err
		}
		diet.Interval = &merged
		return nil
	}
	if diet.Interval.Begin > it.Begin {
		if diet.Left == nil {
			diet.Left = new(DiscreteIntervalEncodingTree)
		}
		return diet.Left.Insert(it)
	} else {
		if diet.Right == nil {
			diet.Right = new(DiscreteIntervalEncodingTree)
		}
		return diet.Right.Insert(it)
	}
}

// Intervals traverses the tree and returns an array of intervals.
func (diet *DiscreteIntervalEncodingTree) Intervals() []Interval {
	result := make([]Interval, 0)
	if diet.Interval == nil {
		return result
	}
	if diet.Left != nil {
		result = append(result, diet.Left.Intervals()...)
	}
	result = append(result, *diet.Interval)
	if diet.Right != nil {
		result = append(result, diet.Right.Intervals()...)
	}
	return result
}

// single instance of Validate to cache struct information
var validate *validator.Validate

func init() {
	validate = validator.New()
}
