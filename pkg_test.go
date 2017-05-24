//go_series is a package for working with time series.
//the package currently only works with float32 values
package go_series

import (
	"testing"
	"fmt"
)

//TestAll tests all functions in the package
func TestAll (t *testing.T) {
	test_set := []float32{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	lq := NewLimitedQueue(10)
	for _, v := range test_set {
		lq.Insert(v)
	}

	fmt.Println(lq.Avg())
	if lq.Avg() != 4.5 {
		t.Error("LimitedQueue error")
	}

	ts := NewTimeSeries(5, 2)

	for _, v := range test_set {
		ts.Insert(v)
	}

	fmt.Println(ts.Avg())
	if ts.Avg() != 4.5 {
		t.Error("TimeSeries error")
	}
}

