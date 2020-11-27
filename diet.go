package diet

type DiscreteIntervalEncodingTree struct {
	Interval *Interval                     // The interval contained by the node
	Left     *DiscreteIntervalEncodingTree // Intervals preceding the node interval
	Right    *DiscreteIntervalEncodingTree // Intervals succeeding the node interval
}

// IsEmpty tests whether the tree contains any intervals.
func (diet *DiscreteIntervalEncodingTree) IsEmpty() bool {
	return diet.Interval == nil
}
