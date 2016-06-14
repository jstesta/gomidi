package parser

import (
	"encoding/binary"
	"io"

	"github.com/jstesta/gomidi/cfg"
)

func readAlienChunk(r io.Reader, c cfg.GomidiConfig) (err error) {

	var length uint32
	err = binary.Read(r, c.ByteOrder, &length)
	if err != nil {
		return
	}

	data := make([]byte, length)
	_, err = io.ReadFull(r, data)
	return
}
