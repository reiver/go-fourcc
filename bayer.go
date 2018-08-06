package fourcc

// Bayer formats - see http://www.siliconimaging.com/RGB%20Bayer.htm
const (
	SBGGR8  = Const("BA81") //  8  BGBG.. GRGR..
	SGBRG8  = Const("GBRG") //  8  GBGB.. RGRG..
	SGRBG8  = Const("GRBG") //  8  GRGR.. BGBG..
	SRGGB8  = Const("RGGB") //  8  RGRG.. GBGB..
	SBGGR10 = Const("BG10") // 10  BGBG.. GRGR..
	SGBRG10 = Const("GB10") // 10  GBGB.. RGRG..
	SGRBG10 = Const("BA10") // 10  GRGR.. BGBG..
	SRGGB10 = Const("RG10") // 10  RGRG.. GBGB..

        // 10bit raw bayer packed, 5 bytes for every 4 pixels
	SBGGR10P = Const("pBAA")
	SGBRG10P = Const("pGAA")
	SGRBG10P = Const("pgAA")
	SRGGB10P = Const("pRAA")

        // 10bit raw bayer a-law compressed to 8 bits
	SBGGR10ALAW8 = Const("aBA8")
	SGBRG10ALAW8 = Const("aGA8")
	SGRBG10ALAW8 = Const("agA8")
	SRGGB10ALAW8 = Const("aRA8")

        // 10bit raw bayer DPCM compressed to 8 bits
	SBGGR10DPCM8 = Const("bBA8")
	SGBRG10DPCM8 = Const("bGA8")
	SGRBG10DPCM8 = Const("BD10")
	SRGGB10DPCM8 = Const("bRA8")
	SBGGR12      = Const("BG12") // 12  BGBG.. GRGR..
	SGBRG12      = Const("GB12") // 12  GBGB.. RGRG..
	SGRBG12      = Const("BA12") // 12  GRGR.. BGBG..
	SRGGB12      = Const("RG12") // 12  RGRG.. GBGB..

        // 12bit raw bayer packed, 6 bytes for every 4 pixels
	SBGGR12P = Const("pBCC")
	SGBRG12P = Const("pGCC")
	SGRBG12P = Const("pgCC")
	SRGGB12P = Const("pRCC")
	SBGGR16  = Const("BYR2") // 16  BGBG.. GRGR..
	SGBRG16  = Const("GB16") // 16  GBGB.. RGRG..
	SGRBG16  = Const("GR16") // 16  GRGR.. BGBG..
	SRGGB16  = Const("RG16") // 16  RGRG.. GBGB..
)
