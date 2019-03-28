// Package clock implements a clock type with a new-up methods
//	and methods for adding & subtracting
package clock

import "fmt"

// Clock is a type that holds a times hour and minute.
type Clock struct {
	m int
}

// New returns a new Clock type
func New(h, m int) Clock {
	nm := (h*60 + m) % 1440
	if nm < 0 {
		nm += 1440
	}
	return Clock{m: nm}
}

// String returns a string representation of the clock
func (c *Clock) String() string {
	h := c.m / 60
	m := c.m % 60
	return fmt.Sprintf("%02v:%02v", h, m)
}

// Add accepts a clock object pointer, and adds indicated hours minutes
//	to it's fields
func (c Clock) Add(m int) Clock {
	nm := (c.m + m) % 1440
	return Clock{m: nm}
}

// Subtract accepts a clock object, and subtracts indicated hours minutes
//	to it's fields
func (c Clock) Subtract(m int) Clock {
	nm := (c.m - m) % 1440
	if nm < 0 {
		nm += 1440
	}
	return Clock{m: nm}
}
