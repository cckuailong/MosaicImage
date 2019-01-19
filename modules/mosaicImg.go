package modules

import (
	"MosaicImg/utils"
	"errors"
	"image"
)

func GetAvgHsv(img image.Image) (float64, float64, float64, error) {
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
