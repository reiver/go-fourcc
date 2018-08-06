package fourcc

// Three non contiguous planes - Y, Cb, Cr
const (
	YUV420M = Const("YM12") // depth: 12; description: YUV420 planar
	YVU420M = Const("YM21") // depth: 12; description: YVU420 planar
	YUV422M = Const("YM16") // depth: 16; description: YUV422 planar
	YVU422M = Const("YM61") // depth: 16; description: YVU422 planar
	YUV444M = Const("YM24") // depth: 24; description: YUV444 planar
	YVU444M = Const("YM42") // depth: 24; description: YVU444 planar
)
