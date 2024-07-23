package service

import(
	"context"
	"tiktok-simple/idl/kitex_gen/video"
)

type VideoService struct{}

var VideoSrv *VideoService

// 懒汉单例获取UserSrv
func GetVideoSrv() *VideoService {
	if VideoSrv == nil {
		VideoSrv = new(VideoService)
	}
	return VideoSrv
}

func (videosrv *VideoService) Feed(ctx context.Context, req *video.FeedRequest) (resp *video.FeedResponse, err error) {
	// TODO: Your code here...
	return
}

func (videosrv *VideoService) PublishAction(ctx context.Context, req *video.PublishActionRequest) (resp *video.PublishActionResponse, err error) {
	// TODO: Your code here...
	return
}

func (videosrv *VideoService) PublishList(ctx context.Context, req *video.PublishListRequest) (resp *video.PublishListResponse, err error) {
	// TODO: Your code here...
	return
}