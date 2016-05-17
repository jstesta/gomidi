package gomidi

import (
	"bytes"
	"encoding/binary"
	"errors"
	"io"
	"io/ioutil"

	"github.com/go-kit/kit/log"
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

func ReadMidiFromReader(r io.Reader, cfg cfg.GomidiConfig) (m *midi.Midi, err error) {

	if cfg.ByteOrder == nil {
		cfg.ByteOrder = binary.BigEndian
	}

	if cfg.LogContext == nil {
		cfg.LogContext = log.NewContext(log.NewNopLogger())
	}

	c, err := parser.ReadChunk(r, cfg)
	if err != nil {
		return nil, err
	}

	var tracks []midi.Track

	if header, ok := c.(*midi.Header); ok {
		tracks = make([]midi.Track, 0, header.NumberOfTracks)
		var i uint16
		for i = 0; i < header.NumberOfTracks; i++ {
			c, err := parser.ReadChunk(r, cfg)
			if err != nil {
				return nil, err
			}
			if track, ok := c.(*midi.Track); ok {
				tracks = append(tracks, *track)
			}
		}

		return midi.NewMidi(header, tracks), nil
	}

	return nil, errors.New("invalid midi file")
}
