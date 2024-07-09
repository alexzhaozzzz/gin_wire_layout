package redisx

//var redisClient *redisx.Client
//
//// InitRedis 初始化redisClient
//func InitRedis(config config.Config) {
//	redisCfg := config.RedisConfig
//	redisClient = redisx.NewClient(&redisx.Options{
//		DB:           redisCfg.Db,
//		Addr:         redisCfg.Addr,
//		Password:     redisCfg.Password,
//		PoolSize:     redisCfg.PoolSize,
//		MinIdleConns: redisCfg.MinIdleConns,
//		IdleTimeout:  time.Duration(redisCfg.IdleTimeout) * time.Second,
//	})
//	_, err := redisClient.Ping(context.TODO()).Result()
//	if err != nil {
//		panic(err)
//	}
//}
//
//// GetRedisClient 获取redis client
//func GetRedisClient() *redisx.Client {
//	if nil == redisClient {
//		panic("Please initialize the Redis client first!")
//	}
//	return redisClient
//}
//
//// CloseRedis 关闭redis client
//func CloseRedis() {
//	if nil != redisClient {
//		_ = redisClient.Close()
//	}
//}
