local path = arg[0]:match("(.*[/\\])")
dofile(path.."reelgen.lua")

local symset = {
	1, --  1 wild
	2, --  2 dolphin
	2, --  3 turtle
	4, --  4 fish
	4, --  5 seahorse
	4, --  6 starfish
	4, --  7 ace
	4, --  8 king
	4, --  9 queen
	4, -- 10 jack
	4, -- 11 ten
	4, -- 12 nine
	1, -- 13 scatter
}

local neighbours = {
	--1, 2, 3, 4, 5, 6, 7, 8, 9,10,11,12,13,
	{ 2, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 2,}, --  1
	{ 1, 2, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0,}, --  2
	{ 1, 1, 2, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0,}, --  3
	{ 1, 1, 1, 2, 1, 1, 0, 0, 0, 0, 0, 0, 0,}, --  4
	{ 1, 1, 1, 1, 2, 1, 0, 0, 0, 0, 0, 0, 0,}, --  5
	{ 1, 1, 1, 1, 1, 2, 0, 0, 0, 0, 0, 0, 0,}, --  6
	{ 0, 0, 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0,}, --  7
	{ 0, 0, 0, 0, 0, 0, 0, 2, 0, 0, 0, 0, 0,}, --  8
	{ 0, 0, 0, 0, 0, 0, 0, 0, 2, 0, 0, 0, 0,}, --  9
	{ 0, 0, 0, 0, 0, 0, 0, 0, 0, 2, 0, 0, 0,}, -- 10
	{ 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2, 0, 0,}, -- 11
	{ 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2, 0,}, -- 12
	{ 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2,}, -- 13
}

math.randomseed(os.time())
local reel = MakeReel(symset)
print("reel length: " .. #reel)
ShuffleReel(reel)
local iter = CorrectReel(reel, neighbours)
if iter > 1 then
	if iter >= 1000 then
		print"too many neighbours shuffle iterations"
	else
		print(iter.." iterations")
	end
end
RrintReel(reel)
