package midi

type Event interface {
	DeltaTime() int
	Data() []byte
}
