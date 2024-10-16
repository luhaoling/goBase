local val=redis.call('get',KEYS[1])
-- 加锁重试时，要判断自己上一次是不是加锁成功了
if val==false then
    -- key 不存在
    return redis.call('set',KEYS[1],ARGV[1],'EX',ARGV[2])
elseif val==ARGV[1] then
    -- 刷新过期时间
    redis.call('expire',KEYS[1],ARGV[2])
    return "OK"
else
    -- 此时别人持有锁
    return ""
end