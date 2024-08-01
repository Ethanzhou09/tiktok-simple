package service

import(
	"context"
	"tiktok-simple/idl/kitex_gen/video"
	"time"
	"fmt"
	"bytes"
	"tiktok-simple/pkg/jwt"
	"tiktok-simple/dal/db"
	"tiktok-simple/pkg/minio"
	"tiktok-simple/pkg/ffmpeg"
	"tiktok-simple/idl/kitex_gen/user"
)

type VideoService struct{}

var VideoSrv *VideoService

const limit = 30

// 懒汉单例获取UserSrv
func GetVideoSrv() *VideoService {
	if VideoSrv == nil {
		VideoSrv = new(VideoService)
	}
	return VideoSrv
}

func (videosrv *VideoService) Feed(ctx context.Context, req *video.FeedRequest) (resp *video.FeedResponse, err error) {
	nextTime := time.Now().UnixMilli()
	var userID int64 = -1
	// 验证token有效性
	if req.Token != "" {
		claims, err := jwt.ParseToken(req.Token)
		if err != nil {
			res := &video.FeedResponse{
				StatusCode: -1,
				StatusMsg:  "token 解析错误",
			}
			return res, nil
		}
		userID = int64(claims.Id)
	}
	// 调用数据库查询 video_list
	videos, err := db.MGetVideos(ctx, limit, &req.LatestTime)
	if err != nil {
		res := &video.FeedResponse{
			StatusCode: -1,
			StatusMsg:  "视频获取失败：服务器内部错误",
		}
		return res, nil
	}
	videoList := make([]*video.Video, 0)
	for _, r := range videos {
		author, err := db.GetUserById(ctx, int64(r.AuthorID))
		if err != nil {
			return nil, err
		}
		relation, err := db.GetRelationByUserIDs(ctx, userID, int64(author.ID))
		if err != nil {
			res := &video.FeedResponse{
				StatusCode: -1,
				StatusMsg:  "视频获取失败：服务器内部错误",
			}
			return res, nil
		}
		favorite, err := db.GetFavoriteVideoRelationByUserVideoID(ctx, userID, int64(r.ID))
		if err != nil {
			res := &video.FeedResponse{
				StatusCode: -1,
				StatusMsg:  "视频获取失败：服务器内部错误",
			}
			return res, nil
		}
		playUrl, err := minio.GetFileTemporaryURL(minio.VideoBucketName, r.PlayUrl)
		if err != nil {
			res := &video.FeedResponse{
				StatusCode: -1,
				StatusMsg:  "服务器内部错误：视频获取失败",
			}
			return res, nil
		}
		coverUrl, err := minio.GetFileTemporaryURL(minio.CoverBucketName, r.CoverUrl)
		if err != nil {
			res := &video.FeedResponse{
				StatusCode: -1,
				StatusMsg:  "服务器内部错误：封面获取失败",
			}
			return res, nil
		}
		avatarUrl, err := minio.GetFileTemporaryURL(minio.AvatarBucketName, author.Avatar)
		if err != nil {
			res := &video.FeedResponse{
				StatusCode: -1,
				StatusMsg:  "服务器内部错误：头像获取失败",
			}
			return res, nil
		}
		backgroundUrl, err := minio.GetFileTemporaryURL(minio.BackgroundImageBucketName, author.BackgroundImage)
		if err != nil {
			res := &video.FeedResponse{
				StatusCode: -1,
				StatusMsg:  "服务器内部错误：背景图获取失败",
			}
			return res, nil
		}

		videoList = append(videoList, &video.Video{
			Id: int64(r.ID),
			Author: &user.User{
				Id:              int64(author.ID),
				Name:            author.UserName,
				FollowCount:     int64(author.FollowingCount),
				FollowerCount:   int64(author.FollowerCount),
				IsFollow:        relation != nil,
				Avatar:          avatarUrl,
				BackgroundImage: backgroundUrl,
				Signature:       author.Signature,
				TotalFavorited:  int64(author.TotalFavorited),
				WorkCount:       int64(author.WorkCount),
				FavoriteCount:   int64(author.FavoriteCount),
			},
			PlayUrl:       playUrl,
			CoverUrl:      coverUrl,
			FavoriteCount: int64(r.FavoriteCount),
			CommentCount:  int64(r.CommentCount),
			IsFavorite:    favorite != nil,
			Title:         r.Title,
		})
	}
	if len(videos) != 0 {
		nextTime = videos[len(videos)-1].UpdatedAt.UnixMilli()
	}
	res := &video.FeedResponse{
		StatusCode: 0,
		StatusMsg:  "success",
		VideoList:  videoList,
		NextTime:   nextTime,
	}
	return res, nil
}

func (videosrv *VideoService) PublishAction(ctx context.Context, req *video.PublishActionRequest) (resp *video.PublishActionResponse, err error) {
	// TODO: Your code here...
	claims, err := jwt.ParseToken(req.Token)
	if err != nil {
		res := &video.PublishActionResponse{
			StatusCode: -1,
			StatusMsg:  "token 解析错误",
		}
		return res, nil
	}
	userID := claims.Id

	if len(req.Title) == 0 || len(req.Title) > 32 {
		res := &video.PublishActionResponse{
			StatusCode: -1,
			StatusMsg:  "标题不能为空且不能超过32个字符",
		}
		return res, nil
	}

	// 限制文件上传大小
	maxSize := 50
	size := len(req.Data)
	if size > maxSize*1000*1000 {
		res := &video.PublishActionResponse{
			StatusCode: -1,
			StatusMsg:  fmt.Sprintf("该视频文件大于%dMB，上传受限", maxSize),
		}
		return res, nil
	}

	createTimestamp := time.Now().UnixMilli()
	videoTitle, coverTitle := fmt.Sprintf("%d_%s_%d.mp4", userID, req.Title, createTimestamp), fmt.Sprintf("%d_%s_%d.png", userID, req.Title, createTimestamp)

	// 插入数据库
	v := &db.Video{
		Title:    req.Title,
		PlayUrl:  videoTitle,
		CoverUrl: coverTitle,
		AuthorID: uint(userID),
	}
	err = db.CreateVideo(ctx, v)
	if err != nil {
		res := &video.PublishActionResponse{
			StatusCode: -1,
			StatusMsg:  "视频发布失败，服务器内部错误",
		}
		return res, nil
	}

	go func() {
		err := VideoPublish(req.Data, videoTitle, coverTitle)
		if err != nil {
			// 发生错误，则删除插入的记录
			e := db.DelVideoByID(ctx, int64(v.ID), int64(userID))
			if e != nil {
			}
		}
	}()
	res := &video.PublishActionResponse{
		StatusCode: 0,
		StatusMsg:  "创建记录成功，等待后台上传完成",
	}
	return res, nil
}

func (videosrv *VideoService) PublishList(ctx context.Context, req *video.PublishListRequest) (resp *video.PublishListResponse, err error) {
	userID := req.UserId

	results, err := db.GetVideosByUserID(ctx, userID)
	if err != nil {
		res := &video.PublishListResponse{
			StatusCode: -1,
			StatusMsg:  "发布列表获取失败：服务器内部错误",
		}
		return res, nil
	}
	videos := make([]*video.Video, 0)
	for _, r := range results {
		author, err := db.GetUserById(ctx, int64(r.AuthorID))
		if err != nil {
			res := &video.PublishListResponse{
				StatusCode: -1,
				StatusMsg:  "发布列表获取失败：服务器内部错误",
			}
			return res, nil
		}
		follow, err := db.GetRelationByUserIDs(ctx, userID, int64(author.ID))
		if err != nil {
			res := &video.PublishListResponse{
				StatusCode: -1,
				StatusMsg:  "发布列表获取失败：服务器内部错误",
			}
			return res, nil
		}
		favorite, err := db.GetFavoriteVideoRelationByUserVideoID(ctx, userID, int64(r.ID))
		if err != nil {
			res := &video.PublishListResponse{
				StatusCode: -1,
				StatusMsg:  "发布列表获取失败：服务器内部错误",
			}
			return res, nil
		}
		playUrl, err := minio.GetFileTemporaryURL(minio.VideoBucketName, r.PlayUrl)
		if err != nil {
			res := &video.PublishListResponse{
				StatusCode: -1,
				StatusMsg:  "服务器内部错误：视频获取失败",
			}
			return res, nil
		}
		coverUrl, err := minio.GetFileTemporaryURL(minio.CoverBucketName, r.CoverUrl)
		if err != nil {
			res := &video.PublishListResponse{
				StatusCode: -1,
				StatusMsg:  "服务器内部错误：封面获取失败",
			}
			return res, nil
		}
		avatarUrl, err := minio.GetFileTemporaryURL(minio.AvatarBucketName, author.Avatar)
		if err != nil {
			res := &video.PublishListResponse{
				StatusCode: -1,
				StatusMsg:  "服务器内部错误：发布者头像获取失败",
			}
			return res, nil
		}
		backgroundUrl, err := minio.GetFileTemporaryURL(minio.BackgroundImageBucketName, author.BackgroundImage)
		if err != nil {
			res := &video.PublishListResponse{
				StatusCode: -1,
				StatusMsg:  "背景图获取失败",
			}
			return res, nil
		}

		videos = append(videos, &video.Video{
			Id: int64(r.ID),
			Author: &user.User{
				Id:              int64(author.ID),
				Name:            author.UserName,
				FollowerCount:   int64(author.FollowerCount),
				FollowCount:     int64(author.FollowingCount),
				IsFollow:        follow != nil,
				Avatar:          avatarUrl,
				BackgroundImage: backgroundUrl,
				Signature:       author.Signature,
				TotalFavorited:  int64(author.TotalFavorited),
				WorkCount:       int64(author.WorkCount),
				FavoriteCount:   int64(author.FavoriteCount),
			},
			PlayUrl:       playUrl,
			CoverUrl:      coverUrl,
			FavoriteCount: int64(r.FavoriteCount),
			CommentCount:  int64(r.CommentCount),
			IsFavorite:    favorite != nil,
			Title:         r.Title,
		})
	}

	res := &video.PublishListResponse{
		StatusCode: 0,
		StatusMsg:  "success",
		VideoList:  videos,
	}
	return res, nil
}

// uploadVideo 上传视频至 Minio
func uploadVideo(data []byte, videoTitle string) (string, error) {
	// 将视频数据上传至minio
	reader := bytes.NewReader(data)
	contentType := "application/mp4"

	uploadSize, err := minio.UploadFileByIO(minio.VideoBucketName, videoTitle, reader, int64(len(data)), contentType)
	if err != nil {
		return "", err
	}
	fmt.Println("视频文件大小为：", uploadSize)

	// 获取上传文件的路径
	playUrl, err := minio.GetFileTemporaryURL(minio.VideoBucketName, videoTitle)
	if err != nil {
		return "", err
	}
	return playUrl, nil
}

// uploadCover 截取并上传封面至 Minio
func uploadCover(playUrl string, coverTitle string) error {
	// 截取第一帧并将图像上传至minio
	imgBuffer, err := ffmpeg.GetSnapshotImageBuffer(playUrl, 1)
	if err != nil {
		return err
	}
	var imgByte []byte
	imgBuffer.Write(imgByte)
	contentType := "image/png"

	uploadSize, err := minio.UploadFileByIO(minio.CoverBucketName, coverTitle, imgBuffer, int64(imgBuffer.Len()), contentType)
	if err != nil {
		return err
	}

	// 获取上传文件的路径
	_, err = minio.GetFileTemporaryURL(minio.CoverBucketName, coverTitle)
	if err != nil {
		return err
	}
	fmt.Println("封面文件大小为：", uploadSize)

	return nil
}

// VideoPublish 上传视频并获取封面
func VideoPublish(data []byte, videoTitle string, coverTitle string) error {
	playUrl, err := uploadVideo(data, videoTitle)
	if err != nil {
		return err
	}
	err = uploadCover(playUrl, coverTitle)
	if err != nil {
		return err
	}
	return nil
}