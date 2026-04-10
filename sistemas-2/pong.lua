local ball_x = 316
local ball_y = 236
local ball_size = 8
local ball_velocity_x = 4
local ball_velocity_y = 3

local paddle_width = 10
local paddle_height = 64
local paddle_velocity = 5

local left_paddle_x = 10
local right_paddle_x = 620
local left_paddle_y = 208
local right_paddle_y = 208

local pressed = {}

function key(name, is_pressed)
    pressed[name] = is_pressed
end

function tick()
    if pressed["w"] then
        left_paddle_y = left_paddle_y - paddle_velocity
    end
    if pressed["s"] then
        left_paddle_y = left_paddle_y + paddle_velocity
    end
    if pressed["up"] then
        right_paddle_y = right_paddle_y - paddle_velocity
    end
    if pressed["down"] then
        right_paddle_y = right_paddle_y + paddle_velocity
    end

    left_paddle_y = math.max(0, math.min(480 - paddle_height, left_paddle_y))
    right_paddle_y = math.max(0, math.min(480 - paddle_height, right_paddle_y))

    ball_x = ball_x + ball_velocity_x
    ball_y = ball_y + ball_velocity_y

    if ball_y <= 0 then
        ball_velocity_y = math.abs(ball_velocity_y)
    end
    if ball_y + ball_size >= 480 then
        ball_velocity_y = -math.abs(ball_velocity_y)
    end

    if ball_x <= left_paddle_x + paddle_width
    and ball_y + ball_size >= left_paddle_y
    and ball_y <= left_paddle_y + paddle_height then
        ball_velocity_x = math.abs(ball_velocity_x)
    end

    if ball_x + ball_size >= right_paddle_x
    and ball_y + ball_size >= right_paddle_y
    and ball_y <= right_paddle_y + paddle_height then
        ball_velocity_x = -math.abs(ball_velocity_x)
    end

    if ball_x < 0 or ball_x + ball_size > 640 then
        ball_x = 316
        ball_y = 236
    end

    rect(left_paddle_x, left_paddle_y, paddle_width, paddle_height)
    rect(right_paddle_x, right_paddle_y, paddle_width, paddle_height)
    rect(ball_x, ball_y, ball_size, ball_size)
end
