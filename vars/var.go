package vars

import (
	"MosaicImg/logger"
	"MosaicImg/settings"
	"MosaicImg/utils"
	"os"
)

type (
		DOWNLOAD struct {
			Web_url  string
			Key_word string
			Img_url  string
			Img_dir  string
		}
		IMGINFO struct{
			Img_width    int
			Slice_size   int
			Slice_width  int
			Slice_height int
			Diff         int
			Repeat       int
		}
	)

var(
	Download DOWNLOAD
	ImgInfo IMGINFO
	Headers map[string]string
)

func init(){
	get_img_info := settings.Cfg.Section("GET_IMG")
	Download.Web_url = get_img_info.Key("WEB_URL").MustString("http://huaban.com/explore/")
	Download.Key_word = get_img_info.Key("KEY_WORD").MustString("longmao")
	Download.Img_url = get_img_info.Key("IMG_URL").MustString("http://img.hb.aicdn.com/")
	Download.Img_dir = get_img_info.Key("IMG_DIR").MustString("imgs")

	mosaic_info := settings.Cfg.Section("MOSAIC_IMG")
	ImgInfo.Img_width = mosaic_info.Key("IMG_WIDTH").MustInt(6000)
	ImgInfo.Slice_size = mosaic_info.Key("SLICE_SIZE").MustInt(60)
	ImgInfo.Slice_width = mosaic_info.Key("SLICE_WIDTH").MustInt(60)
	ImgInfo.Slice_height = mosaic_info.Key("SLICE_HEIGHT").MustInt(60)
	ImgInfo.Diff = mosaic_info.Key("DIFF").MustInt(10000)
	ImgInfo.Repeat = mosaic_info.Key("REPEAT").MustInt(6)

	Headers = make(map[string]string)
	Headers["Accept"] = "application/ "
	Headers["User-Agent"] = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.100 Safari/537.36"

	if exist, _:=utils.PathExists(Download.Img_dir);!exist{
		err := os.Mkdir(Download.Img_dir, 0755)
		if err != nil{
			logger.Log.Println("[ Error ] Create Img_dir Err, Check Your 'info.conf'")
		}
	}
}