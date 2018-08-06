package fourcc_test

import (
	"github.com/reiver/go-fourcc"

	"encoding/json"
	"fmt"
)

func ExampleType_MarshalText() {

	type ScreenInfo struct {
		Width      uint64      `json:"width"`
		Height     uint64      `json:"height"`
		Colorspace fourcc.Type `json:"colorspace"`
	}

	var info ScreenInfo

	info.Width  = 640
	info.Height = 480
	info.Colorspace = fourcc.FOURCC('V','P','8','0')


	data, err := json.Marshal(&info) // calls MarshalText, to convert FOURCC into a text representation.
	if nil != err {
		fmt.Printf("Problem creating JSON data: (%T) %q", err, err)
		return
	}

	fmt.Printf("%s", data)

	// Output:
	// {"width":640,"height":480,"colorspace":"VP80"}
}
