package modules

import (
	"MosaicImg/logger"
	"MosaicImg/utils"
	"MosaicImg/vars"
	"github.com/bitly/go-simplejson"
	"io/ioutil"
	"net/url"
	"regexp"
)

// Get the Search result with the KeyWord
func getSearchRes(params url.Values)string{
	resp, err := utils.Http_req(vars.Download.Web_url+vars.Download.Key_word, params, "GET", vars.Headers)
	if err != nil{
		logger.Log.Println("[ Error ] Get Search Result Err")
	}
	body, _ := ioutil.ReadAll(resp.Body)
	return string(body)
}

// Get the key from the Search result
func getKey(text string)[]string{
	keys := []string{}
	re := regexp.MustCompile(`app.page\["pins"\] = (.+?)\];`)
	data := re.FindString(text)
	data = data[19:len(data)-1]
	list, _ :=  simplejson.NewJson([]byte(data))
	items, _ := list.Array()
	for i, _ := range(items){
		keys = append(keys, list.GetIndex(i).Get("file").Get("key").MustString())
	}
	return keys
}

// Download Images
func downloadImg(keys []string){
	for _, key := range(keys){
		uri := vars.Download.Img_url + key
		resp, err := utils.Http_req(uri, nil, "GET", vars.Headers)
		if err != nil{
			logger.Log.Println("[ Error ] Download Err happens, Key: "+key)
			continue
		}
		body, _ := ioutil.ReadAll(resp.Body)
		suffix := utils.GetImgSuffix(body[:10])
		img_name := utils.MD5(key)+suffix
		if ioutil.WriteFile(vars.Download.Img_dir+img_name, body, 644) != nil{
			logger.Log.Println("[ Warning ] Download Image Err, URI: "+uri)
			continue
		}
	}

}

// Get the Images
func GetImg(page int){
	params := url.Values{
		"jn8pvql9":{""},
		"page":{string(page)},
		"per_page":{"20"},
		"wfl":{"1"},
	}
	text := getSearchRes(params)
	keys := getKey(text)
	downloadImg(keys)
}
