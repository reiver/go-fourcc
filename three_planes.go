package fourcc

// Three planes - Y Cb, Cr
const (
	YUV410  = Const("YUV9") // depth:  9; description: YUV 4:1:0
	YVU410  = Const("YVU9") // depth:  9; description: YVU 4:1:0
	YUV411P = Const("411P") // depth: 12; description: YVU411 planar
	YUV420  = Const("YU12") // depth: 12; description: YUV 4:2:0
	YVU420  = Const("YV12") // depth: 12; description: YVU 4:2:0
	YUV422P = Const("422P") // depth: 16; description: YVU422 planar
)
