package client

const (
	// Video Codecs
	VideoCodecH264 = "h264"
	VideoCodecAV1  = "av1"
	VideoCodecVP9  = "vp9"

	// Video Qualities
	VideoQuality144  = "144"
	VideoQuality240  = "240"
	VideoQuality360  = "360"
	VideoQuality480  = "480"
	VideoQuality720  = "720"
	VideoQuality1080 = "1080"
	VideoQuality1440 = "1440"
	VideoQuality4K   = "2160"
	VideoQualityMax  = "max"

	// Audio Formats
	AudioFormatMP3  = "mp3"
	AudioFormatOpus = "opus"
	AudioFormatOgg  = "ogg"
	AudioFormatWav  = "wav"
	AudioFormatBest = "best"

	// Statuses

	StatusError     = "error"
	StatusRedirect  = "redirect"
	StatusStream    = "stream"
	StatusSuccess   = "success"
	StatusRatelimit = "rate-limit"
	StatusPicker    = "picker"
)
