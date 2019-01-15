package main

import (
	"MosaicImg/utils"
	"fmt"
)

func main(){
	//modules.GetImg(1)
	a := []byte{1,2,3}
	b := []byte{1,2,3}
	if utils.IsArrayEqual(a,b,3){
		fmt.Println("Equal")
	}
}
