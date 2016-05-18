package parser

import (
	"encoding/binary"
	"io"

	"github.com/jstesta/gomidi/cfg"
	"github.com/jstesta/gomidi/midi"
)

const HEADER_CHUNK_LITERAL = "MThd"
const TRACK_CHUNK_LITERAL = "MTrk"

func ReadChunk(b io.Reader, cfg cfg.GomidiConfig) (c midi.Chunk, err error) {

	chunkType := make([]byte, 4)
	err = binary.Read(b, cfg.ByteOrder, chunkType)
	if err != nil {
		return nil, err
	}

	switch string(chunkType) {

	case HEADER_CHUNK_LITERAL:
		c, err = readHeaderChunk(b, cfg)
		cfg.LogContext.Log("chunk", c, "err", err)
		return

	case TRACK_CHUNK_LITERAL:
		c, err = readTrackChunk(b, cfg)
		cfg.LogContext.Log("chunk", c, "err", err)
		return

	default:
		err = readAlienChunk(b, cfg)
		return nil, err

	}
}
