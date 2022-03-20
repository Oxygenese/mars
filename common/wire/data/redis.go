package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"github.com/mars-projects/mars/conf"
)

// ProviderRedisSet is data providers.
var ProviderRedisSet = wire.NewSet(NewRedisClient)

// NewRedisClient .
func NewRedisClient(c *conf.Data, logger log.Logger) (*redis.Client, func(), error) {
	l := log.NewHelper(logger)
	rdb := redis.NewClient(&redis.Options{
		Addr:         c.Redis.Addr,
		Password:     c.Redis.Password,
		DB:           int(c.Redis.Db),
		WriteTimeout: c.Redis.WriteTimeout.AsDuration(),
		ReadTimeout:  c.Redis.ReadTimeout.AsDuration(),
	})
	return rdb, func() {
		l.Info("message", "closing the redis resources")
	}, nil
}
