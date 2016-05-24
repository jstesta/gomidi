package parser_test

import (
	"bytes"
	"io"
	"testing"

	"github.com/jstesta/gomidi/cfg"
	"github.com/jstesta/gomidi/midi"
	"github.com/jstesta/gomidi/parser"
)

var successHeaderTests = []struct {
	n        []byte
	expected *midi.Header
}{
	{[]byte{'M', 'T', 'h', 'd', 0x0, 0x0, 0x0, 0x6, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}, midi.NewHeader(6, 0, 0, 0)},
	{[]byte{'M', 'T', 'h', 'd', 0x0, 0x0, 0x0, 0x6, 0x0, 0x1, 0x0, 0x1, 0x0, 0xf0}, midi.NewHeader(6, 1, 1, 240)},
}

func TestParseHeaderOK(t *testing.T) {

	for _, tt := range successHeaderTests {
		r := bytes.NewReader(tt.n)

		actual, err := parser.ReadHeader(r, cfg.GomidiConfig{})
		if err != nil {
			t.Errorf("parser.ReadHeader: %v", err.Error())
		}
		if *actual != *tt.expected {
			t.Errorf("parser.ReadHeader: expected %v, actual %v", tt.expected, actual)
		}
	}
}

func TestParseHeaderInvalidChunkType(t *testing.T) {

	tt := []byte{'M', 'T', 'h', 0x0, 0x0, 0x0, 0x6, 0x0, 0x1, 0x0, 0x1, 0x0, 0xf0}
	r := bytes.NewReader(tt)

	actual, err := parser.ReadHeader(r, cfg.GomidiConfig{})

	if err == nil {
		t.Errorf("parser.ReadHeader: expected error but succeeded, input=%d, n=%d", tt, actual)
	}
}

func TestParseHeaderUnexpectedEOF(t *testing.T) {

	tt := []byte{'M', 'T', 'h', 'd', 0x0, 0x0, 0x0}
	r := bytes.NewReader(tt)

	actual, err := parser.ReadHeader(r, cfg.GomidiConfig{})

	if err == nil {
		t.Errorf("parser.ReadHeader: expected error but succeeded, input=%d, n=%d", tt, actual)
	}

	if err != io.ErrUnexpectedEOF {
		t.Errorf("parser.ReadHeader: expected io.ErrUnexpectedEOF but was %v", err)
	}
}

func TestParseHeaderEOF(t *testing.T) {

	tt := []byte{'M', 'T', 'h', 'd', 0x0, 0x0, 0x0, 0x6}
	r := bytes.NewReader(tt)

	actual, err := parser.ReadHeader(r, cfg.GomidiConfig{})

	if err == nil {
		t.Errorf("parser.ReadHeader: expected error but succeeded, input=%d, n=%d", tt, actual)
	}

	if err != io.EOF {
		t.Errorf("parser.ReadHeader: expected io.EOF but was %v", err)
	}
}

// TODO test for failure when header fails integrity checks
//func TestParseHeaderLengthNot6(t *testing.T) {
//
//	tt := []byte{'M', 'T', 'h', 'd', 0x0, 0x0, 0x0, 0x1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}
//	r := bytes.NewReader(tt)
//
//	actual, err := parser.ReadHeader(r, cfg.GomidiConfig{})
//
//	if err == nil {
//		t.Errorf("parser.ReadHeader: expected header integrity error but succeeded, input=%v, n=%v", tt, actual)
//	}
//}
//
//func TestParseHeaderFormat0HasMoreThanOneTrack(t *testing.T) {
//
//	tt := []byte{'M', 'T', 'h', 'd', 0x0, 0x0, 0x0, 0x6, 0x0, 0x0, 0x0, 0x2, 0x0, 0x0}
//	r := bytes.NewReader(tt)
//
//	actual, err := parser.ReadHeader(r, cfg.GomidiConfig{})
//
//	if err == nil {
//		t.Errorf("parser.ReadHeader: expected header integrity error but succeeded, input=%v, n=%v", tt, actual)
//	}
//}
