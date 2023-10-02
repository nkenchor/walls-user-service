package helper

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"strings"
	logger "walls-user-service/internal/core/helper/log-helper"

	"github.com/redis/go-redis/v9"
)

type RedisClient struct {
	client *redis.Client
}

func NewRedisClient(client *redis.Client) *RedisClient {
	return &RedisClient{
		client: client,
	}
}

func (r *RedisClient) SubscribeToEvent(ctx context.Context, event interface{}, eventHandler func(context.Context, interface{})) error {
	// Get the channel name from the event object's type

	pubSub := r.client.PSubscribe(ctx, event.(string))
	defer pubSub.Close()

	ch := pubSub.Channel()
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case msg := <-ch:
			var eventData interface{}

			err := json.Unmarshal([]byte(msg.Payload), &eventData)
			if err != nil {
				log.Printf("Error decoding event: %v\n", err)
				logger.LogEvent("ERROR", "Error decoding event: "+err.Error())
				continue
			}

			eventHandler(ctx, eventData) // Pass the appropriate UserRepository instance here
		}
	}

}

func (r *RedisClient) PublishEvent(ctx context.Context, event interface{}, eventType ...string) error {
	eventBytes, err := json.Marshal(event)
	if err != nil {
		return err
	}

	// Determine the channel name
	var channel string = strings.ToUpper(reflect.TypeOf(event).Name())
	fmt.Println("publishing to channel:", channel)
	if len(eventType) != 0 {
		channel = fmt.Sprintf("%s:%s", channel, strings.ToUpper(eventType[0]))
	}

	err = r.client.Publish(ctx, channel, string(eventBytes)).Err()
	if err != nil {
		return err
	}

	return nil
}
