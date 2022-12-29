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

func TestHospitalUpload(t *testing.T) {
	t.Log("========== HospitalUpload ==========")
	var p = alipay.HospitalUploadReq{}
	p.IsvPid = "2088541479049757"
	p.RequestId = "20221229100001"

	p.HospitalList = append(p.HospitalList, &alipay.HospitalData{
		HospitalName:          "北京怀柔医院",
		HospitalId:            "100001",
		HospitalAlias:         "怀柔医院",
		HospitalStandardCode:  "27110001",
		HospitalProvince:      "北京",
		HospitalCity:          "北京市",
		HospitalDistrict:      "怀柔区",
		HospitalAddr:          "北京市怀柔区永泰北街9号院",
		HospitalLgt:           "116.4133800",
		HospitalLat:           "39.9109200",
		HospitalType:          "公立",
		HospitalGrade:         "三级",
		HospitalWorkTime:      "每天0:00放7天后的号",
		HospitalTel:           "010-60686699",
		HospitalLogo:          "https://hryy.olancloud.com/misc/public/hospital/logo.png",
		CountryKeyDepartment:  "",
		ProvinceKeyDepartment: "",
		KeyDepartment:         "",
		HospitalOptag:         "",
	})

	r, err := json.Marshal(p)
	fmt.Println(string(r))

	rsp, err := client.HospitalUpload(p)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(rsp.Content)
}

func TestDepartmentUpload(t *testing.T) {
	t.Log("========== DepartmentUpload ==========")
	var p = alipay.DepartmentUploadReq{}
	p.RequestId = "20221229200001"
	p.IsvPid = "2088541479049757"
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
}
