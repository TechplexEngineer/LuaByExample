-- `for` is one of Lua's looping constructs. Here are
-- some types of `for` loops.


-- A three expression loop is called a _numeric for_ loop
-- i=1 is the initial condition, 5 is the end condition, 1 is the step increment
-- the third expression is optional, when omitted 1 is used as the step.
for i=1, 5, 1 do
    print('first: ', i)
end
-- Note the variable declared in the loop is local to the body of the loop
-- despite the lack of the local keyword.
print(i) -- will print nil

-- `break` can be used to end the loop early
for i=1, 10, 1 do
    print('brk: ', i)
    if i == 5 then
        break
    end
end

-- A _generic for_ loop supports traversing elements returned from an iterator
-- print all values of array `arr`
arr = {'a', 'b', 'c', 'd'}
for i,v in ipairs(arr) do
    print('arr: i, v', i, v)
end

-- Lua does not have a `continue` construct. A workaround is to use goto.
-- prints odd numbers in [1,5] inclusive
for i=1, 5 do
  if i % 2 == 0 then goto continue end
  print('odd: ', i)
  ::continue::
end

