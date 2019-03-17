package testing

import (
	"fmt"
	"github.com/Farteen/travelfinance/util"
	"testing"
)

func TestMd5(t *testing.T) {
	abcStr := "123456"
	resultStr := util.MD5String(abcStr)
	//md5Str := string(md5Bytes)
	fmt.Printf("%v", resultStr)
	//fmt.Println(md5Str)
}
