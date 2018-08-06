package fourcc

const (

	RGB332  = Const("RGB1") // depth:  8; description: RGB-3-3-2
	RGB444  = Const("R444") // depth: 16; description: xxxxrrrr ggggbbbb
	ARGB444 = Const("AR12") // depth: 16; description: aaaarrrr ggggbbbb
	XRGB444 = Const("XR12") // depth: 16; description: xxxxrrrr ggggbbbb
	RGB555  = Const("RGBO") // depth: 16; description: RGB-5-5-5
	ARGB555 = Const("AR15") // depth: 16; description: ARGB-1-5-5-5
	XRGB555 = Const("XR15") // depth: 16; description: XRGB-1-5-5-5
	RGB565  = Const("RGBP") // depth: 16; description: RGB-5-6-5
	RGB555X = Const("RGBQ") // depth: 16; description: RGB-5-5-5 BE

	RGB565X = Const("RGBR") // depth: 16; description: RGB-5-6-5 BE
	BGR666  = Const("BGRH") // depth: 18; description: BGR-6-6-6
	BGR24   = Const("BGR3") // depth: 24; description: BGR-8-8-8
	RGB24   = Const("RGB3") // depth: 24; description: RGB-8-8-8
	BGR32   = Const("BGR4") // depth: 32; description: BGR-8-8-8-8
	ABGR32  = Const("AR24") // depth: 32; description: BGRA-8-8-8-8
	XBGR32  = Const("XR24") // depth: 32; description: BGRX-8-8-8-8
	RGB32   = Const("RGB4") // depth: 32; description: RGB-8-8-8-8
	ARGB32  = Const("BA24") // depth: 32; description: AR GB-8-8-8-8
	XRGB32  = Const("BX24") // depth: 32; description: XRGB-8-8-8-8
)
