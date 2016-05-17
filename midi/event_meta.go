package midi

import "fmt"

type MetaEvent struct {
	deltaTime int
	metaType  byte
	data      []byte
}

func (e *MetaEvent) DeltaTime() int {
	return e.deltaTime
}

func (e *MetaEvent) MetaType() byte {
	return e.metaType
}

func (e *MetaEvent) Data() []byte {
	return e.data
}

func (e *MetaEvent) String() string {

	var format string
	switch {
	case 0x1 <= e.MetaType() && e.MetaType() <= 0x7:
		format = "MetaEvent [MetaType=%X, DeltaTime=%d, Data=%s]"
	default:
		format = "MetaEvent [MetaType=%X, DeltaTime=%d, Data=%d]"
	}

	return fmt.Sprintf(format,
		e.MetaType(),
		e.DeltaTime(),
		e.Data())
}

func NewMetaEvent(delta int, metaType byte, data []byte) *MetaEvent {
	return &MetaEvent{delta, metaType, data}
}
