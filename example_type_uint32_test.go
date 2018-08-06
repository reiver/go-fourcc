package fourcc_test

import (
	"github.com/reiver/go-fourcc"

	"fmt"
)

func ExampleType_Uint32() {

	var colorspace fourcc.Type = fourcc.FOURCC('P','A','L','8')

	// Since:
	//
	// ASCII 'P' = 0x50
	// ASCII 'A' = 0x41
	// ASCII 'L' = 0x4C
	// ASCII '8' = 0x38
	//
	// Then, the FOURCC uint32 code, as a hexadecimal, should be those in reverse order.
	// I.e.,:
	//
	// 38,4C,41,50

	u32, err := colorspace.Uint32()
	if nil != err {
		fmt.Printf("THIS SHOULD NEVER HAPPEN: (%T) %v", err, err)
		return
	}

	fmt.Printf("%X", u32)

	// Output:
	// 384C4150
}
