--- Day 1: Secret Entrance

local input = {}

for line in io.lines("input.txt") do
  table.insert(input, line)
end

local function part_one()
  local dial, total = 50, 0

  for _, line in ipairs(input) do
    -- Seperate each line to direction and number of turn
    -- Example, R45 to 'R' and 45
    local direction, numOfTurn = string.sub(line, 1, 1), tonumber(string.sub(line, 2))

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

print("Part 1:", total)

end

local function part_two()
  local dial, total = 50, 0

  for _, line in ipairs(input) do
    -- Seperate each line to direction and number of turn
    -- Example, R45 to 'R' and 45
    local direction, numOfTurn = string.sub(line, 1, 1), tonumber(string.sub(line, 2))

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

print("Part 2:", total)
end


part_one()
part_two()

