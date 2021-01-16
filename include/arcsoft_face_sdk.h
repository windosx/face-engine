/*******************************************************************************
* Copyright(c) ArcSoft, All right reserved.
*
* This file is ArcSoft's property. It contains ArcSoft's trade secret, proprietary
* and confidential information.
*
* DO NOT DISTRIBUTE, DO NOT DUPLICATE OR TRANSMIT IN ANY FORM WITHOUT PROPER
* AUTHORIZATION.
*
* If you are not an intended recipient of this file, you must not copy,
* distribute, modify, or take any action in reliance on it.
*
* If you have received this file in error, please immediately notify ArcSoft and
* permanently delete the original and any copy of any file and any printout
* thereof.
*********************************************************************************/

#ifndef _ARCSOFT_SDK_ASF_H_
#define _ARCSOFT_SDK_ASF_H_

#include "amcomdef.h"
#include "asvloffscreen.h"

#ifdef __cplusplus
extern "C" {
#endif

#define ASF_NONE				0x00000000	//������
#define ASF_FACE_DETECT			0x00000001	//�˴�detect������tracking����detection��������֮һ�������ѡ����detect mode ȷ��
#define ASF_FACERECOGNITION		0x00000004	//��������
#define ASF_AGE					0x00000008	//����
#define ASF_GENDER				0x00000010	//�Ա�
#define ASF_FACE3DANGLE			0x00000020	//3D�Ƕ�
#define ASF_FACELANDMARK		0x00000040	//��ͷ������
#define ASF_LIVENESS			0x00000080	//RGB����
#define ASF_IMAGEQUALITY		0x00000200	//ͼ���������
#define ASF_IR_LIVENESS			0x00000400	//IR����
#define ASF_FACESHELTER			0x00000800  //�����ڵ�
#define ASF_MASKDETECT			0x00001000	//���ּ��
#define ASF_UPDATE_FACEDATA		0x00002000	//������Ϣ


#define ASF_MAX_DETECTFACENUM   10          //�ð汾���֧��ͬʱ���10������

	//���ģʽ
	typedef enum __tag_ASF_DetectMode{
		ASF_DETECT_MODE_VIDEO = 0x00000000,		//Videoģʽ��һ�����ڶ�֡�������
		ASF_DETECT_MODE_IMAGE = 0xFFFFFFFF		//Imageģʽ��һ�����ھ�̬ͼ�ĵ��μ��
	} ASF_DetectMode;

	//���ʱ�������Ƕȵ����ȼ������ĵ��г�ʼ���ӿ�����ͼʾ˵������ο�
	typedef enum __tag_ASF_OrientPriority {
		ASF_OP_0_ONLY = 0x1,		// ����Ԥ����������
		ASF_OP_90_ONLY = 0x2,		// ����0����ʱ����ת90��ķ���
		ASF_OP_270_ONLY = 0x3,		// ����0����ʱ����ת270��ķ���
		ASF_OP_180_ONLY = 0x4,		// ����0����ת180��ķ�����ʱ�롢˳ʱ��Ч��һ����
		ASF_OP_ALL_OUT = 0x5		// ȫ�Ƕ�
	} ASF_OrientPriority;

	//orientation �Ƕȣ���ʱ�뷽��
	typedef enum __tag_ASF_OrientCode {
		ASF_OC_0 = 0x1,			// 0 degree 
		ASF_OC_90 = 0x2,		// 90 degree 
		ASF_OC_270 = 0x3,		// 270 degree 
		ASF_OC_180 = 0x4,   	// 180 degree 
		ASF_OC_30 = 0x5,		// 30 degree 
		ASF_OC_60 = 0x6,		// 60 degree 
		ASF_OC_120 = 0x7,		// 120 degree 
		ASF_OC_150 = 0x8,		// 150 degree 
		ASF_OC_210 = 0x9,		// 210 degree 
		ASF_OC_240 = 0xa,		// 240 degree 
		ASF_OC_300 = 0xb,		// 300 degree 
		ASF_OC_330 = 0xc		// 330 degree 
	} ASF_OrientCode;

	//���ģ��
	typedef enum __tag_ASF_DetectModel {
		ASF_DETECT_MODEL_RGB = 0x1	//RGBͼ����ģ��
		//Ԥ����չ�������ģ��
	} ASF_DetectModel;

	//�����ȶԿ�ѡ��ģ��
	typedef enum __tag_ASF_CompareModel{
		ASF_LIFE_PHOTO = 0x1,	//����������֮��������ȶԣ��Ƽ���ֵ0.80
		ASF_ID_PHOTO = 0x2		//����֤���ջ���������֤����֮��������ȶԣ��Ƽ���ֵ0.82
	} ASF_CompareModel;

	typedef enum __tag_ASF_RegisterOrNot{
		ASF_RECOGNITION = 0x0,  //����ʶ��������������ȡ
		ASF_REGISTER = 0x1      //����ע��������������ȡ
	} ASF_RegisterOrNot;

	//�汾��Ϣ
	typedef struct {
		MPChar Version;				// �汾��
		MPChar BuildDate;			// ��������
		MPChar CopyRight;			// Copyright
	}ASF_VERSION, *LPASF_VERSION;

	//ͼ������
	typedef LPASVLOFFSCREEN LPASF_ImageData;

	//������Ϣ
	typedef struct{
		MPVoid		data;		// ������Ϣ
		MInt32		dataSize;	// ������Ϣ����
	} ASF_FaceDataInfo, *LPASF_FaceDataInfo;

	//��������Ϣ
	typedef struct SingleFaceInfo {
		MRECT		faceRect;		        // ��������Ϣ
		MInt32		faceOrient;		        // ����ͼ��ĽǶȣ����Բο� ArcFaceCompare_OrientCode
		ASF_FaceDataInfo faceDataInfo;		// ����������Ϣ
	} ASF_SingleFaceInfo, *LPASF_SingleFaceInfo;

	//��������Ϣ
	typedef struct MultiFaceInfo {
		MRECT* 		faceRect;			        // ��������Ϣ
		MInt32*		faceOrient;			        // ����ͼ��ĽǶȣ����Բο� ArcFaceCompare_OrientCode .
		MInt32		faceNum;			        // ��⵽����������
		MInt32*     faceID;				        // face ID��IMAGEģʽ�²�����FaceID
		MFloat*		wearGlasses;		        // ���۾����Ŷ�[0-1],�Ƽ���ֵ0.5
		MInt32*		leftEyeClosed;	            // ����״̬ 0 δ���ۣ�1 ����
		MInt32*		rightEyeClosed;	            // ����״̬ 0 δ���ۣ�1 ����
		MInt32*	    faceShelter;	            // "1" ��ʾ �ڵ�, "0" ��ʾ  δ�ڵ�, "-1" ��ʾ��ȷ��
		LPASF_FaceDataInfo faceDataInfoList;	// ����������Ϣ
	}ASF_MultiFaceInfo, *LPASF_MultiFaceInfo;

	// �����ļ���Ϣ
	typedef struct ActiveFileInfo {
		MPChar startTime;		//��ʼʱ��
		MPChar endTime;			//��ֹʱ��
		MPChar activeKey;		//������
		MPChar platform;		//ƽ̨
		MPChar sdkType;			//sdk����
		MPChar appId;			//APPID
		MPChar sdkKey;			//SDKKEY
		MPChar sdkVersion;		//SDK�汾��
		MPChar fileVersion;		//�����ļ��汾��
	}ASF_ActiveFileInfo, *LPASF_ActiveFileInfo;

	/*******************************************************************************************
	* ��ȡ�����ļ���Ϣ�ӿ�
	*******************************************************************************************/
	MRESULT ASFGetActiveFileInfo(
		LPASF_ActiveFileInfo  activeFileInfo  // [out] �����ļ���Ϣ
		);

	/*******************************************************************************************
	* ���߼���ӿ�
	*******************************************************************************************/
	MRESULT ASFOnlineActivation(
		MPChar				AppId,			// [in]  APPID	��������
		MPChar				SDKKey,			// [in]  SDKKEY	��������
		MPChar				ActiveKey		// [in]  ActiveKey	��������
		);

	/*******************************************************************************************
	* ��ȡ�豸��Ϣ�ӿ�
	*******************************************************************************************/
	MRESULT ASFGetActiveDeviceInfo(
		MPChar*				deviceInfo		// [out] �ɼ����豸��Ϣ�����ڵ����������������߼������������Ȩ�ļ�
		);

	/*******************************************************************************************
	* ���߼���ӿ�
	*******************************************************************************************/
	MRESULT ASFOfflineActivation(
		MPChar				filePath		// [in]  �����ļ�·��(������Ȩ�ļ�)����Ҫ��дȨ��
		);

	/************************************************************************
	* ��ʼ������
	************************************************************************/
	MRESULT ASFInitEngine(
		ASF_DetectMode		detectMode,					// [in]	AF_DETECT_MODE_VIDEO ��Ƶģʽ������������ͷԤ������Ƶ�ļ�ʶ��
														//		AF_DETECT_MODE_IMAGE ͼƬģʽ�������ھ�̬ͼƬ��ʶ��
		ASF_OrientPriority	detectFaceOrientPriority,	// [in]	��������ĽǶ�����ֵ���ο� ArcFaceCompare_OrientPriority
		MInt32				detectFaceMaxNum,			// [in] �����Ҫ������������
		MInt32				combinedMask,				// [in] �û�ѡ����Ҫ���Ĺ�����ϣ��ɵ�������
		MHandle*			hEngine						// [out] ��ʼ�����ص�����handle
		);

	/************************************************************************
	* ȡֵ��Χ[0-1]�� Ĭ����ֵ:0.8�� �û����Ը���ʵ�����������ڵ���Χ
	************************************************************************/
	MRESULT ASFSetFaceShelterParam(
		MHandle hEngine,					// [in] ����handle
		MFloat  ShelterThreshhold			// [in] �ڵ���ֵ
		);

	/******************************************************
	* VIDEOģʽ:����׷�� IMAGEģʽ:�������
	******************************************************/
	MRESULT ASFDetectFaces(
		MHandle				hEngine,							// [in] ����handle
		MInt32				width,								// [in] ͼƬ����
		MInt32				height,								// [in] ͼƬ�߶�
		MInt32				format,								// [in] ��ɫ�ռ��ʽ
		MUInt8*				imgData,							// [in] ͼƬ����
		LPASF_MultiFaceInfo	detectedFaces,						// [out]��⵽��������Ϣ 
		ASF_DetectModel 	detectModel							// [in] Ԥ���ֶΣ���ǰ�汾ʹ��Ĭ�ϲ�������
		);

	/******************************************************
	* VIDEOģʽ:����׷�� IMAGEģʽ:�������
	* ͼ�������Խṹ����ʽ���룬�Բ��ø����ֽڶ��뷽ʽ��ͼ������Ը���
	******************************************************/
	MRESULT ASFDetectFacesEx(
		MHandle				hEngine,							// [in] ����handle
		LPASF_ImageData		imgData,							// [in] ͼƬ����
		LPASF_MultiFaceInfo	detectedFaces,						// [out] ��⵽��������Ϣ
		ASF_DetectModel 	detectModel							// [in]	Ԥ���ֶΣ���ǰ�汾ʹ��Ĭ�ϲ�������
		);

	/******************************************************
	* �����޸ĺ�������򣬸���������Ϣ��������˫Ŀ�������������
	* ע�⣺LPASF_MultiFaceInfo�ڸýӿ��м������Ҳ�ǳ���
	******************************************************/
	MRESULT ASFUpdateFaceData(
		MHandle				hEngine,				// [in] ����handle
		MInt32				width, 					// [in] ͼƬ����
		MInt32				height, 				// [in] ͼƬ�߶�
		MInt32				format,					// [in] ��ɫ�ռ��ʽ
		MUInt8 *			imgData, 				// [in] ͼƬ����
		LPASF_MultiFaceInfo detectedFaces			// [in/out]��⵽��������Ϣ 
		);

	/******************************************************
	* �����޸ĺ�������򣬸���������Ϣ��������˫Ŀ�������������
	* ע�⣺LPASF_MultiFaceInfo�ڸýӿ��м������Ҳ�ǳ���
	* ͼ�������Խṹ����ʽ���룬�Բ��ø����ֽڶ��뷽ʽ��ͼ������Ը���
	******************************************************/
	MRESULT ASFUpdateFaceDataEx(
		MHandle                 hEngine,							// [in] ����handle
		LPASF_ImageData		    imgData,							// [in] ͼ������
		LPASF_MultiFaceInfo	    detectedFaces						// [in/out] ��⵽��������Ϣ
		);

	/******************************************************
	* ͼ��������⣬��RGBģʽ�� ʶ����ֵ��0.49 ע����ֵ��0.63  ����ģʽ��ʶ����ֵ��0.29��
	******************************************************/
	MRESULT ASFImageQualityDetect(
		MHandle					hEngine,							// [in] ����handle
		MInt32					width,								// [in] ͼƬ����
		MInt32					height,								// [in] ͼƬ�߶�
		MInt32					format,								// [in] ��ɫ�ռ��ʽ
		MUInt8 *				imgData,							// [in] ͼƬ����
		LPASF_SingleFaceInfo	faceInfo,							// [in] ����λ����Ϣ 
		MInt32					isMask,								// [in] ��֧�ִ���1��0��-1�������� 1��������Ϊδ������
		MFloat*					confidenceLevel,					// [out] ͼ�����������
		ASF_DetectModel			detectModel							// [in]	Ԥ���ֶΣ���ǰ�汾ʹ��Ĭ�ϲ�������
		);

	/******************************************************
	* ͼ��������⣬��RGBģʽ�� ʶ����ֵ��0.49 ע����ֵ��0.63  ����ģʽ��ʶ����ֵ��0.29��
	* ͼ�������Խṹ����ʽ���룬�Բ��ø����ֽڶ��뷽ʽ��ͼ������Ը���
	******************************************************/
	MRESULT ASFImageQualityDetectEx(
		MHandle					hEngine,							// [in] ����handle
		LPASF_ImageData			imgData,							// [in] ͼƬ����
		LPASF_SingleFaceInfo	faceInfo,							// [in] ����λ����Ϣ 
		MInt32					isMask,                             // [in] ��֧�ִ���1��0��-1�������� 1��������Ϊδ������
		MFloat*              	confidenceLevel,					// [out] ͼ�����������
		ASF_DetectModel			detectModel							// [in]	Ԥ���ֶΣ���ǰ�汾ʹ��Ĭ�ϲ�������
		);

	/************************************************************************
	* ����/�Ա�/����3D�Ƕ�/����/�ڵ�/��ͷ���򣨸ýӿڽ�֧��RGBͼ�񣩣����֧��4��������Ϣ��⣬�������ַ���δ֪
	* RGB�����֧�ֵ�������⣬�ýӿڲ�֧�ּ��IR����
	************************************************************************/
	MRESULT ASFProcess(
		MHandle				hEngine,			// [in] ����handle
		MInt32				width,				// [in] ͼƬ����
		MInt32				height,				// [in] ͼƬ�߶�
		MInt32				format,				// [in] ��ɫ�ռ��ʽ
		MUInt8*				imgData,			// [in] ͼƬ����
		LPASF_MultiFaceInfo	detectedFaces,		// [in] ������Ϣ���û����ݴ����Ĺ���ѡ����Ҫʹ�õ�������
		MInt32				combinedMask		// [in] ֻ֧�ֳ�ʼ��ʱ��ָ����Ҫ���Ĺ��ܣ���processʱ��һ��������Ѿ�ָ���Ĺ��ܼ��м���ɸѡ
												//      �����ʼ����ʱ��ָ�����������Ա���process��ʱ�����ֻ������䣬���ǲ��ܼ���������Ա�֮��Ĺ���    
		);

	/************************************************************************
	* ����/�Ա�/����3D�Ƕ�/����/�ڵ�/��ͷ���򣨸ýӿڽ�֧��RGBͼ�񣩣����֧��4��������Ϣ��⣬�������ַ���δ֪
	* RGB�����֧�ֵ�������⣬�ýӿڲ�֧�ּ��IR����
	* ͼ�������Խṹ����ʽ���룬�Բ��ø����ֽڶ��뷽ʽ��ͼ������Ը���
	************************************************************************/
	MRESULT ASFProcessEx(
		MHandle				hEngine,			// [in] ����handle
		LPASF_ImageData		imgData,			// [in] ͼƬ����
		LPASF_MultiFaceInfo detectedFaces,		// [in] ������Ϣ���û����ݴ����Ĺ���ѡ����Ҫʹ�õ�������
		MInt32				combinedMask		// [in] ֻ֧�ֳ�ʼ��ʱ��ָ����Ҫ���Ĺ��ܣ���processʱ��һ��������Ѿ�ָ���Ĺ��ܼ��м���ɸѡ
		//      �����ʼ����ʱ��ָ�����������Ա���process��ʱ�����ֻ������䣬���ǲ��ܼ���������Ա�֮��Ĺ��� 
		);

	/************************************************************************
	* �ýӿ�Ŀǰ��֧�ֵ�����IR�����⣬Ĭ��ȡ��һ������
	************************************************************************/
	MRESULT ASFProcess_IR(
		MHandle				hEngine,			// [in] ����handle
		MInt32				width,				// [in] ͼƬ����
		MInt32				height,				// [in] ͼƬ�߶�
		MInt32				format,				// [in] ��ɫ�ռ��ʽ
		MUInt8*				imgData,			// [in] ͼƬ����
		LPASF_MultiFaceInfo	detectedFaces,		// [in] ������Ϣ���û����ݴ����Ĺ���ѡ����Ҫʹ�õ�������
		MInt32				combinedMask		// [in] Ŀǰֻ֧�ִ���ASF_IR_LIVENESS���ԵĴ��룬�ҳ�ʼ���ӿ���Ҫ���� 
		);

	/************************************************************************
	* �ýӿ�Ŀǰ��֧�ֵ�����IR�����⣬Ĭ��ȡ��һ������
	* ͼ�������Խṹ����ʽ���룬�Բ��ø����ֽڶ��뷽ʽ��ͼ������Ը���
	************************************************************************/
	MRESULT ASFProcessEx_IR(
		MHandle				hEngine,			// [in] ����handle
		LPASF_ImageData		imgData,			// [in] ͼƬ����
		LPASF_MultiFaceInfo detectedFaces,		// [in] ������Ϣ���û����ݴ����Ĺ���ѡ����Ҫʹ�õ�������
		MInt32				combinedMask		// [in] Ŀǰֻ֧�ִ���ASF_IR_LIVENESS���ԵĴ��룬�ҳ�ʼ���ӿ���Ҫ����
		);

	//******************************** ����ʶ����� *********************************************
	typedef struct FaceFeature {
		MByte*		feature;		// ����������Ϣ
		MInt32		featureSize;	// ����������Ϣ����	
	}ASF_FaceFeature, *LPASF_FaceFeature;

	/************************************************************************
	* ������������ȡ
	************************************************************************/
	MRESULT ASFFaceFeatureExtract(
		MHandle					hEngine,							// [in]	����handle
		MInt32					width,								// [in] ͼƬ����
		MInt32					height,								// [in] ͼƬ�߶�
		MInt32					format,								// [in] ��ɫ�ռ��ʽ
		MUInt8*					imgData,							// [in] ͼƬ����
		LPASF_SingleFaceInfo	faceInfo,							// [in] ��������λ�úͽǶ���Ϣ
		ASF_RegisterOrNot		registerOrNot,						// [in] ע�� 1 ʶ��Ϊ 0
		MInt32					mask,								// [in] ������ 1������0
		LPASF_FaceFeature		feature								// [out] ��������
		);

	/************************************************************************
	* ������������ȡ
	* ͼ�������Խṹ����ʽ���룬�Բ��ø����ֽڶ��뷽ʽ��ͼ������Ը���
	************************************************************************/
	MRESULT ASFFaceFeatureExtractEx(
		MHandle					hEngine,							// [in]	����handle
		LPASF_ImageData			imgData,							// [in] ͼ������
		LPASF_SingleFaceInfo	faceInfo,							// [in] ��������λ�úͽǶ���Ϣ
		ASF_RegisterOrNot		registerOrNot,						// [in] ע�� 1 ʶ��Ϊ 0
		MInt32					mask,								// [in] ������ 1������0
		LPASF_FaceFeature		feature								// [out] ��������
		);

	/************************************************************************
	* ���������ȶԣ��Ƽ���ֵ ASF_LIFE_PHOTO��0.80  ASF_ID_PHOTO��0.82
	************************************************************************/
	MRESULT ASFFaceFeatureCompare(
		MHandle					hEngine,						// [in] ����handle
		LPASF_FaceFeature		feature1,						// [in] ���Ƚ���������1
		LPASF_FaceFeature		feature2,						// [in] ���Ƚ���������2
		MFloat*					confidenceLevel,				// [out] �ȽϽ�������Ŷ���ֵ
		ASF_CompareModel		compareModel					// [in] ASF_LIFE_PHOTO������������֮��������ȶ�
																//		ASF_ID_PHOTO������֤���ջ�֤���պ�������֮��������ȶ�
		);

	//******************************** ������� **********************************************
	typedef struct AgeInfo {
		MInt32*	ageArray;				// "0" ������ȷ��������0����ֵ������������������
		MInt32	num;					// ������������
	}ASF_AgeInfo, *LPASF_AgeInfo;

	/************************************************************************
	* ��ȡ������Ϣ
	************************************************************************/
	MRESULT ASFGetAge(
		MHandle hEngine,				// [in] ����handle
		LPASF_AgeInfo ageInfo			// [out] ��⵽��������Ϣ
		);

	//******************************** �Ա���� **********************************************
	typedef struct GenderInfo {
		MInt32*	genderArray;			// "0" ��ʾ ����, "1" ��ʾ Ů��, "-1" ��ʾ��ȷ��
		MInt32	num;					// ������������	
	}ASF_GenderInfo, *LPASF_GenderInfo;

	/************************************************************************
	* ��ȡ�Ա���Ϣ
	************************************************************************/
	MRESULT ASFGetGender(
		MHandle hEngine,				// [in] ����handle
		LPASF_GenderInfo genderInfo		// [out] ��⵽���Ա���Ϣ
		);

	//******************************** ����3D �Ƕ���Ϣ ***********************************
	typedef struct Face3DAngle {
		MFloat* roll;
		MFloat* yaw;
		MFloat* pitch;
		MInt32* status;
		MInt32 num;
	}ASF_Face3DAngle, *LPASF_Face3DAngle;

	/************************************************************************
	* ��ȡ3D�Ƕ���Ϣ
	************************************************************************/
	MRESULT ASFGetFace3DAngle(
		MHandle hEngine,					// [in] ����handle
		LPASF_Face3DAngle p3DAngleInfo		// [out] ��⵽����3D �Ƕ���Ϣ
		);

	//******************************** ������Ϣ ***********************************
	typedef struct LivenessThreshold {
		MFloat		thresholdmodel_BGR;
		MFloat		thresholdmodel_IR;
	}ASF_LivenessThreshold, *LPASF_LivenessThreshold;

	/************************************************************************
	* ȡֵ��Χ[0-1]��Ĭ��ֵ BGR:0.5 IR:0.7�� �û����Ը���ʵ���������ò�ͬ����ֵ
	************************************************************************/
	MRESULT ASFSetLivenessParam(
		MHandle					hEngine,		// [in] ����handle
		LPASF_LivenessThreshold threshold		// [in] �������Ŷ�
		);

	typedef struct LivenessInfo {
		MInt32*	isLive;			// [out] �ж��Ƿ����ˣ� 0�������ˣ�
		//						1�����ˣ�
		//						-1����ȷ���� 
		//						-2:����������>1��
		//                      -3: ������С
		//                      -4: �Ƕȹ���
		//                      -5: ���������߽� 
		//					    -6: ���ͼ����
		//					    -7: ����ͼ̫����
		MInt32 num;
	}ASF_LivenessInfo, *LPASF_LivenessInfo;

	/************************************************************************
	* ��ȡRGB������
	************************************************************************/
	MRESULT ASFGetLivenessScore(
		MHandle hEngine,					// [in] ����handle
		LPASF_LivenessInfo livenessInfo		// [out] ���RGB������
		);

	/************************************************************************
	* ��ȡIR������
	************************************************************************/
	MRESULT ASFGetLivenessScore_IR(
		MHandle				hEngine,			// [in] ����handle
		LPASF_LivenessInfo	 irLivenessInfo		// [out] ��⵽IR������
		);

	//******************************** ���ּ����� **********************************************
	typedef struct MaskInfo
	{
		MInt32*	maskArray;				// "0" ����û�д����֣�"1"���������� ,"-1"����ȷ��
		MInt32	num;					// ������������
	}ASF_MaskInfo, *LPASF_MaskInfo;

	/************************************************************************
	* ��ȡ���ּ��Ľ��
	************************************************************************/
	MRESULT ASFGetMask(
		MHandle hEngine,				// [in] ����handle
		LPASF_MaskInfo maskInfo			// [out] ��⵽�Ŀ��ּ�����
		);

	//******************************** ��ͷ��������� **********************************************

	#define LANDMARKS_NUM 4	//���������

	typedef struct
	{
		MFloat x;
		MFloat y;
	}ASF_FaceLandmark, *LPASF_FaceLandmark;

	typedef struct LandMarkInfo
	{
		ASF_FaceLandmark *point;		//��ͷ��λ
		MInt32 num;						//��������
	}ASF_LandMarkInfo, *LPASF_LandMarkInfo;

	/************************************************************************
	* ��ȡ��ͷ������������ǰֻ֧��0, 90, 180, 270�ȽǼ�⣩
	************************************************************************/
	MRESULT ASFGetFaceLandMark(
		MHandle				engine,			 // [in] ����handle
		LPASF_LandMarkInfo  LandMarkInfo	 // [out]������ͷ�����飬ÿ��������ͷ����ͨ���ĸ����ʾ
		);

	/************************************************************************
	* ��������
	************************************************************************/
	MRESULT ASFUninitEngine(
		MHandle hEngine
		);

	/************************************************************************
	* ��ȡ�汾��Ϣ
	************************************************************************/
	const ASF_VERSION ASFGetVersion();

#ifdef __cplusplus
}
#endif
#endif