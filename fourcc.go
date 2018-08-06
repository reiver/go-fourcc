package fourcc

// FOURCC turns 4 characters (where characters are uint8), into a (compact) machine readble version of a FOURCC.
func FOURCC(a uint8, b uint8, c uint8, d uint8) Type {
	return Type{
		value: Uint32(a, b, c, d),
	}
}
