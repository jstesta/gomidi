package parser

import (
	"encoding/binary"
	"errors"
	"io"

	"github.com/jstesta/gomidi/cfg"
	"github.com/jstesta/gomidi/midi"
	"github.com/jstesta/gomidi/util"
	"github.com/jstesta/gomidi/vlq"
)

func readTrackChunk(r io.Reader, cfg cfg.GomidiConfig) (c midi.Chunk, err error) {

	ctx := cfg.LogContext.With("reader", "track")

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
			ctx.Log("event", event)
			events = append(events, event)
			previousEvent = event

		case 0xFF:
			event, br, err := readMetaEvent(r, deltaTime, cfg)
			if err != nil {
				return nil, err
			}
			bytesRead += br
			ctx.Log("event", event)
			events = append(events, event)
			previousEvent = event

		default:
			event, br, err := readMidiEvent(r, deltaTime, eventType, previousEvent, nil, cfg)
			if err != nil {
				return nil, err
			}
			bytesRead += br
			ctx.Log("event", event)
			events = append(events, event)
			previousEvent = event

		}

		if bytesRead > int(length) {
			return nil, errors.New("parsed too many bytes")
		}
	}

	return &midi.Track{events}, nil
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

	e = midi.NewSysexEvent(deltaTime, data)

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
	bytesRead += length

	e = midi.NewMetaEvent(deltaTime, metaEventType, data)

	return
}

func readMidiEvent(r io.Reader, deltaTime int, status byte, prev midi.Event, prefix []byte, cfg cfg.GomidiConfig) (e midi.Event, bytesRead int, err error) {

	subStatus := status >> 4
	ctx := cfg.LogContext.With("reader", "midi", "status", toolbox.ToBitString(status), "subStatus", toolbox.ToBitString(subStatus))

	if subStatus < 0xF && subStatus >= 0x8 {

		switch subStatus {

		case 0x8, 0x9, 0xA, 0xB, 0xE:
			n, data, err := readSome(r, prefix, 2)
			if err != nil {
				return nil, bytesRead, err
			}
			bytesRead += n
			e = midi.NewMidiEvent(deltaTime, status, data)

		case 0xC, 0xD:
			n, data, err := readSome(r, prefix, 1)
			if err != nil {
				return nil, bytesRead, err
			}
			bytesRead += n
			e = midi.NewMidiEvent(deltaTime, status, data)

		default:
			ctx.Log("warning", "unexpected channel voice msg")
		}
	} else if subStatus == 0xF {
		switch status & 0xF {

		case 0x0:
			n, w, err := readSome(r, prefix, 1)
			if err != nil {
				return nil, bytesRead, err
			}
			bytesRead += n

			mada := w[0] != 0xF7
			for mada {
				n, w, err := readSome(r, prefix, 1)
				if err != nil {
					return nil, bytesRead, err
				}
				bytesRead += n
				mada = w[0] != 0xF7
			}

		case 0x2:
			n, data, err := readSome(r, prefix, 2)
			if err != nil {
				return nil, bytesRead, err
			}
			bytesRead += n
			e = midi.NewMidiEvent(deltaTime, status, data)

		case 0x3:
			n, data, err := readSome(r, prefix, 1)
			if err != nil {
				return nil, bytesRead, err
			}
			bytesRead += n
			e = midi.NewMidiEvent(deltaTime, status, data)

		default:
			ctx.Log("aaaa", "other")
		}
	} else {
		if prev == nil {
			return
		}

		if prevMidiEvent, ok := prev.(*midi.MidiEvent); ok {
			return readMidiEvent(r, deltaTime, prevMidiEvent.Status(), nil, []byte{status}, cfg)
		}

		ctx.Log("action", "this is bad! fixme")
	}

	return
}

func readSome(r io.Reader, prefix []byte, count int) (n int, data []byte, err error) {
	data = make([]byte, count)

	if prefix == nil {

		n, err = io.ReadFull(r, data)
		return
	} else {

		copy(data, prefix)
		if len(prefix) == count {
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
