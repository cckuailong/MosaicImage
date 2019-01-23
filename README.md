# MosiacImage

## MosaicImage是什么
大家一定看过楚门的世界，他的宣传海报就是马赛克图片。
所谓马赛克图片，就是一张图片由千千万万个小图片组成。
本项目可以使用户只需输入想构成马赛克图片的小图的类别和原始图片（要做成马赛克图片的原图），即可自动制作出马赛克图片

## 功能
1. 自动获取下载某一类别的图片（大量图片）

2. 使用下载的图片，制作马赛克图片

## 配置参数
```
[GET_IMG]
WEB_URL=http://huaban.com/explore/      // 图片源为花瓣网
KEY_WORD=longmao                        // 小图的类别
IMG_URL=http://img.hb.aicdn.com/        // 默认
IMG_DIR=./imgs/                         // 下载的小图存放目录

[MOSAIC_IMG]
IMG_WIDTH=6000                          // 马赛克图片宽度
IMG_HEIGHT=6000                         // 马赛克图片高度
SLICE_SIZE=60                           // 小图大小
SLICE_WIDTH=60                          // 小图宽度
SLICE_HEIGHT=60                         // 小图高度
DIFF=10                                 // 图片相似度， 越低效果越好，但是需要图片数量越大
REPEAT=6                                // 允许重复使用单个小图的次数
```

## 使用说明
更改配置文件中的KEY_WORD字段为要下载的小图类别，执行main.go即可运行
