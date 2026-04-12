package repository

import (
	"context"
	"encoding/json"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisRepository struct {
	Redis *redis.Client
}

func NewRedisRepository(redis *redis.Client) *RedisRepository {
	return &RedisRepository{
		Redis: redis,
	}
}

func (r *RedisRepository) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	bytes, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return r.Redis.Set(ctx, key, bytes, expiration).Err()
}

func (r *RedisRepository) Get(ctx context.Context, key string, dest interface{}) error {
	val, err := r.Redis.Get(ctx, key).Result()
	if err != nil {
		return err
	}
	return json.Unmarshal([]byte(val), dest)
}

func (r *RedisRepository) Delete(ctx context.Context, key string) error {
	return r.Redis.Del(ctx, key).Err()
}

func (r *RedisRepository) DeleteByPrefix(ctx context.Context, prefix string) error {
	var cursor uint64
	var keys []string
	var err error

	for {
		keys, cursor, err = r.Redis.Scan(ctx, cursor, prefix+"*", 100).Result()
		if err != nil {
			return err
		}

		if len(keys) > 0 {
			if err := r.Redis.Del(ctx, keys...).Err(); err != nil {
				return err
			}
		}

		if cursor == 0 {
			break
		}
	}

	return nil
}
