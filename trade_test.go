// @Title  trade_test.go
// @Description  todo
// @Author  xushuai  2022/12/29 3:20 PM

package alipay_test

import (
	"encoding/json"
	"fmt"
	"github.com/chinagocoder/alipay"
	"testing"
)

func TestTradeRefund(t *testing.T) {
	t.Log("========== OrderSyncQuery ==========")
	var p = alipay.TradeRefundReq{}
	p.TradeNo = "2022122922001419931442635791"
	p.RefundAmount = 57.80
	p.RefundReason = "七天无理退货"
	p.OutRequestNo = "TH2022122900000001"

	r, err := json.Marshal(p)
	fmt.Println(string(r))

	rsp, err := client.TradeRefund(p)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(rsp.Content)
	//if rsp.Content.Code != alipay.CodeSuccess {
	//	t.Fatal(rsp.Content.Msg, rsp.Content.SubMsg)
	//}
	//t.Log(rsp.Content.OrderId)
}
