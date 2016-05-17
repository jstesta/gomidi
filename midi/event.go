package midi

//go:generate stringer -type=EventType

type EventType int

const (
	SYSEX EventType = iota
	META  EventType = iota
	MIDI  EventType = iota
)

type Event interface {
	Type() EventType
	DeltaTime() int
	Data() []byte
}
