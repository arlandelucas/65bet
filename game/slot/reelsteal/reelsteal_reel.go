package reelsteal

import (
	slot "github.com/slotopol/server/game/slot"
)

// *bonus reels calculations*
// reels lengths [44, 44, 44, 44, 44], total reshuffles 164916224
// symbols: 294.63(lined) + 0(scatter) = 294.631437%
// free spins 56221440, q = 0.34091, sq = 1/(1-q) = 1.517241
// free games frequency: 1/3.3615
// RTP = sq*rtp(sym) = 1.5172*294.63 = 447.027007%
// *regular reels calculations*
// reels lengths [44, 44, 44, 44, 44], total reshuffles 164916224
// symbols: 58.926(lined) + 8.7891(scatter) = 67.715377%
// free spins 7146225, q = 0.043332, sq = 1/(1-q) = 1.045295
// free games frequency: 1/350.35
// RTP = 67.715(sym) + 0.043332*447.03(fg) = 87.086155%
var Reels870 = slot.Reels5x{
	{10, 11, 9, 12, 8, 9, 1, 8, 7, 5, 8, 6, 12, 10, 11, 9, 7, 12, 5, 11, 10, 7, 12, 9, 11, 4, 7, 10, 11, 6, 10, 2, 5, 11, 3, 10, 9, 12, 6, 4, 3, 12, 8, 9},
	{11, 10, 9, 7, 11, 6, 7, 3, 5, 11, 10, 9, 8, 7, 10, 11, 8, 12, 5, 2, 8, 12, 10, 11, 12, 6, 8, 10, 4, 12, 9, 11, 7, 4, 9, 5, 12, 10, 9, 3, 6, 9, 1, 12},
	{7, 5, 12, 11, 2, 9, 8, 6, 10, 11, 5, 9, 12, 10, 8, 4, 11, 3, 6, 9, 10, 7, 8, 1, 9, 11, 12, 10, 11, 5, 10, 3, 8, 6, 12, 7, 9, 11, 12, 9, 7, 10, 4, 12},
	{9, 10, 6, 12, 9, 7, 5, 1, 4, 9, 10, 11, 9, 3, 11, 6, 7, 11, 5, 12, 11, 7, 8, 12, 2, 10, 12, 6, 8, 12, 10, 9, 8, 7, 10, 11, 9, 10, 5, 11, 3, 12, 4, 8},
	{10, 11, 12, 3, 7, 5, 4, 1, 11, 10, 7, 6, 9, 5, 8, 9, 10, 11, 12, 9, 8, 10, 6, 2, 4, 3, 6, 11, 12, 8, 10, 9, 12, 11, 10, 12, 7, 9, 11, 8, 5, 7, 9, 12},
}

// *bonus reels calculations*
// reels lengths [43, 44, 44, 44, 43], total reshuffles 157505216
// symbols: 296.09(lined) + 0(scatter) = 296.090162%
// free spins 54194448, q = 0.34408, sq = 1/(1-q) = 1.524577
// free games frequency: 1/3.3347
// RTP = sq*rtp(sym) = 1.5246*296.09 = 451.412238%
// *regular reels calculations*
// reels lengths [43, 44, 44, 44, 43], total reshuffles 157505216
// symbols: 59.218(lined) + 8.9495(scatter) = 68.167582%
// free spins 7010550, q = 0.04451, sq = 1/(1-q) = 1.046583
// free games frequency: 1/341.12
// RTP = 68.168(sym) + 0.04451*451.41(fg) = 88.259920%
var Reels882 = slot.Reels5x{
	{6, 12, 9, 8, 5, 6, 7, 10, 11, 8, 10, 11, 12, 10, 4, 7, 9, 12, 10, 11, 7, 1, 3, 11, 4, 8, 10, 11, 5, 12, 10, 2, 12, 9, 7, 6, 5, 9, 12, 8, 9, 11, 3},
	{11, 10, 9, 7, 11, 6, 7, 3, 5, 11, 10, 9, 8, 7, 10, 11, 8, 12, 5, 2, 8, 12, 10, 11, 12, 6, 8, 10, 4, 12, 9, 11, 7, 4, 9, 5, 12, 10, 9, 3, 6, 9, 1, 12},
	{7, 5, 12, 11, 2, 9, 8, 6, 10, 11, 5, 9, 12, 10, 8, 4, 11, 3, 6, 9, 10, 7, 8, 1, 9, 11, 12, 10, 11, 5, 10, 3, 8, 6, 12, 7, 9, 11, 12, 9, 7, 10, 4, 12},
	{9, 10, 6, 12, 9, 7, 5, 1, 4, 9, 10, 11, 9, 3, 11, 6, 7, 11, 5, 12, 11, 7, 8, 12, 2, 10, 12, 6, 8, 12, 10, 9, 8, 7, 10, 11, 9, 10, 5, 11, 3, 12, 4, 8},
	{12, 10, 9, 11, 7, 10, 8, 11, 10, 12, 11, 6, 12, 8, 3, 2, 7, 12, 4, 10, 9, 6, 12, 10, 8, 1, 10, 9, 8, 7, 5, 3, 6, 11, 7, 4, 11, 5, 9, 12, 11, 5, 9},
}

// *bonus reels calculations*
// reels lengths [43, 44, 43, 44, 43], total reshuffles 153925552
// symbols: 298.03(lined) + 0(scatter) = 298.026510%
// free spins 53206824, q = 0.34567, sq = 1/(1-q) = 1.528271
// free games frequency: 1/3.3215
// RTP = sq*rtp(sym) = 1.5283*298.03 = 455.465393%
// *regular reels calculations*
// reels lengths [43, 44, 43, 44, 43], total reshuffles 153925552
// symbols: 59.605(lined) + 9.0304(scatter) = 68.635744%
// free spins 6943320, q = 0.045108, sq = 1/(1-q) = 1.047239
// free games frequency: 1/336.62
// RTP = 68.636(sym) + 0.045108*455.47(fg) = 89.181014%
var Reels891 = slot.Reels5x{
	{6, 12, 9, 8, 5, 6, 7, 10, 11, 8, 10, 11, 12, 10, 4, 7, 9, 12, 10, 11, 7, 1, 3, 11, 4, 8, 10, 11, 5, 12, 10, 2, 12, 9, 7, 6, 5, 9, 12, 8, 9, 11, 3},
	{11, 10, 9, 7, 11, 6, 7, 3, 5, 11, 10, 9, 8, 7, 10, 11, 8, 12, 5, 2, 8, 12, 10, 11, 12, 6, 8, 10, 4, 12, 9, 11, 7, 4, 9, 5, 12, 10, 9, 3, 6, 9, 1, 12},
	{10, 7, 8, 4, 9, 6, 4, 5, 10, 12, 9, 11, 7, 12, 6, 9, 8, 10, 11, 9, 12, 7, 11, 8, 5, 7, 11, 3, 12, 10, 8, 12, 1, 11, 10, 6, 11, 3, 10, 2, 9, 12, 5},
	{9, 10, 6, 12, 9, 7, 5, 1, 4, 9, 10, 11, 9, 3, 11, 6, 7, 11, 5, 12, 11, 7, 8, 12, 2, 10, 12, 6, 8, 12, 10, 9, 8, 7, 10, 11, 9, 10, 5, 11, 3, 12, 4, 8},
	{12, 10, 9, 11, 7, 10, 8, 11, 10, 12, 11, 6, 12, 8, 3, 2, 7, 12, 4, 10, 9, 6, 12, 10, 8, 1, 10, 9, 8, 7, 5, 3, 6, 11, 7, 4, 11, 5, 9, 12, 11, 5, 9},
}

// *bonus reels calculations*
// reels lengths [43, 43, 43, 44, 43], total reshuffles 150427244
// symbols: 300.49(lined) + 0(scatter) = 300.487756%
// free spins 52236099, q = 0.34725, sq = 1/(1-q) = 1.531984
// free games frequency: 1/3.3085
// RTP = sq*rtp(sym) = 1.532*300.49 = 460.342376%
// *regular reels calculations*
// reels lengths [43, 43, 43, 44, 43], total reshuffles 150427244
// symbols: 60.098(lined) + 9.1118(scatter) = 69.209328%
// free spins 6876495, q = 0.045713, sq = 1/(1-q) = 1.047903
// free games frequency: 1/332.18
// RTP = 69.209(sym) + 0.045713*460.34(fg) = 90.253003%
var Reels902 = slot.Reels5x{
	{6, 12, 9, 8, 5, 6, 7, 10, 11, 8, 10, 11, 12, 10, 4, 7, 9, 12, 10, 11, 7, 1, 3, 11, 4, 8, 10, 11, 5, 12, 10, 2, 12, 9, 7, 6, 5, 9, 12, 8, 9, 11, 3},
	{12, 11, 9, 5, 6, 9, 8, 10, 7, 11, 10, 12, 5, 11, 3, 12, 8, 9, 11, 4, 12, 7, 10, 11, 3, 4, 10, 2, 9, 5, 10, 11, 8, 6, 12, 8, 7, 9, 12, 6, 10, 7, 1},
	{10, 7, 8, 4, 9, 6, 4, 5, 10, 12, 9, 11, 7, 12, 6, 9, 8, 10, 11, 9, 12, 7, 11, 8, 5, 7, 11, 3, 12, 10, 8, 12, 1, 11, 10, 6, 11, 3, 10, 2, 9, 12, 5},
	{9, 10, 6, 12, 9, 7, 5, 1, 4, 9, 10, 11, 9, 3, 11, 6, 7, 11, 5, 12, 11, 7, 8, 12, 2, 10, 12, 6, 8, 12, 10, 9, 8, 7, 10, 11, 9, 10, 5, 11, 3, 12, 4, 8},
	{12, 10, 9, 11, 7, 10, 8, 11, 10, 12, 11, 6, 12, 8, 3, 2, 7, 12, 4, 10, 9, 6, 12, 10, 8, 1, 10, 9, 8, 7, 5, 3, 6, 11, 7, 4, 11, 5, 9, 12, 11, 5, 9},
}

// *bonus reels calculations*
// reels lengths [43, 43, 43, 43, 43], total reshuffles 147008443
// symbols: 301.85(lined) + 0(scatter) = 301.848884%
// free spins 51282015, q = 0.34884, sq = 1/(1-q) = 1.535714
// free games frequency: 1/3.2955
// RTP = sq*rtp(sym) = 1.5357*301.85 = 463.553644%
// *regular reels calculations*
// reels lengths [43, 43, 43, 43, 43], total reshuffles 147008443
// symbols: 60.37(lined) + 9.1936(scatter) = 69.563330%
// free spins 6810075, q = 0.046324, sq = 1/(1-q) = 1.048575
// free games frequency: 1/327.82
// RTP = 69.563(sym) + 0.046324*463.55(fg) = 91.037166%
var Reels910 = slot.Reels5x{
	{6, 12, 9, 8, 5, 6, 7, 10, 11, 8, 10, 11, 12, 10, 4, 7, 9, 12, 10, 11, 7, 1, 3, 11, 4, 8, 10, 11, 5, 12, 10, 2, 12, 9, 7, 6, 5, 9, 12, 8, 9, 11, 3},
	{12, 11, 9, 5, 6, 9, 8, 10, 7, 11, 10, 12, 5, 11, 3, 12, 8, 9, 11, 4, 12, 7, 10, 11, 3, 4, 10, 2, 9, 5, 10, 11, 8, 6, 12, 8, 7, 9, 12, 6, 10, 7, 1},
	{10, 7, 8, 4, 9, 6, 4, 5, 10, 12, 9, 11, 7, 12, 6, 9, 8, 10, 11, 9, 12, 7, 11, 8, 5, 7, 11, 3, 12, 10, 8, 12, 1, 11, 10, 6, 11, 3, 10, 2, 9, 12, 5},
	{8, 9, 10, 3, 1, 5, 12, 7, 4, 10, 7, 11, 12, 10, 11, 6, 12, 9, 11, 5, 12, 8, 2, 5, 8, 9, 10, 4, 11, 6, 3, 11, 8, 12, 9, 7, 10, 9, 12, 6, 11, 10, 7},
	{12, 10, 9, 11, 7, 10, 8, 11, 10, 12, 11, 6, 12, 8, 3, 2, 7, 12, 4, 10, 9, 6, 12, 10, 8, 1, 10, 9, 8, 7, 5, 3, 6, 11, 7, 4, 11, 5, 9, 12, 11, 5, 9},
}

// *bonus reels calculations*
// reels lengths [43, 43, 42, 43, 43], total reshuffles 143589642
// symbols: 304.15(lined) + 0(scatter) = 304.152806%
// free spins 50327931, q = 0.3505, sq = 1/(1-q) = 1.539642
// free games frequency: 1/3.2821
// RTP = sq*rtp(sym) = 1.5396*304.15 = 468.286417%
// *regular reels calculations*
// reels lengths [43, 43, 42, 43, 43], total reshuffles 143589642
// symbols: 60.831(lined) + 9.2792(scatter) = 70.109785%
// free spins 6743655, q = 0.046965, sq = 1/(1-q) = 1.049279
// free games frequency: 1/323.37
// RTP = 70.11(sym) + 0.046965*468.29(fg) = 92.102751%
var Reels921 = slot.Reels5x{
	{6, 12, 9, 8, 5, 6, 7, 10, 11, 8, 10, 11, 12, 10, 4, 7, 9, 12, 10, 11, 7, 1, 3, 11, 4, 8, 10, 11, 5, 12, 10, 2, 12, 9, 7, 6, 5, 9, 12, 8, 9, 11, 3},
	{12, 11, 9, 5, 6, 9, 8, 10, 7, 11, 10, 12, 5, 11, 3, 12, 8, 9, 11, 4, 12, 7, 10, 11, 3, 4, 10, 2, 9, 5, 10, 11, 8, 6, 12, 8, 7, 9, 12, 6, 10, 7, 1},
	{5, 11, 12, 6, 11, 4, 12, 8, 6, 9, 12, 1, 6, 7, 2, 10, 7, 3, 11, 9, 8, 11, 10, 9, 5, 10, 8, 7, 11, 9, 10, 4, 7, 12, 5, 3, 12, 10, 11, 12, 8, 9},
	{8, 9, 10, 3, 1, 5, 12, 7, 4, 10, 7, 11, 12, 10, 11, 6, 12, 9, 11, 5, 12, 8, 2, 5, 8, 9, 10, 4, 11, 6, 3, 11, 8, 12, 9, 7, 10, 9, 12, 6, 11, 10, 7},
	{12, 10, 9, 11, 7, 10, 8, 11, 10, 12, 11, 6, 12, 8, 3, 2, 7, 12, 4, 10, 9, 6, 12, 10, 8, 1, 10, 9, 8, 7, 5, 3, 6, 11, 7, 4, 11, 5, 9, 12, 11, 5, 9},
}

// *bonus reels calculations*
// reels lengths [43, 43, 42, 42, 42], total reshuffles 136988712
// symbols: 305.87(lined) + 0(scatter) = 305.867716%
// free spins 48469428, q = 0.35382, sq = 1/(1-q) = 1.547558
// free games frequency: 1/3.2556
// RTP = sq*rtp(sym) = 1.5476*305.87 = 473.347983%
// *regular reels calculations*
// reels lengths [43, 43, 42, 42, 42], total reshuffles 136988712
// symbols: 61.174(lined) + 9.452(scatter) = 70.625564%
// free spins 6612030, q = 0.048267, sq = 1/(1-q) = 1.050715
// free games frequency: 1/314.68
// RTP = 70.626(sym) + 0.048267*473.35(fg) = 93.472636%
var Reels934 = slot.Reels5x{
	{6, 12, 9, 8, 5, 6, 7, 10, 11, 8, 10, 11, 12, 10, 4, 7, 9, 12, 10, 11, 7, 1, 3, 11, 4, 8, 10, 11, 5, 12, 10, 2, 12, 9, 7, 6, 5, 9, 12, 8, 9, 11, 3},
	{12, 11, 9, 5, 6, 9, 8, 10, 7, 11, 10, 12, 5, 11, 3, 12, 8, 9, 11, 4, 12, 7, 10, 11, 3, 4, 10, 2, 9, 5, 10, 11, 8, 6, 12, 8, 7, 9, 12, 6, 10, 7, 1},
	{5, 11, 12, 6, 11, 4, 12, 8, 6, 9, 12, 1, 6, 7, 2, 10, 7, 3, 11, 9, 8, 11, 10, 9, 5, 10, 8, 7, 11, 9, 10, 4, 7, 12, 5, 3, 12, 10, 11, 12, 8, 9},
	{9, 5, 6, 11, 8, 3, 1, 7, 9, 11, 12, 9, 4, 10, 2, 5, 12, 8, 11, 5, 10, 8, 3, 10, 12, 9, 10, 6, 7, 11, 12, 7, 10, 12, 6, 11, 8, 4, 9, 11, 12, 7},
	{1, 10, 3, 12, 10, 8, 6, 11, 9, 12, 8, 4, 11, 8, 9, 5, 7, 11, 12, 8, 11, 12, 7, 10, 3, 2, 5, 9, 10, 12, 11, 6, 4, 12, 9, 7, 6, 10, 11, 9, 7, 5},
}

// *bonus reels calculations*
// reels lengths [43, 42, 42, 42, 43], total reshuffles 136988712
// symbols: 308.41(lined) + 0(scatter) = 308.406619%
// free spins 48469428, q = 0.35382, sq = 1/(1-q) = 1.547558
// free games frequency: 1/3.2556
// RTP = sq*rtp(sym) = 1.5476*308.41 = 477.277081%
// *regular reels calculations*
// reels lengths [43, 42, 42, 42, 43], total reshuffles 136988712
// symbols: 61.681(lined) + 9.452(scatter) = 71.133344%
// free spins 6612030, q = 0.048267, sq = 1/(1-q) = 1.050715
// free games frequency: 1/314.68
// RTP = 71.133(sym) + 0.048267*477.28(fg) = 94.170063%
var Reels941 = slot.Reels5x{
	{6, 12, 9, 8, 5, 6, 7, 10, 11, 8, 10, 11, 12, 10, 4, 7, 9, 12, 10, 11, 7, 1, 3, 11, 4, 8, 10, 11, 5, 12, 10, 2, 12, 9, 7, 6, 5, 9, 12, 8, 9, 11, 3},
	{2, 12, 3, 10, 11, 7, 8, 9, 11, 6, 4, 7, 12, 8, 5, 11, 9, 8, 12, 5, 10, 9, 8, 11, 12, 9, 6, 10, 1, 3, 11, 12, 7, 9, 5, 4, 10, 6, 11, 10, 12, 7},
	{5, 11, 12, 6, 11, 4, 12, 8, 6, 9, 12, 1, 6, 7, 2, 10, 7, 3, 11, 9, 8, 11, 10, 9, 5, 10, 8, 7, 11, 9, 10, 4, 7, 12, 5, 3, 12, 10, 11, 12, 8, 9},
	{9, 5, 6, 11, 8, 3, 1, 7, 9, 11, 12, 9, 4, 10, 2, 5, 12, 8, 11, 5, 10, 8, 3, 10, 12, 9, 10, 6, 7, 11, 12, 7, 10, 12, 6, 11, 8, 4, 9, 11, 12, 7},
	{12, 10, 9, 11, 7, 10, 8, 11, 10, 12, 11, 6, 12, 8, 3, 2, 7, 12, 4, 10, 9, 6, 12, 10, 8, 1, 10, 9, 8, 7, 5, 3, 6, 11, 7, 4, 11, 5, 9, 12, 11, 5, 9},
}

// *bonus reels calculations*
// reels lengths [42, 42, 42, 43, 42], total reshuffles 133802928
// symbols: 310.91(lined) + 0(scatter) = 310.914590%
// free spins 47564496, q = 0.35548, sq = 1/(1-q) = 1.551546
// free games frequency: 1/3.2425
// RTP = sq*rtp(sym) = 1.5515*310.91 = 482.398410%
// *regular reels calculations*
// reels lengths [42, 42, 42, 43, 42], total reshuffles 133802928
// symbols: 62.183(lined) + 9.5391(scatter) = 71.722064%
// free spins 6546825, q = 0.048929, sq = 1/(1-q) = 1.051446
// free games frequency: 1/310.45
// RTP = 71.722(sym) + 0.048929*482.4(fg) = 95.325269%
var Reels953 = slot.Reels5x{
	{11, 7, 9, 12, 8, 11, 10, 6, 5, 11, 10, 5, 11, 12, 9, 10, 7, 4, 11, 9, 8, 10, 12, 2, 10, 12, 6, 3, 7, 11, 5, 4, 3, 12, 8, 9, 6, 1, 7, 12, 8, 9},
	{2, 12, 3, 10, 11, 7, 8, 9, 11, 6, 4, 7, 12, 8, 5, 11, 9, 8, 12, 5, 10, 9, 8, 11, 12, 9, 6, 10, 1, 3, 11, 12, 7, 9, 5, 4, 10, 6, 11, 10, 12, 7},
	{5, 11, 12, 6, 11, 4, 12, 8, 6, 9, 12, 1, 6, 7, 2, 10, 7, 3, 11, 9, 8, 11, 10, 9, 5, 10, 8, 7, 11, 9, 10, 4, 7, 12, 5, 3, 12, 10, 11, 12, 8, 9},
	{8, 9, 10, 3, 1, 5, 12, 7, 4, 10, 7, 11, 12, 10, 11, 6, 12, 9, 11, 5, 12, 8, 2, 5, 8, 9, 10, 4, 11, 6, 3, 11, 8, 12, 9, 7, 10, 9, 12, 6, 11, 10, 7},
	{1, 10, 3, 12, 10, 8, 6, 11, 9, 12, 8, 4, 11, 8, 9, 5, 7, 11, 12, 8, 11, 12, 7, 10, 3, 2, 5, 9, 10, 12, 11, 6, 4, 12, 9, 7, 6, 10, 11, 9, 7, 5},
}

// *bonus reels calculations*
// reels lengths [42, 42, 42, 42, 42], total reshuffles 130691232
// symbols: 312.66(lined) + 0(scatter) = 312.662937%
// free spins 46675440, q = 0.35714, sq = 1/(1-q) = 1.555556
// free games frequency: 1/3.2296
// RTP = sq*rtp(sym) = 1.5556*312.66 = 486.364569%
// *regular reels calculations*
// reels lengths [42, 42, 42, 42, 42], total reshuffles 130691232
// symbols: 62.533(lined) + 9.6268(scatter) = 72.159343%
// free spins 6482025, q = 0.049598, sq = 1/(1-q) = 1.052186
// free games frequency: 1/306.28
// RTP = 72.159(sym) + 0.049598*486.36(fg) = 96.282057%
var Reels962 = slot.Reels5x{
	{11, 7, 9, 12, 8, 11, 10, 6, 5, 11, 10, 5, 11, 12, 9, 10, 7, 4, 11, 9, 8, 10, 12, 2, 10, 12, 6, 3, 7, 11, 5, 4, 3, 12, 8, 9, 6, 1, 7, 12, 8, 9},
	{2, 12, 3, 10, 11, 7, 8, 9, 11, 6, 4, 7, 12, 8, 5, 11, 9, 8, 12, 5, 10, 9, 8, 11, 12, 9, 6, 10, 1, 3, 11, 12, 7, 9, 5, 4, 10, 6, 11, 10, 12, 7},
	{5, 11, 12, 6, 11, 4, 12, 8, 6, 9, 12, 1, 6, 7, 2, 10, 7, 3, 11, 9, 8, 11, 10, 9, 5, 10, 8, 7, 11, 9, 10, 4, 7, 12, 5, 3, 12, 10, 11, 12, 8, 9},
	{9, 5, 6, 11, 8, 3, 1, 7, 9, 11, 12, 9, 4, 10, 2, 5, 12, 8, 11, 5, 10, 8, 3, 10, 12, 9, 10, 6, 7, 11, 12, 7, 10, 12, 6, 11, 8, 4, 9, 11, 12, 7},
	{1, 10, 3, 12, 10, 8, 6, 11, 9, 12, 8, 4, 11, 8, 9, 5, 7, 11, 12, 8, 11, 12, 7, 10, 3, 2, 5, 9, 10, 12, 11, 6, 4, 12, 9, 7, 6, 10, 11, 9, 7, 5},
}

// *bonus reels calculations*
// reels lengths [42, 42, 42, 40, 42], total reshuffles 124467840
// symbols: 315.29(lined) + 0(scatter) = 315.289295%
// free spins 44897328, q = 0.36071, sq = 1/(1-q) = 1.564246
// free games frequency: 1/3.2021
// RTP = sq*rtp(sym) = 1.5642*315.29 = 493.189958%
// *regular reels calculations*
// reels lengths [42, 42, 42, 40, 42], total reshuffles 124467840
// symbols: 63.058(lined) + 9.8151(scatter) = 72.872975%
// free spins 6352425, q = 0.051037, sq = 1/(1-q) = 1.053782
// free games frequency: 1/297.68
// RTP = 72.873(sym) + 0.051037*493.19(fg) = 98.043752%
var Reels980 = slot.Reels5x{
	{11, 7, 9, 12, 8, 11, 10, 6, 5, 11, 10, 5, 11, 12, 9, 10, 7, 4, 11, 9, 8, 10, 12, 2, 10, 12, 6, 3, 7, 11, 5, 4, 3, 12, 8, 9, 6, 1, 7, 12, 8, 9},
	{2, 12, 3, 10, 11, 7, 8, 9, 11, 6, 4, 7, 12, 8, 5, 11, 9, 8, 12, 5, 10, 9, 8, 11, 12, 9, 6, 10, 1, 3, 11, 12, 7, 9, 5, 4, 10, 6, 11, 10, 12, 7},
	{5, 11, 12, 6, 11, 4, 12, 8, 6, 9, 12, 1, 6, 7, 2, 10, 7, 3, 11, 9, 8, 11, 10, 9, 5, 10, 8, 7, 11, 9, 10, 4, 7, 12, 5, 3, 12, 10, 11, 12, 8, 9},
	{1, 9, 5, 8, 6, 12, 10, 9, 11, 12, 9, 10, 3, 7, 12, 10, 8, 12, 10, 11, 4, 7, 11, 9, 7, 6, 5, 7, 4, 9, 12, 11, 3, 8, 11, 5, 2, 6, 10, 8},
	{1, 10, 3, 12, 10, 8, 6, 11, 9, 12, 8, 4, 11, 8, 9, 5, 7, 11, 12, 8, 11, 12, 7, 10, 3, 2, 5, 9, 10, 12, 11, 6, 4, 12, 9, 7, 6, 10, 11, 9, 7, 5},
}

// *bonus reels calculations*
// reels lengths [42, 42, 40, 42, 42], total reshuffles 124467840
// symbols: 321.05(lined) + 0(scatter) = 321.051362%
// free spins 44897328, q = 0.36071, sq = 1/(1-q) = 1.564246
// free games frequency: 1/3.2021
// RTP = sq*rtp(sym) = 1.5642*321.05 = 502.203247%
// *regular reels calculations*
// reels lengths [42, 42, 40, 42, 42], total reshuffles 124467840
// symbols: 64.21(lined) + 9.8151(scatter) = 74.025388%
// free spins 6352425, q = 0.051037, sq = 1/(1-q) = 1.053782
// free games frequency: 1/297.68
// RTP = 74.025(sym) + 0.051037*502.2(fg) = 99.656174%
var Reels996 = slot.Reels5x{
	{11, 7, 9, 12, 8, 11, 10, 6, 5, 11, 10, 5, 11, 12, 9, 10, 7, 4, 11, 9, 8, 10, 12, 2, 10, 12, 6, 3, 7, 11, 5, 4, 3, 12, 8, 9, 6, 1, 7, 12, 8, 9},
	{2, 12, 3, 10, 11, 7, 8, 9, 11, 6, 4, 7, 12, 8, 5, 11, 9, 8, 12, 5, 10, 9, 8, 11, 12, 9, 6, 10, 1, 3, 11, 12, 7, 9, 5, 4, 10, 6, 11, 10, 12, 7},
	{5, 10, 11, 9, 3, 10, 2, 8, 9, 10, 8, 11, 7, 6, 11, 8, 12, 6, 9, 4, 12, 9, 5, 7, 10, 12, 9, 7, 1, 4, 12, 8, 11, 5, 10, 6, 12, 7, 3, 11},
	{9, 5, 6, 11, 8, 3, 1, 7, 9, 11, 12, 9, 4, 10, 2, 5, 12, 8, 11, 5, 10, 8, 3, 10, 12, 9, 10, 6, 7, 11, 12, 7, 10, 12, 6, 11, 8, 4, 9, 11, 12, 7},
	{1, 10, 3, 12, 10, 8, 6, 11, 9, 12, 8, 4, 11, 8, 9, 5, 7, 11, 12, 8, 11, 12, 7, 10, 3, 2, 5, 9, 10, 12, 11, 6, 4, 12, 9, 7, 6, 10, 11, 9, 7, 5},
}

// *bonus reels calculations*
// reels lengths [40, 40, 40, 40, 40], total reshuffles 102400000
// symbols: 346.73(lined) + 0(scatter) = 346.733862%
// free spins 38400000, q = 0.375, sq = 1/(1-q) = 1.600000
// free games frequency: 1/3.0978
// RTP = sq*rtp(sym) = 1.6*346.73 = 554.774180%
// *regular reels calculations*
// reels lengths [40, 40, 40, 40, 40], total reshuffles 102400000
// symbols: 69.347(lined) + 10.591(scatter) = 79.937725%
// free spins 5850225, q = 0.057131, sq = 1/(1-q) = 1.060593
// free games frequency: 1/266.07
// RTP = 79.938(sym) + 0.057131*554.77(fg) = 111.632586%
var Reels112 = slot.Reels5x{
	{1, 12, 7, 6, 10, 12, 9, 11, 6, 2, 3, 11, 5, 8, 4, 3, 12, 10, 8, 6, 9, 8, 4, 9, 11, 7, 10, 8, 9, 10, 12, 9, 11, 5, 7, 11, 12, 5, 10, 7},
	{11, 5, 8, 10, 3, 12, 7, 11, 8, 6, 12, 11, 2, 9, 10, 5, 8, 7, 10, 9, 4, 6, 5, 12, 7, 10, 9, 11, 10, 12, 9, 8, 7, 12, 4, 6, 11, 1, 9, 3},
	{5, 10, 11, 9, 3, 10, 2, 8, 9, 10, 8, 11, 7, 6, 11, 8, 12, 6, 9, 4, 12, 9, 5, 7, 10, 12, 9, 7, 1, 4, 12, 8, 11, 5, 10, 6, 12, 7, 3, 11},
	{1, 9, 5, 8, 6, 12, 10, 9, 11, 12, 9, 10, 3, 7, 12, 10, 8, 12, 10, 11, 4, 7, 11, 9, 7, 6, 5, 7, 4, 9, 12, 11, 3, 8, 11, 5, 2, 6, 10, 8},
	{11, 8, 9, 12, 11, 9, 4, 5, 10, 1, 6, 3, 9, 11, 10, 12, 7, 10, 8, 5, 12, 9, 2, 4, 7, 9, 8, 10, 11, 8, 6, 12, 7, 5, 12, 6, 11, 3, 7, 10},
}

var Reels50 = slot.Reels5x{
	{6, 12, 9, 8, 5, 6, 7, 10, 11, 8, 10, 11, 12, 10, 4, 7, 9, 12, 10, 11, 7, 1, 3, 11, 4, 8, 10, 11, 5, 12, 10, 2, 12, 9, 7, 6, 5, 9, 12, 8, 9, 11, 3},
	{11, 10, 9, 7, 11, 6, 7, 3, 5, 11, 10, 9, 8, 7, 10, 11, 8, 12, 5, 2, 8, 12, 10, 11, 12, 6, 8, 10, 4, 12, 9, 11, 7, 4, 9, 5, 12, 10, 9, 3, 6, 9, 1, 12},
	{7, 5, 12, 11, 2, 9, 8, 6, 10, 11, 5, 9, 12, 10, 8, 4, 11, 3, 6, 9, 10, 7, 8, 1, 9, 11, 12, 10, 11, 5, 10, 3, 8, 6, 12, 7, 9, 11, 12, 9, 7, 10, 4, 12},
	{9, 10, 6, 12, 9, 7, 5, 1, 4, 9, 10, 11, 9, 3, 11, 6, 7, 11, 5, 12, 11, 7, 8, 12, 2, 10, 12, 6, 8, 12, 10, 9, 8, 7, 10, 11, 9, 10, 5, 11, 3, 12, 4, 8},
	{12, 10, 9, 11, 7, 10, 8, 11, 10, 12, 11, 6, 12, 8, 3, 2, 7, 12, 4, 10, 9, 6, 12, 10, 8, 1, 10, 9, 8, 7, 5, 3, 6, 11, 7, 4, 11, 5, 9, 12, 11, 5, 9},
}

// Map with available reels.
var ReelsMap = map[float64]*slot.Reels5x{
	87.086155:  &Reels870,
	88.259920:  &Reels882,
	89.181014:  &Reels891,
	90.253003:  &Reels902,
	91.037166:  &Reels910,
	92.102751:  &Reels921,
	93.472636:  &Reels934,
	94.170063:  &Reels941,
	95.325269:  &Reels953,
	96.282057:  &Reels962,
	98.043752:  &Reels980,
	99.656174:  &Reels996,
	111.632586: &Reels112,
}
