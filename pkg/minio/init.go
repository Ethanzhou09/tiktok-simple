package minio

import (
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"log"
	"tiktok-simple/pkg/viper"
)

var (
	MinioClient               *minio.Client
	ExpireTime                int
	AvatarBucketName          string
	BackgroundImageBucketName string
)

func Init() {
	viper.InitMinioCfg()
	miniocfg := viper.GetMinioCfg()
	// 初始化minio client对象。
	s3client, err := minio.New(miniocfg.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(miniocfg.AccessKeyId, miniocfg.SecretAccessKey, ""),
		Secure: miniocfg.UseSSL,
	})
	if err != nil {
		log.Fatalln(err)
	}
	MinioClient = s3client
	ExpireTime = miniocfg.ExpireTime
	AvatarBucketName = miniocfg.AvatarBucketName
	BackgroundImageBucketName = miniocfg.BackgroundImageBucketName
}
