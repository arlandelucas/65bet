package hotclover

// See: https://demo.agtsoftware.com/games/agt/hotclover100

import (
	"github.com/slotopol/server/game/slot"
)

// Lined payment.
var LinePay = [10][5]float64{
	{},                     //  1 wild
	{},                     //  2 scatter
	{0, 10, 50, 200, 3000}, //  3 seven
	{0, 0, 40, 100, 600},   //  4 strawberry
	{0, 0, 40, 100, 400},   //  5 grapes
	{0, 0, 20, 50, 200},    //  6 bar
	{0, 0, 12, 30, 100},    //  7 plum
	{0, 0, 12, 30, 100},    //  8 orange
	{0, 0, 8, 28, 80},      //  9 lemon
	{0, 0, 8, 28, 80},      // 10 cherry
}

// Scatters payment.
var ScatPay = [5]float64{0, 0, 3, 20, 100} // 2 scatter

// Bet lines
var BetLines = slot.BetLinesAgt5x4[:100]

type Game struct {
	slot.Slot5x4 `yaml:",inline"`
}

// Declare conformity with SlotGame interface.
var _ slot.SlotGame = (*Game)(nil)

func NewGame() *Game {
	return &Game{
		Slot5x4: slot.Slot5x4{
			Sel: len(BetLines),
			Bet: 1,
		},
	}
}

const wild, scat = 1, 2

func (g *Game) Scanner(screen slot.Screen, wins *slot.Wins) {
	var scrn5x4 = screen.(*slot.Screen5x4)
	g.ScanLined(scrn5x4, wins)
	g.ScanScatters(scrn5x4, wins)
}

// Lined symbols calculation.
func (g *Game) ScanLined(screen *slot.Screen5x4, wins *slot.Wins) {
	var reelwild [5]bool
	var x, y slot.Pos
	for x = 2; x <= 4; x++ {
		for y = 1; y <= 4; y++ {
			if screen.At(x, y) == wild {
				reelwild[x-1] = true
				break
			}
		}
	}

	for li := 1; li <= g.Sel; li++ {
		var line = BetLines[li-1]

		var numl slot.Pos = 5
		var syml = screen.Pos(1, line)
		var x slot.Pos
		for x = 2; x <= 5; x++ {
			var sx = screen.Pos(x, line)
			if reelwild[x-1] {
				continue
			} else if sx != syml {
				numl = x - 1
				break
			}
		}

		if pay := LinePay[syml-1][numl-1]; pay > 0 {
			*wins = append(*wins, slot.WinItem{
				Pay:  g.Bet * pay,
				Mult: 1,
				Sym:  syml,
				Num:  numl,
				Line: li,
				XY:   line.CopyL(numl),
			})
		}
	}
}

// Scatters calculation.
func (g *Game) ScanScatters(screen *slot.Screen5x4, wins *slot.Wins) {
	if count := screen.ScatNum(scat); count >= 3 {
		var pay = ScatPay[count-1]
		*wins = append(*wins, slot.WinItem{
			Pay:  g.Bet * float64(g.Sel) * pay,
			Mult: 1,
			Sym:  scat,
			Num:  count,
			XY:   screen.ScatPos(scat),
		})
	}
}

func (g *Game) Spin(screen slot.Screen, mrtp float64) {
	var reels, _ = slot.FindReels(ReelsMap, mrtp)
	screen.Spin(reels)
}

func (g *Game) SetSel(sel int) error {
	return slot.ErrNoFeature
}