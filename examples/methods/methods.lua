-- lua supports _methods_ defined on tables.


-- items declared in this table are shared between all instances
-- but can be overridden
local My_Psuedo_Class = {
    shared = "string is shared"
}

function My_Psuedo_Class:new()

    -- items declared in this table are unique between instances
    local instance = {
        unique = "string is unique"
    }

    setmetatable(instance, { __index = My_Psuedo_Class })

    return instance
end

-- declaration of a method that changes internal state
function My_Psuedo_Class:method(new_val)
    self.shared = new_val
    self.unique = new_val
end

-- another method that prints internal state
function My_Psuedo_Class:print()
    print('shared: ', self.shared, ' unique: '..self.unique)
end

local instance_a = My_Psuedo_Class:new()
local instance_b = My_Psuedo_Class:new()


instance_a:print()
instance_b:print()
print()

-- we can change the shared state
My_Psuedo_Class.shared = "new shared state"

instance_a:print()
instance_b:print()
print()

instance_a:method('overwrite first')
instance_b:method('overwrite second')

instance_a:print()
instance_b:print()
print()

-- once the members have overridden the shared state
-- changes will only affect new instances
My_Psuedo_Class.shared = "only in inst. c"
local instance_c = My_Psuedo_Class:new()

instance_a:print()
instance_b:print()
instance_c:print()