package vlq_test

import (
	"bytes"
	"github.com/jstesta/gomidi/util/vlq"
	"testing"
)

var successTests = []struct {
	n        []byte
	expected int
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
		b := bytes.NewReader(tt.n)

		actual, read, err := vlq.ReadVLQ(b)
		if err != nil {
			t.Errorf("Read: %s", err.Error())
		}
		if read != len(tt.n) {
			t.Errorf("Read: expected %d bytes read, actual %d", len(tt.n), read)
		}
		if actual != tt.expected {
			t.Errorf("Read: expected %d, actual %d", tt.expected, actual)
		}
	}
}

func TestReadValueTooLarge(t *testing.T) {

	tt := []byte{0xff, 0xff, 0xff, 0xff, 0x7f}
	b := bytes.NewReader(tt)

	n, _, err := vlq.ReadVLQ(b)
	if err == nil {
		t.Errorf("Read: expected overflow error but succeeded, input=%d, n=%d", tt, n)
	}
}

func TestUnexpectedEndOfInput(t *testing.T) {

	tt := []byte{0xff}
	b := bytes.NewReader(tt)

	n, _, err := vlq.ReadVLQ(b)
	if err == nil {
		t.Errorf("Read: expected unexpected end of input error but succeeded, input=%d, n=%d", tt, n)
	}
}

func TestEmptyInput(t *testing.T) {

	tt := []byte{}
	b := bytes.NewReader(tt)

	n, _, err := vlq.ReadVLQ(b)
	if err == nil {
		t.Errorf("Read: expected unexpected end of input error but succeeded, input=%d, n=%d", tt, n)
	}
}
