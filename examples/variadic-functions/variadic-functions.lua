-- [_Variadic functions_](http://en.wikipedia.org/wiki/Variadic_function)
-- can be called with any number of trailing arguments.
-- For example, `print` is a common variadic function.


-- Here's a function that will take an arbitrary quantity
-- of `number`s as arguments.
local function sum(...)
	print("items: ", ...)
	total = 0
	for idx, num in ipairs({...}) do
	    total = total + num
	end
	print("total: ",total)
end



-- Variadic functions can be called in the usual way
-- with individual arguments.
sum(1, 2)
sum(1, 2, 3)

-- If you already have multiple args in a slice,
-- apply them to a variadic function using
-- `table.unpack()`.
nums = {1, 2, 3, 4}
sum(table.unpack(nums))

