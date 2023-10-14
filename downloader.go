package main

import (
	"fmt"
	"strings"

	"github.com/timelessnesses/gobalt/client"
	"github.com/timelessnesses/gobalt/settings"
	"github.com/urfave/cli/v2"
)

func convert_settings_to_actual_client_setting(c settings.Settings) client.Setting {
	fmt.Println(c.AFormat)
	vCodec, err := client.NewVideoCodec(c.VCodec)
	vQuality, err2 := client.NewVideoQuality(c.VQuality)
	aFormat, err3 := client.NewAudioFormat(strings.ToLower(c.AFormat))
	if err != nil || err2 != nil || err3 != nil {
		if err != nil {
			panic(err)
		} else if err2 != nil {
			panic(err2)
		} else if err3 != nil {
			panic(err3)
		}

	}
	return client.Setting{
		VCodec:          vCodec,
		VQuality:        vQuality,
		AFormat:         aFormat,
		IsAudioOnly:     c.IsAudioOnly,
		IsNoTTWatermark: c.IsNoTTWatermark,
		IsTTFullAudio:   c.IsTTFullAudio,
		IsAudioMuted:    c.IsAudioMuted,
		DubLang:         c.DubLang,
		DisableMetadata: c.DisableMetadata,
	}
}

func DownloaderDo(ctx *cli.Context) error {
	config := settings.GetSettings(ctx.String("configPath"))
	c := client.NewClient(config.Endpoint)

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

	url := ctx.String("url")
	s := convert_settings_to_actual_client_setting(config)
	s.Url = url
	result, err := c.GetInfo(s)
	if err != nil {
		panic(err)
	}
	fmt.Println("Downloading...")
	c.Download(result, ctx.Path("out"))
	return nil
}
