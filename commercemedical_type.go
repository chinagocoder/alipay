// @Title  commercemedical_type.go
// @Description  医疗信息
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

type InquiryDoctorUploadReq struct {
	AppAuthToken string               `json:"-"`                     // 可选
	PlatformCode string               `json:"platform_code"`         // 医生执业问诊平台编码
	OutRequestNo string               `json:"out_request_no"`        // 外部请求号
	DoctorList   []*InquiryDoctorData `json:"doctor_list,omitempty"` // 医生列表
}

type InquiryDoctorData struct {
	DoctorName                    string `json:"doctor_name"`                      // 医生名称
	Gender                        string `json:"gender"`                           // 医生性别 男: MALE 女: FEMALE
	Avatar                        string `json:"avator"`                           // 头像
	PracticingDoctorCertificateNo string `json:"practicing_doctor_certificate_no"` // 医生执医许可证编号（医生执医许可证编号和身份证号至少填一项）
	IdNo                          string `json:"id_no"`                            // 身份证号（医生执医许可证编号和身份证号至少填一项）
	HospitalName                  string `json:"hospital_name"`                    // 执业医院名称
	DepartmentName                string `json:"department_name"`                  // 医生职称 医师: PHYSICIANS 主治医师: ATTENDING_PHYSICIANS 副主任医师: DEPUTY_CHIEF_PHYSICIANS 主任医师: CHIEF_PHYSICIANS
	Title                         string `json:"title"`                            // 科室名称
	DoctorDesc                    string `json:"doctor_desc"`                      // 医生个人简介
	SkilledDesc                   string `json:"skilled_desc"`                     // 医生擅长描述
	SkilledDisease                string `json:"skilled_disease"`                  // 医生擅长疾病（多个擅长疾病时用英文逗号分隔）
	SpecialTags                   string `json:"special_tags"`                     // 医生专业标签（多个标签时用英文逗号分隔）
	PracticeYear                  string `json:"practice_year"`                    // 执医年限 1-5年: WITHIN_FIVE_YEARS 5-10年: WITHIN_TEN_YEARS 10年以上: OVER_TEN_YEARS
	ChangeType                    string `json:"change_type"`                      // 变更类型 新增: INSERT 修改: UPDATE 删除: DELETE
}

func (r InquiryDoctorUploadReq) APIName() string {
	return "alipay.commerce.medical.industrydata.inquirydoctor.upload"
}

func (r InquiryDoctorUploadReq) Params() map[string]string {
	var m = make(map[string]string)
	m["app_auth_token"] = r.AppAuthToken
	return m
}

type InquiryDoctorUploadResp struct {
	Content struct {
		Code    Code   `json:"code"`
		Msg     string `json:"msg"`
		SubCode string `json:"sub_code"`
		SubMsg  string `json:"sub_msg"`
	} `json:"alipay_commerce_medical_industrydata_inquirydoctor_upload_response"`
	Sign string `json:"sign"`
}
