package rpc

import (
	"context"
	"tiktok-simple/idl/kitex_gen/comment"
	"errors"
)

func CommentAction(ctx context.Context, req *comment.CommentActionRequest) (res *comment.CommentActionResponse, err error) {
	if req == nil {
		return nil, errors.New("comment action req is nil")
	}
	return CommentClient.CommentAction(ctx, req)
}

func CommentList(ctx context.Context, req *comment.CommentListRequest) (res *comment.CommentListResponse, err error) {
	if req == nil {
		return nil, errors.New("comment list req is nil")
	}
	return CommentClient.CommentList(ctx, req)
}