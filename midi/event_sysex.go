package midi

import "fmt"

const (
	SYSEX_SYSTEM_EXCLUSIVE      = iota
	SYSEX_UNDEFINED_1           = iota
	SYSEX_SONG_POSITION_POINTER = iota
	SYSEX_SONG_SELECT           = iota
	SYSEX_UNDEFINED_4           = iota
	SYSEX_UNDEFINED_5           = iota
	SYSEX_TUNE_REQUEST          = iota
	SYSEX_END_OF_EXCLUSIVE      = iota
	SYSEX_TIMING_CLOCK          = iota
	SYSEX_UNDEFINED_9           = iota
	SYSEX_START                 = iota
	SYSEX_CONTINUE              = iota
	SYSEX_STOP                  = iota
	SYSEX_UNDEFINED_13          = iota
	SYSEX_ACTIVE_SENSING        = iota
	SYSEX_RESET                 = iota
)

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

func NewSysexEvent(d int, l int, data []byte) *SysexEvent {
	return &SysexEvent{d, l, data}
}
