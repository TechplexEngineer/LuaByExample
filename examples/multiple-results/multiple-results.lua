-- Lua has built-in support for _multiple results_.


-- This function returns two number results
local function values()
	return 3, 7
end

-- Here we use the 2 different return values from the
-- call with _multiple assignment_.
a, b = values()
print(a)
print(b)

-- If you only want a subset of the returned values,
-- use the blank identifier `_`.
_, c = values()
print(c)

-- The [Multiple Results](https://www.lua.org/pil/5.1.html) page in the Programming in Lua book has
-- more examples that cover the many edge cases
