package chunk

import (
	"encoding/binary"
	"errors"
	"io"
	"log"

	"github.com/jstesta/gomidi/midi"
	"github.com/jstesta/gomidi/vlq"
)

func ReadTrackChunk(r io.Reader) (c midi.Chunk, err error) {

	var length uint32
	err = binary.Read(r, binary.BigEndian, &length)
	if err != nil {
		return
	}
	log.Printf("track.ReadTrackChunk length %d bytes", length)

	// track how many bytes have been read for exit condition
	var bytesRead int
	var events []midi.Event = make([]midi.Event, 0)
	var previousEvent midi.Event = nil

	for bytesRead < int(length) {

		deltaTime, br, err := vlq.ReadVLQ(r)
		if err != nil {
			return nil, err
		}
		bytesRead += br
		log.Print("track.ReadTrackChunk deltaTime ", deltaTime)

		var b [1]byte
		_, err = r.Read(b[:])
		if err != nil {
			return nil, err
		}
		eventType := b[0]
		bytesRead++

		switch eventType {

		case 0xF0, 0xF7:
			log.Print("track.ReadTrackChunk found sysex event")
			event, br, err := readSysexEvent(r, deltaTime)
			if err != nil {
				return nil, err
			}
			bytesRead += br
			log.Print("track.ReadTrackChunk parsed sysex event ", event)
			events = append(events, event)
			previousEvent = event

		case 0xFF:
			log.Print("track.ReadTrackChunk found meta event")
			event, br, err := readMetaEvent(r, deltaTime)
			if err != nil {
				return nil, err
			}
			bytesRead += br
			log.Print("track.ReadTrackChunk parsed meta event ", event)
			events = append(events, event)
			previousEvent = event

		default:
			log.Print("track.ReadTrackChunk found midi event")
			event, br, err := readMidiEvent(r, deltaTime, eventType, previousEvent)
			if err != nil {
				return nil, err
			}
			bytesRead += br
			log.Print("track.ReadTrackChunk parsed midi event ", event)
			events = append(events, event)
			previousEvent = event

		}

		log.Printf("track.ReadTrackChunk bytesRead: %d", bytesRead)

		if bytesRead > int(length) {
			return nil, errors.New("parsed too many bytes")
		}
	}

	return &midi.Track{events}, nil
}

func readSysexEvent(r io.Reader, deltaTime int) (e midi.Event, bytesRead int, err error) {

	length, br, err := vlq.ReadVLQ(r)
	if err != nil {
		return
	}
	bytesRead += br

	log.Printf("track.ReadTrackChunk[sysex] reading %d bytes of bulk data", length)
	data := make([]byte, length)
	err = binary.Read(r, binary.BigEndian, data)
	bytesRead += length

	e = midi.NewSysexEvent(deltaTime, data)

	return
}

func readMetaEvent(r io.Reader, deltaTime int) (e midi.Event, bytesRead int, err error) {

	var b [1]byte
	_, err = r.Read(b[:])
	if err != nil {
		return
	}
	metaEventType := b[0]
	bytesRead++
	log.Printf("track.ReadTrackChunk[meta] meta event type: %X", metaEventType)

	length, br, err := vlq.ReadVLQ(r)
	if err != nil {
		return
	}
	bytesRead += br

	log.Printf("track.ReadTrackChunk[meta] reading %d bytes of bulk data", length)
	data := make([]byte, length)
	err = binary.Read(r, binary.BigEndian, data)
	bytesRead += length

	e = midi.NewMetaEvent(deltaTime, metaEventType, data)

	return
}

func readMidiEvent(r io.Reader, deltaTime int, status byte, prev midi.Event) (e midi.Event, bytesRead int, err error) {

	switch status >> 4 {
	case 0x8, 0x9, 0xA, 0xB, 0xE:
		data := make([]byte, 2)
		err = binary.Read(r, binary.BigEndian, data)
		if err != nil {
			return
		}
		bytesRead += 2
		e = midi.NewMidiEvent(deltaTime, status, data)
	case 0xC, 0xD:
		data := make([]byte, 1)
		err = binary.Read(r, binary.BigEndian, data)
		if err != nil {
			return
		}
		bytesRead += 1
		e = midi.NewMidiEvent(deltaTime, status, data)
	default:
		if prev == nil {
			log.Printf("skipped unknown MIDI event, type %X", status>>4)
		}

		if prevMidiEvent, ok := prev.(*midi.MidiEvent); ok {
			log.Printf("no status MIDI event, using previous status of: %X", prevMidiEvent.Status())
			return readMidiEvent(r, deltaTime, prevMidiEvent.Status(), nil)
		}

		log.Print("this is bad! fixme")
	}

	return
}
