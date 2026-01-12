--- Day 2: Gift Shop ---

---Split text in half: e.g. meow to me ow, LOL to LO L
local function split_half(str)
  local left_len = math.ceil(#str/2)
  local first_half = str:sub(1, left_len)
  local second_half = str:sub(left_len + 1, -1)

  return first_half, second_half
end

local function part_one(input)
  local invalid_ids = 0
  for _, value in ipairs(input) do

    --Get the two ids separated by "-" 
    local start_str, end_str = string.match(value, "([^-]+)-([^-]+)")
    local start_id = assert(tonumber(start_str))
    local end_id = assert(tonumber(end_str))

    local current_id = start_id
    while current_id <= end_id do
      --Logic for checking if ID is invalid
      local first_half, second_half = split_half(tostring(current_id))
      if first_half == second_half then
        invalid_ids = invalid_ids + current_id
      end

      current_id = current_id + 1

    end
  end

  return invalid_ids

end


local function part_two(input)

  local invalid_ids = 0
  for _, value in ipairs(input) do
    --Get the two ids separated by "-" 
    local start_str, end_str = string.match(value, "([^-]+)-([^-]+)")
    local start_id = assert(tonumber(start_str))
    local end_id = assert(tonumber(end_str))

    local current_id = start_id
    while current_id <= end_id do
      local id = assert(tostring(current_id))

      -- Loop through all possible pattern and stop half the string
      for pattern_len = 1, #id // 2 do

        -- Execute logic
        if #id % pattern_len == 0 then
          local part = id:sub(1, pattern_len)
          local repeat_by = #id // pattern_len
          local result = string.rep(part, repeat_by)

          if result == id then
            invalid_ids = invalid_ids + current_id
            break
          end

        end

      end

    current_id = current_id + 1

    end

  end

  return invalid_ids

end

local file = assert(io.open("input.txt", "r"))
local content = file:read("a")
file:close()

local input = {}
--Match non-comma and non-whitespace characters
for word in string.gmatch(content, "[^,%s]+") do
  table.insert(input, word)
end

print("Part 1: ", part_one(input))
print("Part 2: ", part_two(input))

