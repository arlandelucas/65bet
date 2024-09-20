package tikiwonders

import (
	"math"

	slot "github.com/slotopol/server/game/slot"
)

// reels lengths [31, 31, 31, 31, 31], total reshuffles 28629151
// symbols: 43.995(lined) + 19.079(scatter) = 63.074399%
// free spins 2350890, q = 0.082115, sq = 1/(1-q) = 1.089461
// free games frequency: 1/128.23
// RTP = 63.074(sym) + 0.082115*206.15(fg) = 80.002571%
var Reels80 = slot.Reels5x{
	{6, 10, 11, 2, 3, 6, 9, 5, 10, 8, 7, 9, 4, 1, 3, 6, 4, 11, 8, 7, 10, 9, 5, 7, 3, 11, 8, 5, 10, 4, 11},
	{11, 6, 7, 11, 4, 8, 3, 6, 10, 5, 11, 3, 5, 2, 10, 4, 9, 8, 11, 9, 7, 10, 8, 9, 7, 10, 1, 4, 3, 6, 5},
	{4, 8, 11, 7, 8, 1, 5, 4, 9, 8, 4, 6, 11, 10, 3, 7, 10, 11, 5, 10, 6, 2, 7, 3, 11, 6, 9, 5, 3, 10, 9},
	{11, 2, 10, 6, 3, 8, 7, 10, 4, 8, 10, 9, 7, 6, 11, 5, 3, 9, 5, 1, 10, 4, 11, 5, 8, 9, 11, 7, 4, 3, 6},
	{9, 5, 3, 8, 11, 6, 2, 10, 9, 7, 5, 6, 3, 4, 10, 8, 7, 9, 3, 6, 7, 8, 11, 10, 4, 5, 10, 11, 1, 4, 11},
}

// reels lengths [30, 31, 31, 31, 31], total reshuffles 27705630
// symbols: 44.971(lined) + 19.352(scatter) = 64.322385%
// free spins 2319030, q = 0.083702, sq = 1/(1-q) = 1.091349
// free games frequency: 1/125.84
// RTP = 64.322(sym) + 0.083702*210.59(fg) = 81.949662%
var Reels82 = slot.Reels5x{
	{11, 3, 7, 6, 1, 8, 7, 3, 11, 4, 9, 11, 5, 6, 10, 4, 8, 10, 9, 7, 10, 9, 8, 5, 3, 11, 2, 6, 4, 5},
	{11, 6, 7, 11, 4, 8, 3, 6, 10, 5, 11, 3, 5, 2, 10, 4, 9, 8, 11, 9, 7, 10, 8, 9, 7, 10, 1, 4, 3, 6, 5},
	{4, 8, 11, 7, 8, 1, 5, 4, 9, 8, 4, 6, 11, 10, 3, 7, 10, 11, 5, 10, 6, 2, 7, 3, 11, 6, 9, 5, 3, 10, 9},
	{11, 2, 10, 6, 3, 8, 7, 10, 4, 8, 10, 9, 7, 6, 11, 5, 3, 9, 5, 1, 10, 4, 11, 5, 8, 9, 11, 7, 4, 3, 6},
	{9, 5, 3, 8, 11, 6, 2, 10, 9, 7, 5, 6, 3, 4, 10, 8, 7, 9, 3, 6, 7, 8, 11, 10, 4, 5, 10, 11, 1, 4, 11},
}

// reels lengths [30, 30, 31, 31, 31], total reshuffles 26811900
// symbols: 46.064(lined) + 19.627(scatter) = 65.690921%
// free spins 2287440, q = 0.085314, sq = 1/(1-q) = 1.093272
// free games frequency: 1/123.51
// RTP = 65.691(sym) + 0.085314*215.45(fg) = 84.072248%
var Reels84 = slot.Reels5x{
	{11, 3, 7, 6, 1, 8, 7, 3, 11, 4, 9, 11, 5, 6, 10, 4, 8, 10, 9, 7, 10, 9, 8, 5, 3, 11, 2, 6, 4, 5},
	{6, 10, 3, 7, 11, 6, 7, 10, 11, 5, 2, 4, 10, 11, 9, 3, 1, 6, 8, 4, 9, 8, 5, 11, 9, 4, 5, 8, 3, 7},
	{4, 8, 11, 7, 8, 1, 5, 4, 9, 8, 4, 6, 11, 10, 3, 7, 10, 11, 5, 10, 6, 2, 7, 3, 11, 6, 9, 5, 3, 10, 9},
	{11, 2, 10, 6, 3, 8, 7, 10, 4, 8, 10, 9, 7, 6, 11, 5, 3, 9, 5, 1, 10, 4, 11, 5, 8, 9, 11, 7, 4, 3, 6},
	{9, 5, 3, 8, 11, 6, 2, 10, 9, 7, 5, 6, 3, 4, 10, 8, 7, 9, 3, 6, 7, 8, 11, 10, 4, 5, 10, 11, 1, 4, 11},
}

// reels lengths [31, 30, 30, 30, 30], total reshuffles 25110000
// symbols: 46.587(lined) + 20.186(scatter) = 66.772983%
// free spins 2225070, q = 0.088613, sq = 1/(1-q) = 1.097229
// free games frequency: 1/119
// RTP = 66.773(sym) + 0.088613*219.8(fg) = 86.249716%
var Reels86 = slot.Reels5x{
	{6, 10, 11, 2, 3, 6, 9, 5, 10, 8, 7, 9, 4, 1, 3, 6, 4, 11, 8, 7, 10, 9, 5, 7, 3, 11, 8, 5, 10, 4, 11},
	{6, 10, 3, 7, 11, 6, 7, 10, 11, 5, 2, 4, 10, 11, 9, 3, 1, 6, 8, 4, 9, 8, 5, 11, 9, 4, 5, 8, 3, 7},
	{6, 5, 8, 4, 7, 1, 10, 11, 5, 10, 9, 11, 7, 6, 3, 4, 8, 11, 9, 3, 4, 6, 7, 10, 8, 11, 5, 9, 3, 2},
	{8, 6, 3, 9, 11, 7, 6, 9, 11, 6, 4, 3, 9, 5, 10, 11, 7, 3, 1, 4, 5, 11, 8, 10, 2, 8, 10, 4, 5, 7},
	{10, 8, 4, 11, 2, 4, 3, 11, 7, 4, 10, 5, 9, 7, 6, 9, 5, 8, 10, 5, 1, 8, 6, 7, 3, 11, 9, 6, 3, 11},
}

// reels lengths [30, 30, 30, 30, 31], total reshuffles 25110000
// symbols: 47.663(lined) + 20.186(scatter) = 67.848331%
// free spins 2225070, q = 0.088613, sq = 1/(1-q) = 1.097229
// free games frequency: 1/119
// RTP = 67.848(sym) + 0.088613*223.34(fg) = 87.638729%
var Reels88 = slot.Reels5x{
	{11, 3, 7, 6, 1, 8, 7, 3, 11, 4, 9, 11, 5, 6, 10, 4, 8, 10, 9, 7, 10, 9, 8, 5, 3, 11, 2, 6, 4, 5},
	{6, 10, 3, 7, 11, 6, 7, 10, 11, 5, 2, 4, 10, 11, 9, 3, 1, 6, 8, 4, 9, 8, 5, 11, 9, 4, 5, 8, 3, 7},
	{6, 5, 8, 4, 7, 1, 10, 11, 5, 10, 9, 11, 7, 6, 3, 4, 8, 11, 9, 3, 4, 6, 7, 10, 8, 11, 5, 9, 3, 2},
	{8, 6, 3, 9, 11, 7, 6, 9, 11, 6, 4, 3, 9, 5, 10, 11, 7, 3, 1, 4, 5, 11, 8, 10, 2, 8, 10, 4, 5, 7},
	{9, 5, 3, 8, 11, 6, 2, 10, 9, 7, 5, 6, 3, 4, 10, 8, 7, 9, 3, 6, 7, 8, 11, 10, 4, 5, 10, 11, 1, 4, 11},
}

// reels lengths [30, 30, 30, 30, 30], total reshuffles 24300000
// symbols: 47.856(lined) + 20.47(scatter) = 68.325638%
// free spins 2194290, q = 0.0903, sq = 1/(1-q) = 1.099263
// free games frequency: 1/116.82
// RTP = 68.326(sym) + 0.0903*225.32(fg) = 88.672362%
var Reels89 = slot.Reels5x{
	{11, 3, 7, 6, 1, 8, 7, 3, 11, 4, 9, 11, 5, 6, 10, 4, 8, 10, 9, 7, 10, 9, 8, 5, 3, 11, 2, 6, 4, 5},
	{6, 10, 3, 7, 11, 6, 7, 10, 11, 5, 2, 4, 10, 11, 9, 3, 1, 6, 8, 4, 9, 8, 5, 11, 9, 4, 5, 8, 3, 7},
	{6, 5, 8, 4, 7, 1, 10, 11, 5, 10, 9, 11, 7, 6, 3, 4, 8, 11, 9, 3, 4, 6, 7, 10, 8, 11, 5, 9, 3, 2},
	{8, 6, 3, 9, 11, 7, 6, 9, 11, 6, 4, 3, 9, 5, 10, 11, 7, 3, 1, 4, 5, 11, 8, 10, 2, 8, 10, 4, 5, 7},
	{10, 8, 4, 11, 2, 4, 3, 11, 7, 4, 10, 5, 9, 7, 6, 9, 5, 8, 10, 5, 1, 8, 6, 7, 3, 11, 9, 6, 3, 11},
}

// reels lengths [29, 30, 30, 30, 30], total reshuffles 23490000
// symbols: 48.947(lined) + 20.774(scatter) = 69.720877%
// free spins 2163510, q = 0.092103, sq = 1/(1-q) = 1.101447
// free games frequency: 1/114.58
// RTP = 69.721(sym) + 0.092103*230.38(fg) = 90.939814%
var Reels91 = slot.Reels5x{
	{9, 11, 5, 10, 9, 1, 6, 3, 5, 4, 8, 7, 11, 2, 3, 8, 5, 10, 11, 7, 6, 10, 4, 7, 9, 6, 8, 4, 3},
	{6, 10, 3, 7, 11, 6, 7, 10, 11, 5, 2, 4, 10, 11, 9, 3, 1, 6, 8, 4, 9, 8, 5, 11, 9, 4, 5, 8, 3, 7},
	{6, 5, 8, 4, 7, 1, 10, 11, 5, 10, 9, 11, 7, 6, 3, 4, 8, 11, 9, 3, 4, 6, 7, 10, 8, 11, 5, 9, 3, 2},
	{8, 6, 3, 9, 11, 7, 6, 9, 11, 6, 4, 3, 9, 5, 10, 11, 7, 3, 1, 4, 5, 11, 8, 10, 2, 8, 10, 4, 5, 7},
	{10, 8, 4, 11, 2, 4, 3, 11, 7, 4, 10, 5, 9, 7, 6, 9, 5, 8, 10, 5, 1, 8, 6, 7, 3, 11, 9, 6, 3, 11},
}

// reels lengths [29, 30, 30, 30, 29], total reshuffles 22707000
// symbols: 49.091(lined) + 21.081(scatter) = 70.172013%
// free spins 2133000, q = 0.093936, sq = 1/(1-q) = 1.103675
// free games frequency: 1/112.39
// RTP = 70.172(sym) + 0.093936*232.34(fg) = 91.997167%
var Reels92 = slot.Reels5x{
	{9, 11, 5, 10, 9, 1, 6, 3, 5, 4, 8, 7, 11, 2, 3, 8, 5, 10, 11, 7, 6, 10, 4, 7, 9, 6, 8, 4, 3},
	{6, 10, 3, 7, 11, 6, 7, 10, 11, 5, 2, 4, 10, 11, 9, 3, 1, 6, 8, 4, 9, 8, 5, 11, 9, 4, 5, 8, 3, 7},
	{6, 5, 8, 4, 7, 1, 10, 11, 5, 10, 9, 11, 7, 6, 3, 4, 8, 11, 9, 3, 4, 6, 7, 10, 8, 11, 5, 9, 3, 2},
	{8, 6, 3, 9, 11, 7, 6, 9, 11, 6, 4, 3, 9, 5, 10, 11, 7, 3, 1, 4, 5, 11, 8, 10, 2, 8, 10, 4, 5, 7},
	{4, 10, 6, 4, 5, 9, 8, 10, 1, 5, 10, 7, 11, 8, 9, 6, 11, 3, 8, 2, 7, 3, 9, 7, 4, 6, 3, 5, 11},
}

// reels lengths [30, 30, 29, 29, 29], total reshuffles 21950100
// symbols: 49.441(lined) + 21.392(scatter) = 70.833422%
// free spins 2102760, q = 0.095797, sq = 1/(1-q) = 1.105947
// free games frequency: 1/110.25
// RTP = 70.833(sym) + 0.095797*235.01(fg) = 93.347122%
var Reels93 = slot.Reels5x{
	{11, 3, 7, 6, 1, 8, 7, 3, 11, 4, 9, 11, 5, 6, 10, 4, 8, 10, 9, 7, 10, 9, 8, 5, 3, 11, 2, 6, 4, 5},
	{6, 10, 3, 7, 11, 6, 7, 10, 11, 5, 2, 4, 10, 11, 9, 3, 1, 6, 8, 4, 9, 8, 5, 11, 9, 4, 5, 8, 3, 7},
	{2, 5, 10, 11, 4, 9, 7, 8, 6, 5, 11, 4, 3, 7, 9, 4, 10, 8, 6, 11, 3, 5, 1, 3, 6, 9, 7, 10, 8},
	{3, 11, 7, 2, 11, 4, 9, 5, 10, 7, 9, 8, 7, 6, 8, 3, 6, 4, 8, 1, 11, 10, 3, 5, 9, 6, 5, 10, 4},
	{4, 10, 6, 4, 5, 9, 8, 10, 1, 5, 10, 7, 11, 8, 9, 6, 11, 3, 8, 2, 7, 3, 9, 7, 4, 6, 3, 5, 11},
}

// reels lengths [29, 30, 29, 30, 29], total reshuffles 21950100
// symbols: 50.334(lined) + 21.392(scatter) = 71.726106%
// free spins 2102760, q = 0.095797, sq = 1/(1-q) = 1.105947
// free games frequency: 1/110.25
// RTP = 71.726(sym) + 0.095797*237.98(fg) = 94.523537%
var Reels95 = slot.Reels5x{
	{9, 11, 5, 10, 9, 1, 6, 3, 5, 4, 8, 7, 11, 2, 3, 8, 5, 10, 11, 7, 6, 10, 4, 7, 9, 6, 8, 4, 3},
	{6, 10, 3, 7, 11, 6, 7, 10, 11, 5, 2, 4, 10, 11, 9, 3, 1, 6, 8, 4, 9, 8, 5, 11, 9, 4, 5, 8, 3, 7},
	{2, 5, 10, 11, 4, 9, 7, 8, 6, 5, 11, 4, 3, 7, 9, 4, 10, 8, 6, 11, 3, 5, 1, 3, 6, 9, 7, 10, 8},
	{8, 6, 3, 9, 11, 7, 6, 9, 11, 6, 4, 3, 9, 5, 10, 11, 7, 3, 1, 4, 5, 11, 8, 10, 2, 8, 10, 4, 5, 7},
	{4, 10, 6, 4, 5, 9, 8, 10, 1, 5, 10, 7, 11, 8, 9, 6, 11, 3, 8, 2, 7, 3, 9, 7, 4, 6, 3, 5, 11},
}

// reels lengths [29, 29, 29, 30, 30], total reshuffles 21950100
// symbols: 51.51(lined) + 21.392(scatter) = 72.901590%
// free spins 2102760, q = 0.095797, sq = 1/(1-q) = 1.105947
// free games frequency: 1/110.25
// RTP = 72.902(sym) + 0.095797*241.88(fg) = 96.072637%
var Reels96 = slot.Reels5x{
	{9, 11, 5, 10, 9, 1, 6, 3, 5, 4, 8, 7, 11, 2, 3, 8, 5, 10, 11, 7, 6, 10, 4, 7, 9, 6, 8, 4, 3},
	{5, 3, 9, 7, 10, 4, 11, 8, 10, 5, 3, 11, 1, 6, 8, 4, 6, 10, 11, 7, 4, 6, 9, 8, 3, 7, 5, 2, 9},
	{2, 5, 10, 11, 4, 9, 7, 8, 6, 5, 11, 4, 3, 7, 9, 4, 10, 8, 6, 11, 3, 5, 1, 3, 6, 9, 7, 10, 8},
	{8, 6, 3, 9, 11, 7, 6, 9, 11, 6, 4, 3, 9, 5, 10, 11, 7, 3, 1, 4, 5, 11, 8, 10, 2, 8, 10, 4, 5, 7},
	{10, 8, 4, 11, 2, 4, 3, 11, 7, 4, 10, 5, 9, 7, 6, 9, 5, 8, 10, 5, 1, 8, 6, 7, 3, 11, 9, 6, 3, 11},
}

// reels lengths [29, 29, 29, 29, 30], total reshuffles 21218430
// symbols: 51.963(lined) + 21.706(scatter) = 73.669541%
// free spins 2072790, q = 0.097688, sq = 1/(1-q) = 1.108264
// free games frequency: 1/108.16
// RTP = 73.67(sym) + 0.097688*244.94(fg) = 97.596893%
var Reels98 = slot.Reels5x{
	{9, 11, 5, 10, 9, 1, 6, 3, 5, 4, 8, 7, 11, 2, 3, 8, 5, 10, 11, 7, 6, 10, 4, 7, 9, 6, 8, 4, 3},
	{5, 3, 9, 7, 10, 4, 11, 8, 10, 5, 3, 11, 1, 6, 8, 4, 6, 10, 11, 7, 4, 6, 9, 8, 3, 7, 5, 2, 9},
	{2, 5, 10, 11, 4, 9, 7, 8, 6, 5, 11, 4, 3, 7, 9, 4, 10, 8, 6, 11, 3, 5, 1, 3, 6, 9, 7, 10, 8},
	{3, 11, 7, 2, 11, 4, 9, 5, 10, 7, 9, 8, 7, 6, 8, 3, 6, 4, 8, 1, 11, 10, 3, 5, 9, 6, 5, 10, 4},
	{10, 8, 4, 11, 2, 4, 3, 11, 7, 4, 10, 5, 9, 7, 6, 9, 5, 8, 10, 5, 1, 8, 6, 7, 3, 11, 9, 6, 3, 11},
}

// reels lengths [29, 29, 29, 29, 29], total reshuffles 20511149
// symbols: 52.182(lined) + 22.024(scatter) = 74.206057%
// free spins 2043090, q = 0.099609, sq = 1/(1-q) = 1.110628
// free games frequency: 1/106.11
// RTP = 74.206(sym) + 0.099609*247.25(fg) = 98.833928%
var Reels99 = slot.Reels5x{
	{9, 11, 5, 10, 9, 1, 6, 3, 5, 4, 8, 7, 11, 2, 3, 8, 5, 10, 11, 7, 6, 10, 4, 7, 9, 6, 8, 4, 3},
	{5, 3, 9, 7, 10, 4, 11, 8, 10, 5, 3, 11, 1, 6, 8, 4, 6, 10, 11, 7, 4, 6, 9, 8, 3, 7, 5, 2, 9},
	{2, 5, 10, 11, 4, 9, 7, 8, 6, 5, 11, 4, 3, 7, 9, 4, 10, 8, 6, 11, 3, 5, 1, 3, 6, 9, 7, 10, 8},
	{3, 11, 7, 2, 11, 4, 9, 5, 10, 7, 9, 8, 7, 6, 8, 3, 6, 4, 8, 1, 11, 10, 3, 5, 9, 6, 5, 10, 4},
	{4, 10, 6, 4, 5, 9, 8, 10, 1, 5, 10, 7, 11, 8, 9, 6, 11, 3, 8, 2, 7, 3, 9, 7, 4, 6, 3, 5, 11},
}

// reels lengths [26, 26, 26, 26, 26], total reshuffles 11881376
// symbols: 43.104(lined) + 27.979(scatter) = 71.083181%
// free spins 1621890, q = 0.13651, sq = 1/(1-q) = 1.158087
// free games frequency: 1/77.968
// RTP = 71.083(sym) + 0.13651*246.96(fg) = 104.795133%
var Reels105 = slot.Reels5x{
	{6, 9, 5, 7, 2, 10, 5, 4, 8, 9, 11, 7, 6, 8, 9, 10, 3, 1, 8, 10, 11, 6, 4, 3, 11, 7},
	{10, 8, 9, 3, 1, 11, 4, 7, 3, 5, 9, 2, 8, 6, 4, 7, 11, 10, 8, 6, 7, 9, 10, 11, 6, 5},
	{11, 9, 5, 10, 3, 8, 10, 7, 6, 11, 9, 1, 4, 6, 11, 5, 9, 7, 8, 6, 2, 10, 4, 7, 8, 3},
	{8, 1, 7, 4, 9, 5, 6, 9, 5, 10, 8, 6, 3, 11, 2, 3, 4, 7, 8, 6, 10, 11, 9, 7, 11, 10},
	{9, 4, 8, 5, 7, 8, 6, 10, 11, 5, 10, 11, 3, 9, 10, 6, 8, 9, 7, 6, 1, 3, 4, 7, 2, 11},
}

// reels lengths [25, 25, 25, 25, 25], total reshuffles 9765625
// symbols: 43.278(lined) + 30.536(scatter) = 73.814630%
// free spins 1492290, q = 0.15281, sq = 1/(1-q) = 1.180373
// free games frequency: 1/69.838
// RTP = 73.815(sym) + 0.15281*261.39(fg) = 113.757230%
var Reels114 = slot.Reels5x{
	{2, 9, 3, 7, 8, 11, 9, 10, 11, 7, 10, 4, 1, 5, 7, 4, 5, 3, 6, 8, 9, 11, 6, 10, 8},
	{8, 6, 3, 11, 9, 10, 1, 5, 8, 11, 4, 2, 9, 5, 4, 7, 10, 6, 8, 7, 9, 3, 10, 11, 7},
	{8, 11, 5, 9, 3, 1, 10, 5, 9, 4, 7, 3, 11, 4, 8, 6, 7, 10, 6, 7, 2, 8, 11, 10, 9},
	{8, 10, 11, 3, 8, 7, 9, 4, 6, 2, 9, 10, 6, 11, 7, 5, 1, 9, 8, 10, 3, 11, 5, 7, 4},
	{3, 6, 5, 9, 11, 1, 7, 3, 11, 10, 4, 8, 9, 7, 4, 10, 8, 7, 5, 6, 2, 10, 8, 11, 9},
}

// Map with available reels.
var ReelsMap = map[float64]*slot.Reels5x{
	80.002571:  &Reels80,
	81.949662:  &Reels82,
	84.072248:  &Reels84,
	86.249716:  &Reels86,
	87.638729:  &Reels88,
	88.672362:  &Reels89,
	90.939814:  &Reels91,
	91.997167:  &Reels92,
	93.415147:  &Reels93,
	94.523537:  &Reels95,
	96.072637:  &Reels96,
	98.833928:  &Reels99,
	104.795133: &Reels105,
	113.757230: &Reels114,
}

func FindReels(mrtp float64) (rtp float64, reels slot.Reels) {
	for p, r := range ReelsMap {
		if math.Abs(mrtp-p) < math.Abs(mrtp-rtp) {
			rtp, reels = p, r
		}
	}
	return
}