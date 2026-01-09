--- Day 3: Lobby ---
local file = assert(io.open("input.txt", "r"))

local input = {}
if file then
  for line in file:lines() do
    table.insert(input, line)
  end
  file:close()
else
  print("Error: Could not open file for reading.")
end

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

local function part_one()
  local total = 0
  for _, value in ipairs(input) do
    local combination = find_highest_combination(value, "", 2)
    total = total + assert(tonumber(combination))
  end
  print("Part 1:", total)
end


local function part_two()
  local total = 0

  for _, value in ipairs(input) do
    local combination = find_highest_combination(value, "", 12)
    total = total + assert(tonumber(combination))
  end

  print("Part 2:", total)

end




part_one()
part_two()

