package fourcc

// Uint32 turns 4 characters (where characters are uint8), into a (compact) machine readble version of a FOURCC (as an uint32).
//
// If you want stronger protection from the type system, use the FOURCC() func instead.
func Uint32(a uint8, b uint8, c uint8, d uint8) uint32 {

	var value uint32 = (uint32(a) | (uint32(b) << 8) | (uint32(c) << 16) | (uint32(d) << 24))

	return value
}
