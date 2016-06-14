package main

import (
	"flag"
	"log"
	"os"

	"github.com/jstesta/gomidi"
	"github.com/jstesta/gomidi/cfg"
)

func main() {
	var (
		midiFile = flag.String("input", "", "Filesystem location of MIDI file to parse")
	)
	flag.Parse()

	f, err := os.Create("out.log")
	if err != nil {
		log.Fatal(err)
	}
	logger := log.New(f, "gomidi ", log.LUTC|log.LstdFlags)
	//logger := log.New(os.Stdout, "gomidi ", log.LUTC|log.LstdFlags)
	logger.Printf("midiFile: %v", *midiFile)

	m, err := gomidi.ReadMidiFromFile(*midiFile, cfg.GomidiConfig{
		Log: logger,
	})
	if err != nil {
		logger.Fatal(err)
	}
	logger.Printf("midi: %v", m)
}
