// @Title  commercemedical_type.go
// @Description  医院科室信息上传接口 https://opendocs.alipay.com/pre-apis/02x1v2
// @Author  xushuai  2022/12/20 3:11 PM

package alipay

type DepartmentUploadReq struct {
	AppAuthToken   string            `json:"-"`                         // 可选
	IsvPid         string            `json:"isv_pid"`                   // isvpid
	RequestId      string            `json:"request_id"`                // 请求id
	DepartmentList []*DepartmentData `json:"department_list,omitempty"` // 请求id
}

type DepartmentData struct {
	DepartmentName        string `json:"department_name"`         // 科室名称
	DepartmentId          string `json:"department_id"`           // 科室id
	DepartmentType        string `json:"department_type"`         // 科室类型
	HospitalName          string `json:"hospital_name"`           // 科室所属医院名称
	ParentDepartmentName  string `json:"parent_department_name"`  // 所属上级科室名称
	ParentDepartmentId    string `json:"parent_department_id"`    // 所属上级科室id
	DepartmentTag         string `json:"department_tag"`          // 科室标签
	DepartmentDiseaseInfo string `json:"department_disease_info"` // 呼吸道疾病
	DepartmentUrl         string `json:"department_url"`          // 科室挂号服务跳转
}

func (r DepartmentUploadReq) APIName() string {
	return "alipay.commerce.medical.industrydata.department.upload"
}

func (r DepartmentUploadReq) Params() map[string]string {
	var m = make(map[string]string)
	m["app_auth_token"] = r.AppAuthToken
	return m
}

type DepartmentUploadResp struct {
	Content struct {
		Data []*CommerceDataUploadResponseData `json:"data"`
	} `json:"alipay_commerce_medical_industrydata_hospital_upload_response"`
	Sign string `json:"sign"`
}
