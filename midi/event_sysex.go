package midi

import "fmt"

type SysexEvent struct {
	deltaTime int
	data      []byte
}

func (e *SysexEvent) DeltaTime() int {
	return e.deltaTime
}

func (e *SysexEvent) Data() []byte {
	return e.data
}

func (e *SysexEvent) String() string {
	return fmt.Sprintf("SysexEvent [DeltaTime=%d, Data=%d]",
		e.DeltaTime(),
		e.Data())
}

func NewSysexEvent(x int, y []byte) *SysexEvent {
	return &SysexEvent{x, y}
}
