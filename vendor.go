package fourcc

// Vendor-specific formats
const (
	CPIA1        = Const("CPIA") // cpia1 YUV
	WNVA         = Const("WNVA") // Winnov hw compress
	SN9C10X      = Const("S910") // SN9C10x compression
	SN9C20X_I420 = Const("S920") // SN9C20x YUV 4:2:0
	PWC1         = Const("PWC1") // pwc older webcam
	PWC2         = Const("PWC2") // pwc newer webcam
	ET61X251     = Const("E625") // ET61X251 compression
	SPCA501      = Const("S501") // YUYV per line
	SPCA505      = Const("S505") // YYUV per line
	SPCA508      = Const("S508") // YUVY per line
	SPCA561      = Const("S561") // compressed GBRG bayer
	PAC207       = Const("P207") // compressed BGGR bayer
	MR97310A     = Const("M310") // compressed BGGR bayer
	JL2005BCD    = Const("JL20") // compressed RGGB bayer
	SN9C2028     = Const("SONX") // compressed GBRG bayer
	SQ905C       = Const("905C") // compressed RGGB bayer
	PJPG         = Const("PJPG") // Pixart 73xx JPEG
	OV511        = Const("O511") // ov511 JPEG
	OV518        = Const("O518") // ov518 JPEG
	STV0680      = Const("S680") // stv0680 bayer
	TM6000       = Const("TM60") // tm5600/tm60x0
	CIT_YYVYUY   = Const("CITV") // one line of Y then 1 line of VYUY
	KONICA420    = Const("KONI") // YUV420 planar in blocks of 256 pixels
	JPGL         = Const("JPGL") // JPEG-Lite
	SE401        = Const("S401") // se401 janggu compressed rgb
	S5C_UYVY_JPG = Const("S5CI") // S5C73M3 interleaved UYVY/JPEG
	Y8I          = Const("Y8I ") // Greyscale 8-bit L/R interleaved
	Y12I         = Const("Y12I") // Greyscale 12-bit L/R interleaved
	Z16          = Const("Z16 ") // Depth data 16-bit
	MT21C        = Const("MT21") // Mediatek compressed block mode
	INZI         = Const("INZI") // Intel Planar Greyscale 10-bit and Depth 16-bit
)
