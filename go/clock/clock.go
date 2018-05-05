package clock

import "fmt"

// Clock interface declare Add and Subtract functions
type Clock interface {
	Add(int) Clock
	Subtract(int) Clock
}

// HoursAndMinutes struct contains Hours and Minutes integers
type HoursAndMinutes struct {
	Hours   int
	Minutes int
}

// New returns a new HoursAndMinutes struct
func New(hours, minutes int) HoursAndMinutes {
	hours, minutes = standardize(hours, minutes)

	return HoursAndMinutes{
		Hours:   hours,
		Minutes: minutes,
	}
}

// Add implements the Clock interface for HoursAndMinutes
func (clock *HoursAndMinutes) Add(time int) HoursAndMinutes {
	return New(
		clock.Hours,
		clock.Minutes+time,
	)
}

// Subtract implements the Clock interface for HoursAndMinutes
func (clock *HoursAndMinutes) Subtract(time int) HoursAndMinutes {
	return New(
		clock.Hours,
		clock.Minutes-time,
	)
}

// String implements the stringer interface for HoursAndMinutes
func (clock HoursAndMinutes) String() string {
	return fmt.Sprintf("%02v:%02v", clock.Hours, clock.Minutes)
}

func standardize(hours, minutes int) (int, int) {
	hours += minutes / 60
	minutes = minutes % 60
	hours = hours % 24

	if minutes < 0 {
		minutes += 60
		hours--
	}

	if hours < 0 {
		hours += 24
	}

	return hours, minutes
}
