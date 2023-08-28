package db

import (
	"context"

	"github.com/go-redis/redis/v8"
)

type IAuthCache interface {
	Temp()
}

type authCache struct {
	c   *redis.Client
	ctx context.Context
}

func InitAuthCache(c *redis.Client, ctx context.Context) IAuthCache {
	return &authCache{c, ctx}
}

func (ac *authCache) Temp() {
	// Temp.
}
