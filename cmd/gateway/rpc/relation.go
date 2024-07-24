package rpc

import (
	"context"
	"tiktok-simple/idl/kitex_gen/relation"
	"errors"
)

func RelationAction(ctx context.Context, req *relation.RelationActionRequest) (res *relation.RelationActionResponse, err error) {
	if req == nil {
		return nil, errors.New("relation action req is nil")
	}
	return RelationClient.RelationAction(ctx, req)
}

func RelationFollowList(ctx context.Context, req *relation.RelationFollowListRequest) (res *relation.RelationFollowListResponse, err error) {
	if req == nil {
		return nil, errors.New("relation follow list req is nil")
	}
	return RelationClient.RelationFollowList(ctx, req)
}

func RelationFollowerList(ctx context.Context, req *relation.RelationFollowerListRequest) (res *relation.RelationFollowerListResponse, err error) {
	if req == nil {
		return nil, errors.New("relation follower list req is nil")
	}
	return RelationClient.RelationFollowerList(ctx, req)
}

func RelationFriendList(ctx context.Context, req *relation.RelationFriendListRequest) (res *relation.RelationFriendListResponse, err error) {
	if req == nil {
		return nil, errors.New("relation friend list req is nil")
	}
	return RelationClient.RelationFriendList(ctx, req)
}
