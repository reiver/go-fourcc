package fourcc_test

import (
	"github.com/reiver/go-fourcc"

	"fmt"
)

func ExampleConst_FOURCC() {

	const (
		PAL8 = fourcc.Const("PAL8")
	)

	var colorspace fourcc.Type
	var err error

	colorspace, err = PAL8.FOURCC()
	if nil != err {
		fmt.Printf("Problem with FOURCC Const: (%T) %q", err, err)
		return
	}

	s := colorspace.String()

	u32, _ := colorspace.Uint32() // <--- We ignored the error here (to make the example easier to read), but you should NOT do that.

	fmt.Println(s)
	fmt.Printf("%X", u32)

	// Output:
	// PAL8
	// 384C4150
}
