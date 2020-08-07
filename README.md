# 虹软人脸检测SDK
![LICENSE](https://img.shields.io/github/license/windosx/face-engine)
![Supported platform](https://img.shields.io/badge/platform-win%20%7C%20linux-lightgrey)
[![Build Status](https://travis-ci.org/windosx/face-engine.svg?branch=master)](https://travis-ci.org/windosx/face-engine)
[![GoDoc](http://godoc.org/github.com/windosx/face-engine?status.svg)](http://godoc.org/github.com/windosx/face-engine)

基于虹软SDK V3

目前已适配的版本：
[v2.2.0](https://github.com/windosx/face-engine/tree/v2.2.0)
[v3.0.0](https://github.com/windosx/face-engine/tree/v3.0.5)

开始使用
---
* 安装方式一（不使用go mod）：`go get -d github.com/windosx/face-engine`
* 安装方式二（使用go mod）：`go get github.com/windosx/face-engine/v3`
* 编译：`go build -o test github.com/windosx/face-engine/v3/cmd`
* 测试：`./test`
* 库文件：Linux下将`libarcsoft_face.so`, `libarcsoft_face_engine.so`放入`/usr/lib`目录
* Windows下将`libarcsoft_face.dll`, `libarcsoft_face_engine.dll`放入`%WINDIR%/System32`目录或执行文件同目录下
* **编译环境依赖GCC** Linux环境下编译需要先安装gcc，Windows下编译需要先安装MinGW-w64（32位下用mingw），并将`libarcsoft_face.dll`, `libarcsoft_face_engine.dll`放入MingGW-w64安装目录下的`lib`目录中
