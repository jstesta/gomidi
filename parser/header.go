package parser

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"

	"github.com/jstesta/gomidi/cfg"
	"github.com/jstesta/gomidi/midi"
)

func readHeader(r io.Reader, c cfg.GomidiConfig) (h *midi.Header, err error) {

	chunkType := make([]byte, 4)
	_, err = io.ReadFull(r, chunkType)
	if err != nil {
		return
	}

	switch string(chunkType) {

	case HEADER_CHUNK_LITERAL:
		return readHeaderChunk(r, c)

	default:
		return nil, errors.New(fmt.Sprintf("wanted Header chunk but found %s", chunkType))
	}
}

func readHeaderChunk(r io.Reader, c cfg.GomidiConfig) (h *midi.Header, err error) {

	type header struct {
		Length         int32
		Format         int16
		NumberOfTracks int16
		Division       int16
	}

	var data header
	err = binary.Read(r, c.ByteOrder, &data)
	if err != nil {
		return
	}

	h = midi.NewHeader(
		int(data.Length),
		int(data.Format),
		int(data.NumberOfTracks),
		int(data.Division))
	return
}
