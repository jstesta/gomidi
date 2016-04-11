package vlq

import "fmt"

type VLQReadError struct {
	Input []byte
	Msg   string
	Err   error
}

func (e *VLQReadError) Error() string {

	format := "variable-length quantity read %d: %s"
	if e.Err == nil {
		return fmt.Sprintf(format, e.Input, e.Msg)
	}
	return fmt.Sprintf(format+" [%s]", e.Input, e.Msg, e.Err.Error())
}
