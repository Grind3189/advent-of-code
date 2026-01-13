--- Day 5: Cafeteria ---

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

  local invalid_ranges_index = {}
  local num_of_non_overlap = 0
  for i, current_range in ipairs(fresh_ids) do
    local r1_min ,r1_max = current_range:match("(%d+)-(%d+)")
    r1_min,r1_max = assert(tonumber(r1_min)), assert(tonumber(r1_max))

    -- Check if the current range is one of the invalid ranges
    -- Invalid ranges are ranges that is inside the previous ranges
    for _, invalid_index in ipairs(invalid_ranges_index) do
      if i == invalid_index then
        goto continue
      end
    end

    if i == #fresh_ids then
      goto compute
    end

    for j = i + 1, #fresh_ids do
      local r2_min ,r2_max = fresh_ids[j]:match("(%d+)-(%d+)")
      r2_min, r2_max = assert(tonumber(r2_min)), assert(tonumber(r2_max))
      local overlap_start = math.max(r1_min, r2_min)
      local overlap_end = math.min(r1_max, r2_max)

      if overlap_start <= overlap_end then
        -- Ranges overlap, check what kind of overlap and
        -- execute the logic

        if overlap_start == r1_min and overlap_end == r1_max then
          -- Current range is in range 2, so we skip this range
          goto continue
        elseif overlap_start == r2_min and overlap_end == r2_max then
          -- Range 2 is in current range, so range 2 is invalid
          table.insert(invalid_ranges_index, j)
        elseif overlap_start == r2_min then
          -- Overlap is at the beginning of range 2
          -- Example: R1 = "1,2,3", R2 = "3,4,5"
          -- R2 would become "4,5"
          r2_min = overlap_end + 1
          fresh_ids[j] = r2_min .. "-" .. r2_max
        else
          -- Overlap is at the end of range 2
          -- Example: R1 = "3,4,5", R2 = "1,2,3"
          -- R2 would become "1,2"
          r2_max = overlap_start - 1
          fresh_ids[j] = r2_min .. "-" .. r2_max
        end

      end

    end

    ::compute::
    -- Compute how many numbers are non overlapping
    -- E.G. "4,5,6" would be "3"
    num_of_non_overlap = num_of_non_overlap + (r1_max - r1_min) + 1

    ::continue::
  end

  return num_of_non_overlap

end

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

-- Put into separate table
local fresh_ids = split_by_newline(fresh_str_ids)
local available_ids = split_by_newline(available_str_ids)

print("Part 1: ", part_one(available_ids, fresh_ids))
print("Part 2: ", part_two(fresh_ids))

