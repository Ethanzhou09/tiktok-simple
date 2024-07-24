package service

import(
	"context"
	"tiktok-simple/idl/kitex_gen/message"
)

type MessageService struct{}

var MessageSrv *MessageService

// 懒汉单例获取UserSrv
func GetMessageSrv() *MessageService {
	if MessageSrv == nil {
		MessageSrv = new(MessageService)
	}
	return MessageSrv
}

func (messageSrv *MessageService) MessageChat(ctx context.Context, req *message.MessageChatRequest) (resp *message.MessageChatResponse, err error){

	return
}

func (messageSrv *MessageService) MessageAction(ctx context.Context, req *message.MessageActionRequest) (resp *message.MessageActionResponse, err error){

	return
}