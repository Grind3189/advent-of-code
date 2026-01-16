--- Day 6: Trash Compactor ---
local inspect = require("inspect")

local function part_one(input)
  local nums, operators = input.nums, input.operators

  local grand_total = 0
  for i, group_of_num_str in ipairs(nums) do
    local operator = operators[i]
    local total = 0
    for j, num_str in ipairs(group_of_num_str) do
      -- Remove string spaces and convert to number
      local num = assert(num_str:gsub("%s", ""))

      if j ~= 1 then
        -- Check what operator to perform
        if operator == "*" then
          total = total * num
        elseif operator == "+" then
          total = total + num
        else
          print("Unknown operator: ", operator)
        end
      else
        -- This is the first element just assign num to total
        total = num
      end

    end

    grand_total = grand_total + total
  end

  return grand_total
end

local function part_two(input)
  local nums, operators = input.nums, input.operators

  local grand_total = 0
  for i, group_of_num_str in ipairs(nums) do
    -- Convert group of numbers strings, Each digit should be in its own column
    -- with the most significant digit at the top
    -- Example: 
    -- "125"
    -- " 40"
    -- Result:
    -- "1, 24, 50"
    local parsed_num_strs = {}
    for _, num_str in ipairs(group_of_num_str) do
      for j = 1, #num_str do
        if parsed_num_strs[j] == nil then
          parsed_num_strs[j] = ""
        end

        local char = string.sub(num_str, j, j)
        if char ~= " " then
          parsed_num_strs[j] = parsed_num_strs[j] .. char
        end
      end
    end

    local operator = operators[i]
    local total = 0
    for j, num_str in ipairs(parsed_num_strs) do
      local num = assert(tonumber(num_str))
      if j ~= 1 then
        -- Execute logic
        if operator == "*" then
          total = total * num
        elseif operator == "+" then
          total = total + num
        else
          print("Unknown operator")
          break
        end
      else
        -- This is the first element just assign num to total
        total = num
      end

    end

    grand_total = grand_total + total

  end

  return grand_total

end

-- Get operators and their index,
-- index will be use to know the starting index of a group
-- example input:
-- "123 328"
-- "+   *  "
-- Since operators are always alin left we know that
-- the index of "*" is a start of a group of number
-- in our example:
-- index: 5, value: "*"
-- 123 328
-- 3 is a starting group
local function get_operators()
  local operators = {}
  for line in io.lines("input.txt") do

    if string.find(line, "[%*%+]") then
      for i = 1, #line do
        local char = string.sub(line, i, i)
        if char ~= " " then
          operators[i] = char
        end
      end

    end
  end

  return operators
end


-- Put group of numbers and operators into table
-- Group of numbers includes spaces but not the separator and
-- arrange it to their proper group
local function get_input()
  local input = {
    nums = {},
    operators = {}
  }

  local operators = get_operators()
  for line in io.lines("input.txt") do
    
    if string.find(line, "[%*%+]") then
      break
    end

    -- Get numbers in a row excluding only space separator 
    local num_str = ""
    local num_strs = {}
    for i = 1, #line do
      local char = string.sub(line, i, i)

      if i == #line then
        -- Append and insert the last character
        table.insert(num_strs, num_str .. char)
        break
      end

      if char == " " and operators[i+1] ~= nil then
        -- Character is a separator, insert the num into table
        table.insert(num_strs, num_str)
        num_str = ""
        goto continue
      end

      num_str = num_str .. char
      ::continue::
    end

    -- Arrange numbers from columns to row 
    -- Example: 
    -- "123, 328"
    -- "45, 64"
    -- Result:
    -- "123, 45"
    -- "328, 64"
    for index, value in ipairs(num_strs) do
      if input.nums[index] == nil then
        input.nums[index] = {}
      end
      table.insert(input.nums[index], value)
    end
  end

  -- Operators table is sparse array, convert it to
  -- a simple table before inserting to input
  local indices = {}
  for k in pairs(operators) do
    table.insert(indices, k)
  end
  table.sort(indices)
  for _, index in ipairs(indices) do
    table.insert(input.operators, operators[index])
  end

  return input

end

print("Part One: ", part_one(get_input()))
print("Part Two: ", part_two(get_input()))
