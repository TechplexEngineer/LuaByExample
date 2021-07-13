-- Lua does not have built in JSON support. However there are many
-- open source libraries available from the community. The Lua Wiki
-- has a [listing of json libraries](http://lua-users.org/wiki/JsonModules)

-- The examples here use the [jf-JSON library](http://regex.info/blog/lua/json)

-- The require function searches LUA_PATH and LUA_CPATH
-- For more about require see the Section [8.1 of the PIL guide](https://www.lua.org/pil/8.1.html)
local JSON = require 'jfjson'
local inspect = require 'inspect'

local lua_value = JSON:decode('[ "Larry", "Curly", "Moe" ]')
print('json array -> lua', inspect(lua_value))


local raw_json_text = JSON:encode({ "a", "b", "c"})
print('lua -> json array: ', raw_json_text)

local raw_json_map = JSON:encode({
    ['key'] = 'value'
})
print('lua -> json map: ', raw_json_map)
