package modules

import (
	"MosaicImg/utils"
	"MosaicImg/vars"
	"errors"
	"fmt"
	"github.com/disintegration/imaging"
	"image"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type HSVP struct {
	h float64
	s float64
	v float64
	p string
}

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

func renameAllWithRsv(){
	complete := make(chan int)
	list := []string{}
	err := filepath.Walk(vars.Download.Img_dir, func(path string, f os.FileInfo, err error) error {
		if f == nil{
			return err
		}
		if f.IsDir(){
			return nil
		}
		list = append(list, path)
		return nil
	})
	if err != nil {
		fmt.Printf("filepath.Walk() returned %v\n", err)
	}
	// handle each img
	for _,p := range(list){
		go renameSingleImg(p, complete)
	}
	for i:=0;i<len(list);i++{
		<- complete
	}
}

func renameSingleImg(path string, complete chan<- int){
	img, err0 := imaging.Open(path)
	if err0 != nil{
		fmt.Printf("Img open %v\n", err0)
	}
	new_img := imaging.Resize(img, vars.ImgInfo.Slice_width, vars.ImgInfo.Slice_width, imaging.Lanczos)
	h,s, v, err1:=getAvgHsv(new_img)
	if err1 != nil {
		fmt.Printf("Get avg hsv %v\n", err1)
	}
	err2 := imaging.Save(new_img, fmt.Sprintf("%v%v_%v_%v%v",vars.Download.Img_dir, h, s, v, path[len(path)-4:]))
	if err2 != nil {
		fmt.Printf("Save img %v\n", err1)
	}
	err3 := os.Remove(path)
	if err3 != nil {
		fmt.Printf("Remove pre img %v\n", err1)
	}
	complete <- 0
}

func GetHsvList()[]HSVP{
	list := []HSVP{}
	err := filepath.Walk(vars.Download.Img_dir, func(path string, f os.FileInfo, err error) error {
		if f == nil{
			return err
		}
		if f.IsDir(){
			return nil
		}
		pos := strings.Index(path, vars.Download.Img_dir)
		pos += len(vars.Download.Img_dir)
		filename := path[pos:len(path)-4]
		hsv := strings.Split(filename, "_")
		hsvp := HSVP{}
		hsvp.h, err = strconv.ParseFloat(hsv[0], 64)
		hsvp.s, err = strconv.ParseFloat(hsv[1], 64)
		hsvp.v, err = strconv.ParseFloat(hsv[2], 64)
		hsvp.p = path[len(path)-4:]
		list = append(list, hsvp)
		return nil
	})
	if err != nil {
		fmt.Printf("filepath.Walk() returned %v\n", err)
	}
	return list
}

func FindClosestHsv(hsvp HSVP, hsvp_l []HSVP)string{
	for _, 
}
