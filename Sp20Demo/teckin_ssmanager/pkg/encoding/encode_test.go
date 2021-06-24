package encoding

import (
	"fmt"
	"strings"
	"testing"
)

func TestBase64Encode(t *testing.T){
	//str := "e10adc3949ba59abbe56e057f20f883e"
	en := "1f2a999adb12113e2e1a77ac6081a8999"
	fmt.Println(EncodeBase64("1:"+StringMD5(en)))
	fmt.Println(StringMD5(en))
	dn := "MTo0MTBhZmU0ODE1ODJmMDk3ZTg4OTk3ZjE3ZmVlODgwOQ=="
	fmt.Println(DecodeBase64(dn))
	a := strings.Split("1:410afe481582f097e88997f17fee8809", ":")
	fmt.Println(a)
}
