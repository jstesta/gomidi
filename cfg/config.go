package cfg

import (
	"encoding/binary"

	"github.com/go-kit/kit/log"
)

type GomidiConfig struct {
	ByteOrder  binary.ByteOrder
	LogContext *log.Context
}
