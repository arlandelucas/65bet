local path = arg[0]:match("(.*[/\\])")
dofile(path.."reelgen.lua")

local symset = {
	1, --  1 wild
	1, --  2 scatter
	1, --  3 sabers
	3, --  4 map
	4, --  5 anchor
	5, --  6 ace
	5, --  7 king
	5, --  8 queen
	5, --  9 jack
	6, -- 10 ten
	6, -- 11 nine
}

local neighbours = {
	--1, 2, 3, 4, 5, 6, 7, 8, 9,10,11,
	{ 2, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0,}, --  1
	{ 2, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0,}, --  2
	{ 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 0,}, --  3
	{ 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0,}, --  4
	{ 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0,}, --  5
	{ 0, 0, 0, 0, 0, 2, 0, 0, 0, 0, 0,}, --  6
	{ 0, 0, 0, 0, 0, 0, 2, 0, 0, 0, 0,}, --  7
	{ 0, 0, 0, 0, 0, 0, 0, 2, 0, 0, 0,}, --  8
	{ 0, 0, 0, 0, 0, 0, 0, 0, 2, 0, 0,}, --  9
	{ 0, 0, 0, 0, 0, 0, 0, 0, 0, 2, 0,}, -- 10
	{ 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2,}, -- 11
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
