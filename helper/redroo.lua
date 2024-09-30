local path = arg[0]:match("(.*[/\\])")
dofile(path.."reelgen.lua")

local symset = {
	2, --  1 wild
	2, --  2 scatter
	5, --  3 redroo
	5, --  4 girl
	5, --  5 boy
	5, --  6 dog
	5, --  7 parrot
	6, --  8 ace
	6, --  9 king
	6, -- 10 queen
	6, -- 11 jack
	6, -- 12 ten
	6, -- 13 nine
}

local neighbours = {
	--1, 2, 3, 4, 5, 6, 7, 8, 9,10,11,12,13,
	{ 3, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2,}, --  1
	{ 3, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,}, --  2
	{ 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,}, --  3
	{ 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,}, --  4
	{ 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,}, --  5
	{ 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,}, --  6
	{ 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,}, --  7
	{ 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,}, --  8
	{ 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,}, --  9
	{ 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,}, -- 10
	{ 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,}, -- 11
	{ 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,}, -- 12
	{ 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,}, -- 13
}

math.randomseed(os.time())
local reel = MakeReel(symset)
print("reel length: " .. #reel)
ShuffleReel(reel)
local iter = CorrectReel(reel, neighbours)
RrintReel(reel, iter)
