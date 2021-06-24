package encoding

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"time"
)


//当前时间转换的md5值
func HexTimeNow() string {
	h := md5.New()
	h.Write([]byte(time.Now().String()))
	return hex.EncodeToString(h.Sum(nil))
}

//字符串类型转md5
func StringMD5(str string)string{
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}


//base64 加密解密
func EncodeBase64(data string) string {
	return base64.StdEncoding.EncodeToString([]byte(data))
}

func DecodeBase64(sEnc string) string {
	sDec, err := base64.StdEncoding.DecodeString(sEnc)
	if err != nil {
		fmt.Printf("Error decoding string: %s ", err.Error())
		return ""
	}
	return string(sDec)
}

//func Sha256(key string) string{
//	sum := sha256.Sum256([]byte(key))
//	return
//}