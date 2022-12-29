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
		Code    Code                           `json:"code"`
		Msg     string                         `json:"msg"`
		SubCode string                         `json:"sub_code"`
		SubMsg  string                         `json:"sub_msg"`
		Data    CommerceDataUploadResponseData `json:"data"`
	} `json:"alipay_commerce_medical_industrydata_department_upload_response"`
	Sign string `json:"sign"`
}

type HospitalUploadReq struct {
	AppAuthToken string          `json:"-"`                       // 可选
	IsvPid       string          `json:"isv_pid"`                 // isvpid
	RequestId    string          `json:"request_id"`              // 请求id
	HospitalList []*HospitalData `json:"hospital_list,omitempty"` // 医院列表
}

type HospitalData struct {
	HospitalName          string `json:"hospital_name"`           // 医院名称
	HospitalId            string `json:"hospital_id"`             // 医院id
	HospitalAlias         string `json:"hospital_alias"`          // 医院别名
	HospitalStandardCode  string `json:"hospital_standard_code"`  // 医疗定点机构编号
	HospitalProvince      string `json:"hospital_province"`       // 医院所在省份
	HospitalCity          string `json:"hospital_city"`           // 医院所在城市
	HospitalDistrict      string `json:"hospital_district"`       // 医院所在城市区划
	HospitalAddr          string `json:"hospital_addr"`           // 医院地址
	HospitalLgt           string `json:"hospital_lgt"`            // 医院经度
	HospitalLat           string `json:"hospital_lat"`            // 医院纬度
	HospitalType          string `json:"hospital_type"`           // 医院类型
	HospitalGrade         string `json:"hospital_grade"`          // 医院等级
	HospitalWorkTime      string `json:"hospital_work_time"`      // 医院放号时间
	HospitalTel           string `json:"hospital_tel"`            // 医院对公电话
	HospitalLogo          string `json:"hospital_logo"`           // png格式的医院logo
	CountryKeyDepartment  string `json:"country_key_department"`  // 国家重点科室
	ProvinceKeyDepartment string `json:"province_key_department"` // 省级重点科室
	KeyDepartment         string `json:"key_department"`          // 特色科室
	HospitalOptag         string `json:"hospital_optag"`          // 医院标签
}

func (r HospitalUploadReq) APIName() string {
	return "alipay.commerce.medical.industrydata.hospital.upload"
}

func (r HospitalUploadReq) Params() map[string]string {
	var m = make(map[string]string)
	m["app_auth_token"] = r.AppAuthToken
	return m
}

type HospitalUploadResp struct {
	Content struct {
		Code    Code                           `json:"code"`
		Msg     string                         `json:"msg"`
		SubCode string                         `json:"sub_code"`
		SubMsg  string                         `json:"sub_msg"`
		Data    CommerceDataUploadResponseData `json:"data"`
	} `json:"alipay_commerce_medical_industrydata_hospital_upload_response"`
	Sign string `json:"sign"`
}
