package midi

const (
	_                        = iota
	DIVISION_METRICAL        = iota
	DIVISION_TIME_CODE_BASED = iota
)

type Division interface {
	Type() int
}
