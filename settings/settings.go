package settings

import (
	"MosaicImg/logger"
	"gopkg.in/ini.v1"
)

var(
	Cfg *ini.File
)

func init(){
	var err error
	src := "conf/info.conf"
	Cfg, err = ini.Load(src)
	if err != nil{
		logger.Log.Println(err)
	}
}
