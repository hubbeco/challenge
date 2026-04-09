local image_path = "dvd.png"
local image_width = 200
local image_height = 100

local position_x = 100
local position_y = 100
local velocity_x = 3
local velocity_y = 2

function tick()
    position_x = position_x + velocity_x
    position_y = position_y + velocity_y

    if position_x <= 0 then
        velocity_x = math.abs(velocity_x)
    end
    if position_x + image_width >= 640 then
        velocity_x = -math.abs(velocity_x)
    end
    if position_y <= 0 then
        velocity_y = math.abs(velocity_y)
    end
    if position_y + image_height >= 480 then
        velocity_y = -math.abs(velocity_y)
    end

    png(position_x, position_y, image_path)
end
