package midi

import "fmt"

type MetricalDivision struct {
	resolution int
}

func (d *MetricalDivision) Type() int {
	return DIVISION_METRICAL
}

func (d *MetricalDivision) Resolution() int {
	return d.resolution
}

func (d *MetricalDivision) String() string {
	return fmt.Sprintf(
		"MetricalDivision [Resolution=%d]",
		d.Resolution())
}

func NewMetricalDivision(d int) *MetricalDivision {

	resolution := d & 0x7FFF

	return &MetricalDivision{resolution}
}
