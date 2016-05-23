package parser

import (
	"encoding/binary"
	"io"

	"github.com/go-kit/kit/log"
	"github.com/jstesta/gomidi/cfg"
)

func readAlienChunk(r io.Reader, cfg cfg.GomidiConfig) (err error) {

	if cfg.ByteOrder == nil {
		cfg.ByteOrder = binary.BigEndian
	}

	if cfg.LogContext == nil {
		cfg.LogContext = log.NewContext(log.NewNopLogger())
	}

	var length uint32
	err = binary.Read(r, cfg.ByteOrder, &length)
	if err != nil {
		return
	}

	data := make([]byte, length)
	_, err = io.ReadFull(r, data)
	return
}
