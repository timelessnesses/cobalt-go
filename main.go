package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/timelessnesses/cobalt-go/client"
	"github.com/timelessnesses/cobalt-go/settings"
	"github.com/urfave/cli/v2"
)

func main() {

	download := cli.Command{
		Name:    "download",
		Aliases: []string{"d"},
		Usage:   "Download content from the URL and custom settings provided.",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "url",
				Usage:    "URL",
				Required: true,
			},
			&cli.StringFlag{
				Name:  "vCodec",
				Usage: "Video Codecs (h264, vp9, av1)",
				Action: func(ctx *cli.Context, s string) error {
					s = strings.ToLower(s)
					_, err := client.NewVideoCodec(s)
					if err != nil {
						return err
					}
					return nil
				},
			},
			&cli.StringFlag{
				Name:  "vQuality",
				Usage: "Video Qualities (Specify max for maximum possible quality)",
				Action: func(ctx *cli.Context, s string) error {
					_, e := client.NewVideoQuality(strings.ToLower(s))
					if e != nil {
						return e
					}
					return nil
				},
				Value: "max",
			},
			&cli.StringFlag{
				Name:  "aFormat",
				Usage: "Audio Formats (MP3, Opus, Wav, Ogg, Best)",
				Action: func(ctx *cli.Context, s string) error {
					_, e := client.NewAudioFormat(strings.ToLower(s))
					if e != nil {
						return e
					}
					return nil
				},
				Value: "Best",
			},
			&cli.BoolFlag{
				Name:  "isAudioOnly",
				Usage: "Specify forcing audio only or not. (Default: False)",
				Value: false,
			},
			&cli.BoolFlag{
				Name:  "isNoTTWatermark",
				Usage: "Specify forcing no TikTok watermark. (Default: False)",
				Value: false,
			},
			&cli.BoolFlag{
				Name:  "isTTFullAudio",
				Usage: "Specify forcing full audio from TikTok or not. (Default: False)",
				Value: false,
			},
			&cli.BoolFlag{
				Name:  "isAudioMuted",
				Usage: "Specify forcing muting audio or not. (Default: False)",
				Value: false,
			},
			&cli.BoolFlag{
				Name:  "dubLang",
				Usage: "Specify forcing dubbing languages or not. (Default: False)",
				Value: false,
			},
			&cli.BoolFlag{
				Name:  "disableMetadata",
				Usage: "Specify forcing no metadatas or not. (Default: False)",
				Value: false,
			},
		},
	}

	setting := cli.Command{
		Name:  "setting",
		Usage: "Edit or see current configuration.",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "endpoint",
				Usage: "An endpoint (Included with schemas )",
			},
			&cli.StringFlag{
				Name:  "vCodec",
				Usage: "Video Codecs (h264, vp9, av1)",
				Action: func(ctx *cli.Context, s string) error {
					s = strings.ToLower(s)
					_, err := client.NewVideoCodec(s)
					if err != nil {
						return err
					}
					return nil
				},
			},
			&cli.StringFlag{
				Name:  "vQuality",
				Usage: "Video Qualities (Specify max for maximum possible quality)",
				Action: func(ctx *cli.Context, s string) error {
					_, e := client.NewVideoQuality(strings.ToLower(s))
					if e != nil {
						return e
					}
					return nil
				},
			},
			&cli.StringFlag{
				Name:  "aFormat",
				Usage: "Audio Formats (MP3, Opus, Wav, Ogg, Best)",
				Action: func(ctx *cli.Context, s string) error {
					_, e := client.NewAudioFormat(strings.ToLower(s))
					if e != nil {
						return e
					}
					return nil
				},
			},
			&cli.BoolFlag{
				Name:  "isAudioOnly",
				Usage: "Specify forcing audio only or not.",
			},
			&cli.BoolFlag{
				Name:  "isNoTTWatermark",
				Usage: "Specify forcing no TikTok watermark.",
			},
			&cli.BoolFlag{
				Name:  "isTTFullAudio",
				Usage: "Specify forcing full audio from TikTok or not.",
			},
			&cli.BoolFlag{
				Name:  "isAudioMuted",
				Usage: "Specify forcing muting audio or not.",
			},
			&cli.BoolFlag{
				Name:  "dubLang",
				Usage: "Specify forcing dubbing languages or not.",
			},
			&cli.BoolFlag{
				Name:  "disableMetadata",
				Usage: "Specify forcing no metadatas or not.",
			},
			&cli.StringFlag{
				Name:  "configPath",
				Usage: "Specify config.json path for this program to use.",
				Value: "./config.json",
			},
		},
		Action: func(ctx *cli.Context) error {
			if ctx.NumFlags() == 0 || ctx.NumFlags() == 1 {
				jsonized, err := json.Marshal(settings.GetSettings(ctx.String("configPath")))
				if err != nil {
					return err
				}
				fmt.Println("Configuration (in JSON)")
				fmt.Println(jsonized)
			}
			if ctx.NumFlags() >= 2 {
				settings.Save(ctx)
			}
			return nil
		},
	}

	app := cli.App{
		Name:  "cobalt-go",
		Usage: "A CLI tool for interacting with cobalt's API. Written in Golang.",
		Commands: []*cli.Command{
			&download,
			&setting,
		},
	}
	if err := app.Run(os.Args); err != nil {
		panic(err)
	}
}
