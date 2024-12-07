package egypt

// See: https://demo.agtsoftware.com/games/agt/egypt

import (
	"math/rand/v2"

	"github.com/slotopol/server/game/slot"
)

// Lined payment.
var LinePay = [12][5]float64{
	{0, 10, 250, 2500, 9000}, //  1 wild
	{},                       //  2 scatter
	{0, 2, 25, 125, 750},     //  3 anubis
	{0, 2, 25, 125, 750},     //  4 sphinx
	{0, 0, 20, 100, 400},     //  5 snake
	{0, 0, 15, 75, 250},      //  6 mummy
	{0, 0, 15, 75, 250},      //  7 wall
	{0, 0, 10, 50, 125},      //  8 cat
	{0, 0, 10, 50, 125},      //  9 ace
	{0, 0, 5, 25, 100},       // 10 king
	{0, 0, 5, 10, 50},        // 11 queen
	{0, 0, 5, 10, 25},        // 12 jack
}

// Scatters payment.
var ScatPay = [5]float64{0, 2, 5, 20, 500} // 2 scatter

// Bet lines
var BetLines = slot.BetLinesAgt5x3[:15]

type Game struct {
	slot.Slot5x3 `yaml:",inline"`
	Mini         [3]slot.Sym `json:"mini" yaml:"mini" xml:"mini"`
}

// Declare conformity with SlotGame interface.
var _ slot.SlotGame = (*Game)(nil)

func NewGame() *Game {
	return &Game{
		Slot5x3: slot.Slot5x3{
			Sel: len(BetLines),
			Bet: 1,
		},
		Mini: [3]slot.Sym{1, 2, 3},
	}
}

const wild, scat = 1, 2

func (g *Game) Scanner(screen slot.Screen, wins *slot.Wins) {
	g.ScanLined(screen, wins)
	g.ScanScatters(screen, wins)
}

// Lined symbols calculation.
func (g *Game) ScanLined(screen slot.Screen, wins *slot.Wins) {
	for li := 1; li <= g.Sel; li++ {
		var line = BetLines[li-1]

		var numw, numl slot.Pos = 0, 5
		var syml slot.Sym
		var x slot.Pos
		for x = 1; x <= 5; x++ {
			var sx = screen.Pos(x, line)
			if sx == wild {
				if syml == 0 {
					numw = x
				}
			} else if syml == 0 && sx != scat {
				syml = sx
			} else if sx != syml {
				numl = x - 1
				break
			}
		}

		var payw, payl float64
		if numw > 0 {
			payw = LinePay[wild-1][numw-1]
		}
		if numl > 0 && syml > 0 {
			payl = LinePay[syml-1][numl-1]
		}
		if payl > payw {
			var m = g.mult()
			*wins = append(*wins, slot.WinItem{
				Pay:  g.Bet * payl,
				Mult: m,
				Sym:  syml,
				Num:  numl,
				Line: li,
				XY:   line.CopyL(numl),
			})
		} else if payw > 0 {
			var m = g.mult()
			*wins = append(*wins, slot.WinItem{
				Pay:  g.Bet * payw,
				Mult: m,
				Sym:  wild,
				Num:  numw,
				Line: li,
				XY:   line.CopyL(numw),
			})
		}
	}
}

// Scatters calculation.
func (g *Game) ScanScatters(screen slot.Screen, wins *slot.Wins) {
	if count := screen.ScatNum(scat); count >= 2 {
		var pay, m = ScatPay[count-1], g.mult()
		*wins = append(*wins, slot.WinItem{
			Pay:  g.Bet * float64(g.Sel) * pay,
			Mult: m,
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

func (g *Game) Prepare() {
	g.Mini[0] = rand.N[slot.Sym](3) + 1
	g.Mini[1] = rand.N[slot.Sym](3) + 1
	g.Mini[2] = rand.N[slot.Sym](3) + 1
}

func (g *Game) mult() float64 {
	if g.Mini[0] != g.Mini[1] || g.Mini[1] != g.Mini[2] {
		return 1
	}
	switch g.Mini[0] {
	case 1:
		return 3
	case 2:
		return 6
	case 3:
		return 9
	}
	panic("no way here")
}

func (g *Game) SetSel(sel int) error {
	return g.SetSelNum(sel, len(BetLines))
}