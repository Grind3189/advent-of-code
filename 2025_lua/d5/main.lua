--- Day 5: Cafeteria ---
local inspect = require("inspect")

-- Read file and remove whitespace at the end of string
local file = assert(io.open("input.txt", "r"))
local content = file:read("a"):gsub("%s+$", "")
file:close()

-- Split by blankline to get
-- Fresh and Available ingredient IDs 
local blank_line_pos = content:find("\n\n")
local fresh_str_ids = content:sub(1, blank_line_pos - 1)
local available_str_ids = content:sub(blank_line_pos + 2)

local function split_by_newline(str)
  local lines = {}
  for line in str:gmatch("[^\n]+") do
    table.insert(lines, line)
  end
  return lines
end


local function part_one(available_ids, fresh_ids)
  local total = 0
  for _, id in ipairs(available_ids) do
    id = assert(tonumber(id))

    for _, range_str in ipairs(fresh_ids) do

      -- Get both numbers from the range
      -- "3-5" -> 3 and 5
      local min,max = range_str:match("(%d+)-(%d+)")
      min,max = assert(tonumber(min)), assert(tonumber(max))

      -- Check if the current if falls within the range
      if id <= max and id >= min then
        total = total + 1
        break
      end

    end

  end
  return total
end

local function part_two(fresh_ids)
      local r1_min ,r1_max = fresh_ids[2]:match("(%d+)-(%d+)")
      local r2_min ,r2_max = fresh_ids[4]:match("(%d+)-(%d+)")
      r1_min,r1_max= assert(tonumber(r1_min)), assert(tonumber(r1_max))
      r2_min,r2_max= assert(tonumber(r2_min)), assert(tonumber(r2_max))
      print("R1 Min:", r1_min, "R1 Max:", r1_max)
      print("R2 Min:", r2_min, "R2 Max:", r2_max)

      local overlap_start = math.max(r1_min, r2_min)
      local overlap_end = math.min(r1_max, r2_max)
      print("overlap_start:", overlap_start)
      print("overlap_end:", overlap_end)

end


-- Put into separate table
local fresh_ids = split_by_newline(fresh_str_ids)
local available_ids = split_by_newline(available_str_ids)

-- print("Part 1: ", part_one(available_ids, fresh_ids))
part_two(fresh_ids)

