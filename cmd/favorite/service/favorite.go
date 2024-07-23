package service

import(
	"context"
	"tiktok-simple/idl/kitex_gen/favorite"
)

type FavoriteService struct{}

var FavoriteSrv *FavoriteService

// 懒汉单例获取UserSrv
func GetFavoriteSrv() *FavoriteService {
	if FavoriteSrv == nil {
		FavoriteSrv = new(FavoriteService)
	}
	return FavoriteSrv
}

func (favoriteSrv *FavoriteService) FavoriteAction(ctx context.Context, req *favorite.FavoriteActionRequest) (resp *favorite.FavoriteActionResponse, err error) {
	
	return
}

func (favoriteSrv *FavoriteService) FavoriteList(ctx context.Context, req *favorite.FavoriteListRequest) (resp *favorite.FavoriteListResponse, err error) {

	return
}