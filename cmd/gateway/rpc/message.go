package rpc

import (
	"context"
	"tiktok-simple/idl/kitex_gen/message"
	"errors"
)

func MessageChat(ctx context.Context, req *message.MessageChatRequest) (res *message.MessageChatResponse, err error) {
	if req == nil {
		return nil, errors.New("message chat req is nil")
	}
	return MessageClient.MessageChat(ctx, req)
}

func MessageAction(ctx context.Context, req *message.MessageActionRequest) (res *message.MessageActionResponse, err error) {
	if req == nil {
		return nil, errors.New("message action req is nil")
	}
	return MessageClient.MessageAction(ctx, req)
}
