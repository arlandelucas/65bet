package slotopol

import (
	"github.com/slotopol/server/game/slot"
)

// Original reels.
// reels lengths [32, 32, 32, 32, 32], total reshuffles 33554432
// symbols: 48.848(lined) + 39.546(scatter) = 88.394135%
// spin9 bonuses: count 2700, rtp = 7.676482%
// monopoly bonuses: count 4320, rtp = 3.689938%
// jackpots: count 32, frequency 1/1048576
// RTP = 88.394(sym) + 7.6765(mje9) + 3.6899(mjm) = 99.760556%
var Reels100 = slot.Reels5x{
	{13, 1, 5, 12, 13, 11, 12, 11, 13, 8, 2, 12, 13, 3, 4, 6, 13, 2, 5, 10, 13, 9, 7, 8, 13, 10, 7, 9, 13, 3, 4, 6},
	{9, 5, 10, 13, 9, 6, 3, 4, 13, 2, 12, 8, 12, 13, 11, 12, 11, 13, 5, 7, 10, 6, 3, 4, 13, 2, 12, 8, 13, 7, 1, 12},
	{12, 13, 11, 12, 11, 13, 5, 10, 9, 7, 1, 12, 13, 3, 8, 6, 12, 13, 8, 4, 12, 2, 5, 10, 13, 7, 2, 13, 6, 3, 4, 9},
	{12, 1, 2, 13, 6, 5, 12, 4, 8, 12, 13, 3, 10, 9, 7, 13, 11, 11, 11, 11, 13, 5, 12, 9, 8, 6, 13, 3, 10, 2, 7, 4},
	{13, 11, 13, 12, 6, 4, 12, 3, 2, 5, 12, 10, 7, 12, 8, 1, 9, 12, 8, 9, 12, 4, 3, 12, 2, 5, 12, 10, 7, 13, 12, 6},
}

// Map with available reels.
var ReelsMap = map[float64]*slot.Reels5x{
	99.760556: &Reels100, // original reels
}