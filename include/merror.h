/*----------------------------------------------------------------------------------------------
*
* This file is ArcSoft's property. It contains ArcSoft's trade secret, proprietary and
* confidential information.
*
* The information and code contained in this file is only for authorized ArcSoft employees
* to design, create, modify, or review.
*
* DO NOT DISTRIBUTE, DO NOT DUPLICATE OR TRANSMIT IN ANY FORM WITHOUT PROPER AUTHORIZATION.
*
* If you are not an intended recipient of this file, you must not copy, distribute, modify,
* or take any action in reliance on it.
*
* If you have received this file in error, please immediately notify ArcSoft and
* permanently delete the original and any copy of any file and any printout thereof.
*
*-------------------------------------------------------------------------------------------------*/


#ifndef __MERROR_H__
#define __MERROR_H__


#define MERR_NONE                        (0)
#define MOK                                (0)


#define MERR_BASIC_BASE                    0X0001                            //通用错误类型
#define MERR_UNKNOWN                    MERR_BASIC_BASE                    //错误原因不明
#define MERR_INVALID_PARAM                (MERR_BASIC_BASE+1)                //无效的参数
#define MERR_UNSUPPORTED                (MERR_BASIC_BASE+2)                //引擎不支持
#define MERR_NO_MEMORY                    (MERR_BASIC_BASE+3)                //内存不足
#define MERR_BAD_STATE                    (MERR_BASIC_BASE+4)                //状态错误
#define MERR_USER_CANCEL                (MERR_BASIC_BASE+5)                //用户取消相关操作
#define MERR_EXPIRED                    (MERR_BASIC_BASE+6)                //操作时间过期
#define MERR_USER_PAUSE                    (MERR_BASIC_BASE+7)                //用户暂停操作
#define MERR_BUFFER_OVERFLOW            (MERR_BASIC_BASE+8)                //缓冲上溢
#define MERR_BUFFER_UNDERFLOW            (MERR_BASIC_BASE+9)                //缓冲下溢
#define MERR_NO_DISKSPACE                (MERR_BASIC_BASE+10)            //存贮空间不足
#define    MERR_COMPONENT_NOT_EXIST        (MERR_BASIC_BASE+11)            //组件不存在
#define    MERR_GLOBAL_DATA_NOT_EXIST        (MERR_BASIC_BASE+12)            //全局数据不存在


#define MERR_FSDK_BASE                            0X7000                    //Free SDK通用错误类型
#define MERR_FSDK_INVALID_APP_ID                (MERR_FSDK_BASE+1)        //无效的App Id
#define MERR_FSDK_INVALID_SDK_ID                (MERR_FSDK_BASE+2)        //无效的SDK key
#define MERR_FSDK_INVALID_ID_PAIR                (MERR_FSDK_BASE+3)        //AppId和SDKKey不匹配
#define MERR_FSDK_MISMATCH_ID_AND_SDK            (MERR_FSDK_BASE+4)        //SDKKey 和使用的SDK 不匹配
#define MERR_FSDK_SYSTEM_VERSION_UNSUPPORTED    (MERR_FSDK_BASE+5)        //系统版本不被当前SDK所支持
#define MERR_FSDK_LICENCE_EXPIRED                (MERR_FSDK_BASE+6)        //SDK有效期过期，需要重新下载更新


#define MERR_FSDK_FR_ERROR_BASE                    0x12000                            //Face Recognition错误类型
#define MERR_FSDK_FR_INVALID_MEMORY_INFO        (MERR_FSDK_FR_ERROR_BASE+1)        //无效的输入内存
#define MERR_FSDK_FR_INVALID_IMAGE_INFO            (MERR_FSDK_FR_ERROR_BASE+2)        //无效的输入图像参数
#define MERR_FSDK_FR_INVALID_FACE_INFO            (MERR_FSDK_FR_ERROR_BASE+3)        //无效的脸部信息
#define MERR_FSDK_FR_NO_GPU_AVAILABLE            (MERR_FSDK_FR_ERROR_BASE+4)        //当前设备无GPU可用
#define MERR_FSDK_FR_MISMATCHED_FEATURE_LEVEL    (MERR_FSDK_FR_ERROR_BASE+5)        //待比较的两个人脸特征的版本不一致


#define MERR_FSDK_FACEFEATURE_ERROR_BASE            0x14000                                    //人脸特征检测错误类型
#define MERR_FSDK_FACEFEATURE_UNKNOWN                (MERR_FSDK_FACEFEATURE_ERROR_BASE+1)    //人脸特征检测错误未知
#define MERR_FSDK_FACEFEATURE_MEMORY                (MERR_FSDK_FACEFEATURE_ERROR_BASE+2)    //人脸特征检测内存错误
#define MERR_FSDK_FACEFEATURE_INVALID_FORMAT        (MERR_FSDK_FACEFEATURE_ERROR_BASE+3)    //人脸特征检测格式错误
#define MERR_FSDK_FACEFEATURE_INVALID_PARAM            (MERR_FSDK_FACEFEATURE_ERROR_BASE+4)    //人脸特征检测参数错误
#define MERR_FSDK_FACEFEATURE_LOW_CONFIDENCE_LEVEL    (MERR_FSDK_FACEFEATURE_ERROR_BASE+5)    //人脸特征检测结果置信度低

#define MERR_ASF_EX_BASE                                0x15000                            //ASF错误类型
#define MERR_ASF_EX_FEATURE_UNSUPPORTED_ON_INIT            (MERR_ASF_EX_BASE+1)            //Engine不支持的检测属性
#define MERR_ASF_EX_FEATURE_UNINITED                    (MERR_ASF_EX_BASE+2)            //需要检测的属性未初始化
#define MERR_ASF_EX_FEATURE_UNPROCESSED                    (MERR_ASF_EX_BASE+3)            //待获取的属性未在process中处理过
#define MERR_ASF_EX_FEATURE_UNSUPPORTED_ON_PROCESS        (MERR_ASF_EX_BASE+4)            //PROCESS不支持的检测属性组合，例如FR，有自己独立的处理函数
#define MERR_ASF_EX_INVALID_IMAGE_INFO                    (MERR_ASF_EX_BASE+5)            //无效的输入图像
#define MERR_ASF_EX_INVALID_FACE_INFO                    (MERR_ASF_EX_BASE+6)            //无效的脸部信息

#define MERR_ASF_BASE                                    0x16000                            //人脸比对基础错误类型
#define MERR_ASF_ACTIVATION_FAIL                        (MERR_ASF_BASE+1)                //SDK激活失败,请打开读写权限
#define MERR_ASF_ALREADY_ACTIVATED                        (MERR_ASF_BASE+2)                //SDK已激活
#define MERR_ASF_NOT_ACTIVATED                            (MERR_ASF_BASE+3)                //SDK未激活
#define MERR_ASF_SCALE_NOT_SUPPORT                        (MERR_ASF_BASE+4)                //detectFaceScaleVal 不支持
#define MERR_ASF_ACTIVEFILE_SDKTYPE_MISMATCH            (MERR_ASF_BASE+5)                //激活文件与SDK类型不匹配，请确认使用的sdk
#define MERR_ASF_DEVICE_MISMATCH                        (MERR_ASF_BASE+6)                //设备不匹配
#define MERR_ASF_UNIQUE_IDENTIFIER_ILLEGAL                (MERR_ASF_BASE+7)                //唯一标识不合法
#define MERR_ASF_PARAM_NULL                                (MERR_ASF_BASE+8)                //参数为空
#define MERR_ASF_LIVENESS_EXPIRED                        (MERR_ASF_BASE+9)                //活体已过期
#define MERR_ASF_VERSION_NOT_SUPPORT                    (MERR_ASF_BASE+10)                //版本不支持
#define MERR_ASF_SIGN_ERROR                                (MERR_ASF_BASE+11)                //签名错误
#define MERR_ASF_DATABASE_ERROR                            (MERR_ASF_BASE+12)                //激活信息保存异常
#define MERR_ASF_UNIQUE_CHECKOUT_FAIL                    (MERR_ASF_BASE+13)                //唯一标识符校验失败
#define MERR_ASF_COLOR_SPACE_NOT_SUPPORT                (MERR_ASF_BASE+14)                //颜色空间不支持
#define    MERR_ASF_IMAGE_WIDTH_HEIGHT_NOT_SUPPORT            (MERR_ASF_BASE+15)                //图片宽高不支持，宽度需四字节对齐

#define MERR_ASF_BASE_EXTEND                            0x16010                            //人脸比对基础错误类型
#define MERR_ASF_READ_PHONE_STATE_DENIED                (MERR_ASF_BASE_EXTEND)            //android.permission.READ_PHONE_STATE权限被拒绝
#define    MERR_ASF_ACTIVATION_DATA_DESTROYED                (MERR_ASF_BASE_EXTEND+1)        //激活数据被破坏,请删除激活文件，重新进行激活
#define    MERR_ASF_SERVER_UNKNOWN_ERROR                    (MERR_ASF_BASE_EXTEND+2)        //服务端未知错误
#define MERR_ASF_INTERNET_DENIED                        (MERR_ASF_BASE_EXTEND+3)        //INTERNET权限被拒绝
#define MERR_ASF_ACTIVEFILE_SDK_MISMATCH                (MERR_ASF_BASE_EXTEND+4)        //激活文件与SDK版本不匹配,请重新激活
#define MERR_ASF_DEVICEINFO_LESS                        (MERR_ASF_BASE_EXTEND+5)        //设备信息太少，不足以生成设备指纹
#define MERR_ASF_LOCAL_TIME_NOT_CALIBRATED                (MERR_ASF_BASE_EXTEND+6)        //客户端时间与服务器时间（即北京时间）前后相差在30分钟以上
#define MERR_ASF_APPID_DATA_DECRYPT                        (MERR_ASF_BASE_EXTEND+7)        //数据校验异常
#define MERR_ASF_APPID_APPKEY_SDK_MISMATCH                (MERR_ASF_BASE_EXTEND+8)        //传入的AppId和AppKey与使用的SDK版本不一致
#define MERR_ASF_NO_REQUEST                                (MERR_ASF_BASE_EXTEND+9)        //短时间大量请求会被禁止请求,30分钟之后解封
#define MERR_ASF_ACTIVE_FILE_NO_EXIST                    (MERR_ASF_BASE_EXTEND+10)        //激活文件不存在
#define MERR_ASF_DETECT_MODEL_UNSUPPORTED                (MERR_ASF_BASE_EXTEND+11)        //检测模型不支持，请查看对应接口说明，使用当前支持的检测模型
#define MERR_ASF_CURRENT_DEVICE_TIME_INCORRECT            (MERR_ASF_BASE_EXTEND+12)        //当前设备时间不正确，请调整设备时间

#define MERR_ASF_NETWORK_BASE                            0x17000                            //网络错误类型
#define MERR_ASF_NETWORK_COULDNT_RESOLVE_HOST            (MERR_ASF_NETWORK_BASE+1)        //无法解析主机地址
#define MERR_ASF_NETWORK_COULDNT_CONNECT_SERVER            (MERR_ASF_NETWORK_BASE+2)        //无法连接服务器
#define MERR_ASF_NETWORK_CONNECT_TIMEOUT                (MERR_ASF_NETWORK_BASE+3)        //网络连接超时
#define MERR_ASF_NETWORK_UNKNOWN_ERROR                    (MERR_ASF_NETWORK_BASE+4)        //网络未知错误


#endif

