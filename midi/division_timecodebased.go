package midi

import "fmt"

type TimeCodeBasedDivision struct {
	format     int
	resolution int
}

func (d *TimeCodeBasedDivision) Type() int {
	return DIVISION_TIME_CODE_BASED
}

func (d *TimeCodeBasedDivision) Format() int {
	return d.format
}

func (d *TimeCodeBasedDivision) Resolution() int {
	return d.resolution
}

func (d *TimeCodeBasedDivision) String() string {
	return fmt.Sprintf(
		"TimeCodeBasedDivision [Format=%d, Resolution=%d]",
		d.Format(),
		d.Resolution())
}

func NewTimeCodeBasedDivision(d int) *TimeCodeBasedDivision {

	format := d >> 8
	format &= 0x7F
	format ^= 0xFF
	format += 1

	resolution := d & 0xFF

	return &TimeCodeBasedDivision{format, resolution}
}
