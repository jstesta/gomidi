package midi

import "fmt"

type Header struct {
	length         int
	format         int
	numberOfTracks int
	division       int
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

func (h *Header) Division() int {
	return h.division
}

func (h *Header) String() string {

	return fmt.Sprintf(
		"Header [Length=%d, Format=%d, NumberOfTracks=%d, Division=%d]",
		h.Length(),
		h.Format(),
		h.NumberOfTracks(),
		h.Division())
}

func NewHeader(l int, f int, n int, d int) *Header {
	return &Header{l, f, n, d}
}
