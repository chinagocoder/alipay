package alipay_test

import (
	"encoding/json"
	"fmt"
	alipay "github.com/chinagocoder/alipay"
	"testing"
	"time"
)

func TestClient_OrderSync(t *testing.T) {
	t.Log("========== OrderSyncQuery ==========")
	var p = alipay.OrderSyncReq{}
	p.OrderType = "SERVICE_ORDER"
	p.OutBizNo = "2022112201010007"
	p.OrderCreateTime = "2022-11-22 08:00:00"
	p.OrderModifiedTime = time.Now().Format("2006-01-02 15:04:05")
	p.BuyerId = "2088002227719932"
	p.SourceApp = "Alipay"
	p.Amount = 100

	p.ItemOrderList = append(p.ItemOrderList, &alipay.ItemOrderInfo{ItemName: "预约挂号"})

	p.ExtInfo = append(p.ExtInfo, &alipay.ExtInfo{ExtKey: "business_info", ExtValue: "{\"call_num_url\":\"/pages/order/orderDetail/orderDetail?orderId\",\"scheduled_time\":\"2022-11-29 15:00:00\",\"patient\":\"王*\",\"hospital\":\"怀柔妇幼保健院\",\"department\":\"内科\",\"dept_loc\":\"1号楼3层门诊6区\"}"})
	p.ExtInfo = append(p.ExtInfo, &alipay.ExtInfo{ExtKey: "merchant_biz_type", ExtValue: "MEDICAL_APPOINTMENT"})
	p.ExtInfo = append(p.ExtInfo, &alipay.ExtInfo{ExtKey: "merchant_order_status", ExtValue: "PREORDER_SUCCESS"})
	p.ExtInfo = append(p.ExtInfo, &alipay.ExtInfo{ExtKey: "tiny_app_id", ExtValue: appID})
	p.ExtInfo = append(p.ExtInfo, &alipay.ExtInfo{ExtKey: "merchant_order_link_page", ExtValue: "/pages/order/orderDetail/orderDetail?orderId=202108310xxxxxxx"})

	r, err := json.Marshal(p)
	fmt.Println(string(r))

	rsp, err := client.OrderSync(p)
	if err != nil {
		t.Fatal(err)
	}
	if rsp.Content.Code != alipay.CodeSuccess {
		t.Fatal(rsp.Content.Msg, rsp.Content.SubMsg)
	}
	t.Log(rsp.Content.OrderId)
}
