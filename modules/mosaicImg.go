package modules

import (
	"MosaicImg/logger"
	"MosaicImg/utils"
	"MosaicImg/vars"
	"errors"
	"fmt"
	"github.com/disintegration/imaging"
	"image"
	"image/color"
	"image/draw"
	"math"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type HSVPR struct {
	h float64
	s float64
	v float64
	p string
	rep int
}

// Cal single image's average HSV
func getAvgHsv(img image.Image) (float64, float64, float64, error) {
	var h_sum, s_sum, v_sum float64
	count := 0
	width := img.Bounds().Dx()
	height := img.Bounds().Dy()
	for i:=0;i<width;i++{
		for j:=0;j<height;j++{
			pixel := img.At(i,j)
			r, g, b ,_ := pixel.RGBA()
			h, s, v := utils.Rgb2Hsv(r>>8,g>>8,b>>8)
			h /= 360
			h_sum += h
			s_sum += s
			v_sum += v
			count++
		}
	}
	if count > 0{
		h_avg := utils.Round(h_sum/float64(count), 3)
		s_avg := utils.Round(s_sum/float64(count), 3)
		v_avg := utils.Round(v_sum/float64(count), 3)
		return h_avg, s_avg, v_avg, nil
	}else{
		return 0, 0, 0, errors.New("Read Img Err")
	}
}

// Rename all the images in the img_dir
func RenameAllWithRsv(){
	complete := make(chan int)
	list := []string{}
	err := filepath.Walk(vars.Download.Img_dir, func(path string, f os.FileInfo, err error) error {
		if f == nil{
			return err
		}
		if f.IsDir(){
			return nil
		}
		if path != "imgs\\default.jpg"{
			list = append(list, path)
		}
		return nil
	})
	if err != nil {
		logger.Log.Printf("[ Error ] Read the img dir Error")
	}
	// handle each img
	for _,p := range(list){
		go renameSingleImg(p, complete)
	}
	for i:=0;i<len(list);i++{
		<- complete
	}
}

// Rename the image with its average HSV
func renameSingleImg(path string, complete chan int){
	img, err0 := imaging.Open(path)
	if err0 != nil{
		logger.Log.Printf("[ Warning ] Open single image Err")
		os.Remove(path)
		complete <- 0
		return
	}
	new_img := imaging.Resize(img, vars.ImgInfo.Slice_width, vars.ImgInfo.Slice_height, imaging.Lanczos)
	h,s, v, err1:=getAvgHsv(new_img)
	if err1 != nil {
		logger.Log.Printf("[ Warning ] Get single image avg HSV Err")
		os.Remove(path)
		complete <- 0
		return
	}
	err2 := imaging.Save(new_img, fmt.Sprintf("%v%v_%v_%v%v",vars.Download.Img_dir, h, s, v, path[len(path)-4:]))
	if err2 != nil {
		logger.Log.Printf("[ Warning ] Save single image Err")
		os.Remove(path)
		complete <- 0
		return
	}
	err3 := os.Remove(path)
	if err3 != nil {
		logger.Log.Printf("[ Warning ] Remove single pre image Err")
		return
	}
	complete <- 0
}

// Get all the images's HSV, return a list
func GetHsvList()[]HSVPR{
	list := []HSVPR{}
	err := filepath.Walk(vars.Download.Img_dir, func(path string, f os.FileInfo, err error) error {
		if f == nil{
			return err
		}
		if f.IsDir(){
			return nil
		}
		if path == "imgs\\default.jpg"{
			return nil
		}
		pos := strings.Index(path, vars.Download.Img_dir)
		pos += len(vars.Download.Img_dir)
		filename := path[pos:len(path)-4]
		hsv := strings.Split(filename, "_")
		hsvpr := HSVPR{}
		hsvpr.h, err = strconv.ParseFloat(hsv[0], 64)
		hsvpr.s, err = strconv.ParseFloat(hsv[1], 64)
		hsvpr.v, err = strconv.ParseFloat(hsv[2], 64)
		hsvpr.p = path[len(path)-4:]
		hsvpr.rep = 0
		list = append(list, hsvpr)
		return nil
	})
	if err != nil {
		logger.Log.Printf("[ Error ] Get HSV List Err")
	}
	return list
}

// Find the most similar Slice with HSV
func findClosestHsv(hsvpr HSVPR, list []HSVPR)string{
	var similar HSVPR
	selected := false
	for i, item := range(list){
		diff := math.Sqrt(math.Pow(hsvpr.h-item.h,2) + math.Pow(hsvpr.s-item.s,2) + math.Pow(hsvpr.v-item.v,2))
		if diff < vars.ImgInfo.Diff && item.rep < vars.ImgInfo.Repeat{
			similar = item
			selected = true
			list[i].rep += 1
			break
		}
	}
	if !selected{
		return "default.jpg"
	}
	return fmt.Sprintf("%v_%v_%v%v", similar.h, similar.s, similar.v, similar.p)
}

// Create the Mosaic image
func CreateMosaicImg(img image.Image, list []HSVPR) image.Image{
	var err error
	cnt := 1
	slice_h := vars.ImgInfo.Slice_height
	slice_w := vars.ImgInfo.Slice_width
	hsvpr := HSVPR{}
	canvas := imaging.New(vars.ImgInfo.Img_width, vars.ImgInfo.Img_height, color.White)
	wid_cnt := vars.ImgInfo.Img_width / vars.ImgInfo.Slice_width
	h_cnt := vars.ImgInfo.Img_height / vars.ImgInfo.Slice_height
	logger.Log.Println("[ Info ] Total Slices:", h_cnt*wid_cnt)
	for h:=0;h<h_cnt;h++{
		for w:=0;w<wid_cnt;w++{
			logger.Log.Println("[ Info ] Schedule:", cnt)
			cnt++
			rec := image.Rectangle{image.Point{h*slice_h,w*slice_w},image.Point{(h+1)*slice_h,(w+1)*slice_w}}
			slice := imaging.Crop(img, rec)
			hsvpr.h, hsvpr.s, hsvpr.v, err = getAvgHsv(slice)
			if err != nil{
				logger.Log.Println("[ Error ] Get single image avg HSV Err")
			}
			sim_img_name := findClosestHsv(hsvpr, list)
			fmt.Println(sim_img_name)
			slice_img, err := imaging.Open(vars.Download.Img_dir+sim_img_name)
			if err != nil{
				logger.Log.Println("[ Error ] Open Slice image Err")
			}
			draw.Draw(canvas, rec, slice_img, image.Point{0,0},draw.Src)
		}
	}
	return canvas
}


