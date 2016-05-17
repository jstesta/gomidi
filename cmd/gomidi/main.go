package main

import (
	"encoding/binary"
	kitlog "github.com/go-kit/kit/log"
	"github.com/jstesta/gomidi"
	"github.com/jstesta/gomidi/cfg"
	stdlog "log"
	"os"
)

func main() {

	logger := kitlog.NewLogfmtLogger(os.Stderr)
	stdlog.SetOutput(kitlog.NewStdlibAdapter(logger))

	ctx := kitlog.NewContext(logger).WithPrefix("ts", kitlog.DefaultTimestampUTC)

	m, err := gomidi.ReadMidiFromFile("resource/walzamin.mid", cfg.GomidiConfig{
		ByteOrder:  binary.BigEndian,
		LogContext: ctx,
	})
	if err != nil {
		stdlog.Fatal(err)
	}
	ctx.Log("midi", m)
}
