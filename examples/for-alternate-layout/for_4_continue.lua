-- Lua does not have a `continue` construct. A workaround is to use goto.
-- prints odd numbers in [|1,10|]
for i=1, 5 do
  if i % 2 == 0 then goto continue end
  print('odd: ', i)
  ::continue::
end


