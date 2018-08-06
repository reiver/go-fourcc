package fourcc

// two planes -- one Y, one Cr + Cb interleaved
const (
	NV12 = Const("NV12") // depth: 12; description: Y/CbCr 4:2:0
	NV21 = Const("NV21") // depth: 12; description: Y/CrCb 4:2:0
	NV16 = Const("NV16") // depth: 16; description: Y/CbCr 4:2:2
	NV61 = Const("NV61") // depth: 16; description: Y/CrCb 4:2:2
	NV24 = Const("NV24") // depth: 24; description: Y/CbCr 4:4:4
	NV42 = Const("NV42") // depth: 24; description: Y/CrCb 4:4:4
)
