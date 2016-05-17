package parser

import (
	"encoding/binary"
	"io"
	"log"

	"github.com/jstesta/gomidi/midi"
)

const MTHD_CHUNK_LITERAL = "MThd"
const MTRK_CHUNK_LITERAL = "MTrk"

func ReadChunk(b io.Reader) (c midi.Chunk, err error) {

	chunkType := make([]byte, 4)
	err = binary.Read(b, binary.BigEndian, chunkType)
	if err != nil {
		return nil, err
	}

	switch string(chunkType) {

	case MTHD_CHUNK_LITERAL:
		c, err = ReadHeaderChunk(b)
		log.Printf("chunk.ReadChunk parsed Header chunk: %v, err: %v", c, err)
		return

	case MTRK_CHUNK_LITERAL:
		c, err = ReadTrackChunk(b)
		log.Printf("chunk.ReadChunk parsed Track chunk: %v, err: %v", c, err)
		return

	default:
		log.Print("chunk.ReadChunk found Alien chunk... skipping")
	}

	return
}
