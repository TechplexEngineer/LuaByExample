-- [Environment variables](http://en.wikipedia.org/wiki/Environment_variable)
-- are a universal mechanism for [conveying configuration
-- information to Unix programs](http://www.12factor.net/config).


-- To get a value for a key, use `os.getenv`. This will return
-- nil if the key isn't present in the environment.
print("FOO:", os.getenv("FOO"))
print("BAR:", os.getenv("BAR"))



