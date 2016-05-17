package midi

import "fmt"

type SysexEvent struct {
	deltaTime int
	data      []byte
}

func (e *SysexEvent) Type() EventType {
	return SYSEX
}

func (e *SysexEvent) DeltaTime() int {
	return e.deltaTime
}

func (e *SysexEvent) Data() []byte {
	return e.data
}

func (e *SysexEvent) String() string {
	return fmt.Sprintf("SysexEvent [Type=%s, DeltaTime=%d, Data=%d]",
		e.Type(),
		e.DeltaTime(),
		e.Data())
}

func NewSysexEvent(x int, y []byte) *SysexEvent {
	return &SysexEvent{x, y}
}
