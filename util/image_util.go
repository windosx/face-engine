package util

import (
	"fmt"
	"github.com/Comdex/imgo"
)

// 获取适配SDK大小的图片，并转为BGR格式的数据
func GetResizedBGR(filename string) (bgrData []uint8) {
	img, err := imgo.DecodeImage(filename)
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	height := imgo.GetImageHeight(img)
	width := imgo.GetImageWidth(img)
	newWidth := width - width%4
	imgMatrix, err := imgo.ResizeForMatrix(filename, newWidth, height)
	if err != nil {
		panic(err)
	}
	for starty := 0; starty < height; starty++ {
		for startx := 0; startx < newWidth; startx++ {
			R := imgMatrix[starty][startx][0]
			G := imgMatrix[starty][startx][1]
			B := imgMatrix[starty][startx][2]
			bgrData = append(bgrData, B, G, R)
		}
	}
	return bgrData
}

// 获取图片宽高
func GetImageWidthAndHeight(filename string) (width, height int) {
	img, err := imgo.DecodeImage(filename)
	if err != nil {
		return 0, 0
	}
	return imgo.GetImageWidth(img), imgo.GetImageHeight(img)
}
