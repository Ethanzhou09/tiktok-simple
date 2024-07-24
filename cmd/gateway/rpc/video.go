package rpc

import(
	"context"
	"tiktok-simple/idl/kitex_gen/video"
	"errors"
)

func Feed(ctx context.Context, req *video.FeedRequest) (resp *video.FeedResponse, err error) {
	if req==nil{
		return nil,errors.New("feed req is nil")
	}
	return VideoClient.Feed(ctx,req)
}

func PublishAction(ctx context.Context, req *video.PublishActionRequest) (resp *video.PublishActionResponse, err error) {
	if req==nil{
		return nil,errors.New("publish action req is nil")
	}
	return VideoClient.PublishAction(ctx,req) 
}

func PublishList(ctx context.Context, req *video.PublishListRequest) (resp *video.PublishListResponse, err error) {
	if req==nil{
		return nil,errors.New("publish list req is nil")
	}
	return VideoClient.PublishList(ctx,req)
}