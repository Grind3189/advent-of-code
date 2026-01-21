--- Day 7: Laboratories ---
local inspect = require("inspect")

local SPLITTER = "^"
-- Replace specific "." of a string with "|"
local function replace_dot(orig_str, index, is_splitter)
  local result = ""
  if is_splitter then
    -- replace left and right position of target index
    local left_pos, right_pos = index - 1, index + 1
    local left_result = string.sub(orig_str, 1, left_pos - 1) .. "|"
    local right_result = "|" .. string.sub(orig_str, right_pos + 1)
    result = left_result .. SPLITTER .. right_result
  else
    -- replace target index with "|"
    local prefix = string.sub(orig_str, 1, index-1)
    local postfix = string.sub(orig_str, index+1)
    result = prefix .. "|" .. postfix
  end

  return result
end

local function part_one(input)
  local total = 0
  for row, line in ipairs(input) do
    if row == #input then
      -- Last row
      break
    end

    -- Loop through each character of a line
    for i = 1, #line do
      local current_char = string.sub(line, i, i)

      if current_char == "S" or current_char == "|" then
        -- Found a Beam, jump 1 row below
        local target_str = input[row+1]
        local target_char = string.sub(target_str, i, i)

        if target_char == "." then
          -- Manipulate input,  replace "." with "|"
          input[row+1] = replace_dot(target_str, i, false)
        end

        if target_char == SPLITTER then
          -- Splitter found, manipulate input, replace left and 
          -- right "." of "^" with "|".
          -- Calculate total
          input[row+1] = replace_dot(target_str, i, true)
          total = total + 1
        end
      end

    end

  end

  return total
end


local function part_two(input)
  -- beam_count is a key value pair
  -- key stores the column number of the beams
  -- value stores the number of active beams in that column
  local beam_count = {}

  for row, line in ipairs(input) do
    if row == #input then
      -- Last row
      break
    end

    -- Loop through each character of a line
    for i = 1, #line do
      local current_char = string.sub(line, i, i)

      if current_char == "S" or current_char == "|" then
        -- Found a Beam, jump 1 row below
        local target_str = input[row+1]
        local target_char = string.sub(target_str, i, i)

        if target_char == "." then
          -- Manipulate input,  replace "." with "|"
          input[row+1] = replace_dot(target_str, i, false)
        end

        if target_char == SPLITTER then
          -- Splitter found, manipulate input, replace left and 
          -- right "." of "^" with "|"
          input[row+1] = replace_dot(target_str, i, true)

          -- At this point we found a beam, and below that is a SPLITTER "^".
          -- Since we're tracking how many active beams or TIMELINES,
          -- we'll add the current_beam_count to the active beam
          -- count of left and right position
          local current_beam_count = beam_count[i] ~= nil and beam_count[i] or 1
          local left_pos, right_pos = i - 1, i + 1
          -- Left beam count
          if beam_count[left_pos] ~= nil then
            beam_count[left_pos] = beam_count[left_pos] + current_beam_count
          else
            beam_count[left_pos] = current_beam_count
          end
          -- Track right beam count
          if beam_count[right_pos] ~= nil then
            beam_count[right_pos] = beam_count[right_pos] + current_beam_count
          else
            beam_count[right_pos] = current_beam_count
          end

          -- This beam count is splitted, reset it to 0
          beam_count[i] = 0
        end

      end

    end

  end

  -- Count the total number of active beam count
  local total = 0
  for _, value in pairs(beam_count) do
    if value > 0 then
      total = total + value
    end
  end

  return total

end

local input = {}
for line in io.lines("input.txt") do
  table.insert(input, line)
end

print("Part 1: ", part_one(input))
print("Part 2: ", part_two(input))
