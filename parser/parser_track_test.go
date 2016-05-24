package parser_test

import (
	"bytes"
	"io"
	"testing"

	"github.com/jstesta/gomidi/cfg"
	"github.com/jstesta/gomidi/parser"
)

func TestParseTrackEOF(t *testing.T) {

	tt := []byte{'M', 'T', 'r', 'k'}
	r := bytes.NewReader(tt)

	actual, err := parser.ReadTrack(r, cfg.GomidiConfig{})

	if err == nil {
		t.Errorf("parser.ReadTrack: expected error but succeeded, input=%d, n=%d", tt, actual)
	}

	if err != io.EOF {
		t.Errorf("parser.ReadTrack: expected io.EOF but was %v", err)
	}
}

func TestParseTrackUnexpectedEOF(t *testing.T) {

	tt := []byte{'M', 'T', 'r', 'k', 0x0, 0x0, 0x0}
	r := bytes.NewReader(tt)

	actual, err := parser.ReadTrack(r, cfg.GomidiConfig{})

	if err == nil {
		t.Errorf("parser.ReadTrack: expected error but succeeded, input=%d, n=%d", tt, actual)
	}

	if err != io.ErrUnexpectedEOF {
		t.Errorf("parser.ReadTrack: expected io.ErrUnexpectedEOF but was %v", err)
	}
}
