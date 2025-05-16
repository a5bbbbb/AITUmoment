package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/a5bbbbb/AITUmoment/core_service/internal/models"
	"github.com/a5bbbbb/AITUmoment/core_service/pkg/redis"
	goredis "github.com/redis/go-redis/v9"
)

const (
	keyPrefix = "user:%d"
)

// User is our entity, client is redis client through which we are going to make queries to redis
type User struct {
	client *redis.Client
	ttl    time.Duration
}

func NewUser(client *redis.Client, ttl time.Duration) *User {
	return &User{
		client: client,
		ttl:    ttl,
	}
}

func (c *User) Set(ctx context.Context, user *models.User) error {
	data, err := json.Marshal(*user)
	if err != nil {
		return fmt.Errorf("failed to marshal client: %w", err)
	}

	return c.client.Unwrap().Set(ctx, c.key(user.Id), data, c.ttl).Err()
}

func (c *User) SetMany(ctx context.Context, users []models.User) error {
	pipe := c.client.Unwrap().Pipeline()
	for _, user := range users {
		data, err := json.Marshal(user)
		if err != nil {
			return fmt.Errorf("failed to marshal client: %w", err)
		}
		pipe.Set(ctx, c.key(user.Id), data, c.ttl)
	}
	_, err := pipe.Exec(ctx)
	if err != nil {
		return fmt.Errorf("failed to set many clients: %w", err)
	}
	return nil
}

func (c *User) Get(ctx context.Context, userId int) (*models.User, error) {
	data, err := c.client.Unwrap().Get(ctx, c.key(userId)).Bytes()
	if err != nil {
		if err == goredis.Nil {
			return &models.User{}, fmt.Errorf("user not found id=%v", userId) // not found
		}

		return &models.User{}, fmt.Errorf("failed to get client: %w", err)
	}

	var user models.User
	err = json.Unmarshal(data, &user)
	if err != nil {
		return &models.User{}, fmt.Errorf("failed to unmarshal client: %w", err)
	}

	return &user, nil
}

func (c *User) Delete(ctx context.Context, userId int) error {
	return c.client.Unwrap().Del(ctx, c.key(userId)).Err()
}

func (c *User) key(userId int) string {
	return fmt.Sprintf(keyPrefix, userId)
}
