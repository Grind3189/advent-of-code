-- Day 4: Printing Department--

-- Parse input.txt add "." placeholder at the first and
-- last position of the table and each line of string
-- so if "@@@" would be ->  
-- "....."
-- ".@@@."
-- "....."
-- now we dont have to worry about index out of bounds
local file = assert(io.open("input.txt", "r"))
local input = {}
if file then
  local line_len = 0
  for line in file:lines() do
    local parsed_line = "." .. line .. "."
    table.insert(input, parsed_line)

    if line_len == 0 then
      line_len = #parsed_line
    end
  end

  table.insert(input, 1, string.rep(".", line_len))
  table.insert(input, string.rep(".", line_len))
else
  print("Error: Could not open file for reading.")
end


local function part_one()
  local total = 0
  -- Loop thorough each value excluding placeholders
  for i = 2, #input - 1 do

    for j = 2, #input[i - 1] do
      local input_char = string.sub(input[i], j, j)

      -- Get the 8 adjacent chars including current
      if input_char == "@" then
        local top = string.sub(input[i-1], j-1, j+1)
        local mid = string.sub(input[i], j-1, j+1)
        local bottom = string.sub(input[i+1], j-1, j+1)

        -- Execute logic
        local adjacent_str = top .. mid .. bottom
        local roll_of_paper = 0
        for k = 1, #adjacent_str do
            local adj_char = string.sub(adjacent_str, k, k)
            if adj_char == "@" then
              roll_of_paper = roll_of_paper + 1
            end
        end
        if roll_of_paper - 1 < 4 then
          total = total + 1
        end
      end
      
    end

  end

  print("Part one: ", total)

end

local function replace_char(str, position)
  local prefix = string.sub(str, 1, position - 1)
  local postfix = string.sub(str, position + 1)

  return prefix .. "." .. postfix
end

local function part_two()
  local go_loop = true
  local total = 0

  while go_loop do
    go_loop = false
    -- Loop thorough each value excluding placeholders
    for i = 2, #input - 1 do

      for j = 2, #input[i - 1] do
        local input_char = string.sub(input[i], j, j)

        -- Get the 8 adjacent chars including current
        if input_char == "@" then
          local top = string.sub(input[i-1], j-1, j+1)
          local mid = string.sub(input[i], j-1, j+1)
          local bottom = string.sub(input[i+1], j-1, j+1)

          -- Execute logic
          local adjacent_str = top .. mid .. bottom
          local roll_of_paper = 0
          for k = 1, #adjacent_str do
              local adj_char = string.sub(adjacent_str, k, k)
              if adj_char == "@" then
                roll_of_paper = roll_of_paper + 1
              end
          end

          -- Excluding the current roll of paper 
          -- check if there are fewer than 4 adjacent roll
          -- replace the char with dot if so and continue looping
          if roll_of_paper - 1 < 4 then
            total = total + 1
            input[i] = replace_char(input[i], j)
            go_loop = true
          end

        end
      end
    end
  end

  print("Part two: ", total)
end

part_one()
part_two()

