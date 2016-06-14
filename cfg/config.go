package cfg

import (
	"encoding/binary"
	"log"
)

type GomidiConfig struct {
	ByteOrder binary.ByteOrder
	Log       *log.Logger
}

var DefaultByteOrder = binary.BigEndian
