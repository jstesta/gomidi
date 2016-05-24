package parser

import (
	"encoding/binary"
	"errors"
	"io"

	"github.com/go-kit/kit/log"
	"github.com/jstesta/gomidi/cfg"
	"github.com/jstesta/gomidi/midi"
	"github.com/jstesta/gomidi/vlq"
)

func ReadTrack(r io.Reader, cfg cfg.GomidiConfig) (c *midi.Track, err error) {

	if cfg.ByteOrder == nil {
		cfg.ByteOrder = binary.BigEndian
	}

	if cfg.LogContext == nil {
		cfg.LogContext = log.NewContext(log.NewNopLogger())
	}

	chunkType := make([]byte, 4)
	_, err = io.ReadFull(r, chunkType)
	if err != nil {
		return nil, err
	}

	switch string(chunkType) {

	case TRACK_CHUNK_LITERAL:
		return readTrackChunk(r, cfg)

	case HEADER_CHUNK_LITERAL:
		return nil, errors.New("wanted Track chunk but found Header")

	default:
		return nil, readAlienChunk(r, cfg)
	}
}

func readTrackChunk(r io.Reader, cfg cfg.GomidiConfig) (c *midi.Track, err error) {

	var length uint32
	err = binary.Read(r, cfg.ByteOrder, &length)
	if err != nil {
		return
	}

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

		b := []byte{0}
		_, err = io.ReadFull(r, b)
		if err != nil {
			return nil, err
		}
		eventType := b[0]
		bytesRead++

		switch eventType {

		case 0xF0, 0xF7:
			event, br, err := readSysexEvent(r, deltaTime, cfg)
			if err != nil {
				return nil, err
			}
			bytesRead += br
			events = append(events, event)
			previousEvent = event

		case 0xFF:
			event, br, err := readMetaEvent(r, deltaTime, cfg)
			if err != nil {
				return nil, err
			}
			bytesRead += br
			events = append(events, event)
			previousEvent = event

		default:
			event, br, err := readMidiEvent(r, deltaTime, eventType, previousEvent, nil, cfg)
			if err != nil {
				return nil, err
			}
			bytesRead += br
			events = append(events, event)
			previousEvent = event

		}

		if bytesRead > int(length) {
			return nil, errors.New("parsed too many bytes")
		}
	}

	return midi.NewTrack(events), nil
}

func readSysexEvent(r io.Reader, deltaTime int, cfg cfg.GomidiConfig) (e midi.Event, bytesRead int, err error) {

	length, br, err := vlq.ReadVLQ(r)
	if err != nil {
		return
	}
	bytesRead += br

	data := make([]byte, length)
	_, err = io.ReadFull(r, data)
	bytesRead += length

	e = midi.NewSysexEvent(deltaTime, length, data)

	return
}

func readMetaEvent(r io.Reader, deltaTime int, cfg cfg.GomidiConfig) (e midi.Event, bytesRead int, err error) {

	b := []byte{0}
	_, err = io.ReadFull(r, b)
	if err != nil {
		return
	}
	metaEventType := b[0]
	bytesRead++

	length, br, err := vlq.ReadVLQ(r)
	if err != nil {
		return
	}
	bytesRead += br

	data := make([]byte, length)
	_, err = io.ReadFull(r, data)
	if err != nil {
		return
	}
	bytesRead += length

	e = midi.NewMetaEvent(deltaTime, metaEventType, length, data)

	return
}

func readMidiEvent(r io.Reader, deltaTime int, status byte, prev midi.Event, prefix []byte, cfg cfg.GomidiConfig) (e midi.Event, bytesRead int, err error) {

	subStatus := status >> 4
	ctx := cfg.LogContext.With("reader", "midi", "status", status, "subStatus", subStatus)

	if subStatus < 0xF && subStatus >= 0x8 {

		switch subStatus {

		case 0x8, 0x9, 0xA, 0xB, 0xE:
			n, data, err := readPrefixed(r, prefix, 2)
			if err != nil {
				return nil, bytesRead, err
			}
			bytesRead += n
			e = midi.NewMidiEvent(deltaTime, status, 2, data)

		case 0xC, 0xD:
			n, data, err := readPrefixed(r, prefix, 1)
			if err != nil {
				return nil, bytesRead, err
			}
			bytesRead += n
			e = midi.NewMidiEvent(deltaTime, status, 1, data)

		default:
			ctx.Log("warning", "unexpected channel voice msg")
		}
	} else if subStatus == 0xF {

		switch status & 0xF {

		case 0x0:
			n, data, err := readPrefixed(r, prefix, 1)
			if err != nil {
				return nil, bytesRead, err
			}
			bytesRead += n

			mada := data[0] != 0xF7
			for mada {
				n, tmp, err := readPrefixed(r, prefix, 1)
				if err != nil {
					return nil, bytesRead, err
				}
				bytesRead += n
				data = append(data, tmp[0])
				mada = tmp[0] != 0xF7
			}
			e = midi.NewMidiEvent(deltaTime, status, len(data), data)

		case 0x2:
			n, data, err := readPrefixed(r, prefix, 2)
			if err != nil {
				return nil, bytesRead, err
			}
			bytesRead += n
			e = midi.NewMidiEvent(deltaTime, status, 2, data)

		case 0x3:
			n, data, err := readPrefixed(r, prefix, 1)
			if err != nil {
				return nil, bytesRead, err
			}
			bytesRead += n
			e = midi.NewMidiEvent(deltaTime, status, 1, data)

		default:
			e = midi.NewMidiEvent(deltaTime, status, 0, make([]byte, 0))
		}
	} else {

		if prev == nil {
			return nil, bytesRead, errors.New("want to assume status, but no previous status exists")
		}

		if prevMidiEvent, ok := prev.(*midi.MidiEvent); ok {
			return readMidiEvent(r, deltaTime, prevMidiEvent.Status(), nil, []byte{status}, cfg)
		}

		return nil, bytesRead, errors.New("want to assume status, but previous event is not midi")
	}

	return
}

// Create and return a byte array of size=count, using prefix to fill the beginning of the array, and
// reading any remaining data from the Reader r
func readPrefixed(r io.Reader, prefix []byte, count int) (n int, data []byte, err error) {

	data = make([]byte, count)

	if prefix == nil {

		n, err = io.ReadFull(r, data)
		return
	} else {

		copy(data, prefix)
		if len(prefix) >= count {
			return
		}

		remaining := count - len(prefix)

		tmp := make([]byte, remaining)
		n, err = io.ReadFull(r, tmp)
		if err != nil {
			return
		}
		copy(data[len(prefix):], tmp)
		return
	}
}
