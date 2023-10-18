package client

import (
	"errors"
)

type Status struct {
	status string
}

func NewStatus(status string) (Status, error) {
	if status != StatusError && status != StatusPicker && status != StatusRatelimit && status != StatusRedirect && status != StatusStream && status != StatusSuccess {
		return Status{}, errors.New("The Status Is Not Supported.")
	}
	return Status{status}, nil
}

type VideoCodec struct {
	vCodec string
}

func NewVideoCodec(codec string) (VideoCodec, error) {
	if codec != VideoCodecAV1 && codec != VideoCodecH264 && codec != VideoCodecVP9 { // ?? what the fuck??
		return VideoCodec{}, errors.New("The Video Codec Is Not Supported.")
	}
	return VideoCodec{codec}, nil
}

type VideoQuality struct {
	vQuality string
}

func NewVideoQuality(quality string) (VideoQuality, error) {
	if quality != VideoQuality144 && quality != VideoQuality240 && quality != VideoQuality360 && quality != VideoQuality480 && quality != VideoQuality720 && quality != VideoQuality1080 && quality != VideoQuality1440 && quality != VideoQuality4K && quality != VideoQualityMax {
		return VideoQuality{}, errors.New("The quality is not supported.")
	}
	return VideoQuality{quality}, nil
}

type AudioFormat struct {
	format string
}

func NewAudioFormat(format string) (AudioFormat, error) {
	if format != AudioFormatMP3 && format != AudioFormatOgg && format != AudioFormatOpus && format != AudioFormatWav && format != AudioFormatBest {
		return AudioFormat{}, errors.New("The audio format is not supported.")
	}
	return AudioFormat{format}, nil
}
