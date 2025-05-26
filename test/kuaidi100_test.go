package test

import (
	"fmt"
	"github.com/linbe-ff/express-go-sdk"
	"testing"
)

func TestKuaidi100Addr(t *testing.T) {

	// 更改为您的Key和Secret等
	kdClient := express.NewKuaiDi100(
		"gkbPFJccxxxx", "ed4133d954ec4cxxxxxx", "57BD27187D74A8778xxxxxxxxxxxxxxxxx",
		"05476ddac4f14e2xxxxxxxxxxxxxx")

	resolution, err := kdClient.AddressResolution(&express.AddressResolutionParam{
		Content: "张三广东省深圳市南山区粤海街道科技南十二路金蝶软件园13088888888",
	})
	if err != nil {
		t.Error(err)
	}
	fmt.Println(resolution)
}
