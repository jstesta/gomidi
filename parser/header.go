package parser

import (
	"encoding/binary"
	"io"

	"github.com/jstesta/gomidi/midi"
)

func readHeaderChunk(b io.Reader) (c midi.Chunk, err error) {

	c = new(midi.Header)
	err = binary.Read(b, binary.BigEndian, c)
	return
}
