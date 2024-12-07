package fruitshop

import (
	"github.com/slotopol/server/game/slot"
)

// *bonus reels calculations*
// reels lengths [49, 51, 51, 50, 49], total reshuffles 312250050
// symbols: rtp(sym) = 108.429994%
// free spins 159657390, q = 0.51131, sq = 1/(1-q) = 2.046298
// free games frequency: 1/2.2859
// RTP = sq*rtp(sym) = 2.0463*108.43 = 221.880076%
// *regular reels calculations*
// reels lengths [49, 51, 51, 50, 49], total reshuffles 312250050
// symbols: rtp(sym) = 54.214997%
// free spins 40443750, q = 0.12952, sq = 1/(1-q) = 1.148796
// free games frequency: 1/7.978
// RTP = 54.215(sym) + 0.12952*221.88(fg) = 82.953703%
var Reels829 = slot.Reels5x{
	{8, 11, 6, 4, 8, 9, 4, 8, 5, 7, 11, 10, 5, 9, 11, 8, 10, 3, 11, 4, 3, 8, 7, 2, 11, 10, 3, 6, 7, 11, 9, 8, 2, 10, 9, 6, 5, 10, 7, 9, 11, 10, 2, 7, 8, 9, 10, 7, 9},
	{7, 10, 5, 9, 2, 4, 10, 8, 5, 6, 10, 2, 9, 8, 11, 10, 9, 7, 4, 11, 9, 7, 11, 6, 10, 4, 8, 1, 3, 7, 11, 9, 8, 3, 1, 7, 10, 8, 2, 11, 10, 9, 11, 7, 8, 11, 9, 3, 6, 5, 8},
	{10, 7, 9, 2, 8, 3, 10, 1, 6, 11, 9, 7, 3, 11, 5, 9, 4, 10, 7, 1, 10, 11, 6, 8, 9, 6, 2, 9, 11, 8, 9, 7, 10, 2, 8, 10, 11, 4, 3, 9, 8, 5, 11, 10, 8, 7, 11, 5, 7, 4, 8},
	{11, 7, 9, 3, 8, 5, 11, 9, 10, 2, 5, 11, 6, 7, 10, 8, 4, 5, 11, 3, 8, 4, 3, 10, 7, 9, 10, 11, 8, 9, 11, 10, 7, 1, 2, 8, 6, 7, 10, 9, 4, 7, 9, 10, 6, 8, 11, 9, 8, 2},
	{4, 3, 2, 9, 3, 10, 8, 7, 9, 2, 8, 7, 11, 10, 4, 9, 10, 6, 7, 11, 9, 2, 11, 5, 10, 11, 9, 10, 11, 5, 6, 10, 8, 7, 5, 8, 4, 9, 11, 8, 9, 10, 8, 7, 3, 8, 7, 11, 6},
}

// *bonus reels calculations*
// reels lengths [49, 50, 50, 50, 49], total reshuffles 300125000
// symbols: rtp(sym) = 109.145586%
// free spins 153733290, q = 0.51223, sq = 1/(1-q) = 2.050150
// free games frequency: 1/2.2793
// RTP = sq*rtp(sym) = 2.0502*109.15 = 223.764850%
// *regular reels calculations*
// reels lengths [49, 50, 50, 50, 49], total reshuffles 300125000
// symbols: rtp(sym) = 54.572793%
// free spins 39892500, q = 0.13292, sq = 1/(1-q) = 1.153296
// free games frequency: 1/7.7778
// RTP = 54.573(sym) + 0.13292*223.76(fg) = 84.315531%
var Reels843 = slot.Reels5x{
	{8, 11, 6, 4, 8, 9, 4, 8, 5, 7, 11, 10, 5, 9, 11, 8, 10, 3, 11, 4, 3, 8, 7, 2, 11, 10, 3, 6, 7, 11, 9, 8, 2, 10, 9, 6, 5, 10, 7, 9, 11, 10, 2, 7, 8, 9, 10, 7, 9},
	{2, 11, 8, 3, 10, 7, 11, 9, 4, 8, 1, 7, 6, 9, 10, 5, 11, 9, 10, 1, 7, 8, 11, 10, 8, 9, 3, 5, 11, 8, 4, 2, 10, 9, 4, 3, 11, 6, 10, 9, 6, 7, 10, 5, 2, 7, 8, 11, 9, 7},
	{6, 8, 11, 10, 4, 9, 11, 2, 3, 10, 7, 9, 4, 11, 7, 6, 1, 5, 9, 8, 10, 11, 7, 2, 10, 9, 5, 10, 9, 4, 8, 10, 3, 8, 7, 2, 1, 7, 9, 11, 6, 7, 11, 8, 5, 9, 8, 11, 3, 10},
	{11, 7, 9, 3, 8, 5, 11, 9, 10, 2, 5, 11, 6, 7, 10, 8, 4, 5, 11, 3, 8, 4, 3, 10, 7, 9, 10, 11, 8, 9, 11, 10, 7, 1, 2, 8, 6, 7, 10, 9, 4, 7, 9, 10, 6, 8, 11, 9, 8, 2},
	{4, 3, 2, 9, 3, 10, 8, 7, 9, 2, 8, 7, 11, 10, 4, 9, 10, 6, 7, 11, 9, 2, 11, 5, 10, 11, 9, 10, 11, 5, 6, 10, 8, 7, 5, 8, 4, 9, 11, 8, 9, 10, 8, 7, 3, 8, 7, 11, 6},
}

// *bonus reels calculations*
// reels lengths [48, 50, 50, 49, 48], total reshuffles 282240000
// symbols: rtp(sym) = 109.512868%
// free spins 144839340, q = 0.51318, sq = 1/(1-q) = 2.054139
// free games frequency: 1/2.2734
// RTP = sq*rtp(sym) = 2.0541*109.51 = 224.954611%
// *regular reels calculations*
// reels lengths [48, 50, 50, 49, 48], total reshuffles 282240000
// symbols: rtp(sym) = 54.756434%
// free spins 38326500, q = 0.13579, sq = 1/(1-q) = 1.157132
// free games frequency: 1/7.619
// RTP = 54.756(sym) + 0.13579*224.95(fg) = 85.303922%
var Reels853 = slot.Reels5x{
	{3, 7, 10, 8, 2, 11, 3, 6, 9, 11, 10, 8, 2, 7, 11, 9, 3, 5, 7, 10, 11, 4, 5, 9, 8, 10, 7, 2, 11, 9, 4, 8, 7, 6, 8, 10, 9, 11, 8, 10, 5, 6, 9, 7, 10, 11, 4, 9},
	{2, 11, 8, 3, 10, 7, 11, 9, 4, 8, 1, 7, 6, 9, 10, 5, 11, 9, 10, 1, 7, 8, 11, 10, 8, 9, 3, 5, 11, 8, 4, 2, 10, 9, 4, 3, 11, 6, 10, 9, 6, 7, 10, 5, 2, 7, 8, 11, 9, 7},
	{6, 8, 11, 10, 4, 9, 11, 2, 3, 10, 7, 9, 4, 11, 7, 6, 1, 5, 9, 8, 10, 11, 7, 2, 10, 9, 5, 10, 9, 4, 8, 10, 3, 8, 7, 2, 1, 7, 9, 11, 6, 7, 11, 8, 5, 9, 8, 11, 3, 10},
	{10, 3, 5, 10, 1, 8, 9, 11, 7, 8, 9, 7, 11, 6, 10, 9, 11, 10, 2, 9, 8, 2, 11, 10, 5, 3, 8, 7, 10, 4, 11, 10, 8, 6, 7, 9, 11, 7, 6, 2, 7, 11, 8, 9, 3, 4, 5, 9, 4},
	{10, 4, 8, 10, 2, 7, 11, 4, 9, 8, 6, 7, 5, 9, 10, 3, 9, 10, 8, 11, 3, 7, 9, 8, 11, 7, 9, 11, 8, 3, 11, 10, 6, 8, 7, 2, 11, 9, 10, 5, 4, 6, 10, 5, 2, 7, 9, 11},
}

// *bonus reels calculations*
// reels lengths [47, 50, 50, 48, 47], total reshuffles 265080000
// symbols: rtp(sym) = 110.377147%
// free spins 135413910, q = 0.51084, sq = 1/(1-q) = 2.044328
// free games frequency: 1/2.2785
// RTP = sq*rtp(sym) = 2.0443*110.38 = 225.647075%
// *regular reels calculations*
// reels lengths [47, 50, 50, 48, 47], total reshuffles 265080000
// symbols: rtp(sym) = 55.188573%
// free spins 36792000, q = 0.1388, sq = 1/(1-q) = 1.161165
// free games frequency: 1/7.4603
// RTP = 55.189(sym) + 0.1388*225.65(fg) = 86.507448%
var Reels865 = slot.Reels5x{
	{6, 7, 11, 9, 8, 10, 7, 8, 3, 5, 11, 7, 9, 8, 3, 9, 2, 10, 5, 3, 10, 7, 9, 11, 4, 5, 8, 2, 9, 4, 11, 9, 10, 11, 6, 10, 4, 6, 7, 2, 11, 10, 7, 8, 11, 10, 8},
	{2, 11, 8, 3, 10, 7, 11, 9, 4, 8, 1, 7, 6, 9, 10, 5, 11, 9, 10, 1, 7, 8, 11, 10, 8, 9, 3, 5, 11, 8, 4, 2, 10, 9, 4, 3, 11, 6, 10, 9, 6, 7, 10, 5, 2, 7, 8, 11, 9, 7},
	{6, 8, 11, 10, 4, 9, 11, 2, 3, 10, 7, 9, 4, 11, 7, 6, 1, 5, 9, 8, 10, 11, 7, 2, 10, 9, 5, 10, 9, 4, 8, 10, 3, 8, 7, 2, 1, 7, 9, 11, 6, 7, 11, 8, 5, 9, 8, 11, 3, 10},
	{10, 7, 11, 2, 10, 8, 11, 9, 8, 7, 6, 9, 10, 5, 11, 3, 6, 9, 7, 8, 11, 7, 10, 3, 9, 5, 10, 11, 7, 2, 8, 4, 10, 9, 7, 8, 6, 5, 8, 11, 9, 10, 4, 11, 2, 4, 3, 1},
	{11, 2, 8, 10, 7, 3, 10, 9, 7, 10, 6, 5, 11, 8, 5, 9, 10, 6, 7, 11, 9, 4, 5, 11, 10, 8, 11, 3, 6, 8, 9, 7, 11, 4, 8, 7, 2, 10, 7, 9, 3, 2, 11, 4, 8, 9, 10},
}

// *bonus reels calculations*
// reels lengths [48, 49, 49, 48, 48], total reshuffles 265531392
// symbols: rtp(sym) = 111.773650%
// free spins 136225620, q = 0.51303, sq = 1/(1-q) = 2.053515
// free games frequency: 1/2.2698
// RTP = sq*rtp(sym) = 2.0535*111.77 = 229.528911%
// *regular reels calculations*
// reels lengths [48, 49, 49, 48, 48], total reshuffles 265531392
// symbols: rtp(sym) = 55.886825%
// free spins 37052100, q = 0.13954, sq = 1/(1-q) = 1.162168
// free games frequency: 1/7.4234
// RTP = 55.887(sym) + 0.13954*229.53(fg) = 87.915159%
var Reels879 = slot.Reels5x{
	{3, 7, 10, 8, 2, 11, 3, 6, 9, 11, 10, 8, 2, 7, 11, 9, 3, 5, 7, 10, 11, 4, 5, 9, 8, 10, 7, 2, 11, 9, 4, 8, 7, 6, 8, 10, 9, 11, 8, 10, 5, 6, 9, 7, 10, 11, 4, 9},
	{8, 9, 11, 5, 9, 10, 3, 7, 2, 10, 11, 4, 10, 3, 8, 9, 11, 8, 9, 7, 8, 6, 1, 5, 7, 10, 4, 11, 6, 9, 7, 10, 8, 7, 11, 10, 9, 11, 7, 5, 6, 2, 3, 4, 11, 10, 8, 1, 2},
	{9, 4, 3, 11, 7, 9, 11, 1, 8, 11, 7, 8, 11, 9, 10, 11, 4, 10, 8, 9, 10, 8, 6, 4, 10, 7, 6, 1, 2, 6, 5, 8, 3, 10, 8, 9, 2, 7, 11, 10, 5, 11, 7, 9, 5, 7, 10, 3, 2},
	{10, 7, 11, 2, 10, 8, 11, 9, 8, 7, 6, 9, 10, 5, 11, 3, 6, 9, 7, 8, 11, 7, 10, 3, 9, 5, 10, 11, 7, 2, 8, 4, 10, 9, 7, 8, 6, 5, 8, 11, 9, 10, 4, 11, 2, 4, 3, 1},
	{10, 4, 8, 10, 2, 7, 11, 4, 9, 8, 6, 7, 5, 9, 10, 3, 9, 10, 8, 11, 3, 7, 9, 8, 11, 7, 9, 11, 8, 3, 11, 10, 6, 8, 7, 2, 11, 9, 10, 5, 4, 6, 10, 5, 2, 7, 9, 11},
}

// *bonus reels calculations*
// reels lengths [47, 49, 49, 48, 47], total reshuffles 254582832
// symbols: rtp(sym) = 112.830688%
// free spins 130758480, q = 0.51362, sq = 1/(1-q) = 2.056000
// free games frequency: 1/2.2653
// RTP = sq*rtp(sym) = 2.056*112.83 = 231.979861%
// *regular reels calculations*
// reels lengths [47, 49, 49, 48, 47], total reshuffles 254582832
// symbols: rtp(sym) = 56.415344%
// free spins 36284400, q = 0.14252, sq = 1/(1-q) = 1.166215
// free games frequency: 1/7.2687
// RTP = 56.415(sym) + 0.14252*231.98(fg) = 89.478257%
var Reels894 = slot.Reels5x{
	{6, 7, 11, 9, 8, 10, 7, 8, 3, 5, 11, 7, 9, 8, 3, 9, 2, 10, 5, 3, 10, 7, 9, 11, 4, 5, 8, 2, 9, 4, 11, 9, 10, 11, 6, 10, 4, 6, 7, 2, 11, 10, 7, 8, 11, 10, 8},
	{8, 9, 11, 5, 9, 10, 3, 7, 2, 10, 11, 4, 10, 3, 8, 9, 11, 8, 9, 7, 8, 6, 1, 5, 7, 10, 4, 11, 6, 9, 7, 10, 8, 7, 11, 10, 9, 11, 7, 5, 6, 2, 3, 4, 11, 10, 8, 1, 2},
	{9, 4, 3, 11, 7, 9, 11, 1, 8, 11, 7, 8, 11, 9, 10, 11, 4, 10, 8, 9, 10, 8, 6, 4, 10, 7, 6, 1, 2, 6, 5, 8, 3, 10, 8, 9, 2, 7, 11, 10, 5, 11, 7, 9, 5, 7, 10, 3, 2},
	{10, 7, 11, 2, 10, 8, 11, 9, 8, 7, 6, 9, 10, 5, 11, 3, 6, 9, 7, 8, 11, 7, 10, 3, 9, 5, 10, 11, 7, 2, 8, 4, 10, 9, 7, 8, 6, 5, 8, 11, 9, 10, 4, 11, 2, 4, 3, 1},
	{11, 2, 8, 10, 7, 3, 10, 9, 7, 10, 6, 5, 11, 8, 5, 9, 10, 6, 7, 11, 9, 4, 5, 11, 10, 8, 11, 3, 6, 8, 9, 7, 11, 4, 8, 7, 2, 10, 7, 9, 3, 2, 11, 4, 8, 9, 10},
}

// *bonus reels calculations*
// reels lengths [46, 49, 49, 47, 46], total reshuffles 238784252
// symbols: rtp(sym) = 114.110473%
// free spins 121876740, q = 0.51041, sq = 1/(1-q) = 2.042506
// free games frequency: 1/2.2726
// RTP = sq*rtp(sym) = 2.0425*114.11 = 233.071285%
// *regular reels calculations*
// reels lengths [46, 49, 49, 47, 46], total reshuffles 238784252
// symbols: rtp(sym) = 57.055237%
// free spins 34802550, q = 0.14575, sq = 1/(1-q) = 1.170616
// free games frequency: 1/7.1141
// RTP = 57.055(sym) + 0.14575*233.07(fg) = 91.025128%
var Reels910 = slot.Reels5x{
	{10, 7, 3, 8, 11, 7, 3, 11, 7, 10, 11, 8, 9, 11, 6, 9, 2, 4, 8, 9, 11, 10, 7, 9, 8, 4, 5, 11, 10, 2, 4, 7, 5, 10, 8, 2, 6, 11, 8, 10, 3, 9, 6, 7, 5, 9},
	{8, 9, 11, 5, 9, 10, 3, 7, 2, 10, 11, 4, 10, 3, 8, 9, 11, 8, 9, 7, 8, 6, 1, 5, 7, 10, 4, 11, 6, 9, 7, 10, 8, 7, 11, 10, 9, 11, 7, 5, 6, 2, 3, 4, 11, 10, 8, 1, 2},
	{9, 4, 3, 11, 7, 9, 11, 1, 8, 11, 7, 8, 11, 9, 10, 11, 4, 10, 8, 9, 10, 8, 6, 4, 10, 7, 6, 1, 2, 6, 5, 8, 3, 10, 8, 9, 2, 7, 11, 10, 5, 11, 7, 9, 5, 7, 10, 3, 2},
	{10, 8, 3, 2, 7, 10, 9, 5, 6, 11, 8, 6, 9, 10, 8, 11, 6, 2, 7, 10, 5, 8, 3, 9, 11, 8, 4, 2, 10, 9, 5, 1, 8, 7, 10, 11, 9, 7, 11, 4, 7, 11, 9, 4, 3, 11, 7},
	{10, 8, 4, 7, 8, 11, 10, 7, 9, 8, 6, 3, 9, 8, 7, 11, 4, 10, 6, 8, 10, 11, 7, 9, 2, 5, 4, 9, 10, 11, 2, 3, 11, 10, 3, 9, 8, 5, 11, 9, 2, 7, 6, 11, 7, 5},
}

// *bonus reels calculations*
// reels lengths [46, 48, 49, 48, 47], total reshuffles 244081152
// symbols: rtp(sym) = 115.033643%
// free spins 125088480, q = 0.51249, sq = 1/(1-q) = 2.051228
// free games frequency: 1/2.2639
// RTP = sq*rtp(sym) = 2.0512*115.03 = 235.960278%
// *regular reels calculations*
// reels lengths [46, 48, 49, 48, 47], total reshuffles 244081152
// symbols: rtp(sym) = 57.516821%
// free spins 36284400, q = 0.14866, sq = 1/(1-q) = 1.174615
// free games frequency: 1/6.9689
// RTP = 57.517(sym) + 0.14866*235.96(fg) = 92.593996%
var Reels925 = slot.Reels5x{
	{10, 7, 3, 8, 11, 7, 3, 11, 7, 10, 11, 8, 9, 11, 6, 9, 2, 4, 8, 9, 11, 10, 7, 9, 8, 4, 5, 11, 10, 2, 4, 7, 5, 10, 8, 2, 6, 11, 8, 10, 3, 9, 6, 7, 5, 9},
	{8, 9, 5, 8, 4, 11, 6, 9, 7, 10, 8, 5, 6, 9, 5, 1, 11, 10, 8, 11, 10, 4, 7, 11, 8, 7, 11, 10, 4, 3, 7, 2, 10, 3, 9, 7, 6, 2, 9, 11, 3, 7, 9, 10, 1, 8, 11, 2},
	{9, 4, 3, 11, 7, 9, 11, 1, 8, 11, 7, 8, 11, 9, 10, 11, 4, 10, 8, 9, 10, 8, 6, 4, 10, 7, 6, 1, 2, 6, 5, 8, 3, 10, 8, 9, 2, 7, 11, 10, 5, 11, 7, 9, 5, 7, 10, 3, 2},
	{10, 7, 11, 2, 10, 8, 11, 9, 8, 7, 6, 9, 10, 5, 11, 3, 6, 9, 7, 8, 11, 7, 10, 3, 9, 5, 10, 11, 7, 2, 8, 4, 10, 9, 7, 8, 6, 5, 8, 11, 9, 10, 4, 11, 2, 4, 3, 1},
	{11, 2, 8, 10, 7, 3, 10, 9, 7, 10, 6, 5, 11, 8, 5, 9, 10, 6, 7, 11, 9, 4, 5, 11, 10, 8, 11, 3, 6, 8, 9, 7, 11, 4, 8, 7, 2, 10, 7, 9, 3, 2, 11, 4, 8, 9, 10},
}

// *bonus reels calculations*
// reels lengths [46, 48, 49, 47, 46], total reshuffles 233911104
// symbols: rtp(sym) = 115.554908%
// free spins 119762640, q = 0.512, sq = 1/(1-q) = 2.049183
// free games frequency: 1/2.2639
// RTP = sq*rtp(sym) = 2.0492*115.55 = 236.793164%
// *regular reels calculations*
// reels lengths [46, 48, 49, 47, 46], total reshuffles 233911104
// symbols: rtp(sym) = 57.777454%
// free spins 34802550, q = 0.14879, sq = 1/(1-q) = 1.174792
// free games frequency: 1/6.9689
// RTP = 57.777(sym) + 0.14879*236.79(fg) = 93.008812%
var Reels930 = slot.Reels5x{
	{10, 7, 3, 8, 11, 7, 3, 11, 7, 10, 11, 8, 9, 11, 6, 9, 2, 4, 8, 9, 11, 10, 7, 9, 8, 4, 5, 11, 10, 2, 4, 7, 5, 10, 8, 2, 6, 11, 8, 10, 3, 9, 6, 7, 5, 9},
	{8, 9, 5, 8, 4, 11, 6, 9, 7, 10, 8, 5, 6, 9, 5, 1, 11, 10, 8, 11, 10, 4, 7, 11, 8, 7, 11, 10, 4, 3, 7, 2, 10, 3, 9, 7, 6, 2, 9, 11, 3, 7, 9, 10, 1, 8, 11, 2},
	{9, 4, 3, 11, 7, 9, 11, 1, 8, 11, 7, 8, 11, 9, 10, 11, 4, 10, 8, 9, 10, 8, 6, 4, 10, 7, 6, 1, 2, 6, 5, 8, 3, 10, 8, 9, 2, 7, 11, 10, 5, 11, 7, 9, 5, 7, 10, 3, 2},
	{10, 8, 3, 2, 7, 10, 9, 5, 6, 11, 8, 6, 9, 10, 8, 11, 6, 2, 7, 10, 5, 8, 3, 9, 11, 8, 4, 2, 10, 9, 5, 1, 8, 7, 10, 11, 9, 7, 11, 4, 7, 11, 9, 4, 3, 11, 7},
	{10, 8, 4, 7, 8, 11, 10, 7, 9, 8, 6, 3, 9, 8, 7, 11, 4, 10, 6, 8, 10, 11, 7, 9, 2, 5, 4, 9, 10, 11, 2, 3, 11, 10, 3, 9, 8, 5, 11, 9, 2, 7, 6, 11, 7, 5},
}

// *bonus reels calculations*
// reels lengths [46, 48, 48, 48, 46], total reshuffles 234012672
// symbols: rtp(sym) = 116.409508%
// free spins 119920500, q = 0.51245, sq = 1/(1-q) = 2.051084
// free games frequency: 1/2.2624
// RTP = sq*rtp(sym) = 2.0511*116.41 = 238.765724%
// *regular reels calculations*
// reels lengths [46, 48, 48, 48, 46], total reshuffles 234012672
// symbols: rtp(sym) = 58.204754%
// free spins 35019900, q = 0.14965, sq = 1/(1-q) = 1.175986
// free games frequency: 1/6.9271
// RTP = 58.205(sym) + 0.14965*238.77(fg) = 93.935946%
var Reels939 = slot.Reels5x{
	{10, 7, 3, 8, 11, 7, 3, 11, 7, 10, 11, 8, 9, 11, 6, 9, 2, 4, 8, 9, 11, 10, 7, 9, 8, 4, 5, 11, 10, 2, 4, 7, 5, 10, 8, 2, 6, 11, 8, 10, 3, 9, 6, 7, 5, 9},
	{8, 9, 5, 8, 4, 11, 6, 9, 7, 10, 8, 5, 6, 9, 5, 1, 11, 10, 8, 11, 10, 4, 7, 11, 8, 7, 11, 10, 4, 3, 7, 2, 10, 3, 9, 7, 6, 2, 9, 11, 3, 7, 9, 10, 1, 8, 11, 2},
	{6, 11, 4, 3, 11, 6, 10, 8, 1, 2, 5, 11, 10, 3, 9, 2, 10, 8, 11, 9, 10, 11, 7, 10, 8, 5, 7, 9, 8, 4, 10, 11, 7, 9, 2, 11, 9, 7, 3, 8, 7, 6, 8, 9, 5, 7, 1, 4},
	{10, 7, 11, 2, 10, 8, 11, 9, 8, 7, 6, 9, 10, 5, 11, 3, 6, 9, 7, 8, 11, 7, 10, 3, 9, 5, 10, 11, 7, 2, 8, 4, 10, 9, 7, 8, 6, 5, 8, 11, 9, 10, 4, 11, 2, 4, 3, 1},
	{10, 8, 4, 7, 8, 11, 10, 7, 9, 8, 6, 3, 9, 8, 7, 11, 4, 10, 6, 8, 10, 11, 7, 9, 2, 5, 4, 9, 10, 11, 2, 3, 11, 10, 3, 9, 8, 5, 11, 9, 2, 7, 6, 11, 7, 5},
}

// *bonus reels calculations*
// reels lengths [46, 48, 48, 47, 46], total reshuffles 229137408
// symbols: rtp(sym) = 116.900258%
// free spins 117396990, q = 0.51234, sq = 1/(1-q) = 2.050622
// free games frequency: 1/2.2624
// RTP = sq*rtp(sym) = 2.0506*116.9 = 239.718291%
// *regular reels calculations*
// reels lengths [46, 48, 48, 47, 46], total reshuffles 229137408
// symbols: rtp(sym) = 58.450129%
// free spins 34316100, q = 0.14976, sq = 1/(1-q) = 1.176141
// free games frequency: 1/6.9271
// RTP = 58.45(sym) + 0.14976*239.72(fg) = 94.350844%
var Reels943 = slot.Reels5x{
	{10, 7, 3, 8, 11, 7, 3, 11, 7, 10, 11, 8, 9, 11, 6, 9, 2, 4, 8, 9, 11, 10, 7, 9, 8, 4, 5, 11, 10, 2, 4, 7, 5, 10, 8, 2, 6, 11, 8, 10, 3, 9, 6, 7, 5, 9},
	{8, 9, 5, 8, 4, 11, 6, 9, 7, 10, 8, 5, 6, 9, 5, 1, 11, 10, 8, 11, 10, 4, 7, 11, 8, 7, 11, 10, 4, 3, 7, 2, 10, 3, 9, 7, 6, 2, 9, 11, 3, 7, 9, 10, 1, 8, 11, 2},
	{6, 11, 4, 3, 11, 6, 10, 8, 1, 2, 5, 11, 10, 3, 9, 2, 10, 8, 11, 9, 10, 11, 7, 10, 8, 5, 7, 9, 8, 4, 10, 11, 7, 9, 2, 11, 9, 7, 3, 8, 7, 6, 8, 9, 5, 7, 1, 4},
	{10, 8, 3, 2, 7, 10, 9, 5, 6, 11, 8, 6, 9, 10, 8, 11, 6, 2, 7, 10, 5, 8, 3, 9, 11, 8, 4, 2, 10, 9, 5, 1, 8, 7, 10, 11, 9, 7, 11, 4, 7, 11, 9, 4, 3, 11, 7},
	{10, 8, 4, 7, 8, 11, 10, 7, 9, 8, 6, 3, 9, 8, 7, 11, 4, 10, 6, 8, 10, 11, 7, 9, 2, 5, 4, 9, 10, 11, 2, 3, 11, 10, 3, 9, 8, 5, 11, 9, 2, 7, 6, 11, 7, 5},
}

// *bonus reels calculations*
// reels lengths [46, 48, 48, 43, 46], total reshuffles 209636352
// symbols: rtp(sym) = 117.943037%
// free spins 107101755, q = 0.51089, sq = 1/(1-q) = 2.044543
// free games frequency: 1/2.2624
// RTP = sq*rtp(sym) = 2.0445*117.94 = 241.139564%
// *regular reels calculations*
// reels lengths [46, 48, 48, 43, 46], total reshuffles 209636352
// symbols: rtp(sym) = 58.971518%
// free spins 31500900, q = 0.15026, sq = 1/(1-q) = 1.176837
// free games frequency: 1/6.9271
// RTP = 58.972(sym) + 0.15026*241.14(fg) = 95.206233%
var Reels952 = slot.Reels5x{
	{10, 7, 3, 8, 11, 7, 3, 11, 7, 10, 11, 8, 9, 11, 6, 9, 2, 4, 8, 9, 11, 10, 7, 9, 8, 4, 5, 11, 10, 2, 4, 7, 5, 10, 8, 2, 6, 11, 8, 10, 3, 9, 6, 7, 5, 9},
	{8, 9, 5, 8, 4, 11, 6, 9, 7, 10, 8, 5, 6, 9, 5, 1, 11, 10, 8, 11, 10, 4, 7, 11, 8, 7, 11, 10, 4, 3, 7, 2, 10, 3, 9, 7, 6, 2, 9, 11, 3, 7, 9, 10, 1, 8, 11, 2},
	{6, 11, 4, 3, 11, 6, 10, 8, 1, 2, 5, 11, 10, 3, 9, 2, 10, 8, 11, 9, 10, 11, 7, 10, 8, 5, 7, 9, 8, 4, 10, 11, 7, 9, 2, 11, 9, 7, 3, 8, 7, 6, 8, 9, 5, 7, 1, 4},
	{11, 5, 9, 2, 3, 7, 8, 10, 4, 8, 10, 7, 9, 2, 10, 11, 4, 8, 11, 10, 8, 11, 3, 9, 7, 8, 11, 5, 6, 10, 7, 9, 4, 3, 6, 5, 9, 11, 6, 2, 10, 7, 1},
	{10, 8, 4, 7, 8, 11, 10, 7, 9, 8, 6, 3, 9, 8, 7, 11, 4, 10, 6, 8, 10, 11, 7, 9, 2, 5, 4, 9, 10, 11, 2, 3, 11, 10, 3, 9, 8, 5, 11, 9, 2, 7, 6, 11, 7, 5},
}

// *bonus reels calculations*
// reels lengths [45, 48, 48, 47, 45], total reshuffles 219283200
// symbols: rtp(sym) = 118.035280%
// free spins 111557250, q = 0.50874, sq = 1/(1-q) = 2.035565
// free games frequency: 1/2.2722
// RTP = sq*rtp(sym) = 2.0356*118.04 = 240.268515%
// *regular reels calculations*
// reels lengths [45, 48, 48, 47, 45], total reshuffles 219283200
// symbols: rtp(sym) = 59.017640%
// free spins 33574500, q = 0.15311, sq = 1/(1-q) = 1.180791
// free games frequency: 1/6.7765
// RTP = 59.018(sym) + 0.15311*240.27(fg) = 95.805206%
var Reels958 = slot.Reels5x{
	{9, 11, 8, 7, 4, 8, 9, 10, 11, 6, 2, 7, 5, 10, 8, 5, 10, 6, 7, 5, 4, 9, 11, 7, 3, 9, 7, 10, 11, 9, 3, 10, 4, 3, 8, 11, 2, 9, 6, 7, 8, 2, 11, 8, 10},
	{8, 9, 5, 8, 4, 11, 6, 9, 7, 10, 8, 5, 6, 9, 5, 1, 11, 10, 8, 11, 10, 4, 7, 11, 8, 7, 11, 10, 4, 3, 7, 2, 10, 3, 9, 7, 6, 2, 9, 11, 3, 7, 9, 10, 1, 8, 11, 2},
	{6, 11, 4, 3, 11, 6, 10, 8, 1, 2, 5, 11, 10, 3, 9, 2, 10, 8, 11, 9, 10, 11, 7, 10, 8, 5, 7, 9, 8, 4, 10, 11, 7, 9, 2, 11, 9, 7, 3, 8, 7, 6, 8, 9, 5, 7, 1, 4},
	{10, 8, 3, 2, 7, 10, 9, 5, 6, 11, 8, 6, 9, 10, 8, 11, 6, 2, 7, 10, 5, 8, 3, 9, 11, 8, 4, 2, 10, 9, 5, 1, 8, 7, 10, 11, 9, 7, 11, 4, 7, 11, 9, 4, 3, 11, 7},
	{10, 7, 11, 8, 7, 10, 5, 4, 11, 9, 4, 10, 3, 6, 8, 11, 10, 5, 11, 9, 8, 2, 7, 10, 9, 8, 6, 5, 2, 4, 9, 7, 3, 9, 7, 8, 6, 7, 2, 8, 11, 9, 10, 3, 11},
}

// *bonus reels calculations*
// reels lengths [46, 46, 48, 47, 46], total reshuffles 219590016
// symbols: rtp(sym) = 118.192915%
// free spins 112968180, q = 0.51445, sq = 1/(1-q) = 2.059522
// free games frequency: 1/2.2481
// RTP = sq*rtp(sym) = 2.0595*118.19 = 243.420907%
// *regular reels calculations*
// reels lengths [46, 46, 48, 47, 46], total reshuffles 219590016
// symbols: rtp(sym) = 59.096457%
// free spins 34316100, q = 0.15627, sq = 1/(1-q) = 1.185218
// free games frequency: 1/6.6384
// RTP = 59.096(sym) + 0.15627*243.42(fg) = 97.136694%
var Reels971 = slot.Reels5x{
	{10, 7, 3, 8, 11, 7, 3, 11, 7, 10, 11, 8, 9, 11, 6, 9, 2, 4, 8, 9, 11, 10, 7, 9, 8, 4, 5, 11, 10, 2, 4, 7, 5, 10, 8, 2, 6, 11, 8, 10, 3, 9, 6, 7, 5, 9},
	{1, 10, 11, 8, 2, 11, 10, 8, 9, 2, 8, 11, 7, 9, 3, 11, 5, 6, 9, 10, 3, 8, 6, 4, 8, 5, 7, 9, 3, 2, 7, 4, 10, 11, 7, 9, 4, 10, 6, 5, 1, 8, 7, 10, 9, 11},
	{6, 11, 4, 3, 11, 6, 10, 8, 1, 2, 5, 11, 10, 3, 9, 2, 10, 8, 11, 9, 10, 11, 7, 10, 8, 5, 7, 9, 8, 4, 10, 11, 7, 9, 2, 11, 9, 7, 3, 8, 7, 6, 8, 9, 5, 7, 1, 4},
	{10, 8, 3, 2, 7, 10, 9, 5, 6, 11, 8, 6, 9, 10, 8, 11, 6, 2, 7, 10, 5, 8, 3, 9, 11, 8, 4, 2, 10, 9, 5, 1, 8, 7, 10, 11, 9, 7, 11, 4, 7, 11, 9, 4, 3, 11, 7},
	{10, 8, 4, 7, 8, 11, 10, 7, 9, 8, 6, 3, 9, 8, 7, 11, 4, 10, 6, 8, 10, 11, 7, 9, 2, 5, 4, 9, 10, 11, 2, 3, 11, 10, 3, 9, 8, 5, 11, 9, 2, 7, 6, 11, 7, 5},
}

// *bonus reels calculations*
// reels lengths [46, 47, 47, 46, 45], total reshuffles 210340980
// symbols: rtp(sym) = 120.361995%
// free spins 107147610, q = 0.5094, sq = 1/(1-q) = 2.038319
// free games frequency: 1/2.2664
// RTP = sq*rtp(sym) = 2.0383*120.36 = 245.336110%
// *regular reels calculations*
// reels lengths [46, 47, 47, 46, 45], total reshuffles 210340980
// symbols: rtp(sym) = 60.180998%
// free spins 32420250, q = 0.15413, sq = 1/(1-q) = 1.182217
// free games frequency: 1/6.7406
// RTP = 60.181(sym) + 0.15413*245.34(fg) = 97.995113%
var Reels979 = slot.Reels5x{
	{10, 7, 3, 8, 11, 7, 3, 11, 7, 10, 11, 8, 9, 11, 6, 9, 2, 4, 8, 9, 11, 10, 7, 9, 8, 4, 5, 11, 10, 2, 4, 7, 5, 10, 8, 2, 6, 11, 8, 10, 3, 9, 6, 7, 5, 9},
	{6, 7, 2, 9, 10, 11, 7, 9, 2, 8, 10, 7, 4, 10, 11, 7, 10, 8, 5, 1, 2, 9, 4, 11, 10, 7, 9, 8, 11, 3, 6, 8, 3, 7, 9, 6, 5, 1, 11, 4, 8, 5, 11, 9, 3, 8, 10},
	{10, 7, 8, 5, 7, 9, 3, 10, 6, 11, 4, 2, 5, 11, 9, 8, 11, 9, 10, 2, 3, 7, 8, 5, 7, 9, 2, 8, 10, 7, 4, 6, 10, 11, 3, 7, 6, 1, 11, 8, 9, 4, 10, 8, 11, 1, 9},
	{5, 8, 2, 7, 10, 2, 9, 4, 6, 8, 7, 9, 10, 2, 8, 7, 11, 5, 10, 9, 8, 7, 11, 8, 3, 7, 4, 3, 7, 11, 9, 10, 11, 8, 9, 10, 3, 6, 11, 4, 10, 9, 5, 6, 11, 1},
	{10, 7, 11, 8, 7, 10, 5, 4, 11, 9, 4, 10, 3, 6, 8, 11, 10, 5, 11, 9, 8, 2, 7, 10, 9, 8, 6, 5, 2, 4, 9, 7, 3, 9, 7, 8, 6, 7, 2, 8, 11, 9, 10, 3, 11},
}

// *bonus reels calculations*
// reels lengths [46, 46, 46, 45, 46], total reshuffles 201485520
// symbols: rtp(sym) = 119.881468%
// free spins 103372260, q = 0.51305, sq = 1/(1-q) = 2.053601
// free games frequency: 1/2.25
// RTP = sq*rtp(sym) = 2.0536*119.88 = 246.188741%
// *regular reels calculations*
// reels lengths [46, 46, 46, 45, 46], total reshuffles 201485520
// symbols: rtp(sym) = 59.940734%
// free spins 31977000, q = 0.15871, sq = 1/(1-q) = 1.188645
// free games frequency: 1/6.5546
// RTP = 59.941(sym) + 0.15871*246.19(fg) = 99.012412%
var Reels990 = slot.Reels5x{
	{10, 7, 3, 8, 11, 7, 3, 11, 7, 10, 11, 8, 9, 11, 6, 9, 2, 4, 8, 9, 11, 10, 7, 9, 8, 4, 5, 11, 10, 2, 4, 7, 5, 10, 8, 2, 6, 11, 8, 10, 3, 9, 6, 7, 5, 9},
	{1, 10, 11, 8, 2, 11, 10, 8, 9, 2, 8, 11, 7, 9, 3, 11, 5, 6, 9, 10, 3, 8, 6, 4, 8, 5, 7, 9, 3, 2, 7, 4, 10, 11, 7, 9, 4, 10, 6, 5, 1, 8, 7, 10, 9, 11},
	{9, 4, 10, 8, 9, 1, 11, 9, 8, 1, 2, 7, 11, 10, 5, 8, 9, 2, 6, 4, 8, 11, 7, 10, 5, 6, 11, 10, 7, 2, 8, 4, 7, 9, 3, 10, 6, 9, 10, 11, 3, 8, 11, 7, 3, 5},
	{9, 3, 8, 10, 11, 8, 9, 10, 5, 9, 8, 10, 6, 9, 3, 11, 7, 2, 4, 9, 8, 7, 6, 11, 7, 1, 2, 8, 6, 3, 8, 5, 10, 11, 2, 7, 4, 10, 11, 7, 5, 9, 10, 11, 4},
	{10, 8, 4, 7, 8, 11, 10, 7, 9, 8, 6, 3, 9, 8, 7, 11, 4, 10, 6, 8, 10, 11, 7, 9, 2, 5, 4, 9, 10, 11, 2, 3, 11, 10, 3, 9, 8, 5, 11, 9, 2, 7, 6, 11, 7, 5},
}

// *bonus reels calculations*
// reels lengths [45, 47, 47, 46, 45], total reshuffles 205768350
// symbols: rtp(sym) = 121.795990%
// free spins 104737050, q = 0.509, sq = 1/(1-q) = 2.036679
// free games frequency: 1/2.2656
// RTP = sq*rtp(sym) = 2.0367*121.8 = 248.059364%
// *regular reels calculations*
// reels lengths [45, 47, 47, 46, 45], total reshuffles 205768350
// symbols: rtp(sym) = 60.897995%
// free spins 32420250, q = 0.15756, sq = 1/(1-q) = 1.187024
// free games frequency: 1/6.594
// RTP = 60.898(sym) + 0.15756*248.06(fg) = 99.981492%
var Reels999 = slot.Reels5x{
	{9, 11, 8, 7, 4, 8, 9, 10, 11, 6, 2, 7, 5, 10, 8, 5, 10, 6, 7, 5, 4, 9, 11, 7, 3, 9, 7, 10, 11, 9, 3, 10, 4, 3, 8, 11, 2, 9, 6, 7, 8, 2, 11, 8, 10},
	{6, 7, 2, 9, 10, 11, 7, 9, 2, 8, 10, 7, 4, 10, 11, 7, 10, 8, 5, 1, 2, 9, 4, 11, 10, 7, 9, 8, 11, 3, 6, 8, 3, 7, 9, 6, 5, 1, 11, 4, 8, 5, 11, 9, 3, 8, 10},
	{10, 7, 8, 5, 7, 9, 3, 10, 6, 11, 4, 2, 5, 11, 9, 8, 11, 9, 10, 2, 3, 7, 8, 5, 7, 9, 2, 8, 10, 7, 4, 6, 10, 11, 3, 7, 6, 1, 11, 8, 9, 4, 10, 8, 11, 1, 9},
	{5, 8, 2, 7, 10, 2, 9, 4, 6, 8, 7, 9, 10, 2, 8, 7, 11, 5, 10, 9, 8, 7, 11, 8, 3, 7, 4, 3, 7, 11, 9, 10, 11, 8, 9, 10, 3, 6, 11, 4, 10, 9, 5, 6, 11, 1},
	{10, 7, 11, 8, 7, 10, 5, 4, 11, 9, 4, 10, 3, 6, 8, 11, 10, 5, 11, 9, 8, 2, 7, 10, 9, 8, 6, 5, 2, 4, 9, 7, 3, 9, 7, 8, 6, 7, 2, 8, 11, 9, 10, 3, 11},
}

// *bonus reels calculations*
// reels lengths [44, 46, 46, 45, 44], total reshuffles 184345920
// symbols: rtp(sym) = 120.707049%
// free spins 94791510, q = 0.5142, sq = 1/(1-q) = 2.058480
// free games frequency: 1/2.24
// RTP = sq*rtp(sym) = 2.0585*120.71 = 248.472990%
// *regular reels calculations*
// reels lengths [44, 46, 46, 45, 44], total reshuffles 184345920
// symbols: rtp(sym) = 60.353525%
// free spins 30595500, q = 0.16597, sq = 1/(1-q) = 1.198995
// free games frequency: 1/6.2696
// RTP = 60.354(sym) + 0.16597*248.47(fg) = 101.592058%
var Reels101 = slot.Reels5x{
	{10, 4, 9, 2, 8, 3, 5, 2, 9, 7, 11, 8, 7, 6, 8, 3, 10, 4, 7, 9, 10, 11, 3, 10, 4, 11, 10, 2, 8, 11, 6, 5, 9, 6, 11, 7, 5, 10, 8, 9, 11, 8, 9, 7},
	{1, 10, 11, 8, 2, 11, 10, 8, 9, 2, 8, 11, 7, 9, 3, 11, 5, 6, 9, 10, 3, 8, 6, 4, 8, 5, 7, 9, 3, 2, 7, 4, 10, 11, 7, 9, 4, 10, 6, 5, 1, 8, 7, 10, 9, 11},
	{9, 4, 10, 8, 9, 1, 11, 9, 8, 1, 2, 7, 11, 10, 5, 8, 9, 2, 6, 4, 8, 11, 7, 10, 5, 6, 11, 10, 7, 2, 8, 4, 7, 9, 3, 10, 6, 9, 10, 11, 3, 8, 11, 7, 3, 5},
	{9, 3, 8, 10, 11, 8, 9, 10, 5, 9, 8, 10, 6, 9, 3, 11, 7, 2, 4, 9, 8, 7, 6, 11, 7, 1, 2, 8, 6, 3, 8, 5, 10, 11, 2, 7, 4, 10, 11, 7, 5, 9, 10, 11, 4},
	{2, 11, 5, 3, 8, 10, 9, 7, 4, 11, 2, 9, 8, 10, 11, 5, 3, 10, 4, 8, 6, 7, 9, 5, 11, 6, 9, 7, 10, 9, 8, 7, 11, 8, 10, 6, 2, 8, 9, 4, 7, 3, 11, 10},
}

// *bonus reels calculations*
// reels lengths [43, 45, 45, 44, 43], total reshuffles 164745900
// symbols: rtp(sym) = 123.288143%
// free spins 85376820, q = 0.51823, sq = 1/(1-q) = 2.075694
// free games frequency: 1/2.2182
// RTP = sq*rtp(sym) = 2.0757*123.29 = 255.908422%
// *regular reels calculations*
// reels lengths [43, 45, 45, 44, 43], total reshuffles 164745900
// symbols: rtp(sym) = 61.644071%
// free spins 28840500, q = 0.17506, sq = 1/(1-q) = 1.212210
// free games frequency: 1/5.9538
// RTP = 61.644(sym) + 0.17506*255.91(fg) = 106.443528%
var Reels106 = slot.Reels5x{
	{8, 10, 6, 11, 9, 7, 11, 3, 4, 11, 9, 5, 10, 8, 7, 10, 11, 8, 2, 5, 8, 3, 6, 7, 4, 2, 9, 6, 10, 8, 4, 9, 7, 10, 9, 7, 11, 2, 5, 9, 11, 10, 3},
	{2, 4, 9, 5, 1, 8, 10, 11, 4, 6, 7, 10, 9, 7, 3, 6, 8, 9, 10, 8, 5, 10, 9, 2, 11, 5, 9, 3, 11, 6, 8, 7, 1, 11, 9, 10, 3, 11, 4, 2, 7, 8, 10, 7, 11},
	{9, 4, 10, 9, 8, 10, 11, 7, 6, 10, 9, 1, 11, 9, 3, 1, 5, 4, 10, 3, 11, 8, 9, 2, 7, 5, 11, 7, 5, 11, 6, 7, 8, 11, 4, 10, 2, 9, 8, 7, 6, 3, 2, 10, 8},
	{8, 1, 10, 7, 11, 2, 9, 3, 2, 5, 10, 4, 11, 9, 6, 7, 8, 10, 9, 4, 8, 5, 10, 8, 7, 11, 9, 6, 7, 8, 9, 10, 3, 6, 9, 11, 3, 5, 7, 11, 2, 4, 11, 10},
	{4, 2, 5, 10, 11, 9, 5, 8, 11, 10, 6, 9, 11, 6, 8, 9, 7, 3, 2, 7, 4, 8, 7, 11, 3, 10, 8, 2, 9, 11, 4, 7, 8, 10, 3, 5, 9, 10, 11, 9, 6, 7, 10},
}

// *bonus reels calculations*
// reels lengths [42, 44, 44, 43, 42], total reshuffles 146849472
// symbols: rtp(sym) = 128.412011%
// free spins 76479120, q = 0.5208, sq = 1/(1-q) = 2.086809
// free games frequency: 1/2.2009
// RTP = sq*rtp(sym) = 2.0868*128.41 = 267.971318%
// *regular reels calculations*
// reels lengths [42, 44, 44, 43, 42], total reshuffles 146849472
// symbols: rtp(sym) = 64.206005%
// free spins 27153900, q = 0.18491, sq = 1/(1-q) = 1.226858
// free games frequency: 1/5.6467
// RTP = 64.206(sym) + 0.18491*267.97(fg) = 113.756516%
var Reels113 = slot.Reels5x{
	{5, 10, 7, 3, 11, 10, 7, 11, 2, 10, 8, 3, 6, 10, 7, 3, 9, 10, 4, 8, 7, 5, 9, 6, 11, 7, 2, 4, 8, 10, 9, 11, 6, 9, 11, 2, 8, 5, 4, 9, 11, 8},
	{7, 5, 8, 4, 10, 11, 6, 9, 3, 10, 11, 7, 5, 11, 2, 9, 4, 8, 9, 10, 11, 7, 6, 11, 7, 8, 2, 1, 4, 9, 3, 8, 1, 10, 5, 7, 10, 3, 9, 8, 2, 6, 11, 10},
	{7, 11, 3, 10, 9, 6, 5, 2, 4, 11, 5, 7, 11, 10, 5, 9, 11, 8, 10, 3, 11, 7, 8, 2, 10, 4, 9, 11, 8, 7, 1, 6, 4, 9, 7, 10, 3, 6, 8, 9, 10, 1, 2, 8},
	{11, 5, 9, 2, 3, 7, 8, 10, 4, 8, 10, 7, 9, 2, 10, 11, 4, 8, 11, 10, 8, 11, 3, 9, 7, 8, 11, 5, 6, 10, 7, 9, 4, 3, 6, 5, 9, 11, 6, 2, 10, 7, 1},
	{7, 10, 11, 7, 5, 10, 9, 8, 6, 11, 2, 5, 9, 8, 10, 7, 11, 8, 6, 4, 3, 11, 6, 7, 2, 10, 5, 2, 10, 4, 9, 10, 4, 7, 8, 11, 3, 8, 9, 11, 3, 9},
}

// Map with available reels.
var ReelsMap = map[float64]*slot.Reels5x{
	82.953703:  &Reels829,
	84.315531:  &Reels843,
	85.303922:  &Reels853,
	86.507448:  &Reels865,
	87.915159:  &Reels879,
	89.478257:  &Reels894,
	91.025128:  &Reels910,
	92.593996:  &Reels925,
	93.008812:  &Reels930,
	93.935946:  &Reels939,
	94.350844:  &Reels943,
	95.206233:  &Reels952,
	95.805206:  &Reels958,
	97.136694:  &Reels971,
	97.995113:  &Reels979,
	99.012412:  &Reels990,
	99.981492:  &Reels999,
	101.592058: &Reels101,
	106.443528: &Reels106,
	113.756516: &Reels113,
}