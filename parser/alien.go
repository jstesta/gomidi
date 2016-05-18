package parser

import (
	"encoding/binary"
	"github.com/jstesta/gomidi/cfg"
	"io"
)

func readAlienChunk(r io.Reader, cfg cfg.GomidiConfig) (err error) {

	var length uint32
	err = binary.Read(r, cfg.ByteOrder, &length)
	if err != nil {
		return
	}

	data := make([]byte, length)
	_, err = r.Read(data)
	return
}
