package settings

import (
	"encoding/json"
	"os"

	"github.com/timelessnesses/gobalt/client"
	"github.com/urfave/cli/v2"
	"github.com/valyala/fastjson"
)

type Settings struct {
	VCodec          string
	VQuality        string
	AFormat         string
	IsAudioOnly     bool
	IsNoTTWatermark bool
	IsTTFullAudio   bool
	IsAudioMuted    bool
	DubLang         bool
	DisableMetadata bool
	Endpoint        string
}

func GetCurrentFolderPath() string {
	dir, e := os.Getwd()
	if e != nil {
		panic(e)
	}
	return dir
}

func Init() error {
	file, err := os.OpenFile(GetCurrentFolderPath()+"/config.json", os.O_CREATE|os.O_RDWR, os.ModeDevice)
	if err != nil {
		panic(err)
	}
	jsonized, _ := json.MarshalIndent(Settings{
		VQuality:        client.VideoQualityMax,
		VCodec:          client.VideoCodecH264,
		AFormat:         client.AudioFormatBest,
		IsAudioOnly:     false,
		IsNoTTWatermark: false,
		IsTTFullAudio:   true,
		IsAudioMuted:    false,
		DubLang:         false,
		DisableMetadata: false,
		Endpoint:        "https://co.wuk.sh",
	}, "", "	")
	file.Write(jsonized)
	return nil
}

func ValidateSettings(s Settings) {
	if _, err := client.NewAudioFormat(s.AFormat); err != nil {
		panic("Validation Error: Audio Format is not supported.")
	}
	if _, err := client.NewVideoCodec(s.VCodec); err != nil {
		panic("Validation Error: Video Codec is not supported.")
	}
	if _, err := client.NewVideoQuality(s.VQuality); err != nil {
		panic("Validation Error: Video Quality is not supported.")
	}
}

func GetSettings(path string) Settings {
	if path == "" {
		path = GetCurrentFolderPath() + "/config.json"
	}
	file, err := os.ReadFile(path)
	if err != nil {
		Init()
		return GetSettings("")
	}
	parsed, err := fastjson.ParseBytes(file)
	if err != nil {
		panic(err)
	}
	r := Settings{
		VCodec:          string(parsed.GetStringBytes("VCodec")),
		VQuality:        string(parsed.GetStringBytes("VQuality")),
		AFormat:         string(parsed.GetStringBytes("AFormat")),
		IsAudioOnly:     parsed.GetBool("IsAudioOnly"),
		IsNoTTWatermark: parsed.GetBool("IsNoTTWatermark"),
		IsTTFullAudio:   parsed.GetBool("IsTTFullAudio"),
		IsAudioMuted:    parsed.GetBool("IsAudioMuted"),
		DubLang:         parsed.GetBool("DubLang"),
		DisableMetadata: parsed.GetBool("DisableMetadata"),
		Endpoint:        string(parsed.GetStringBytes("Endpoint")),
	}
	ValidateSettings(r)
	return r
}

func WriteSettings(settings Settings, path string) error {
	if path == "" {
		path = GetCurrentFolderPath() + "/config.json"
	}
	jsonized, _ := json.MarshalIndent(settings, "", "    ")
	err := os.WriteFile(path, jsonized, 0644)
	if err != nil {
		panic(err)
	}
	return nil
}

func IsSettingsExists(path string) bool {
	_, e := os.ReadFile(path)
	return e == nil
}

func Save(ctx *cli.Context) error {
	config := GetSettings(ctx.String("configPath"))
	if config.AFormat != ctx.String("aFormat") && ctx.String("aFormat") != "" {
		config.AFormat = ctx.String("aFormat")
	}
	if config.Endpoint != ctx.String("endpoint") && ctx.String("endpoint") != "" {
		config.Endpoint = ctx.String("endpoint")
	}
	if config.VCodec != ctx.String("vCodec") && ctx.String("vCodec") != "" {
		config.VCodec = ctx.String("vCodec")
	}
	if config.VQuality != ctx.String("vQuality") && ctx.String("vQuality") != "" {
		config.VQuality = ctx.String("vQuality")
	}
	if (config.DisableMetadata != ctx.Bool("disableMetadata") && !ctx.Bool("disableMetadata")) || (config.DisableMetadata != ctx.Bool("disableMetadata") && ctx.Bool("disableMetadata")) {
		config.DisableMetadata = ctx.Bool("disableMetadata")
	}
	if (config.DubLang != ctx.Bool("dubLang") && !ctx.Bool("dubLang")) || (config.DubLang != ctx.Bool("dubLang") && ctx.Bool("dubLang")) {
		config.DubLang = ctx.Bool("dubLang")
	}
	if (config.IsAudioMuted != ctx.Bool("isAudioMuted") && !ctx.Bool("isAudioMuted")) || (config.IsAudioMuted != ctx.Bool("isAudioMuted") && ctx.Bool("isAudioMuted")) {
		config.IsAudioMuted = ctx.Bool("isAudioMuted")
	}
	if (config.IsAudioOnly != ctx.Bool("isAudioOnly") && !ctx.Bool("isAudioOnly")) || (config.IsAudioOnly != ctx.Bool("isAudioOnly") && ctx.Bool("isAudioOnly")) {
		config.IsAudioOnly = ctx.Bool("isAudioOnly")
	}
	if (config.IsNoTTWatermark != ctx.Bool("isNoTTWatermark") && !ctx.Bool("isNoTTWatermark")) || (config.IsNoTTWatermark != ctx.Bool("isNoTTWatermark") && ctx.Bool("isNoTTWatermark")) {
		config.IsNoTTWatermark = ctx.Bool("isNoTTWatermark")
	}
	if (config.IsTTFullAudio != ctx.Bool("isTTFullAudio") && !ctx.Bool("isTTFullAudio")) || (config.IsTTFullAudio != ctx.Bool("isNoTTWatermark") && ctx.Bool("isTTFullAudio")) {
		config.IsTTFullAudio = ctx.Bool("isTTFullAudio")
	}

	return WriteSettings(config, ctx.String("configPath"))
}
