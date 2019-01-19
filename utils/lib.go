package utils

import (
	"crypto/md5"
	"fmt"
	"io"
	"math"
	"os"
)

// Md5 func
func MD5(s string) (m string) {
	h := md5.New()
	_, _ = io.WriteString(h, s)
	return fmt.Sprintf("%x", h.Sum(nil))
}

// Round the float64
func Round(f float64, n int) float64 {
	pow10_n := math.Pow10(n)
	return math.Trunc((f+0.5/pow10_n)*pow10_n) / pow10_n
}

// Whether the dir exists
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// Whether first len of th two arrays is equal
func IsArrayEqual(a, b []byte, len int) bool{
	for i:=0;i<len;i++{
		if a[i] != b[i]{
			return false
		}
	}
	return true
}

// Get the Image's suffix
func GetImgSuffix(img_head []byte) string{
	if IsArrayEqual([]byte{0xff,0xd8}, img_head, 2){
		return ".jpg"
	}
	if IsArrayEqual([]byte{0x89,0x50,0x4e,0x47,0x0D,0x0A,0x1A,0x0A}, img_head, 8){
		return ".png"
	}
	if IsArrayEqual([]byte{0x42, 0x4d}, img_head, 2){
		return ".bmp"
	}
	if IsArrayEqual([]byte{0x47,0x49,0x46,0x38,0x39,0x61}, img_head, 6){
		return ".gif"
	}
	return "unknown"
}

