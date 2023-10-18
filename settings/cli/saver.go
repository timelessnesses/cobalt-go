package cli

import (
	"encoding/json"
	"fmt"

	"github.com/timelessnesses/gobalt/settings"
	"github.com/urfave/cli/v2"
)

func Save(ctx *cli.Context) error {
	if ctx.NumFlags() == 0 {
		jsonized, err := json.Marshal(settings.GetSettings(ctx.String("configPath")))
		if err != nil {
			return err
		}
		fmt.Println("Configuration (in JSON)")
		fmt.Println(string(jsonized))
	}
	if ctx.NumFlags() >= 1 {
		settings.Save(ctx)
		fmt.Println("Done!")
		settings.ValidateSettings(settings.GetSettings(ctx.String("configPath")))
	}
	return nil
}
