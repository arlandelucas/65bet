//go:build !prod || full || novomatic

package justjewels

import (
	"github.com/slotopol/server/game"
)

var Info = game.GameInfo{
	Aliases: []game.GameAlias{
		{ID: "novomatic/justjewels", Prov: "Novomatic", Name: "Just Jewels"},
		{ID: "novomatic/justjewelsdeluxe", Prov: "Novomatic", Name: "Just Jewels Deluxe"},
	},
	GP: game.GPsel |
		game.GPfgno |
		game.GPscat,
	SX:  5,
	SY:  3,
	SN:  len(LinePay),
	LN:  len(BetLines),
	BN:  0,
	RTP: game.MakeRtpList(ReelsMap),
}

func init() {
	game.GameList = append(game.GameList, &Info)
	for _, ga := range Info.Aliases {
		game.ScanFactory[ga.ID] = CalcStat
		game.GameFactory[ga.ID] = func() any {
			return NewGame()
		}
	}
}
