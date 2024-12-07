local path = arg[0]:match("(.*[/\\])")
dofile(path.."lib/reelgen.lua")

local symset = {
	3, --  1 wild
	1, --  2 scatter (1, 3, 4, 5 reel)
	2, --  3 warrior
	3, --  4 helmet
	3, --  5 shield
	3, --  6 axe
	3, --  7 mug
	3, --  8 ace
	3, --  9 king
	3, -- 10 queen
	3, -- 11 jack
	3, -- 12 ten
	4, -- 13 nine
}

local neighbours = {
	--1, 2, 3, 4, 5, 6, 7, 8, 9,10,11,12,13,
	{ 2, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,}, --  1 wild
	{ 2, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,}, --  2 scatter
	{ 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,}, --  3 warrior
	{ 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0,}, --  4 helmet
	{ 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 0,}, --  5 shield
	{ 0, 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0,}, --  6 axe
	{ 0, 0, 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0,}, --  7 mug
	{ 0, 0, 0, 0, 0, 0, 0, 2, 0, 0, 0, 0, 0,}, --  8 ace
	{ 0, 0, 0, 0, 0, 0, 0, 0, 2, 0, 0, 0, 0,}, --  9 king
	{ 0, 0, 0, 0, 0, 0, 0, 0, 0, 2, 0, 0, 0,}, -- 10 queen
	{ 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2, 0, 0,}, -- 11 jack
	{ 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2, 0,}, -- 12 ten
	{ 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2,}, -- 13 nine
}

math.randomseed(os.time())
local reel, iter = makereel(symset, neighbours)
printreel(reel, iter)