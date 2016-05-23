package midi

import "fmt"

type Track struct {
	events []Event
}

func (t *Track) Events() []Event {
	return t.events
}

func (t *Track) String() string {
	return fmt.Sprintf(
		"Track [Events=%v]",
		t.Events())
}

func NewTrack(e []Event) *Track {
	return &Track{e}
}
