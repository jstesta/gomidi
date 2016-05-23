package parser

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"

	"github.com/go-kit/kit/log"
	"github.com/jstesta/gomidi/cfg"
	"github.com/jstesta/gomidi/midi"
)

func ReadHeader(r io.Reader, cfg cfg.GomidiConfig) (h *midi.Header, err error) {

	if cfg.ByteOrder == nil {
		cfg.ByteOrder = binary.BigEndian
	}

	if cfg.LogContext == nil {
		cfg.LogContext = log.NewContext(log.NewNopLogger())
	}

	chunkType := make([]byte, 4)
	_, err = io.ReadFull(r, chunkType)
	if err != nil {
		return
	}

	switch string(chunkType) {

	case HEADER_CHUNK_LITERAL:
		return readHeaderChunk(r, cfg)

	default:
		return nil, errors.New(fmt.Sprintf("wanted Header chunk but found %s", chunkType))
	}
}

func readHeaderChunk(r io.Reader, cfg cfg.GomidiConfig) (h *midi.Header, err error) {

	var length int32
	err = binary.Read(r, cfg.ByteOrder, &length)
	if err != nil {
		return
	}

	var format int16
	err = binary.Read(r, cfg.ByteOrder, &format)
	if err != nil {
		return
	}

	var numberOfTracks int16
	err = binary.Read(r, cfg.ByteOrder, &numberOfTracks)
	if err != nil {
		return
	}

	var division int16
	err = binary.Read(r, cfg.ByteOrder, &division)
	if err != nil {
		return
	}

	h = midi.NewHeader(int(length), int(format), int(numberOfTracks), int(division))
	return
}
