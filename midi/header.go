package midi

import "fmt"

type Header struct {
	length         int
	format         int
	numberOfTracks int
	division       Division
}

func (h *Header) Length() int {
	return h.length
}

func (h *Header) Format() int {
	return h.format
}

func (h *Header) NumberOfTracks() int {
	return h.numberOfTracks
}

func (h *Header) Division() Division {
	return h.division
}

func (h *Header) String() string {
	return fmt.Sprintf(
		"Header [Length=%d, Format=%d, NumberOfTracks=%d, Division=%v]",
		h.Length(),
		h.Format(),
		h.NumberOfTracks(),
		h.Division())
}

func NewHeader(l int, f int, n int, d int) *Header {

	var division Division
	switch d >> 15 {
	case 1:
		division = NewTimeCodeBasedDivision(d)
	case 0:
		division = NewMetricalDivision(d)
	}

	return &Header{l, f, n, division}
}
