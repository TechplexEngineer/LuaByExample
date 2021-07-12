-- [_Command-line arguments_](http://en.wikipedia.org/wiki/Command-line_interface#Arguments)
-- are a common way to parameterize execution of programs.
-- For example, `$ mkdir /tmp/fred` on the command line invokes the `mkdir` program with a single argument `/tmp/fred`

-- Function to make a string representation of a table
-- https://stackoverflow.com/a/27028488/429544
local function dump(o)
   if type(o) == 'table' then
      local s = '{ '
      for k,v in pairs(o) do
         if type(k) ~= 'number' then k = '"'..k..'"' end
         s = s .. '['..k..'] = ' .. dump(v) .. ','
      end
      return s .. '} '
   else
      return tostring(o)
   end
end

local args = {...}

print(dump(args))


