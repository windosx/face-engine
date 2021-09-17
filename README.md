# 虹软人脸检测SDK Go语言版
![LICENSE](https://img.shields.io/github/license/windosx/face-engine)
![Supported platform](https://img.shields.io/badge/platform-win%20%7C%20linux-lightgrey)
[![Build Status](https://travis-ci.org/windosx/face-engine.svg?branch=master)](https://travis-ci.org/windosx/face-engine)
[![GoDoc](http://godoc.org/github.com/windosx/face-engine?status.svg)](http://godoc.org/github.com/windosx/face-engine)

# 一、项目简介

由于公司是虹软的重度用户，出于工作需要和个人兴趣，寻思着用golang封装一下C++的SDK，利用golang的跨平台特性达到跨平台的效果（当然前提是SDK支持的平台）

目前支持的SDK版本有：[v2.x](https://github.com/windosx/face-engine/tree/v2.2.0) [v3.x](https://github.com/windosx/face-engine/tree/v3.0.8) [v4.0.0](https://github.com/windosx/face-engine)

# 二、编译环境与运行环境的准备

## 1. 安装go版本SDK
推荐使用go module的方式进行安装(需要哪个版本修改版本号即可)：

```bash
go get -u -d github.com/windosx/face-engine/v4
```
## 2. 安装gcc环境
Linux下可以通过对应的包管理器进行安装，Windows下根据32/64位系统进行选择安装MinGW或者MinGW-w64

## 3. 运行环境
 - Linux：将`libarcsoft_face.so`、`libarcsoft_face_engine.so`放入`/usr/lib64`目录下
 - Windows：将`libarcsoft_face.dll`、`libarcsoft_face_engine.dll`放入MinGW安装目录下的lib目录中，同时需要将这两个文件放入`%WINDIR%`或者`%WINDIR/System32%`目录下

# 三、代码结构说明
```bash
├─examples
│  └─example1.go
├─include
│  ├─amcomdef.h
│  └─arcsoft_face_sdk.h
│  └─asvloffscreen.h
│  └─merror.h
├─util
│  └─image_util.go
├─go.mod
├─engine.go
├─mask.jpg
├─test.jpg
examples目录中是示例代码
include是SDK的头文件
util下的image_util.go是使用golang原生API处理图片的工具
engine.go是封装的SDK
两张JPEG格式的图片是给示例代码用的
```
# 四、代码调用示例
```go
package main

import (
	"fmt"

	. "github.com/windosx/face-engine/v4"
	"github.com/windosx/face-engine/v4/util"
)

// 获取处理好的图片信息
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
```
# 五、注意事项
1. 通过`FaceFeatureExtract`或`FaceFeatureExtractEx`方法提取到人脸特征信息，使用完毕后应调用其`Release()`方法释放内存，避免导致内存溢出，这是由于cgo的一些局限性导致的。

# 六、更多示例代码
- [ArcFace+gocv处理视频流](https://github.com/windosx/arcface-gocv-examples)

***（如果您想分享自己的示例，欢迎PR）***

# 七、开源说明
本项目使用MIT协议，如果对您有帮助，请点亮star支持一下 :)
