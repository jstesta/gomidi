package main

import (
	"encoding/binary"
	"flag"
	stdlog "log"
	"os"

	kitlog "github.com/go-kit/kit/log"
	"github.com/jstesta/gomidi"
	"github.com/jstesta/gomidi/cfg"
)

func main() {
	var (
		midiFile = flag.String("input", "", "Filesystem location of MIDI file to parse")
	)
	flag.Parse()

	logger := kitlog.NewLogfmtLogger(os.Stderr)
	stdlog.SetOutput(kitlog.NewStdlibAdapter(logger))

	logger.Log("midiFile", midiFile)

	ctx := kitlog.NewContext(logger).WithPrefix("ts", kitlog.DefaultTimestampUTC)

	m, err := gomidi.ReadMidiFromFile(*midiFile, cfg.GomidiConfig{
		ByteOrder:  binary.BigEndian,
		LogContext: ctx,
	})
	if err != nil {
		stdlog.Fatal(err)
	}
	ctx.Log("midi", m)
}
