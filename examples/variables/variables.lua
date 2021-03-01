-- In Lua, _variables_ are declared on first use as global.

-- `first` is a global variable.
first = "initial"
print(first)

-- You can declare multiple variables at once.
b, c = 1, 2
print(b, c)

-- Often variables should be declared local to avoid polluting the global scope.
local d = true
print(d)

-- Variables have no predefined types; any variable may contain values of any type.
local value = "starts as a string"
print(value)
-- - Assigning a value of a another type
value = 5
print(value)
-- - To unset the value, assign nil
value = nil
print(value)