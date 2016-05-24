package parser_test

import "github.com/jstesta/gomidi/midi"

func setupTrack1Event(d int, t byte, l int, data []byte) *midi.Track {

	return midi.NewTrack([]midi.Event{midi.NewMetaEvent(d, t, l, data)})
}
