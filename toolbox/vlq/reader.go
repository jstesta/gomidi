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
func Read(buffer *bytes.Reader) (n int32, bytesRead uint32, err error) {

	mada := true
	for mada {
		// read the next byte
		b, err := buffer.ReadByte()
		if err != nil {
			return 0, bytesRead, &VLQReadError{
				originalBytes(buffer),
				"io error",
				err}
		}

		// increment read bytes counter
		bytesRead++

		// add the 7 LSBs to the result
		n ^= int32(b & 0x7f)

		// if the MSB is 1, prepare for the next iteration
		if mada = 1 == b>>7; mada {
			n <<= 7
		}

		// simple check for overflow
		if n < 0 {
			return 0, bytesRead, &VLQReadError{
				originalBytes(buffer),
				"exceeded maximum vlq value [0x0FFFFFFF]",
				nil}
		}
	}

	return
}

// rewinds the buffer and returns all bytes in it
func originalBytes(buffer *bytes.Reader) (out []byte) {

	out = make([]byte, buffer.Size())
	buffer.Seek(0, 0)
	buffer.Read(out)
	return
}
