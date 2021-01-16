package faceengine

/*
#cgo CFLAGS		: -I./include
#cgo LDFLAGS	: -larcsoft_face_engine
#include <stdlib.h>
#include "merror.h"
#include "asvloffscreen.h"
#include "arcsoft_face_sdk.h"
*/
import "C"

import (
	"unsafe"
)

// FaceEngine 引擎结构体
type FaceEngine struct {
	handle C.MHandle
}

// MultiFaceInfo 多人脸信息结构体
type MultiFaceInfo struct {
	FaceRect         []Rect               // 人脸框信息
	FaceOrient       []int32              // 输入图像的角度
	FaceNum          int32                // 检测到的人脸个数
	FaceID           []int32              // face ID，IMAGE模式下不返回FaceID
	WearGlasses      float32              // 带眼镜置信度[0-1]，推荐阈值0.5
	LeftEyeClosed    int32                // 左眼状态 0 未闭眼 1 闭眼
	RightEyeClosed   int32                // 右眼状态 0 未闭眼 1 闭眼
	FaceDataInfoList []C.ASF_FaceDataInfo // 多张人脸信息
	native           *C.ASF_MultiFaceInfo
}

// Rect 人脸坐标结构体
type Rect struct {
	Left   int32
	Top    int32
	Right  int32
	Bottom int32
}

// LivenessThreshold 活体置信度结构体
type LivenessThreshold struct {
	ThresholdModelBGR float32
	ThresholdModelIR  float32
}

// ActiveFileInfo 激活文件信息结构体
type ActiveFileInfo struct {
	StartTime   string //开始时间
	EndTime     string //截止时间
	ActiveKey   string //激活码
	Platform    string //平台
	SdkType     string //sdk类型
	AppID       string //APPID
	SdkKey      string //SDKKEY
	SdkVersion  string //SDK版本号
	FileVersion string //激活文件版本号
}

// Version 版本信息结构体
type Version struct {
	Version   string // 版本号
	BuildDate string // 构建日期
	CopyRight string // Copyright
}

// SingleFaceInfo 单人脸信息结构体
type SingleFaceInfo struct {
	FaceRect   Rect  // 人脸框信息
	FaceOrient int32 // 输入图像的角度，可以参考 ArcFaceCompare_OrientCode
	DataInfo   C.ASF_FaceDataInfo
}

// FaceFeature 人脸特征结构体
type FaceFeature struct {
	Feature     []byte // 人脸特征信息
	FeatureSize int32  // 人脸特征信息长度
	native      *C.ASF_FaceFeature
	featurePtr  *C.MByte
}

// ImageData 对应ASVLOFFSCREEN结构体
type ImageData struct {
	PixelArrayFormat C.MUInt32
	Width            int
	Height           int
	ImageData        [4][]uint8
	WidthStep        [4]int
}

// AgeInfo 年龄信息结构体
type AgeInfo struct {
	AgeArray []int32 // "0" 代表不确定，大于0的数值代表检测出来的年龄结果
	Num      int32   // 检测的人脸个数
}

// GenderInfo 性别信息结构体
type GenderInfo struct {
	GenderArray []int32 // "0" 表示 男性, "1" 表示 女性, "-1" 表示不确定
	Num         int32   // 检测的人脸个数
}

// Face3DAngle 人脸3D角度信息结构体
type Face3DAngle struct {
	Roll   []float32
	Yaw    []float32
	Pitch  []float32
	Status []int32 // 0: 正常，其他数值：出错
	Num    int32
}

// LivenessInfo 活体信息
type LivenessInfo struct {
	IsLive []int32 // 0:非真人 1:真人 -1:不确定 -2:传入人脸数>1
	Num    int32
}

// FaceLandMark 特征点信息结构体
type FaceLandMark struct {
	PosX float32
	PosY float32
}

// LandMarkInfo 额头区域点位
type LandMarkInfo struct {
	Point []FaceLandMark
	Num   int32
}

// MaskInfo 口罩信息结构体
type MaskInfo struct {
	MaskArray []int32 // 0代表没有带口罩 1代表带口罩 -1表示不确定
	Num       int32   // 检测的人脸个数
}

// EngineError SDK错误码
type EngineError struct {
	Code int
	Text string
}

const (
	// DetectModeVideo 视频模式
	DetectModeVideo = C.ASF_DETECT_MODE_VIDEO
	// DetectModeImage 图片模式
	DetectModeImage = C.ASF_DETECT_MODE_IMAGE
	// OrientPriority0 不旋转
	OrientPriority0 = C.ASF_OP_0_ONLY
	// OrientPriority90 旋转90度
	OrientPriority90 = C.ASF_OP_90_ONLY
	// OrientPriority270 旋转270度
	OrientPriority270 = C.ASF_OP_270_ONLY
	// OrientPriority180 旋转180度
	OrientPriority180 = C.ASF_OP_180_ONLY
	// OrientPriorityAllOut 角度不限
	OrientPriorityAllOut = C.ASF_OP_ALL_OUT
	// EnableNone 不开启任何功能
	EnableNone = C.ASF_NONE
	// EnableFaceDetect 开启人脸检测
	EnableFaceDetect = C.ASF_FACE_DETECT
	// EnableFaceRecognition 开启人脸识别
	EnableFaceRecognition = C.ASF_FACERECOGNITION
	// EnableAge 开启年龄检测
	EnableAge = C.ASF_AGE
	// EnableGender 开启性别检测
	EnableGender = C.ASF_GENDER
	// EnableFace3DAngle 开启人脸3D角度检测
	EnableFace3DAngle = C.ASF_FACE3DANGLE
	// EnableLiveness 开启活体检测
	EnableLiveness = C.ASF_LIVENESS
	// EnableIRLiveness 开启IR活体检测
	EnableIRLiveness = C.ASF_IR_LIVENESS
	// EnableFaceLandMark 开启人脸特征点检测
	EnableFaceLandMark = C.ASF_FACELANDMARK
	// EnableImageQuality 开启单人脸图片质量检测
	EnableImageQuality = C.ASF_IMAGEQUALITY
	// EnableFaceShelter 开启遮挡检测
	EnableFaceShelter = C.ASF_FACESHELTER
	// EnableMaskDetect 开启口罩检测
	EnableMaskDetect = C.ASF_MASKDETECT
	// EnableUpdateFaceData 开启人脸数据更新
	EnableUpdateFaceData = C.ASF_UPDATE_FACEDATA
	// ColorFormatBGR24 BGR24格式
	ColorFormatBGR24 = C.ASVL_PAF_RGB24_B8G8R8
	// ColorFormatNV12 NV12格式
	ColorFormatNV12 = C.ASVL_PAF_NV12
	// ColorFormatNV21 NV21格式
	ColorFormatNV21 = C.ASVL_PAF_NV21
	// ColorFormatI420 I420格式
	ColorFormatI420 = C.ASVL_PAF_I420
	// ColorFormatYUYV YUYV格式
	ColorFormatYUYV = C.ASVL_PAF_YUYV
	// RecognitionPhoto 识别照片
	RecognitionPhoto = C.ASF_RECOGNITION
	// RegisterPhoto 注册照片
	RegisterPhoto = C.ASF_REGISTER
)

// NewFaceEngine 创建一个新的引擎实例
//
// 如果调用初始化函数失败则返回一个错误
func NewFaceEngine(
	detectMode C.ASF_DetectMode, // 检测模式
	orientPriority C.ASF_OrientPriority, // 检测角度
	maxFaceNum C.MInt32, // 最大人脸数[1-10]
	combinedMask C.MInt32, // 检测选项
) (*FaceEngine, error) {
	engine, err := &FaceEngine{}, error(nil)
	r := C.ASFInitEngine(detectMode, orientPriority, maxFaceNum, combinedMask, &engine.handle)
	if r != C.MOK {
		err = newError(int(r), "初始化引擎失败")
	}
	return engine, err
}

// GetActiveFileInfo 获取激活文件信息接口
func GetActiveFileInfo() (ActiveFileInfo, error) {
	asfActiveFileInfo := &C.ASF_ActiveFileInfo{}
	r := C.ASFGetActiveFileInfo(asfActiveFileInfo)
	if r != C.MOK {
		return ActiveFileInfo{}, newError(int(r), "获取激活文件信息失败")
	}
	info := ActiveFileInfo{
		StartTime:   C.GoString(asfActiveFileInfo.startTime),
		EndTime:     C.GoString(asfActiveFileInfo.endTime),
		Platform:    C.GoString(asfActiveFileInfo.platform),
		SdkType:     C.GoString(asfActiveFileInfo.sdkType),
		AppID:       C.GoString(asfActiveFileInfo.appId),
		SdkKey:      C.GoString(asfActiveFileInfo.sdkKey),
		SdkVersion:  C.GoString(asfActiveFileInfo.sdkVersion),
		FileVersion: C.GoString(asfActiveFileInfo.fileVersion),
	}
	return info, nil
}

// OnlineActivation 在线激活接口
func OnlineActivation(appID, sdkKey, activeKey string) (err error) {
	id := C.CString(appID)
	sk := C.CString(sdkKey)
	ak := C.CString(activeKey)
	defer func() {
		C.free(unsafe.Pointer(id))
		C.free(unsafe.Pointer(sk))
		C.free(unsafe.Pointer(ak))
	}()
	r := C.ASFOnlineActivation(id, sk, ak)
	if r != C.MOK && r != C.MERR_ASF_ALREADY_ACTIVATED {
		err = newError(int(r), "激活SDK失败")
	}
	return
}

// OfflineActivation 离线激活接口
func OfflineActivation(filePath string) (err error) {
	cFilePath := C.CString(filePath)
	defer C.free(unsafe.Pointer(cFilePath))
	r := C.ASFOfflineActivation(cFilePath)
	if r != C.MOK {
		err = newError(int(r), "离线激活失败")
	}
	return
}

// GetActiveDeviceInfo 采集当前设备信息（可离线）
func GetActiveDeviceInfo() ([]byte, error) {
	deviceInfo := C.CString("")
	defer C.free(unsafe.Pointer(deviceInfo))
	r := C.ASFGetActiveDeviceInfo((*C.MPChar)(unsafe.Pointer(&deviceInfo)))
	if r != C.MOK {
		return nil, newError(int(r), "采集当前设备信息失败")
	}
	str := C.GoString(deviceInfo)
	b := make([]byte, len(str))
	copy(b, str)
	return b, nil
}

// DetectFaces 人脸检测，目前不支持IR图像数据检测
func (engine *FaceEngine) DetectFaces(
	width int, // 宽度
	height int, // 高度
	format C.MInt32, // 图像格式
	imgData []byte, // 图片数据
) (faceInfo MultiFaceInfo, err error) {
	asfFaceInfo := &C.ASF_MultiFaceInfo{}
	r := C.ASFDetectFaces(engine.handle,
		C.MInt32(width),
		C.MInt32(height),
		format,
		(*C.MUInt8)(unsafe.Pointer(&imgData[0])),
		asfFaceInfo,
		C.ASF_DETECT_MODEL_RGB,
	)
	if r != C.MOK {
		return faceInfo, newError(int(r), "人脸检测失败")
	}
	faceNum := int32(asfFaceInfo.faceNum)
	faceInfo.FaceNum = faceNum
	if faceNum > 0 {
		faceInfo.FaceRect = (*[10]Rect)(unsafe.Pointer(asfFaceInfo.faceRect))[:faceNum:faceNum]
		faceInfo.FaceOrient = (*[10]int32)(unsafe.Pointer(asfFaceInfo.faceOrient))[:faceNum:faceNum]
	}
	if asfFaceInfo.faceID != nil {
		faceInfo.FaceID = (*[10]int32)(unsafe.Pointer(asfFaceInfo.faceID))[:faceNum:faceNum]
	}
	if asfFaceInfo.faceDataInfoList != nil {
		faceInfo.FaceDataInfoList = (*[10]C.ASF_FaceDataInfo)(unsafe.Pointer(asfFaceInfo.faceDataInfoList))[:faceNum:faceNum]
	}
	faceInfo.native = asfFaceInfo
	return
}

// DetectFacesEx 检测人脸信息
//
// 该接口与 DetectFaces 功能一致，但采用结构体的形式传入图像数据，对更高精度的图像兼容性更好。
func (engine *FaceEngine) DetectFacesEx(imageData ImageData) (faceInfo MultiFaceInfo, err error) {
	asfFaceInfo := &C.ASF_MultiFaceInfo{}
	r := C.ASFDetectFacesEx(engine.handle, imageDataToASVLOFFSCREEN(imageData), asfFaceInfo, C.ASF_DETECT_MODEL_RGB)
	if r != C.MOK {
		return faceInfo, newError(int(r), "人脸检测失败")
	}
	faceNum := int32(asfFaceInfo.faceNum)
	faceInfo.FaceNum = faceNum
	if faceNum > 0 {
		faceInfo.FaceRect = (*[10]Rect)(unsafe.Pointer(asfFaceInfo.faceRect))[:faceNum:faceNum]
		faceInfo.FaceOrient = (*[10]int32)(unsafe.Pointer(asfFaceInfo.faceOrient))[:faceNum:faceNum]
	}
	if asfFaceInfo.faceID != nil {
		faceInfo.FaceID = (*[10]int32)(unsafe.Pointer(asfFaceInfo.faceID))[:faceNum:faceNum]
	}
	if asfFaceInfo.faceDataInfoList != nil {
		faceInfo.FaceDataInfoList = (*[10]C.ASF_FaceDataInfo)(unsafe.Pointer(asfFaceInfo.faceDataInfoList))[:faceNum:faceNum]
	}
	faceInfo.native = asfFaceInfo
	return
}

// Process 年龄/性别/人脸3D角度（该接口仅支持RGB图像），最多支持4张人脸信息检测，超过部分返回未知
// RGB活体仅支持单人脸检测，该接口不支持检测IR活体
func (engine *FaceEngine) Process(
	width int, // 宽度
	height int, // 高度
	format C.MInt32, // 图像格式
	imgData []byte, // 图片数据
	detectedFaces MultiFaceInfo, // 多人脸信息
	combinedMask C.MInt32, // 检测选项
) error {
	r := C.ASFProcess(engine.handle,
		C.MInt32(width),
		C.MInt32(height),
		format,
		(*C.MUInt8)(unsafe.Pointer(&imgData[0])),
		detectedFaces.native,
		combinedMask)
	if r != C.MOK {
		return newError(int(r), "检测人脸信息失败")
	}
	return nil
}

// ProcessEx 人脸信息检测（年龄/性别/人脸3D角度），最多支持4张人脸信息检测，超过部分返回未知（活体仅支持单张人脸检测，超出返回未知），接口仅支持可见光图像检测
//
// 该接口与 Process 功能一致，但采用结构体的形式传入图像数据，对更高精度的图像兼容性更好
func (engine *FaceEngine) ProcessEx(imageData ImageData, faceInfo MultiFaceInfo, combinedMask C.MInt32) error {
	r := C.ASFProcessEx(engine.handle, imageDataToASVLOFFSCREEN(imageData), faceInfo.native, combinedMask)
	if r != C.MOK {
		return newError(int(r), "检测人脸信息失败")
	}
	return nil
}

// ProcessIR 该接口目前仅支持单人脸IR活体检测（不支持年龄、性别、3D角度的检测），默认取第一张人脸
func (engine *FaceEngine) ProcessIR(
	width int, // 宽度
	height int, // 高度
	format C.MInt32, // 图像格式
	imgData []byte, // 图片数据
	detectedFaces MultiFaceInfo, // 多人脸信息
	combinedMask C.MInt32, // 检测选项
) (err error) {
	r := C.ASFProcess(engine.handle,
		C.MInt32(width),
		C.MInt32(height),
		format,
		(*C.MUInt8)(unsafe.Pointer(&imgData[0])),
		detectedFaces.native,
		combinedMask)
	if r != C.MOK {
		return newError(int(r), "检测人脸IR活体信息失败")
	}
	return nil
}

// ProcessExIR 该接口目前仅支持单人脸IR活体检测（不支持年龄、性别、3D角度的检测），默认取第一张人脸
//
// 该接口与 ProcessIR 功能一致，但采用结构体的形式传入图像数据，对更高精度的图像兼容性更好
func (engine *FaceEngine) ProcessExIR(
	imageData ImageData, // 图像数据
	faceInfo MultiFaceInfo, // 多人脸信息
	combinedMask C.MInt32, // 检测选项
) (err error) {
	r := C.ASFProcessEx_IR(engine.handle, imageDataToASVLOFFSCREEN(imageData), faceInfo.native, combinedMask)
	if r != C.MOK {
		return newError(int(r), "检测人脸IR活体信息失败")
	}
	return nil
}

// SetLivenessParam 设置活体置信度
//
// 取值范围[0-1]内部默认数值RGB-0.75，IR-0.7， 用户可以根据实际需求，设置不同的阈值
func (engine *FaceEngine) SetLivenessParam(threshold LivenessThreshold) error {
	asfLivenessThreshold := &C.ASF_LivenessThreshold{
		thresholdmodel_BGR: C.MFloat(threshold.ThresholdModelBGR),
		thresholdmodel_IR:  C.MFloat(threshold.ThresholdModelIR),
	}
	r := C.ASFSetLivenessParam(engine.handle, asfLivenessThreshold)
	if r != C.MOK {
		return newError(int(r), "设置活体置信度失败")
	}
	return nil
}

// GetVersion 获取版本信息
func (engine *FaceEngine) GetVersion() Version {
	info := C.ASFGetVersion()
	return Version{
		Version:   C.GoString(info.Version),
		BuildDate: C.GoString(info.BuildDate),
		CopyRight: C.GoString(info.CopyRight),
	}
}

// FaceFeatureExtract 单人脸特征提取
func (engine *FaceEngine) FaceFeatureExtract(
	width int, // 宽度
	height int, // 高度
	format C.MInt32, // 图像格式
	imgData []byte, // 图片数据
	faceInfo SingleFaceInfo, // 单人脸信息
	registerOrNot int, // 图片模式
	mask int, // 是否带口罩
) (faceFeature FaceFeature, err error) {
	asfFaceFeature := &C.ASF_FaceFeature{}
	asfFaceInfo := &C.ASF_SingleFaceInfo{
		C.MRECT{
			C.MInt32(faceInfo.FaceRect.Left),
			C.MInt32(faceInfo.FaceRect.Top),
			C.MInt32(faceInfo.FaceRect.Right),
			C.MInt32(faceInfo.FaceRect.Bottom)},
		C.MInt32(faceInfo.FaceOrient),
		faceInfo.DataInfo}
	r := C.ASFFaceFeatureExtract(
		engine.handle,
		C.MInt32(width),
		C.MInt32(height),
		format,
		(*C.MUInt8)(unsafe.Pointer(&imgData[0])),
		asfFaceInfo,
		C.ASF_RegisterOrNot(registerOrNot),
		C.MInt32(mask),
		asfFaceFeature)
	if r != C.MOK {
		return FaceFeature{}, newError(int(r), "提取人脸特征失败")
	}
	length := int32(asfFaceFeature.featureSize)
	faceFeature.FeatureSize = length
	faceFeature.Feature = make([]byte, length)
	byteArr := (*[1 << 12]byte)(unsafe.Pointer(asfFaceFeature.feature))[:length:length]
	arr := (*C.MByte)(C.malloc(C.size_t(int32(asfFaceFeature.featureSize))))
	faceFeature.featurePtr = arr
	ps := (*[1 << 12]C.MByte)(unsafe.Pointer(arr))[:length:length]
	for i := 0; i < len(ps); i++ {
		ps[i] = C.MByte(byteArr[i])
		faceFeature.Feature[i] = byteArr[i]
	}
	asfFaceFeature.feature = arr
	faceFeature.native = asfFaceFeature
	return faceFeature, err
}

// FaceFeatureExtractEx 单人脸特征提取
//
// 该接口与 ASFFaceFeatureExtract 功能一致，但采用结构体的形式传入图像数据，对更高精度的图像兼容性更好
func (engine *FaceEngine) FaceFeatureExtractEx(
	imageData ImageData, // 图片数据
	faceInfo SingleFaceInfo, // 单人脸信息
	registerOrNot int, // 图片模式
	mask int, // 是否带口罩
) (feature FaceFeature, err error) {
	asfFaceFeature := &C.ASF_FaceFeature{}
	asfFaceInfo := &C.ASF_SingleFaceInfo{
		C.MRECT{
			C.MInt32(faceInfo.FaceRect.Left),
			C.MInt32(faceInfo.FaceRect.Top),
			C.MInt32(faceInfo.FaceRect.Right),
			C.MInt32(faceInfo.FaceRect.Bottom)},
		C.MInt32(faceInfo.FaceOrient),
		faceInfo.DataInfo}
	r := C.ASFFaceFeatureExtractEx(engine.handle, imageDataToASVLOFFSCREEN(imageData), asfFaceInfo, C.ASF_RegisterOrNot(registerOrNot), C.MInt32(mask), asfFaceFeature)
	if r != C.MOK {
		return FaceFeature{}, newError(int(r), "提取人脸特征失败")
	}
	length := int32(asfFaceFeature.featureSize)
	feature.FeatureSize = length
	feature.Feature = make([]byte, length)
	byteArr := (*[1 << 12]byte)(unsafe.Pointer(asfFaceFeature.feature))[:length:length]
	arr := (*C.MByte)(C.malloc(C.size_t(int32(asfFaceFeature.featureSize))))
	feature.featurePtr = arr
	ps := (*[1 << 12]C.MByte)(unsafe.Pointer(arr))[:length:length]
	for i := 0; i < len(ps); i++ {
		ps[i] = C.MByte(byteArr[i])
		feature.Feature[i] = byteArr[i]
	}
	asfFaceFeature.feature = arr
	feature.native = asfFaceFeature
	return
}

// FaceFeatureCompare 人脸特征比对
func (engine *FaceEngine) FaceFeatureCompare(feature1, feature2 FaceFeature) (confidenceLevel float32, err error) {
	r := C.ASFFaceFeatureCompare(engine.handle,
		feature1.native,
		feature2.native,
		(*C.MFloat)(unsafe.Pointer(&confidenceLevel)),
		C.ASF_DETECT_MODEL_RGB,
	)
	if r != C.MOK {
		err = newError(int(r), "人脸特征比对失败!")
	}
	return
}

// GetAge 获取年龄信息
func (engine *FaceEngine) GetAge() (AgeInfo, error) {
	asfAgeInfo := &C.ASF_AgeInfo{}
	r := C.ASFGetAge((C.MHandle)(engine.handle), asfAgeInfo)
	if r != C.MOK {
		return AgeInfo{}, newError(int(r), "获取年龄信息失败")
	}
	num := int32(asfAgeInfo.num)
	return AgeInfo{
		AgeArray: (*[10]int32)(unsafe.Pointer(asfAgeInfo.ageArray))[:num:num],
		Num:      num,
	}, nil
}

// GetGender 获取性别信息
func (engine *FaceEngine) GetGender() (GenderInfo, error) {
	asfGenderInfo := &C.ASF_GenderInfo{}
	r := C.ASFGetGender((C.MHandle)(engine.handle), asfGenderInfo)
	if r != C.MOK {
		return GenderInfo{}, newError(int(r), "获取性别信息失败")
	}
	num := int32(asfGenderInfo.num)
	return GenderInfo{
		GenderArray: (*[10]int32)(unsafe.Pointer(asfGenderInfo.genderArray))[:num:num],
		Num:         num,
	}, nil
}

// GetFace3DAngle 获取3D角度信息
func (engine *FaceEngine) GetFace3DAngle() (Face3DAngle, error) {
	asfFace3DAngle := &C.ASF_Face3DAngle{}
	r := C.ASFGetFace3DAngle((C.MHandle)(engine.handle), asfFace3DAngle)
	if r != C.MOK {
		return Face3DAngle{}, newError(int(r), "获取3D角度信息失败")
	}
	num := int32(asfFace3DAngle.num)
	return Face3DAngle{
		Roll:   (*[10]float32)(unsafe.Pointer(asfFace3DAngle.roll))[:num:num],
		Yaw:    (*[10]float32)(unsafe.Pointer(asfFace3DAngle.yaw))[:num:num],
		Pitch:  (*[10]float32)(unsafe.Pointer(asfFace3DAngle.pitch))[:num:num],
		Status: (*[10]int32)(unsafe.Pointer(asfFace3DAngle.status))[:num:num],
		Num:    int32(asfFace3DAngle.num),
	}, nil
}

// GetLivenessScore 获取RGB活体结果
func (engine *FaceEngine) GetLivenessScore() (LivenessInfo, error) {
	asfLivenessInfo := &C.ASF_LivenessInfo{}
	r := C.ASFGetLivenessScore((C.MHandle)(engine.handle), asfLivenessInfo)
	if r != C.MOK {
		return LivenessInfo{}, newError(int(r), "获取活体信息失败")
	}
	num := int32(asfLivenessInfo.num)
	return LivenessInfo{
		IsLive: (*[10]int32)(unsafe.Pointer(asfLivenessInfo.isLive))[:num:num],
		Num:    num,
	}, nil
}

// GetLivenessScoreIR 获取IR活体结果
func (engine *FaceEngine) GetLivenessScoreIR() (LivenessInfo, error) {
	asfLivenessInfo := &C.ASF_LivenessInfo{}
	r := C.ASFGetLivenessScore_IR((C.MHandle)(engine.handle), asfLivenessInfo)
	if r != C.MOK {
		return LivenessInfo{}, newError(int(r), "获取活体信息失败")
	}
	num := int32(asfLivenessInfo.num)
	return LivenessInfo{
		IsLive: (*[10]int32)(unsafe.Pointer(asfLivenessInfo.isLive))[:num:num],
		Num:    num,
	}, nil
}

// GetMask 获取口罩信息
func (engine *FaceEngine) GetMask() (maskInfo MaskInfo, err error) {
	asfMaskInfo := &C.ASF_MaskInfo{}
	r := C.ASFGetMask(engine.handle, (*C.ASF_MaskInfo)(unsafe.Pointer(asfMaskInfo)))
	if r != C.MOK {
		return maskInfo, newError(int(r), "获取口罩信息失败")
	}
	num := int32(asfMaskInfo.num)
	return MaskInfo{
		MaskArray: (*[10]int32)(unsafe.Pointer(asfMaskInfo.maskArray))[:num:num],
		Num:       num,
	}, nil
}

// GetFaceLandMarkInfo 获取额头区域位置
func (engine *FaceEngine) GetFaceLandMarkInfo() (landMarkInfo LandMarkInfo, err error) {
	asfLandMarkInfo := &C.ASF_LandMarkInfo{}
	r := C.ASFGetFaceLandMark(engine.handle, asfLandMarkInfo)
	if r != C.MOK {
		return landMarkInfo, newError(int(r), "获取额头区域位置失败")
	}
	num := int32(asfLandMarkInfo.num)
	asfPoint := (*[10]C.ASF_FaceLandmark)(unsafe.Pointer(asfLandMarkInfo.point))[:num:num]
	point := make([]FaceLandMark, num)
	for i := int32(0); i < num; i++ {
		point[i] = FaceLandMark{
			PosX: float32(asfPoint[i].x),
			PosY: float32(asfPoint[i].y),
		}
	}
	return LandMarkInfo{
		Point: point,
		Num:   num,
	}, nil
}

// UpdateFaceData 更新人脸数据
func (engine *FaceEngine) UpdateFaceData(
	width int, // 宽度
	height int, // 高度
	format C.MInt32, // 图像格式
	imgData []byte, // 图片数据
	faceInfo MultiFaceInfo, // 多人脸信息
) (err error) {
	r := C.ASFUpdateFaceData(engine.handle, C.MInt32(width), C.MInt32(height), format, (*C.MUInt8)(unsafe.Pointer(&imgData[0])), faceInfo.native)
	if r != C.MOK {
		err = newError(int(r), "更新人脸数据失败")
	}
	return
}

// ImageQualityDetect 单人脸图片质量检测
func (engine *FaceEngine) ImageQualityDetect(
	width int, // 宽度
	height int, // 高度
	format C.MInt32, // 图像格式
	imgData []byte, // 图片数据
	faceInfo SingleFaceInfo, // 单人脸信息
	isMask int, // 是否带口罩
) (confidenceLevel float32, err error) {
	asfFaceInfo := &C.ASF_SingleFaceInfo{
		C.MRECT{
			C.MInt32(faceInfo.FaceRect.Left),
			C.MInt32(faceInfo.FaceRect.Top),
			C.MInt32(faceInfo.FaceRect.Right),
			C.MInt32(faceInfo.FaceRect.Bottom)},
		C.MInt32(faceInfo.FaceOrient),
		faceInfo.DataInfo}
	r := C.ASFImageQualityDetect(
		engine.handle,
		C.MInt32(width),
		C.MInt32(height),
		format,
		(*C.MUInt8)(unsafe.Pointer(&imgData[0])),
		asfFaceInfo,
		C.MInt32(isMask),
		(*C.MFloat)(unsafe.Pointer(&confidenceLevel)),
		C.ASF_DETECT_MODEL_RGB,
	)
	if r != C.MOK {
		err = newError(int(r), "单人脸图片质量检测失败")
	}
	return confidenceLevel, err
}

// SetFaceShelterParam 设置遮挡算法检测的阈值
func (engine *FaceEngine) SetFaceShelterParam(shelterThreshhold float32) (err error) {
	r := C.ASFSetFaceShelterParam(engine.handle, C.MFloat(shelterThreshhold))
	if r != C.MOK {
		err = newError(int(r), "设置遮挡算法检测阈值失败")
	}
	return
}

// Destroy 销毁引擎
func (engine *FaceEngine) Destroy() (err error) {
	r := C.ASFUninitEngine(engine.handle)
	if r != C.MOK {
		err = newError(int(r), "销毁引擎失败")
	}
	return
}

// GetSingleFaceInfo 从多人脸结构体中提取单人脸信息
func GetSingleFaceInfo(multiFaceInfo MultiFaceInfo) (faceInfo []SingleFaceInfo) {
	faceInfo = make([]SingleFaceInfo, multiFaceInfo.FaceNum)
	for i := 0; i < len(faceInfo); i++ {
		faceInfo[i].FaceRect = Rect{
			Left:   multiFaceInfo.FaceRect[i].Left,
			Top:    multiFaceInfo.FaceRect[i].Top,
			Right:  multiFaceInfo.FaceRect[i].Right,
			Bottom: multiFaceInfo.FaceRect[i].Bottom,
		}
		faceInfo[i].FaceOrient = multiFaceInfo.FaceOrient[i]
		faceInfo[i].DataInfo = multiFaceInfo.FaceDataInfoList[i]
	}
	return
}

// ReadFaceFeatureFromBytes 将字节数据转为人脸特征数据
func ReadFaceFeatureFromBytes(bytes []byte) (feature FaceFeature) {
	asfFaceFeature := &C.ASF_FaceFeature{}
	arr := (*C.MByte)(C.malloc(C.size_t(int32(len(bytes)))))
	featurePtr := (*C.uchar)(unsafe.Pointer(C.CString(string(bytes))))
	asfFaceFeature.feature = featurePtr
	asfFaceFeature.featureSize = C.int(len(bytes))
	return FaceFeature{
		Feature:     bytes,
		FeatureSize: int32(len(bytes)),
		featurePtr:  arr,
		native:      asfFaceFeature,
	}
}

// Release 释放内存
func (feature *FaceFeature) Release() {
	if feature.featurePtr != nil {
		C.free(unsafe.Pointer(feature.featurePtr))
	}
}

// 实现Error接口
func (err EngineError) Error() string {
	return err.Text
}

func newError(code int, text string) EngineError {
	return EngineError{
		Code: code,
		Text: text,
	}
}

func imageDataToASVLOFFSCREEN(imageData ImageData) C.LPASVLOFFSCREEN {
	var pi32Pitch [4](C.MInt32)
	for i := 0; i < 4; i++ {
		pi32Pitch[i] = C.MInt32(imageData.WidthStep[0])
	}
	var ppu8Plane [4](*C.MUInt8)
	for i := 0; i < 4; i++ {
		if len(imageData.ImageData[i]) > 0 {
			ppu8Plane[i] = (*C.MUInt8)(unsafe.Pointer(&imageData.ImageData[i][0]))
		}
	}
	return &C.ASVLOFFSCREEN{
		imageData.PixelArrayFormat,
		C.MInt32(imageData.Width),
		C.MInt32(imageData.Height),
		ppu8Plane,
		pi32Pitch}
}
