package midi

import "fmt"

type SysexEvent struct {
	deltaTime int
	length    int
	data      []byte
}

func (e *SysexEvent) DeltaTime() int {
	return e.deltaTime
}

func (e *SysexEvent) Length() int {
	return e.length
}

func (e *SysexEvent) Data() []byte {
	return e.data
}

func (e *SysexEvent) String() string {
	return fmt.Sprintf("SysexEvent [DeltaTime=%d, Length=%d, Data=%d]",
		e.DeltaTime(),
		e.Length(),
		e.Data())
}

func NewSysexEvent(x int, l int, y []byte) *SysexEvent {
	return &SysexEvent{x, l, y}
}
