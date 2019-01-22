package main

import (
	"MosaicImg/modules"
	"fmt"
)

func main(){
	//img, err := imaging.Open("imgs/0a8e21e1362e4b0c95d14c96be9e302e.jpg")
	//if err != nil{
	//	return
	//}
	//h,s, v, err:=modules.GetAvgHsv(img)
	//fmt.Println(h,s,v)
	fmt.Println(modules.GetHsvList())
}
