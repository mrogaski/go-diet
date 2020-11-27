package diet

import "fmt"

type Interval struct {
	Begin int
	End   int `validate:"gtefield=Begin"`
}

func (i Interval) String() string {
	return fmt.Sprintf("[%d, %d]", i.Begin, i.End)
}
