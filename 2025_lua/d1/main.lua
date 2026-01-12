--- Day 1: Secret Entrance

local function part_one(input)
  local dial, total = 50, 0

  for _, value in ipairs(input) do
    -- Get direction and numOfTurn from value 
    -- Example, R45 to 'R' and 45
    local direction, numOfTurn = string.sub(value, 1, 1), tonumber(string.sub(value, 2))

    -- Execute logic
    local count = 0
    while count < numOfTurn do

      if direction == 'L' then
        dial = dial - 1
        if dial == -1 then
          dial = 99
        end

      elseif direction == 'R' then
        dial = dial + 1
        if dial == 100 then
          dial = 0
        end
      end

      count = count + 1

    end

      -- Calculate if dial points to 0 after each rotation
      if dial == 0 then
          total = total + 1
      end

  end

  return total
end

local function part_two(input)
  local dial, total = 50, 0

  for _, value in ipairs(input) do
    -- Get direction and numOfTurn from value 
    -- Example, R45 to 'R' and 45
    local direction, numOfTurn = string.sub(value, 1, 1), tonumber(string.sub(value, 2))

    -- Execute logic
    local count = 0
    while count < numOfTurn do

      if direction == 'L' then
        dial = dial - 1

        -- Calculate each time dial points to 0 
        if dial == 0 then
          total = total + 1
        end

        if dial == -1 then
          dial = 99
        end

      elseif direction == 'R' then
        dial = dial + 1

        if dial == 100 then
          dial = 0
          total = total + 1
        end
      end

      count = count + 1

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

