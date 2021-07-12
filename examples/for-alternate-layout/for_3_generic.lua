-- A _generic for_ loop supports traversing elements returned from an iterator
-- print all values of array `arr`
arr = {'a', 'b', 'c', 'd'}
for i,v in ipairs(arr) do
    print('arr: i,v', i,v)
end


