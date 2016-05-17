package midi

import (
	"fmt"
)

type Track struct {
	Events []Event
}

func (t *Track) String() string {

	return fmt.Sprintf(
		"Track [Type=%s, Events=%v]",
		t.Type(),
		t.Events)
}

func (h Track) Type() ChunkType {

	return MTRK
}
