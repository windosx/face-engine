package main

import (
	"fmt"
	fe "github.com/windosx/face-engine/v2"
	"github.com/windosx/face-engine/util"
)

var width, height = util.GetImageWidthAndHeight("./test.jpg")
var imageData = util.GetResizedBGR("./test.jpg")

func main() {
	// 激活SDK
	if err := fe.Activation("YourAppId", "YourSDKKey"); err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	// 初始化引擎
	engine, err := fe.NewFaceEngine(fe.DetectModeImage,
		fe.OrientPriority0,
		12,
		50,
		fe.EnableFaceDetect | fe.EnableFaceRecognition | fe.EnableFace3DAngle | fe.EnableLiveness | fe.EnableIRLiveness | fe.EnableAge | fe.EnableGender)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	// 检测人脸
	info, err := engine.ASFDetectFaces(width - width % 4, height, fe.ColorFormatBGR24, imageData)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	// 处理人脸数据
	if err = engine.Process(width - width % 4, height, fe.ColorFormatBGR24, imageData, info, fe.EnableAge | fe.EnableGender | fe.EnableFace3DAngle | fe.EnableLiveness); err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	// 获取年龄
	ageInfo, err := engine.GetAge()
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	fmt.Printf("ageInfo: %v\n", ageInfo)
	// 销毁引擎
	if err = engine.Destroy(); err != nil {
		fmt.Printf("%v\n", err)
		return
	}
}
