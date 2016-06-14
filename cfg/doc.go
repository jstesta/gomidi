// Package cfg defines configuration passed to the MIDI parser.
//
// Configuration Listing:
//
// * ByteOrder
//   Allows overriding the default ByteOrder used when reading
//   numbers from the binary stream.
//   (DEFAULT): binary.BigEndian
//
// * Log
//   Allows overriding the default Logger used to log program
//   messages.
//   (DEFAULT): no logging
package cfg
