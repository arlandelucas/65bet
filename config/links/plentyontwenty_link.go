//go:build !prod || full || novomatic

package config

import (
	"context"

	cfg "github.com/slotopol/server/config"
	"github.com/slotopol/server/game/plentyontwenty"
	"github.com/spf13/pflag"
)

func init() {
	cfg.FlagsSetters = append(cfg.FlagsSetters, func(flags *pflag.FlagSet) {
		flags.Bool("plentyontwenty", false, "'Plenty on Twenty' Novomatic 5x3 slots")
	})
	cfg.ScatIters = append(cfg.ScatIters, func(flags *pflag.FlagSet, ctx context.Context) {
		if is, _ := flags.GetBool("plentyontwenty"); is {
			var rn, _ = flags.GetString("reels")
			plentyontwenty.CalcStat(ctx, rn)
		}
	})

	for _, alias := range []string{
		"plentyontwenty",
	} {
		cfg.GameAliases[alias] = "plentyontwenty"
	}

	cfg.GameFactory["plentyontwenty"] = func(rd string) any {
		if _, ok := plentyontwenty.ReelsMap[rd]; ok {
			return plentyontwenty.NewGame(rd)
		}
		return nil
	}
}