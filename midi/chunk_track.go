package midi

import (
	"fmt"
)

type Track struct {
	Events []Event
}

func (t *Track) String() string {

	return fmt.Sprintf(
		"Track [Events=%v]",
		t.Events)
}
