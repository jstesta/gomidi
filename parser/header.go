package parser

import (
	"encoding/binary"
	"io"

	"errors"
	"fmt"
	"github.com/jstesta/gomidi/cfg"
	"github.com/jstesta/gomidi/midi"
)

func ReadHeader(r io.Reader, cfg cfg.GomidiConfig) (h *midi.Header, err error) {

	chunkType := make([]byte, 4)
	_, err = io.ReadFull(r, chunkType)
	if err != nil {
		return
	}

	switch string(chunkType) {

	case HEADER_CHUNK_LITERAL:
		h = new(midi.Header)
		err = binary.Read(r, cfg.ByteOrder, h)
		return

	default:
		return nil, errors.New(fmt.Sprintf("wanted Header chunk but found %s", chunkType))
	}
}
