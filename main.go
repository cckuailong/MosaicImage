package main

import (
	"MosaicImg/modules"
	"MosaicImg/vars"
	"fmt"
	"github.com/disintegration/imaging"
)

func main(){
	// Get all the small images, here 20 pages
	for i:=1;i<20;i++{
		modules.GetImg(i)
	}
	modules.RenameAllWithRsv()
	// Open the image you want to transform
	img, err := imaging.Open("test.jpg")
	if err != nil{
		return
	}
	// Standard the size of pre image
	new_img := imaging.Resize(img, vars.ImgInfo.Img_width, vars.ImgInfo.Img_height, imaging.Lanczos)
	hsv_list := modules.GetHsvList()
	res_img := modules.CreateMosaicImg(new_img, hsv_list)
	// Save as result image
	err = imaging.Save(res_img, "result.jpg")
	if err != nil{
		fmt.Println("result Save Failed")
	}
}
