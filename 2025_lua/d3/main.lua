--- Day 3: Lobby ---

-- Get the largest combination of numbers from a string recursively
-- For example (Get 2 digit highest combination):
-- 1st iteration: input: "315", combination: "", need_combination: 2(two)
--    Get the 1st highest num from "31" -> "3"
--    "5" would be the new input from "315" cause we're done with "3" and "1"
-- 2nd iteration: input: "5", "combination": "3", need_combination: 1(one)
--    return "35"
local function find_highest_combination(input_str, combination_str, need_combination)
  local rem_combination = need_combination - #combination_str

  -- Base case: if no combination needed, return the combination
  -- or if the combination needed is equal to the length of input
  -- then, append and return
  if rem_combination == 0 then
    return combination_str
  elseif rem_combination == #input_str then
    return combination_str .. input_str
  end


  local largest, index = 0, 0
  for i = 1, #input_str - (rem_combination -1) do

    local num_str = string.sub(input_str, i, i)
    local num = assert(tonumber(num_str))

    if num > largest then
      largest = num
      index = i
    end

  end

  -- recursive step
   return find_highest_combination(
    string.sub(input_str, index+1),
    combination_str .. assert(tostring(largest)),
    need_combination
  )

end

local function part_one(input)
  local total = 0
  for _, value in ipairs(input) do
    local combination = find_highest_combination(value, "", 2)
    total = total + assert(tonumber(combination))
  end

  return total
end


local function part_two(input)
  local total = 0

  for _, value in ipairs(input) do
    local combination = find_highest_combination(value, "", 12)
    total = total + assert(tonumber(combination))
  end

  return total
end


local input = {}
for line in io.lines("input.txt") do
  table.insert(input, line)
end

print("Part 1: ", part_one(input))
print("Part 2: ", part_two(input))

