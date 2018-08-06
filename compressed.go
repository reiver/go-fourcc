package fourcc

// Compressed formats
const (
	MJPEG       = Const("MJPG") // Motion-JPEG
	JPEG        = Const("JPEG") // JFIF JPEG
	DV          = Const("dvsd") // 1394
	MPEG        = Const("MPEG") // MPEG-1/2/4 Multiplexed
	H264        = Const("H264") // H264 with start codes
	H264_NO_SC  = Const("AVC1") // H264 without start codes
	H264_MVC    = Const("M264") // H264 MVC
	H263        = Const("H263") // H263
	MPEG1       = Const("MPG1") // MPEG-1 ES
	MPEG2       = Const("MPG2") // MPEG-2 ES
	MPEG4       = Const("MPG4") // MPEG-4 part 2 ES
	XVID        = Const("XVID") // Xvid
	VC1_ANNEX_G = Const("VC1G") // SMPTE 421M Annex G compliant stream
	VC1_ANNEX_L = Const("VC1L") // SMPTE 421M Annex L compliant stream
	VP8         = Const("VP80") // VP8
	VP9         = Const("VP90") // VP9
)
