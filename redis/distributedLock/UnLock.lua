local key="dis_lock:lock"
local value=redis.call('GET',key)
if value then
    redis.call("DEL",key)
else
    return redis.error_reply("the held lock is not you! Dont' unlock.")
end