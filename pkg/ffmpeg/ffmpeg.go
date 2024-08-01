package ffmpeg

import (
	"bytes"
	"fmt"
	"github.com/disintegration/imaging"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"os"
)

func GetSnapshot(videoPath, snapshotPath string, frameNum int) (ImagePath string, err error) {
	buf := bytes.NewBuffer(nil)
	err = ffmpeg.Input(videoPath).Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", frameNum)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(buf, os.Stdout).
		Run()

	if err != nil {
		return "", err
	}

	img, err := imaging.Decode(buf)
	if err != nil {
		return "", err
	}

	err = imaging.Save(img, snapshotPath+".png")
	if err != nil {
		return "", err
	}

	imgPath := snapshotPath + ".png"

	return imgPath, nil
}

func GetSnapshotImageBuffer(videoPath string, frameNum int) (*bytes.Buffer, error) {

	buf := bytes.NewBuffer(nil)
	err := ffmpeg.Input(videoPath).Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", frameNum)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(buf, os.Stdout).
		Run()

	if err != nil {
		return nil, err
	}
	return buf, nil
}
