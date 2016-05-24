package parser

import (
	"encoding/binary"
	"io"
	"io/ioutil"
	"log"

	"github.com/jstesta/gomidi/cfg"
)

func readAlienChunk(r io.Reader, cfg cfg.GomidiConfig) (err error) {

	if cfg.ByteOrder == nil {
		cfg.ByteOrder = binary.BigEndian
	}

	if cfg.Log == nil {
		cfg.Log = log.New(ioutil.Discard, "", 0)
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
