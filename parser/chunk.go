package parser

import (
	"io"

	"github.com/jstesta/gomidi/cfg"
	"github.com/jstesta/gomidi/midi"
)

const HEADER_CHUNK_LITERAL = "MThd"
const TRACK_CHUNK_LITERAL = "MTrk"

func ReadChunk(r io.Reader, cfg cfg.GomidiConfig) (c midi.Chunk, err error) {

	chunkType := make([]byte, 4)
	_, err = io.ReadFull(r, chunkType)
	if err != nil {
		return nil, err
	}

	switch string(chunkType) {

	case HEADER_CHUNK_LITERAL:
		c, err = readHeaderChunk(r, cfg)
		cfg.LogContext.Log("chunk", c, "err", err)
		return

	case TRACK_CHUNK_LITERAL:
		c, err = readTrackChunk(r, cfg)
		cfg.LogContext.Log("chunk", c, "err", err)
		return

	default:
		err = readAlienChunk(r, cfg)
		return nil, err

	}
}
