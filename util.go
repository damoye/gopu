package main

func getRedisClientKey(client string) string {
	return config.RedisPrefix + ":client:" + client
}

func getRedisMessageKey(messageKey string) string {
	return config.RedisPrefix + ":message:" + messageKey
}
