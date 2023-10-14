package cli

import (
	"encoding/json"
	"fmt"

	"github.com/timelessnesses/cobalt-go/settings"
	"github.com/urfave/cli/v2"
)

func Save(ctx *cli.Context) error {
	if ctx.NumFlags() == 0 || ctx.NumFlags() == 1 {
		jsonized, err := json.Marshal(settings.GetSettings(ctx.String("configPath")))
		if err != nil {
			return err
		}
		fmt.Println("Configuration (in JSON)")
		fmt.Println(string(jsonized))
	}
	if ctx.NumFlags() >= 2 {
		settings.Save(ctx)
	}
	return nil
}
