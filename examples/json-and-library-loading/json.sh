# We can override the LUA_PATH variable and ask
# all require calls to search the locations listed
# in order from left to right. If LUA_PATH is unset
# the default is typically `?;?.lua`
$ LUA_PATH="?;?.lua;lib/?.lua" lua json.lua

