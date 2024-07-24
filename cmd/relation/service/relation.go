package service

import(
	"context"
	"tiktok-simple/idl/kitex_gen/relation"
)

type RelationService struct{}

var RelationSrv *RelationService

// 懒汉单例获取UserSrv
func GetRelationSrv() *RelationService {
	if RelationSrv == nil {
		RelationSrv = new(RelationService)
	}
	return RelationSrv
}

func (relationSrv *RelationService) RelationAction(ctx context.Context, req *relation.RelationActionRequest) (resp *relation.RelationActionResponse, err error){

	return
}

func (relationSrv *RelationService) RelationFollowList(ctx context.Context, req *relation.RelationFollowListRequest) (resp *relation.RelationFollowListResponse, err error){

	return
}

func (relationSrv *RelationService)RelationFollowerList(ctx context.Context, req *relation.RelationFollowerListRequest) (resp *relation.RelationFollowerListResponse, err error){

	return
}

func (relationSrv *RelationService)RelationFriendList(ctx context.Context, req *relation.RelationFriendListRequest) (resp *relation.RelationFriendListResponse, err error){

	return
}