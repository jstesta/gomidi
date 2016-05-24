package cfg

import (
	"encoding/binary"
	"log"
)

type GomidiConfig struct {
	ByteOrder binary.ByteOrder
	Log       *log.Logger
}
