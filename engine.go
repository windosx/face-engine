package face_engine

/*
#cgo CFLAGS	: -I./include
#cgo LDFLAGS: -larcsoft_face_engine
#include <stdlib.h>
#include "merror.h"
#include "asvloffscreen.h"
#include "arcsoft_face_sdk.h"
*/
import "C"

import (
	"fmt"
	"unsafe"
)

// 引擎结构体
type FaceEngine struct {
	handle uintptr
}

// 多人脸信息结构体
type MultiFaceInfo struct {
	FaceRect   []Rect  // 人脸框信息
	FaceOrient []int32 // 输入图像的角度
	FaceNum    int32   // 检测到的人脸个数
	FaceID     []int32 // face ID，IMAGE模式下不返回FaceID
}

// 人脸坐标结构体
type Rect struct {
	Left   int32
	Top    int32
	Right  int32
	Bottom int32
}

// 活体置信度结构体
type LivenessThreshold struct {
	ThresholdModelBGR float32
	ThresholdModelIR  float32
}

// 激活文件信息结构体
type ActiveFileInfo struct {
	StartTime   string //开始时间
	EndTime     string //截止时间
	Platform    string //平台
	SdkType     string //sdk类型
	AppId       string //APPID
	SdkKey      string //SDKKEY
	SdkVersion  string //SDK版本号
	FileVersion string //激活文件版本号
}

// 版本信息结构体
type Version struct {
	Version   string // 版本号
	BuildDate string // 构建日期
	CopyRight string // Copyright
}

// 单人脸信息结构体
type SingleFaceInfo struct {
	FaceRect   Rect  // 人脸框信息
	FaceOrient int32 // 输入图像的角度，可以参考 ArcFaceCompare_OrientCode
}

// 人脸特征结构体
type FaceFeature struct {
	Feature     []byte // 人脸特征信息
	FeatureSize int32  // 人脸特征信息长度
}

// 年龄信息结构体
type AgeInfo struct {
	AgeArray []int32 // "0" 代表不确定，大于0的数值代表检测出来的年龄结果
	Num      int32   // 检测的人脸个数
}

// 性别信息结构体
type GenderInfo struct {
	GenderArray []int32 // "0" 表示 男性, "1" 表示 女性, "-1" 表示不确定
	Num         int32   // 检测的人脸个数
}

// 人脸3D角度信息结构体
type Face3DAngle struct {
	Roll   []float32
	Yaw    []float32
	Pitch  []float32
	Status []int32 // 0: 正常，其他数值：出错
	Num    int32
}

// 活体信息
type LivenessInfo struct {
	IsLive []int32 // 0:非真人 1:真人 -1:不确定 -2:传入人脸数>1
	Num    int32
}

const (
	EnableNone            = C.ASF_NONE
	EnableFaceDetect      = C.ASF_FACE_DETECT
	EnableFaceRecognition = C.ASF_FACERECOGNITION
	EnableAge             = C.ASF_AGE
	EnableGender          = C.ASF_GENDER
	EnableFace3DAngle     = C.ASF_FACE3DANGLE
	EnableLiveness        = C.ASF_LIVENESS
	EnableIRLiveness      = C.ASF_IR_LIVENESS
	ColorFormatBGR24      = C.ASVL_PAF_RGB24_B8G8R8
	ColorFormatNV12       = C.ASVL_PAF_NV12
	ColorFormatNV21       = C.ASVL_PAF_NV21
	ColorFormatI420       = C.ASVL_PAF_I420
	ColorFormatYUYV       = C.ASVL_PAF_YUYV
)

// 创建一个新的引擎实例
//
// 如果调用初始化函数失败则返回一个错误
func NewFaceEngine(dm C.enum_ASF_DetectMode, op C.enum_ASF_OrientPriority, faceScale, maxFaceNum, combinedMask C.MInt32) (*FaceEngine, error) {
	engine, err := &FaceEngine{}, error(nil)
	r := C.ASFInitEngine(dm, op, faceScale, maxFaceNum, combinedMask, (*C.MHandle)(unsafe.Pointer(&engine.handle)))
	if r != C.MOK {
		err = fmt.Errorf("初始化引擎失败!错误码: %d", int(r))
	}
	return engine, err
}

// 获取激活文件信息接口
func GetActiveFileInfo() (ActiveFileInfo, error) {
	asfActiveFileInfo := &C.ASF_ActiveFileInfo{}
	r := C.ASFGetActiveFileInfo(asfActiveFileInfo)
	if r != C.MOK {
		return ActiveFileInfo{}, fmt.Errorf("获取激活文件信息失败!错误码: %d", int(r))
	}
	info := ActiveFileInfo{
		StartTime:   C.GoString(asfActiveFileInfo.startTime),
		EndTime:     C.GoString(asfActiveFileInfo.endTime),
		Platform:    C.GoString(asfActiveFileInfo.platform),
		SdkType:     C.GoString(asfActiveFileInfo.sdkType),
		AppId:       C.GoString(asfActiveFileInfo.appId),
		SdkKey:      C.GoString(asfActiveFileInfo.sdkKey),
		SdkVersion:  C.GoString(asfActiveFileInfo.sdkVersion),
		FileVersion: C.GoString(asfActiveFileInfo.fileVersion),
	}
	return info, nil
}

// 在线激活接口
func OnlineActivation(appId, sdkKey string) (err error) {
	id := C.CString(appId)
	key := C.CString(sdkKey)
	defer func() {
		C.free(unsafe.Pointer(id))
		C.free(unsafe.Pointer(key))
	}()
	r := C.ASFOnlineActivation(id, key)
	if r != C.MOK && r != C.MERR_ASF_ALREADY_ACTIVATED {
		err = fmt.Errorf("激活SDK失败!错误码: %d", int(r))
	}
	return
}

// 在线激活接口，该接口与ASFOnlineActivation接口功能一致，推荐使用该接口
func Activation(appId, sdkKey string) (err error) {
	id := C.CString(appId)
	key := C.CString(sdkKey)
	defer func() {
		C.free(unsafe.Pointer(id))
		C.free(unsafe.Pointer(key))
	}()
	r := C.ASFActivation(id, key)
	if r != C.MOK && r != C.MERR_ASF_ALREADY_ACTIVATED {
		err = fmt.Errorf("激活SDK失败!错误码: %d", int(r))
	}
	return
}

// IMAGE模式:人脸检测
func (engine *FaceEngine) ASFDetectFaces(width, height int, format C.MInt32, imgData []byte) (*C.ASF_MultiFaceInfo, error) {
	faceInfo := &C.ASF_MultiFaceInfo{}
	r := C.ASFDetectFaces(C.MHandle(unsafe.Pointer(engine.handle)),
		C.MInt32(width),
		C.MInt32(height),
		format,
		(*C.MUInt8)(unsafe.Pointer(&imgData[0])),
		faceInfo,
		C.ASF_DETECT_MODEL_RGB,
	)
	if r != C.MOK {
		return faceInfo, fmt.Errorf("人脸检测失败!错误码: %d", int(r))
	}
	return faceInfo, nil
}

// 人脸检测（将返回结果转为Go的结构体类型）,目前不支持IR图像数据检测
func (engine *FaceEngine) DetectFaces(width, height int, format C.MInt32, imgData []byte) (faceInfo MultiFaceInfo, err error) {
	asfFaceInfo, err := engine.ASFDetectFaces(width, height, format, imgData)
	if err != nil {
		return
	}
	faceNum := int32(asfFaceInfo.faceNum)
	faceInfo.FaceRect = (*[50]Rect)(unsafe.Pointer(asfFaceInfo.faceRect))[:faceNum:faceNum]
	faceInfo.FaceOrient = (*[50]int32)(unsafe.Pointer(asfFaceInfo.faceOrient))[:faceNum:faceNum]
	if asfFaceInfo.faceID != nil {
		faceInfo.FaceID = (*[50]int32)(unsafe.Pointer(asfFaceInfo.faceID))[:faceNum:faceNum]
	}
	faceInfo.FaceNum = faceNum
	return
}

// 年龄/性别/人脸3D角度（该接口仅支持RGB图像），最多支持4张人脸信息检测，超过部分返回未知
// RGB活体仅支持单人脸检测，该接口不支持检测IR活体
func (engine *FaceEngine) Process(width, height int, format C.MInt32, imgData []byte, detectedFaces *C.ASF_MultiFaceInfo, combinedMask C.MInt32) error {
	r := C.ASFProcess(C.MHandle(unsafe.Pointer(engine.handle)),
		C.MInt32(width),
		C.MInt32(height),
		format,
		(*C.MUInt8)(unsafe.Pointer(&imgData[0])),
		detectedFaces,
		combinedMask)
	if r != C.MOK {
		return fmt.Errorf("检测人脸信息失败!错误码: %v", int(r))
	}
	return nil
}

// 该接口目前仅支持单人脸IR活体检测（不支持年龄、性别、3D角度的检测），默认取第一张人脸
func (engine *FaceEngine) ProcessIR(width, height int, format C.MInt32, imgData []byte, detectedFaces *C.ASF_MultiFaceInfo, combinedMask C.MInt32) error {
	r := C.ASFProcess(C.MHandle(unsafe.Pointer(engine.handle)),
		C.MInt32(width),
		C.MInt32(height),
		format,
		(*C.MUInt8)(unsafe.Pointer(&imgData[0])),
		detectedFaces,
		combinedMask)
	if r != C.MOK {
		return fmt.Errorf("检测人脸IR活体信息失败!错误码: %v", int(r))
	}
	return nil
}

// 设置活体置信度
//
// 取值范围[0-1]内部默认数值RGB-0.75，IR-0.7， 用户可以根据实际需求，设置不同的阈值
func (engine *FaceEngine) SetLivenessParam(threshold LivenessThreshold) error {
	asfLivenessThreshold := &C.ASF_LivenessThreshold{
		thresholdmodel_BGR: C.MFloat(threshold.ThresholdModelBGR),
		thresholdmodel_IR:  C.MFloat(threshold.ThresholdModelIR),
	}
	r := C.ASFSetLivenessParam(C.MHandle(unsafe.Pointer(engine.handle)), asfLivenessThreshold)
	if r != C.MOK {
		return fmt.Errorf("设置活体置信度失败!错误码: %d", int(r))
	}
	return nil
}

// 获取版本信息
func (engine *FaceEngine) GetVersion() Version {
	info := C.ASFGetVersion()
	return Version{
		Version:   C.GoString(info.Version),
		BuildDate: C.GoString(info.BuildDate),
		CopyRight: C.GoString(info.CopyRight),
	}
}

// 单人脸特征提取
func (engine *FaceEngine) ASFFaceFeatureExtract(width, height int, format C.MInt32, imgData []byte, faceInfo *SingleFaceInfo) (asfFaceFeature *C.ASF_FaceFeature, err error) {
	asfFaceFeature = &C.ASF_FaceFeature{}
	asfFaceInfo := &C.ASF_SingleFaceInfo{
		C.MRECT{
			C.MInt32(faceInfo.FaceRect.Left),
			C.MInt32(faceInfo.FaceRect.Top),
			C.MInt32(faceInfo.FaceRect.Right),
			C.MInt32(faceInfo.FaceRect.Bottom)},
		C.MInt32(faceInfo.FaceOrient)}
	r := C.ASFFaceFeatureExtract(
		C.MHandle(unsafe.Pointer(engine.handle)),
		C.MInt32(width),
		C.MInt32(height),
		format,
		(*C.MUInt8)(unsafe.Pointer(&imgData[0])),
		asfFaceInfo,
		asfFaceFeature)
	if r != C.MOK {
		err = fmt.Errorf("提取人脸特征失败!错误码: %d", int(r))
	}
	length := int32(asfFaceFeature.featureSize)
	byteArr := (*[1 << 28]byte)(unsafe.Pointer(asfFaceFeature.feature))[:length:length]
	arr := (*C.MByte)(C.malloc(C.size_t(int32(asfFaceFeature.featureSize))))
	ps := (*[1 << 28]C.MByte)(unsafe.Pointer(arr))[:length:length]
	for i := 0; i < len(ps); i++ {
		ps[i] = C.MByte(byteArr[i])
	}
	asfFaceFeature.feature = arr
	return asfFaceFeature, err
}

// 人脸特征比对
func (engine *FaceEngine) FaceFeatureCompare(feature1, feature2 *C.ASF_FaceFeature) (float32, error) {
	var confidenceLevel float32 = 0
	r := C.ASFFaceFeatureCompare(C.MHandle(unsafe.Pointer(engine.handle)),
		feature1,
		feature2,
		(*C.MFloat)(unsafe.Pointer(&confidenceLevel)),
		C.ASF_DETECT_MODEL_RGB,
	)
	if r != C.MOK {
		return 0, fmt.Errorf("人脸特征比对失败!错误码:%v", int(r))
	}
	return confidenceLevel, nil
}

// 获取年龄信息
func (engine *FaceEngine) GetAge() (AgeInfo, error) {
	asfAgeInfo := &C.ASF_AgeInfo{}
	r := C.ASFGetAge((C.MHandle)(unsafe.Pointer(engine.handle)), asfAgeInfo)
	if r != C.MOK {
		return AgeInfo{}, fmt.Errorf("获取年龄信息失败!错误码: %d", int(r))
	}
	num := int32(asfAgeInfo.num)
	return AgeInfo{
		AgeArray: (*[50]int32)(unsafe.Pointer(asfAgeInfo.ageArray))[:num:num],
		Num:      num,
	}, nil
}

// 获取性别信息
func (engine *FaceEngine) GetGender() (GenderInfo, error) {
	asfGenderInfo := &C.ASF_GenderInfo{}
	r := C.ASFGetGender((C.MHandle)(unsafe.Pointer(engine.handle)), asfGenderInfo)
	if r != C.MOK {
		return GenderInfo{}, fmt.Errorf("获取性别信息失败!错误码: %d", int(r))
	}
	num := int32(asfGenderInfo.num)
	return GenderInfo{
		GenderArray: (*[50]int32)(unsafe.Pointer(asfGenderInfo.genderArray))[:num:num],
		Num:         num,
	}, nil
}

// 获取3D角度信息
func (engine *FaceEngine) GetFace3DAngle() (Face3DAngle, error) {
	asfFace3DAngle := &C.ASF_Face3DAngle{}
	r := C.ASFGetFace3DAngle((C.MHandle)(unsafe.Pointer(engine.handle)), asfFace3DAngle)
	if r != C.MOK {
		return Face3DAngle{}, fmt.Errorf("获取3D角度信息失败!错误码: %d", int(r))
	}
	num := int32(asfFace3DAngle.num)
	return Face3DAngle{
		Roll:   (*[50]float32)(unsafe.Pointer(asfFace3DAngle.roll))[:num:num],
		Yaw:    (*[50]float32)(unsafe.Pointer(asfFace3DAngle.yaw))[:num:num],
		Pitch:  (*[50]float32)(unsafe.Pointer(asfFace3DAngle.pitch))[:num:num],
		Status: (*[50]int32)(unsafe.Pointer(asfFace3DAngle.status))[:num:num],
		Num:    int32(asfFace3DAngle.num),
	}, nil
}

// 获取RGB活体结果
func (engine *FaceEngine) GetLivenessScore() (LivenessInfo, error) {
	asfLivenessInfo := &C.ASF_LivenessInfo{}
	r := C.ASFGetLivenessScore((C.MHandle)(unsafe.Pointer(engine.handle)), asfLivenessInfo)
	if r != C.MOK {
		return LivenessInfo{}, fmt.Errorf("获取活体信息失败!错误码: %d", int(r))
	}
	num := int32(asfLivenessInfo.num)
	return LivenessInfo{
		IsLive: (*[50]int32)(unsafe.Pointer(asfLivenessInfo.isLive))[:num:num],
		Num:    num,
	}, nil
}

// 获取IR活体结果
func (engine *FaceEngine) GetLivenessScoreIR() (LivenessInfo, error) {
	asfLivenessInfo := &C.ASF_LivenessInfo{}
	r := C.ASFGetLivenessScore_IR((C.MHandle)(unsafe.Pointer(engine.handle)), asfLivenessInfo)
	if r != C.MOK {
		return LivenessInfo{}, fmt.Errorf("获取活体信息失败!错误码: %d", int(r))
	}
	num := int32(asfLivenessInfo.num)
	return LivenessInfo{
		IsLive: (*[50]int32)(unsafe.Pointer(asfLivenessInfo.isLive))[:num:num],
		Num:    num,
	}, nil
}

// 销毁引擎
func (engine *FaceEngine) Destroy() error {
	r := C.ASFUninitEngine(C.MHandle(unsafe.Pointer(engine.handle)))
	if r != C.MOK {
		return fmt.Errorf("销毁引擎失败!错误码: %d", int(r))
	}
	return nil
}
