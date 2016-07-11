package midi

import "fmt"

const (
	SYSEX_SYSTEM_EXCLUSIVE      = 0xF0
	SYSEX_UNDEFINED_1           = 0xF1
	SYSEX_SONG_POSITION_POINTER = 0xF2
	SYSEX_SONG_SELECT           = 0xF3
	SYSEX_UNDEFINED_4           = 0xF4
	SYSEX_UNDEFINED_5           = 0xF5
	SYSEX_TUNE_REQUEST          = 0xF6
	SYSEX_END_OF_EXCLUSIVE      = 0xF7
	SYSEX_TIMING_CLOCK          = 0xF8
	SYSEX_UNDEFINED_9           = 0xF9
	SYSEX_START                 = 0xFA
	SYSEX_CONTINUE              = 0xFB
	SYSEX_STOP                  = 0xFC
	SYSEX_UNDEFINED_13          = 0xFD
	SYSEX_ACTIVE_SENSING        = 0xFE
	SYSEX_RESET                 = 0xFF
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
