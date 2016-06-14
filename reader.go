package gomidi

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"

	"github.com/jstesta/gomidi/cfg"
	"github.com/jstesta/gomidi/midi"
	"github.com/jstesta/gomidi/parser"
)

func ReadMidiFromFile(fn string, cfg cfg.GomidiConfig) (m *midi.Midi, err error) {

	bb, err := ioutil.ReadFile(fn)
	if err != nil {
		return nil, err
	}
	return ReadMidiFromBytes(bb, cfg)
}

func ReadMidiFromBytes(b []byte, cfg cfg.GomidiConfig) (m *midi.Midi, err error) {

	buff := bytes.NewReader(b)
	return ReadMidiFromReader(buff, cfg)
}

func ReadMidiFromReader(r io.Reader, c cfg.GomidiConfig) (m *midi.Midi, err error) {

	if c.ByteOrder == nil {
		c.ByteOrder = cfg.DefaultByteOrder
	}

	if c.Log == nil {
		c.Log = log.New(ioutil.Discard, "", 0)
	}

	return parser.ReadMidi(r, c)
}
