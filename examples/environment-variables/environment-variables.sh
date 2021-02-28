# Running the program shows that we pick up the value
# for `FOO` that we set in the program, but that
# `BAR` is empty.
$ lua environment-variables.lua

# If we set `BAR` in the environment first, the running
# program picks that value up.
$ BAR=2 lua environment-variables.lua

