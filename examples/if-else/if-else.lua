-- Branching with `if` and `else` in Lua:

-- Here's a basic example.
-- Note that you don't need parentheses around conditions
if 7 % 2 == 0 then
    print("7 is even")
else
    print("7 is odd")
end

-- You can have an `if` statement without an else.
if 8 % 4 == 0 then
    print("8 is divisible by 4")
end

-- And an `elseif` example
local num = 9
if num < 0 then
    print(num, "is negative")
elseif num < 10 then
    print(num, "has 1 digit")
else
    print(num, "has multiple digits")
end
