package redisx

import (
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/ryantokmanmokmtm/chat-app-server/app/common/redisx/types"
	"github.com/zeromicro/go-zero/core/jsonx"
	"github.com/zeromicro/go-zero/core/logx"
)

func SubscribeChannel(cli *redis.Client, ctx context.Context, channelName []string) *redis.PubSub {
	return cli.Subscribe(ctx, channelName...)
}

func SendMessageToChannel(cli *redis.Client, ctx context.Context, channelName string, message types.NotificationMessage) {
	logx.WithContext(ctx).Infof("Pushing message %+v to channel %s", message, channelName)
	content, err := jsonx.Marshal(message)
	if err != nil {
		logx.WithContext(ctx).Error(err)
		return
	}
	_, err = cli.Publish(ctx, channelName, content).Result()
	if err != nil {
		logx.WithContext(ctx).Error(err)
	}
}
