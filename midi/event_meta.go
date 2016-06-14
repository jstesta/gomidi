package midi

import "fmt"

// These values were taken from the MIDI spec.
// See: https://www.midi.org/images/downloads/complete_midi_96-1-3.pdf Standard MIDI Files 1.0
const (
	META_SEQUENCE_NUMBER     = iota
	META_TEXT_EVENT          = iota
	META_COPYRIGHT_NOTICE    = iota
	META_SEQUENCE_NAME       = iota
	META_INSTRUMENT_NAME     = iota
	META_LYRIC               = iota
	META_MARKER              = iota
	META_CUE_POINT           = iota
	META_MIDI_CHANNEL_PREFIX = 0x20
	META_END_OF_TRACK        = 0x2F
	META_SET_TEMPO           = 0x51
	META_SMPTE_OFFSET        = 0x54
	META_TIME_SIGNATURE      = 0x58
	META_KEY_SIGNATURE       = 0x59
	META_SEQUENCER_SPECIFIC  = 0x7F
)

type MetaEvent struct {
	deltaTime int
	metaType  byte
	length    int
	data      []byte
}

func (e *MetaEvent) DeltaTime() int {
	return e.deltaTime
}

func (e *MetaEvent) MetaType() byte {
	return e.metaType
}

func (e *MetaEvent) Length() int {
	return e.length
}

func (e *MetaEvent) Data() []byte {
	return e.data
}

func (e *MetaEvent) String() string {
	return fmt.Sprintf("MetaEvent [MetaType=%X, DeltaTime=%d, Length=%d, Data=%v]",
		e.MetaType(),
		e.DeltaTime(),
		e.Length(),
		e.Data())
}

func NewMetaEvent(d int, t byte, l int, data []byte) *MetaEvent {
	return &MetaEvent{d, t, l, data}
}
