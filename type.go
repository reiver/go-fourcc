package fourcc

// Type represents the compact way of storing a FOURCC.
//
// The compact way of storing a FOURCC is as a uint32.
//
// Effectively Type is just an uint32.
type Type struct {
	value uint32
}

// MarshalText makes Type fit the encoding.TextUnmarshaler interface.
func (receiver Type) MarshalText() ([]byte, error) {
	return []byte(receiver.String()), nil
}

// String returns a human readable version of a FOURCC.
func (receiver Type) String() string {
	var buffer [4]byte

	buffer[0] = byte((receiver.value & 0x000000ff) >>  0)
	buffer[1] = byte((receiver.value & 0x0000ff00) >>  8)
	buffer[2] = byte((receiver.value & 0x00ff0000) >> 16)
	buffer[3] = byte((receiver.value & 0xff000000) >> 24)

	return string(buffer[:])
}

// Uint32 returns a (compact) machine readble version of a FOURCC.
func (receiver Type) Uint32() (uint32, error) {
	return receiver.value, nil
}
