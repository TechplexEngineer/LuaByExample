

-- Can be used much like a `for` loop
local i = 1
while i < 5 do
    print(i)
    i = i + 1
end

-- Common use of a `while` loop is an infinite loop.
-- `break` can be used to end the loop early
while true do
   print("This loop will run forever.")
   break
end