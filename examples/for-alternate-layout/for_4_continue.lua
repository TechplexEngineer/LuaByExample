-- Lua does not have a `continue` construct. A workaround is to use goto.
-- This code prints odd numbers in the range [1,5] inclusive.
for i=1, 5 do
  if i % 2 == 0 then goto continue end
  print(i)
  ::continue::
end


