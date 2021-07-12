-- `for` is one of Lua's looping constructs. Here are
-- some types of `for` loops.


-- A three expression loop is called a _numeric for_ loop
-- i=1 is the initial condition, 5 is the end condition, 1 is the step increment
-- the third expression is optional, when omitted 1 is used as the step.
for i=1, 5, 1 do
    print(i)
end
-- Note the variable declared in the loop is local to the body of the loop
-- despite the lack of the local keyword.
print('loop variable scope is local:', i) -- i is nil


