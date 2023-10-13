package settings

import (
	"encoding/json"
	"os"

	"github.com/timelessnesses/cobalt-go/client"
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

func Init() error {
	file, err := os.OpenFile("config.json", os.O_CREATE|os.O_WRONLY, os.ModeDevice)
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

func GetSettings(path string) Settings {
	if path == "" {
		path = "./config.json"
	}
	file, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	parsed, err := fastjson.ParseBytes(file)
	if err != nil {
		panic(err)
	}
	return Settings{
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
}

func WriteSettings(settings Settings, path string) error {
	if path == "" {
		path = "./config.json"
	}
	jsonized, _ := json.Marshal(settings)
	err := os.WriteFile(path, jsonized, os.ModeDevice)
	if err != nil {
		panic(err)
	}
	return nil
}

func IsSettingsExists(path string) bool {
	_, e := os.OpenFile(path, os.O_RDONLY, os.ModeDevice)
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
		config.AFormat = ctx.String("vCodec")
	}
	if config.VQuality != ctx.String("vQuality") && ctx.String("vQuality") != "" {
		config.AFormat = ctx.String("vQuality")
	}
	if (config.DisableMetadata != ctx.Bool("disableMetadata") && !ctx.Bool("disableMetadata")) || (config.DisableMetadata != ctx.Bool("disableMetadata") && ctx.Bool("disableMetadata")) {
		config.DisableMetadata = ctx.Bool("disableMetadata")
	}

	return nil
}
