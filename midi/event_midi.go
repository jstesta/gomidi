package midi

import "fmt"

const (
	MIDI_NOTE_OFF                = 0x8
	MIDI_NOTE_ON                 = 0x9
	MIDI_POLYPHONIC_KEY_PRESSURE = 0xA
	MIDI_CONTROL_CHANGE          = 0xB
	MIDI_PROGRAM_CHANGE          = 0xC
	MIDI_CHANNEL_PRESSURE        = 0xD
	MIDI_PITCH_WHEEL_CHANGE      = 0xE
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
