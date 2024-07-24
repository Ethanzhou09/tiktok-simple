package service

import(
	"context"
	"tiktok-simple/idl/kitex_gen/comment"
)

type CommentService struct{}

var CommentSrv *CommentService

// 懒汉单例获取UserSrv
func GetCommentSrv() *CommentService {
	if CommentSrv == nil {
		CommentSrv = new(CommentService)
	}
	return CommentSrv
}

func (commentSrv *CommentService) CommentAction(ctx context.Context, req *comment.CommentActionRequest) (resp *comment.CommentActionResponse, err error){

	return
}

func (commentSrv *CommentService) CommentList(ctx context.Context, req *comment.CommentListRequest) (resp *comment.CommentListResponse, err error){
	return
}
