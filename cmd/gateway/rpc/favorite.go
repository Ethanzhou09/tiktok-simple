package rpc

import (
	"context"
	"tiktok-simple/idl/kitex_gen/favorite"
	"errors"
)

func FavoriteAction(ctx context.Context, req *favorite.FavoriteActionRequest) (resp *favorite.FavoriteActionResponse, err error) {
	if req==nil {
		return nil, errors.New("favoriteaction req is nil")
	}
	return FavoriteClient.FavoriteAction(ctx,req)
}

func FavoriteList(ctx context.Context, req *favorite.FavoriteListRequest) (resp *favorite.FavoriteListResponse, err error) {
	if req==nil{
		return nil,errors.New("favoritelist req is nil")
	}
	return FavoriteClient.FavoriteList(ctx,req)
}