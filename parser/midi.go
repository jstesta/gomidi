package parser

import (
	"io"

	"github.com/jstesta/gomidi/cfg"
	"github.com/jstesta/gomidi/midi"
)

func ReadMidi(r io.Reader, c cfg.GomidiConfig) (m *midi.Midi, err error) {

	header, err := readHeader(r, c)
	if err != nil {
		return
	}

	var tracks []*midi.Track

	tracks = make([]*midi.Track, 0, header.NumberOfTracks())
	var i int
	for i = 0; i < header.NumberOfTracks(); i++ {
		track, err := readTrack(r, c)
		if err != nil {
			return nil, err
		}
		// track == nil if an alien chunk was found
		if track != nil {
			tracks = append(tracks, track)
		}
	}

	return midi.NewMidi(header, tracks), nil
}
