package bll

import "gopkg.in/redis.v5"
import "github.com/damoye/gopu/config"

var (
	redisClient  = redis.NewClient(&redis.Options{Addr: config.Conf.RedisAddress})
	redisChannel = config.Conf.RedisPrefix + "CHANNEL"
)

type channelMessage struct {
	TaskID int64
	Tokens []string
	Data   string
}
