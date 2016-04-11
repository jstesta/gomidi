package vlq_test

import (
	"bytes"
	"github.com/jstesta/gomidi/toolbox/vlq"
	"testing"
)

var successTests = []struct {
	n        []byte
	expected int32
}{
	{[]byte{0x00}, 0},
	{[]byte{0x40}, 64},
	{[]byte{0x7f}, 127},
	{[]byte{0x81, 0x00}, 128},
	{[]byte{0xc0, 0x00}, 8192},
	{[]byte{0xff, 0x7f}, 16383},
	{[]byte{0x81, 0x80, 0x00}, 16384},
	{[]byte{0xc0, 0x80, 0x00}, 1048576},
	{[]byte{0xff, 0xff, 0x7f}, 2097151},
	{[]byte{0x81, 0x80, 0x80, 0x00}, 2097152},
	{[]byte{0xc0, 0x80, 0x80, 0x00}, 134217728},
	{[]byte{0xff, 0xff, 0xff, 0x7f}, 268435455},
}

func TestReadOK(t *testing.T) {

	for _, tt := range successTests {
		buffer := bytes.NewReader(tt.n)

		actual, err := vlq.Read(buffer)
		if err != nil {
			t.Errorf("Read: %s", err.Error())
		}
		if actual != tt.expected {
			t.Errorf("Read: expected %d, actual %d", tt.expected, actual)
		}
	}
}

func TestReadValueTooLarge(t *testing.T) {

	tt := []byte{0xff, 0xff, 0xff, 0xff, 0x7f}
	buffer := bytes.NewReader(tt)

	_, err := vlq.Read(buffer)
	if err == nil {
		t.Errorf("Read: expected overflow error but succeeded", tt)
	}
}

func TestUnexpectedEndOfInput(t *testing.T) {

	tt := []byte{0xff}
	buffer := bytes.NewReader(tt)

	_, err := vlq.Read(buffer)
	if err == nil {
		t.Errorf("Read: expected unexpected end of input error but succeeded", tt)
	}
}
