-- _Functions_ are central in Lua. We'll learn about
-- functions with a few different examples.


-- Here's a local function that takes two `number`s and returns
-- their sum. Note: functions are global by default.
local function plus (a, b)
	-- lua requires explicit returns, i.e. it won't
	-- automatically return the value of the last
	-- expression.
	return a + b
end

-- Call a function just as you'd expect, with `name(args)`.
print("1+2 =" .. plus(1,2))

-- Because functions are first-class values in Lua, we can store them
-- not only in global variables, but also in local variables and
-- in table fields. As we will see later, the use of functions in
-- table fields is a key ingredient for some advanced uses of Lua,
-- such as packages and object-oriented programming.
local subtract = function (a, b)
    return a - b
end

-- And can be called the same way
print("2-1 =" .. subtract(2,1))

-- Functions without the `local` keyword become available in
-- the global namespace. This distinction becomes important in
-- larger Lua programs where where lua files are required by other lua files.