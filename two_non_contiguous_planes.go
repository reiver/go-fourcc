package fourcc

// Two non contiguous planes - one Y, one Cr + Cb interleaved
const (
	NV12M        = Const("NM12") // depth: 12; description: Y/CbCr 4:2:0
	NV21M        = Const("NM21") // depth: 21; description: Y/CrCb 4:2:0
	NV16M        = Const("NM16") // depth: 16; description: Y/CbCr 4:2:2
	NV61M        = Const("NM61") // depth: 16; description: Y/CrCb 4:2:2
	NV12MT       = Const("TM12") // depth: 12; description: Y/CbCr 4:2:0 64x32 macroblocks
	NV12MT_16X16 = Const("VM12") // depth: 12; description: Y/CbCr 4:2:0 16x16 macroblocks
)
