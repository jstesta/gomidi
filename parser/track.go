package parser

import (
	"encoding/binary"
	"errors"
	"io"

	"github.com/jstesta/gomidi/cfg"
	"github.com/jstesta/gomidi/midi"
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
			event, br, err := readMidiEvent(r, deltaTime, eventType, previousEvent, cfg)
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

	ctx := cfg.LogContext.With("reader", "sysex")

	length, br, err := vlq.ReadVLQ(r)
	if err != nil {
		return
	}
	bytesRead += br

	ctx.Log("datalen", length)
	data := make([]byte, length)
	_, err = io.ReadFull(r, data)
	bytesRead += length

	e = midi.NewSysexEvent(deltaTime, data)

	return
}

func readMetaEvent(r io.Reader, deltaTime int, cfg cfg.GomidiConfig) (e midi.Event, bytesRead int, err error) {

	ctx := cfg.LogContext.With("reader", "meta")

	b := []byte{0}
	_, err = io.ReadFull(r, b)
	if err != nil {
		return
	}
	metaEventType := b[0]
	bytesRead++
	ctx.Log("type", metaEventType)

	length, br, err := vlq.ReadVLQ(r)
	if err != nil {
		return
	}
	bytesRead += br

	ctx.Log("datalen", length)
	data := make([]byte, length)
	_, err = io.ReadFull(r, data)
	bytesRead += length

	e = midi.NewMetaEvent(deltaTime, metaEventType, data)

	return
}

func readMidiEvent(r io.Reader, deltaTime int, status byte, prev midi.Event, cfg cfg.GomidiConfig) (e midi.Event, bytesRead int, err error) {

	ctx := cfg.LogContext.With("reader", "midi")

	switch status >> 4 {

	case 0x8, 0x9, 0xA, 0xB, 0xE:
		data := make([]byte, 2)
		_, err = io.ReadFull(r, data)
		if err != nil {
			return
		}
		bytesRead += 2
		e = midi.NewMidiEvent(deltaTime, status, data)

	case 0xC, 0xD:
		data := make([]byte, 1)
		_, err = io.ReadFull(r, data)
		if err != nil {
			return
		}
		bytesRead += 1
		e = midi.NewMidiEvent(deltaTime, status, data)

	default:
		if prev == nil {
			ctx.Log("skipped unknown MIDI event, type", status>>4)
		}

		if prevMidiEvent, ok := prev.(*midi.MidiEvent); ok {
			ctx.Log("no status MIDI event, using previous status of: %X", prevMidiEvent.Status())
			return readMidiEvent(r, deltaTime, prevMidiEvent.Status(), nil, cfg)
		}

		ctx.Log("this is bad! fixme")

	}

	return
}
