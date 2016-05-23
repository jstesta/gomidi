package midi

import "fmt"

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
	return e.length;
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
