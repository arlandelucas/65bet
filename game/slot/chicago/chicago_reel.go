package chicago

import (
	"math"

	slot "github.com/slotopol/server/game/slot"
)

// reels lengths [35, 35, 35, 35, 35], total reshuffles 52521875
// symbols: 65.165(lined) + 14.402(scatter) = 79.567228%
// free spins 3476196, q = 0.066186, sq = 1/(1-q) = 1.070877
// free games frequency: 1/181.31
// RTP = 79.567(sym) + 0.066186*85.207(fg)*3 = 96.485616%
var ReelsReg96 = slot.Reels5x{
	{11, 10, 4, 8, 6, 7, 3, 8, 5, 10, 2, 12, 11, 4, 9, 13, 6, 11, 2, 7, 6, 9, 5, 8, 3, 12, 4, 7, 2, 10, 1, 9, 5, 12, 3},
	{5, 9, 4, 11, 1, 12, 6, 10, 5, 8, 2, 12, 7, 6, 9, 2, 8, 4, 11, 3, 13, 5, 7, 6, 9, 12, 3, 7, 4, 11, 3, 10, 8, 2, 10},
	{9, 5, 7, 9, 4, 12, 8, 7, 3, 10, 2, 11, 5, 9, 2, 11, 4, 8, 6, 10, 2, 12, 5, 7, 4, 13, 6, 10, 1, 11, 3, 12, 6, 8, 3},
	{12, 9, 2, 7, 5, 10, 4, 13, 6, 12, 4, 9, 7, 11, 4, 12, 3, 7, 2, 11, 1, 8, 3, 10, 5, 8, 6, 9, 5, 10, 6, 8, 3, 11, 2},
	{11, 6, 12, 4, 10, 2, 13, 5, 8, 2, 12, 5, 11, 4, 7, 12, 3, 8, 6, 9, 5, 7, 3, 10, 9, 4, 8, 7, 3, 10, 2, 11, 6, 9, 1},
}

// Map with available reels.
var ReelsMap = map[float64]*slot.Reels5x{
	96.485616: &ReelsReg96,
}

func FindReels(mrtp float64) (rtp float64, reels slot.Reels) {
	for p, r := range ReelsMap {
		if math.Abs(mrtp-p) < math.Abs(mrtp-rtp) {
			rtp, reels = p, r
		}
	}
	return
}