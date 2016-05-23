package midi

import "fmt"

type Midi struct {
	header *Header
	tracks []*Track
}

func (m *Midi) NumberOfTracks() int {
	return int(m.header.NumberOfTracks())
}

func (m *Midi) Division() int {
	return int(m.header.Division())
}

func (m *Midi) Header() *Header {
	return m.header
}

func (m *Midi) Tracks() []*Track {
	return m.tracks
}

func (m *Midi) String() string {
	return fmt.Sprintf("Midi [Header=%v, Tracks=%v]",
		m.header,
		m.tracks)
}

func NewMidi(h *Header, t []*Track) *Midi {
	return &Midi{h, t}
}
