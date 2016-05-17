package midi

import (
	"fmt"
)

type Header struct {
	Length         uint32
	Format         uint16
	NumberOfTracks uint16
	Division       uint16
}

func (h *Header) String() string {

	return fmt.Sprintf(
		"Header [Length=%v, Format=%v, NumberOfTracks=%v, Division=%v]",
		h.Length,
		h.Format,
		h.NumberOfTracks,
		h.Division)
}
