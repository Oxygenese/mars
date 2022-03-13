package data

import (
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"github.com/mars-projects/oauth2/v4"
	"github.com/mars-projects/oauth2/v4/store"
)

// ProviderRedisTokenStoreSet is data providers.
var ProviderRedisTokenStoreSet = wire.NewSet(NewRedisTokenStore)

func NewRedisTokenStore(client *redis.Client) oauth2.TokenStore {
	return store.NewRedisStore(client)
}
