/*
Package vlq implements a library for handling variable-length quantity (vlq)
values specified by the MIDI file format.

See: https://www.midi.org/images/downloads/complete_midi_96-1-3.pdf Standard MIDI Files 1.0
*/
package vlq

import (
	"bytes"
)

// Read reads a single VLQ value from a bytes.Reader
func Read(buffer *bytes.Reader) (n int32, err error) {

	mada := true
	for mada {
		b, err := buffer.ReadByte()
		if err != nil {
			return 0, &VLQReadError{
				originalBytes(buffer),
				"io error",
				err}
		}

		mada = 1 == b>>7
		if mada {
			n ^= int32(b & 0x7f)
			n <<= 7
		} else {
			n ^= int32(b)
		}

		// simple check for overflow
		if n < 0 {
			return 0, &VLQReadError{
				originalBytes(buffer),
				"exceeded maximum vlq value [0x0FFFFFFF]",
				nil}
		}
	}

	return
}

func originalBytes(buffer *bytes.Reader) (out []byte) {
	out = make([]byte, buffer.Size())
	buffer.Seek(0, 0)
	buffer.Read(out)
	return
}
