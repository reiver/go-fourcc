package fourcc_test

import (
	"github.com/reiver/go-fourcc"

	"fmt"
)

func ExampleUint32() {

	var colorspace uint32 = fourcc.Uint32('P','A','L','8')

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

	fmt.Printf("%X", colorspace)

	// Output:
	// 384C4150
}
