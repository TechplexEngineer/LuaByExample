-- Lua has eight basic types: nil, boolean,
-- number, string, userdata, function, thread, and table.
-- Here are a few examples.

-- - Strings, which can be concatenated together with `..`
-- and can be enclosed in double quotes `"` or single `'`
print("lua" .. 'lang')

-- - Numbers and math
print("1+1 =", 1+1)
print("7.0/3.0 =", 7.0/3.0)

-- - Booleans, with boolean operators
print(true and false)
print(true or false)
print(not true)

-- - Parenthesis can be used to control order of operations
print(true or (true and false))

-- Note that nil is false
if nil then
    print('nil is not true')
else
    print('nil is false')
end
