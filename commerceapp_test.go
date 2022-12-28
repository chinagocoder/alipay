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

func TestCommerceAppUpload(t *testing.T) {
	t.Log("========== OrderSyncQuery ==========")
	var p = alipay.CommerceAppUploadReq{}
	p.AuthToken = "hosodnoB42c1f5dc49a841acb45787fffae31X00"
	p.ServiceName = "alipay.commerce.app.data"
	p.TargetId = "2088522069220008"

	p.Content = alipay.CommerceAppUploadContent{
		TenantAppId: "20221213497500041974",
		ActivityId:  "upload_hospital_order",
		Body:        "{\"amount\":\"100.00\",\"pay_amount\":\"100.00\",\"buyer_id\":\"2088522069220008\",\"partner_id\":\"2088001544678234\",\"tiny_app_id\":\"2021003169625235\",\"order_create_time\":\"2022-12-27 20:47:54\",\"order_modified_time\":\"2022-12-28 14:47:54\",\"order_type\":\"HOSPITAL_ORDER\",\"out_biz_no\":\"TG12345678990002\",\"out_biz_type\":\"HOSPITAL_APPOINTMENT\",\"merchant_order_status\":\"MERCHANT_PREORDER_SUCCESS\",\"ext_info\":{\"patient\":\"陈晋辉\",\"department\":\"内科\",\"dept_loc\":\"3楼4诊区\",\"dept_num\":\"2\",\"doctor\":\"王二小\",\"doctor_id\":\"123456\",\"doctor_rank\":\"主任医师\",\"hospital\":\"北京xx医院\",\"hospital_register_id\":\"008431110227510151\",\"medical_num\":\"199\",\"merchant_order_link_page\":\"alipays://platformapi/startapp?appId=2021003169625235&page=pagesE/inpatient-service/inpatient-list\",\"scheduled_time\":\"2020-12-29 20:47:54\"},\"item_order_list\":[{\"sku_id\":\"99988\",\"item_name\":\"预约挂号\",\"quantity\":\"1\",\"unit_price\":\"0.99\"}]}",
	}

	r, err := json.Marshal(p)
	fmt.Println(string(r))

	rsp, err := client.CommerceAppUpload(p)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(rsp)
	//if rsp.Content.Code != alipay.CodeSuccess {
	//	t.Fatal(rsp.Content.Msg, rsp.Content.SubMsg)
	//}
	//t.Log(rsp.Content.OrderId)
}
