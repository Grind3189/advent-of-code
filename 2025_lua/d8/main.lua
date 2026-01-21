--- Day 8: Playground ---
local function calculate_distance(junction_box_one, junction_box_two)
  -- Formula: d = √[(x₂ − x₁)² + (y₂ − y₁)² + (z₂ − z₁)²]
  local x1, y1, z1 = junction_box_one[1], junction_box_one[2], junction_box_one[3]
  local x2, y2, z2 = junction_box_two[1], junction_box_two[2], junction_box_two[3]

  local dx = (x2 - x1) ^ 2
  local dy = (y2 - y1) ^ 2
  local dz = (z2 - z1) ^ 2

  return math.sqrt(dx + dy + dz)
end

local function junction_table_to_str(junction_box)
  return  junction_box[1] .. "," .. junction_box[2] .. "," .. junction_box[3]
end

local function part_one(junction_pairs)
  local circuits_connection = {
    junction_to_ids = {}, -- if we need to find the circuit id using junction box
    id_to_junctions = {} -- if we need to find the list of junction boxes in a circuit id
  }
  for i, junction_pair in ipairs(junction_pairs) do
    -- Stop after 1000 pairs
    if i > 1000 then
      break
    end

    -- Convert junctions to string (original input)
    -- to use as key in circuits table
    local box_one = junction_table_to_str(junction_pair.box_one)
    local box_two = junction_table_to_str(junction_pair.box_two)
    local circuit_one_id = circuits_connection.junction_to_ids[box_one]
    local circuit_two_id = circuits_connection.junction_to_ids[box_two]

    if circuit_one_id ~= nil and circuit_two_id ~= nil and circuit_one_id ~= circuit_two_id then
      -- Both junctions are in another circuit merge
      -- junctions from circuit two to one
      local circuit_two_boxes = circuits_connection.id_to_junctions[circuit_two_id]
      for _, junction in ipairs(circuit_two_boxes) do
        circuits_connection.junction_to_ids[junction] = circuit_one_id
        table.insert(circuits_connection.id_to_junctions[circuit_one_id], junction)
      end
      circuits_connection.id_to_junctions[circuit_two_id] = nil

    elseif circuit_one_id == nil and circuit_two_id ~= nil then
      -- Connect junction one to junction two
      circuits_connection.junction_to_ids[box_one] = circuit_two_id
      table.insert(circuits_connection.id_to_junctions[circuit_two_id], box_one)

    elseif circuit_one_id ~= nil and circuit_two_id == nil then
      -- Connect junction two to junction one
      circuits_connection.junction_to_ids[box_two] = circuit_one_id
      table.insert(circuits_connection.id_to_junctions[circuit_one_id], box_two)

    elseif circuit_one_id == nil and circuit_two_id == nil then
      -- Both junctions are not in a circuit, connect them
      circuits_connection.junction_to_ids[box_one] = i
      circuits_connection.junction_to_ids[box_two] = i
      circuits_connection.id_to_junctions[i] = {box_one, box_two}
    end
  end

  -- Get the three largest circuits 
  local top_three = {}
  for _, value in pairs(circuits_connection.id_to_junctions) do
    if #top_three < 3 then
      table.insert(top_three, #value)
      goto continue
    end

    if #value > top_three[3] then
      top_three[3] = #value
      table.sort(top_three, function (a, b)
        return a > b
      end)
    end

    ::continue::
  end

  return top_three[1] * top_three[2] * top_three[3]
end

local function part_two(junction_pairs, total_boxes)
  local last_connection = {}
  local circuits_connection = {
    junction_to_ids = {}, -- if we need to find the circuit id using junction box
    id_to_junctions = {} -- if we need to find the list of junction boxes in a circuit id
  }
  for i, junction_pair in ipairs(junction_pairs) do
    local target_circuit = i

    -- Convert junctions to string (original input)
    -- to use as key in circuits table
    local box_one = junction_table_to_str(junction_pair.box_one)
    local box_two = junction_table_to_str(junction_pair.box_two)
    local circuit_one_id = circuits_connection.junction_to_ids[box_one]
    local circuit_two_id = circuits_connection.junction_to_ids[box_two]

    if circuit_one_id ~= nil and circuit_two_id ~= nil and circuit_one_id ~= circuit_two_id then
      -- Both junctions are in another circuit merge
      -- junctions from circuit two to one
      local circuit_two_boxes = circuits_connection.id_to_junctions[circuit_two_id]
      for _, junction in ipairs(circuit_two_boxes) do
        circuits_connection.junction_to_ids[junction] = circuit_one_id
        table.insert(circuits_connection.id_to_junctions[circuit_one_id], junction)
      end
      circuits_connection.id_to_junctions[circuit_two_id] = nil
      target_circuit = circuit_one_id

    elseif circuit_one_id == nil and circuit_two_id ~= nil then
      -- Connect junction one to junction two
      circuits_connection.junction_to_ids[box_one] = circuit_two_id
      table.insert(circuits_connection.id_to_junctions[circuit_two_id], box_one)
      target_circuit = circuit_two_id

    elseif circuit_one_id ~= nil and circuit_two_id == nil then
      -- Connect junction two to junction one
      circuits_connection.junction_to_ids[box_two] = circuit_one_id
      table.insert(circuits_connection.id_to_junctions[circuit_one_id], box_two)
      target_circuit = circuit_one_id

    elseif circuit_one_id == nil and circuit_two_id == nil then
      -- Both junctions are not in a circuit, connect them
      circuits_connection.junction_to_ids[box_one] = i
      circuits_connection.junction_to_ids[box_two] = i
      circuits_connection.id_to_junctions[i] = {box_one, box_two}
    end

    -- Check the if all junction boxes is connected
    local circuit = circuits_connection.id_to_junctions[target_circuit]
    local circuit_size = circuit ~= nil and #circuit or nil
    if circuit_size == total_boxes then
      last_connection[1], last_connection[2] = junction_pair.box_one,junction_pair.box_two
      break
    end

  end

  return last_connection[1][1] * last_connection[2][1]
end

-- Parse input.txt
local input = {}
for line in io.lines("input.txt") do
  -- Separate string with comma, and convert to number
  local nums = {}
  for num in string.gmatch(line, "[^%,]+") do
    table.insert(nums, assert(tonumber(num)))
  end

  table.insert(input,nums)
end

-- Get the pairs of junction boxes
local junction_pairs = {}
for i, junction_box in ipairs(input) do
  for j = i + 1, #input do
    local distance = calculate_distance(junction_box, input[j])
    table.insert(junction_pairs, {box_one = junction_box, box_two = input[j], distance = distance})
  end
end

-- Sort the distances of junction pairs
table.sort(junction_pairs, function (a, b)
  return a.distance < b.distance
end)

print("Part 1: ", part_one(junction_pairs))
print("Part 2: ", part_two(junction_pairs, #input))

