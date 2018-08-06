package fourcc_test

import (
	"github.com/reiver/go-fourcc"

	"fmt"
)

func ExampleType_Scan_bytes() {


	var p []byte = []byte{'P','A','L','8'}

	var colorspace fourcc.Type

	colorspace.Scan(p)

	s := colorspace.String()

	u32, _ := colorspace.Uint32() // <--- We ignored the error here (to make the example easier to read), but you should NOT do that.

	fmt.Println(s)
	fmt.Printf("%X", u32)

	// Output:
	// PAL8
	// 384C4150
}

func ExampleType_Scan_int64() {


	var i64 int64 = 0x384C4150

	var colorspace fourcc.Type

	colorspace.Scan(i64)

	s := colorspace.String()

	u32, _ := colorspace.Uint32() // <--- We ignored the error here (to make the example easier to read), but you should NOT do that.

	fmt.Println(s)
	fmt.Printf("%X", u32)

	// Output:
	// PAL8
	// 384C4150
}


func ExampleType_Scan_string() {

	var colorspace fourcc.Type

	colorspace.Scan("PAL8")

	s := colorspace.String()

	u32, _ := colorspace.Uint32() // <--- We ignored the error here (to make the example easier to read), but you should NOT do that.

	fmt.Println(s)
	fmt.Printf("%X", u32)

	// Output:
	// PAL8
	// 384C4150
}

func ExampleType_Scan_uint32() {


	var ui32 uint32 = 0x384C4150

	var colorspace fourcc.Type

	colorspace.Scan(ui32)

	s := colorspace.String()

	u32, _ := colorspace.Uint32() // <--- We ignored the error here (to make the example easier to read), but you should NOT do that.

	fmt.Println(s)
	fmt.Printf("%X", u32)

	// Output:
	// PAL8
	// 384C4150
}

func ExampleType_Scan_uint64() {


	var ui64 uint64 = 0x384C4150

	var colorspace fourcc.Type

	colorspace.Scan(ui64)

	s := colorspace.String()

	u32, _ := colorspace.Uint32() // <--- We ignored the error here (to make the example easier to read), but you should NOT do that.

	fmt.Println(s)
	fmt.Printf("%X", u32)

	// Output:
	// PAL8
	// 384C4150
}
