-- `for` is one of Lua's looping constructs. Here are
-- some basic types of `for` loops.


-- A three expression loop is called a _numeric for_ loop
-- i=1 is the initial condition, 10 is the end condition, 1 is the step increment
-- the third expression is optional, when omitted 1 is used as the step.
for i=1, 10, 1 do
    print(i)
end
-- Note the variable declared in the loop is local to the body of the loop
-- despite the lack of the local keyword.
print(i) -- will print nil

-- `break` can be used to end the loop early
for i=1, 10, 1 do
    print(i)
    if i == 5 then
        break
    end
end

-- A _generic for_ loop supports traversing elements returned from an iterator
-- print all values of array `arr`
for i,v in ipairs(arr) do
    print(v)
end

-- Lua does not have a `continue` construct. A workaround is to use goto.
-- prints odd numbers in [|1,10|]
for i=1,10 do
  if i % 2 == 0 then goto continue end
  print(i)
  ::continue::
end


