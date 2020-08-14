package face_engine

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

// 引擎结构体
type FaceEngine struct {
	handle C.MHandle
}

// 多人脸信息结构体
type MultiFaceInfo struct {
	FaceRect   []Rect  // 人脸框信息
	FaceOrient []int32 // 输入图像的角度
	FaceNum    int32   // 检测到的人脸个数
	FaceID     []int32 // face ID，IMAGE模式下不返回FaceID
	native     *C.ASF_MultiFaceInfo
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
	native      *C.ASF_FaceFeature
	featurePtr  *C.MByte
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

// SDK错误码
type EngineError struct {
	Code int
	Text string
}

const (
	DetectModeVideo       = C.ASF_DETECT_MODE_VIDEO
	DetectModeImage       = C.ASF_DETECT_MODE_IMAGE
	OrientPriority0       = C.ASF_OP_0_ONLY
	OrientPriority90      = C.ASF_OP_90_ONLY
	OrientPriority270     = C.ASF_OP_270_ONLY
	OrientPriority180     = C.ASF_OP_180_ONLY
	OrientPriorityAllOut  = C.ASF_OP_ALL_OUT
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
func NewFaceEngine(dm, op uint32, faceScale, maxFaceNum, combinedMask C.MInt32) (*FaceEngine, error) {
	engine, err := &FaceEngine{}, error(nil)
	r := C.ASFInitEngine(dm, op, faceScale, maxFaceNum, combinedMask, (*C.MHandle)(unsafe.Pointer(&engine.handle)))
	if r != C.MOK {
		err = newError(int(r), "初始化引擎失败")
	}
	return engine, err
}

// 获取激活文件信息接口
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
		err = newError(int(r), "激活SDK失败")
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
		err = newError(int(r), "激活SDK失败")
	}
	return
}

// 人脸检测，目前不支持IR图像数据检测
func (engine *FaceEngine) DetectFaces(width, height int, format C.MInt32, imgData []byte) (faceInfo MultiFaceInfo, err error) {
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
	faceInfo.FaceRect = (*[50]Rect)(unsafe.Pointer(asfFaceInfo.faceRect))[:faceNum:faceNum]
	faceInfo.FaceOrient = (*[50]int32)(unsafe.Pointer(asfFaceInfo.faceOrient))[:faceNum:faceNum]
	if asfFaceInfo.faceID != nil {
		faceInfo.FaceID = (*[50]int32)(unsafe.Pointer(asfFaceInfo.faceID))[:faceNum:faceNum]
	}
	faceInfo.FaceNum = faceNum
	faceInfo.native = asfFaceInfo
	return
}

// 年龄/性别/人脸3D角度（该接口仅支持RGB图像），最多支持4张人脸信息检测，超过部分返回未知
// RGB活体仅支持单人脸检测，该接口不支持检测IR活体
func (engine *FaceEngine) Process(width, height int, format C.MInt32, imgData []byte, detectedFaces MultiFaceInfo, combinedMask C.MInt32) error {
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

// 该接口目前仅支持单人脸IR活体检测（不支持年龄、性别、3D角度的检测），默认取第一张人脸
func (engine *FaceEngine) ProcessIR(width, height int, format C.MInt32, imgData []byte, detectedFaces MultiFaceInfo, combinedMask C.MInt32) error {
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

// 设置活体置信度
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
func (engine *FaceEngine) FaceFeatureExtract(width, height int, format C.MInt32, imgData []byte, faceInfo SingleFaceInfo) (faceFeature FaceFeature, err error) {
	asfFaceFeature := &C.ASF_FaceFeature{}
	asfFaceInfo := &C.ASF_SingleFaceInfo{
		C.MRECT{
			C.MInt32(faceInfo.FaceRect.Left),
			C.MInt32(faceInfo.FaceRect.Top),
			C.MInt32(faceInfo.FaceRect.Right),
			C.MInt32(faceInfo.FaceRect.Bottom)},
		C.MInt32(faceInfo.FaceOrient)}
	r := C.ASFFaceFeatureExtract(
		engine.handle,
		C.MInt32(width),
		C.MInt32(height),
		format,
		(*C.MUInt8)(unsafe.Pointer(&imgData[0])),
		asfFaceInfo,
		asfFaceFeature)
	if r != C.MOK {
		return FaceFeature{}, newError(int(r), "提取人脸特征失败")
	}
	length := int32(asfFaceFeature.featureSize)
	faceFeature.FeatureSize = length
	faceFeature.Feature = make([]byte, length)
	byteArr := (*[1 << 28]byte)(unsafe.Pointer(asfFaceFeature.feature))[:length:length]
	arr := (*C.MByte)(C.malloc(C.size_t(int32(asfFaceFeature.featureSize))))
	faceFeature.featurePtr = arr
	ps := (*[1 << 28]C.MByte)(unsafe.Pointer(arr))[:length:length]
	for i := 0; i < len(ps); i++ {
		ps[i] = C.MByte(byteArr[i])
		faceFeature.Feature[i] = byteArr[i]
	}
	asfFaceFeature.feature = arr
	faceFeature.native = asfFaceFeature
	return faceFeature, err
}



//读取提取后保存的特征，特征可通过文件形式保存至本地或数据库的blob类型字段，并通过这个函数返回FaceFeature结构体
//调用：
//var FromFile FaceFeature
//f, err := os.OpenFile("./face.data", os.O_RDONLY,0600)
//defer f.Close()
//t,_ := ioutil.ReadAll(f)
//FromFile.FaceTureByByte(t)
//FromFile即可放入FaceFeatureCompare函数与其他特征进行对比
func (this *FaceFeature) FaceTureByByte(bytes []byte)(err error){		//(featture FaceFeature,err error)
	asfFaceFeature := &C.ASF_FaceFeature{}
	arr := (*C.MByte)(C.malloc(C.size_t(int32(len(bytes)))))
	featurePtr := (*C.uchar)(unsafe.Pointer(C.CString(string(bytes))))
	asfFaceFeature.feature = featurePtr
	asfFaceFeature.featureSize = C.int(len(bytes))
	this.Feature = bytes
	this.FeatureSize = int32(len(bytes))
	this.featurePtr = arr
	this.native = asfFaceFeature
	return
}




// 人脸特征比对
func (engine *FaceEngine) FaceFeatureCompare(feature1, feature2 FaceFeature) (float32, error) {
	var confidenceLevel float32 = 0
	r := C.ASFFaceFeatureCompare(engine.handle,
		feature1.native,
		feature2.native,
		(*C.MFloat)(unsafe.Pointer(&confidenceLevel)),
		C.ASF_DETECT_MODEL_RGB,
	)
	if r != C.MOK {
		return 0, newError(int(r), "人脸特征比对失败!")
	}
	return confidenceLevel, nil
}

// 获取年龄信息
func (engine *FaceEngine) GetAge() (AgeInfo, error) {
	asfAgeInfo := &C.ASF_AgeInfo{}
	r := C.ASFGetAge((C.MHandle)(engine.handle), asfAgeInfo)
	if r != C.MOK {
		return AgeInfo{}, newError(int(r), "获取年龄信息失败")
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
	r := C.ASFGetGender((C.MHandle)(engine.handle), asfGenderInfo)
	if r != C.MOK {
		return GenderInfo{}, newError(int(r), "获取性别信息失败")
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
	r := C.ASFGetFace3DAngle((C.MHandle)(engine.handle), asfFace3DAngle)
	if r != C.MOK {
		return Face3DAngle{}, newError(int(r), "获取3D角度信息失败")
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
	r := C.ASFGetLivenessScore((C.MHandle)(engine.handle), asfLivenessInfo)
	if r != C.MOK {
		return LivenessInfo{}, newError(int(r), "获取活体信息失败")
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
	r := C.ASFGetLivenessScore_IR((C.MHandle)(engine.handle), asfLivenessInfo)
	if r != C.MOK {
		return LivenessInfo{}, newError(int(r), "获取活体信息失败")
	}
	num := int32(asfLivenessInfo.num)
	return LivenessInfo{
		IsLive: (*[50]int32)(unsafe.Pointer(asfLivenessInfo.isLive))[:num:num],
		Num:    num,
	}, nil
}

// 销毁引擎
func (engine *FaceEngine) Destroy() error {
	r := C.ASFUninitEngine(engine.handle)
	if r != C.MOK {
		return newError(int(r), "销毁引擎失败")
	}
	return nil
}

// 从多人脸结构体中提取单人脸信息
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
	}
	return
}

// 释放内存
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
