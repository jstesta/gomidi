package gomidi

import (
	"bytes"
	"errors"
	"io"
	"io/ioutil"

	"github.com/jstesta/gomidi/parser"
	"github.com/jstesta/gomidi/midi"
)

func ReadMidiFromFile(fn string) (m *midi.Midi, err error) {

	bb, err := ioutil.ReadFile(fn)
	if err != nil {
		return nil, err
	}
	return ReadMidiFromBytes(bb)
}

func ReadMidiFromBytes(b []byte) (m *midi.Midi, err error) {

	buff := bytes.NewReader(b)
	return ReadMidiFromReader(buff)
}

func ReadMidiFromReader(r io.Reader) (m *midi.Midi, err error) {

	c, err := parser.ReadChunk(r)
	if err != nil {
		return nil, err
	}

	var tracks []midi.Track

	if header, ok := c.(*midi.Header); ok {
		tracks = make([]midi.Track, 0, header.NumberOfTracks)
		var i uint16
		for i = 0; i < header.NumberOfTracks; i++ {
			c, err := parser.ReadChunk(r)
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
