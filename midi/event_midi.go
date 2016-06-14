package midi

import "fmt"

const (
	MIDI_NOTE_OFF                = iota
	MIDI_NOTE_ON                 = iota
	MIDI_POLYPHONIC_KEY_PRESSURE = iota
	MIDI_CONTROL_CHANGE          = iota
	MIDI_PROGRAM_CHANGE          = iota
	MIDI_CHANNEL_PRESSURE        = iota
	MIDI_PITCH_WHEEL_CHANGE      = iota
)

type MidiEvent struct {
	deltaTime int
	status    byte
	length    int
	data      []byte
}

func (e *MidiEvent) DeltaTime() int {
	return e.deltaTime
}

func (e *MidiEvent) Length() int {
	return e.length
}

func (e *MidiEvent) Data() []byte {
	return e.data
}

func (e *MidiEvent) Status() byte {
	return e.status
}

func (e *MidiEvent) String() string {
	return fmt.Sprintf("MidiEvent [Status=%X, DeltaTime=%d, Length=%d, Data=%d]",
		e.Status(),
		e.DeltaTime(),
		e.Length(),
		e.Data())
}

func NewMidiEvent(t int, s byte, l int, d []byte) *MidiEvent {
	return &MidiEvent{t, s, l, d}
}
