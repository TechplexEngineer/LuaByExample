-- Lua has eight basic types: nil, boolean,
-- number, string, userdata, function, thread, and table

-- Strings, which can be added together with `..`
print("lua" .. "lang")

-- Numbers and math
print("1+1 =", 1+1)
print("7.0/3.0 =", 7.0/3.0)

-- Booleans, with boolean operators
print(true and false)
print(true or false)
print(not true)

-- Parenthesis can be used to control order of operations
print(true or (true and false))

-- Note that nil is false
if nil then
    print('true')
else
    print('false')
end
