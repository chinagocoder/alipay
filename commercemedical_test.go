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

func TestClient_DepartmentUpload(t *testing.T) {
	t.Log("========== OrderSyncQuery ==========")
	var p = alipay.DepartmentUploadReq{}
	p.RequestId = "1111111111111"
	p.IsvPid = "12222222222"
	p.DepartmentList = append(p.DepartmentList, &alipay.DepartmentData{
		DepartmentName: "妇科(主任医师)",
		DepartmentId:   "013",
		DepartmentType: "一级科室",
		HospitalName:   "测试医院",
	})

	r, err := json.Marshal(p)
	fmt.Println(string(r))

	rsp, err := client.DepartmentUpload(p)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(rsp.Content)
	//if rsp.Content.Code != alipay.CodeSuccess {
	//	t.Fatal(rsp.Content.Msg, rsp.Content.SubMsg)
	//}
	//t.Log(rsp.Content.OrderId)
}
