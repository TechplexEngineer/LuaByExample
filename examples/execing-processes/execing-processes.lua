-- To interact with other programs on the system and perform actions
-- that are not directly supported by lua, one can delecate the action
-- to the operating system shell with `os.execute()` or `io.popen()`


-- `os.execute()` executes a command using the operating system shell,
-- and returns if the command was successful or not.
-- If you need the command's output see io.popen
local success = os.execute('mkdir -p /tmp/example')
if success then
	print('successfully created directory')
end


-- `io.popen()` Runs a command, and returns a handle which can be used to read the command's output.
-- Success or failure of the command is reported as the first return value when
-- closing the handle.
local handle = io.popen("date --date='@2147483647'")
if handle == nil then
	return
end
local stdout = handle:read()
success = handle:close()
if success then
	print('output is:', stdout)
else
	print('error when executing command')
end

