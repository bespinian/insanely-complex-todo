package events

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
)

const REDIS_CHANNEL_PREFIX = "events"

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

	channel := fmt.Sprintf("%s.%s", REDIS_CHANNEL_PREFIX, msg.Type)
	if err = q.client.Publish(ctx, channel, message).Err(); err != nil {
		return errors.Wrap(err, "pushing message to Redis failed")
	}

	// events können dann mit `PSUBSCRIBE events.*` abgehört werden.

	return nil
}
