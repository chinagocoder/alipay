// @Title  commercemedical.go
// @Description  医院科室信息上传接口 https://opendocs.alipay.com/pre-apis/02x1v2
// @Author  xushuai  2022/12/20 3:10 PM

package alipay_test

import (
	"encoding/json"
	"fmt"
	"github.com/chinagocoder/alipay"
	"testing"
)

func TestAuthCode(t *testing.T) {
	t.Log("========== OrderSyncQuery ==========")
	var p = alipay.AuthCodeReq{}
	p.GrantType = "authorization_code"
	p.Code = "dc82be88fea84948bf28c25f72d3ZX93"
	p.RefreshToken = "201208134b203fe6c11548bcabd8da5bb087a83b"

	r, err := json.Marshal(p)
	fmt.Println(string(r))

	rsp, err := client.AuthCode(p)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(rsp.Content)
	//if rsp.Content.Code != alipay.CodeSuccess {
	//	t.Fatal(rsp.Content.Msg, rsp.Content.SubMsg)
	//}
	//t.Log(rsp.Content.OrderId)
}
