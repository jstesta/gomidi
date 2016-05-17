package parser

import (
	"encoding/binary"
	"io"

	"github.com/jstesta/gomidi/cfg"
	"github.com/jstesta/gomidi/midi"
)

func readHeaderChunk(b io.Reader, cfg cfg.GomidiConfig) (c midi.Chunk, err error) {

	c = new(midi.Header)
	err = binary.Read(b, cfg.ByteOrder, c)
	return
}
