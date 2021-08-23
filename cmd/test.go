package main

import (
	"fmt"

	. "github.com/windosx/face-engine/v3"
	"github.com/windosx/face-engine/v3/util"
)

var imageInfo = util.GetResizedImageInfo("./test.jpg")

func main() {
	// 激活SDK
	if err := Activation("Your App ID", "Your SDK key"); err != nil {
		fmt.Printf("%#v\n", err)
		return
	}
	// 初始化引擎
	engine, err := NewFaceEngine(DetectModeImage,
		OrientPriority0,
		12,
		50,
		EnableFaceDetect|EnableFaceRecognition|EnableFace3DAngle|EnableLiveness|EnableIRLiveness|EnableAge|EnableGender)
	if err != nil {
		fmt.Printf("%#v\n", err)
		return
	}
	// 检测人脸
	info, err := engine.DetectFaces(imageInfo.Width, imageInfo.Height, ColorFormatBGR24, imageInfo.DataUInt8)
	if err != nil {
		fmt.Printf("%#v\n", err)
		return
	}
	// 获取单人脸信息
	singleFaceInfoArr := GetSingleFaceInfo(info)
	// 获取人脸特征码
	f1, err := engine.FaceFeatureExtract(imageInfo.Width, imageInfo.Height, ColorFormatBGR24, imageInfo.DataUInt8, singleFaceInfoArr[0])
	if err != nil {
		fmt.Printf("%#v\n", err)
		return
	}
	f2, err := engine.FaceFeatureExtract(imageInfo.Width, imageInfo.Height, ColorFormatBGR24, imageInfo.DataUInt8, singleFaceInfoArr[1])
	if err != nil {
		fmt.Printf("%#v\n", err)
		return
	}
	// 人脸比对
	confidence, err := engine.FaceFeatureCompare(f1, f2)
	if err != nil {
		fmt.Printf("%#v\n", err)
		return
	}
	fmt.Printf("人脸相似度为：%.2f\n", confidence)
	f1.Release()
	f2.Release()
	// 处理人脸数据
	if err = engine.Process(imageInfo.Width, imageInfo.Height, ColorFormatBGR24, imageInfo.DataUInt8, info, EnableAge|EnableGender|EnableFace3DAngle|EnableLiveness); err != nil {
		fmt.Printf("%#v\n", err)
		return
	}
	// 获取年龄
	ageInfo, err := engine.GetAge()
	if err != nil {
		fmt.Printf("%#v\n", err)
		return
	}
	fmt.Printf("ageInfo: %v\n", ageInfo)
	// 销毁引擎
	if err = engine.Destroy(); err != nil {
		fmt.Printf("%#v\n", err)
		return
	}
}
