package fourcc

// Luminance+Chrominance formats
const (
	YUYV   = Const("YUYV") // depth: 16; description: YUV 4:2:2
	YYUV   = Const("YYUV") // depth: 16; description: YUV 4:2:2
	YVYU   = Const("YVYU") // depth: 16; description: YVU 4:2:2
	UYVY   = Const("UYVY") // depth: 16; description: YUV 4:2:2
	VYUY   = Const("VYUY") // depth: 16; description: YUV 4:2:2
	Y41P   = Const("Y41P") // depth: 12; description: YUV 4:1:1
	YUV444 = Const("Y444") // depth: 16; description: xxxxyyyy uuuuvvvv
	YUV555 = Const("YUVO") // depth: 16; description: YUV-5-5-5
	YUV565 = Const("YUVP") // depth: 16; description: YUV-5-6-5
	YUV32  = Const("YUV4") // depth: 32; description: YUV-8-8-8-8
	HI240  = Const("HI24") // depth:  8; description: 8-bit color
	HM12   = Const("HM12") // depth:  8; description: YUV 4:2:0 16x16 macroblocks
	M420   = Const("M420") // depth: 12; description: YUV 4:2:0 2 lines y, 1 line uv interleaved

)
