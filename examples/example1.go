package main

import (
	"fmt"

	. "github.com/windosx/face-engine/v4"
	"github.com/windosx/face-engine/v4/util"
)

var imageInfo = util.GetResizedImageInfo("./mask.jpg")

func main() {
	// 激活SDK
	if err := OnlineActivation("YourAppID", "YourSDKKey", "YourActiveCode"); err != nil {
		fmt.Printf("%#v\n", err)
		return
	}
	// 初始化引擎
	engine, err := NewFaceEngine(DetectModeImage,
		OrientPriority0,
		10, // 4.0最大支持10个人脸
		EnableFaceDetect|EnableFaceRecognition|EnableFace3DAngle|EnableLiveness|EnableIRLiveness|EnableAge|EnableGender|EnableMaskDetect|EnableFaceLandMark)
	if err != nil {
		fmt.Printf("%#v\n", err)
		return
	}
	deviceInfo, err := GetActiveDeviceInfo()
	if err != nil {
		fmt.Printf("%#v\n", err)
	}
	fmt.Printf("设备信息：%s\n", deviceInfo)
	// 检测人脸
	info, err := engine.DetectFaces(imageInfo.Width, imageInfo.Height, ColorFormatBGR24, imageInfo.DataUInt8)
	if err != nil {
		fmt.Printf("%#v\n", err)
		return
	}
	// 处理人脸数据
	if err = engine.Process(imageInfo.Width, imageInfo.Height, ColorFormatBGR24, imageInfo.DataUInt8, info, EnableAge|EnableGender|EnableFace3DAngle|EnableLiveness|EnableMaskDetect|EnableFaceLandMark); err != nil {
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
	// 获取口罩信息
	maskInfo, err := engine.GetMask()
	if err != nil {
		fmt.Printf("%#v\n", err)
		return
	}
	fmt.Printf("口罩信息：%#v\n", maskInfo)
	// 获取额头点位
	landMark, err := engine.GetFaceLandMarkInfo()
	if err != nil {
		fmt.Printf("%#v\n", err)
		return
	}
	fmt.Printf("额头点位：%#v\n", landMark)
	// 销毁引擎
	if err = engine.Destroy(); err != nil {
		fmt.Printf("%#v\n", err)
		return
	}
}
