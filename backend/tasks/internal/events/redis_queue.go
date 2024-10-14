package events

import (
	"context"

	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
)

const REDIS_KEY = "task_events"

type RedisQueue struct {
	client *redis.Client
}

func NewRedisQueue(client *redis.Client) *RedisQueue {
	return &RedisQueue{client}
}

func (q *RedisQueue) QueueMessage(ctx context.Context, msg Event) error {
	message, err := msg.Serialize()
	if err != nil {
		return errors.Wrap(err, "message could not be serialized")
	}

	if err = q.client.RPush(ctx, REDIS_KEY, message).Err(); err != nil {
		return errors.Wrap(err, "pushing message to Redis failed")
	}

	return nil
}
