package fourcc

import (
	"fmt"
)

// Const exists as a workaround for Go not supporting const structs.
//
// This, you can use this to create a const FOURCC.
type Const string

func (receiver Const) FOURCC() (Type, error) {
	value, err := receiver.Uint32()
	if nil != err {
		return Type{}, err
	}

	return Type{
		value: value,
	}, nil
}

func (receiver Const) Uint32() (uint32, error) {
	if 4 != len(receiver) {
		return 0, fmt.Errorf("fourcc: not a valid FOURCC code: %q", receiver)
	}

	a := receiver[0]
	b := receiver[1]
	c := receiver[2]
	d := receiver[3]

	var value uint32 = Uint32(a, b, c, d)

	return value, nil
}
