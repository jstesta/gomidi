package main

import (
	kitlog "github.com/go-kit/kit/log"
	"github.com/jstesta/gomidi"
	stdlog "log"
	"os"
)

func main() {

	logger := kitlog.NewLogfmtLogger(os.Stderr)
	stdlog.SetOutput(kitlog.NewStdlibAdapter(logger))

	m, err := gomidi.ReadMidiFromFile("resource/walzamin.mid")
	if err != nil {
		stdlog.Fatal(err)
	}
	logger.Log("midi", m)
}
