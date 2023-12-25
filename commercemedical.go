// @Title  commercemedical.go
// @Description  医疗信息
// @Author  xushuai  2022/12/20 3:10 PM

package alipay

// DepartmentUpload
// @Description 医院科室信息上传接口 https://opendocs.alipay.com/pre-apis/02x1v2
// @receiver  c
// @param  param
// @return  result
// @return  err
func (c *Client) DepartmentUpload(param DepartmentUploadReq) (result *DepartmentUploadResp, err error) {
	err = c.doRequest("POST", param, &result)
	return result, err
}

// HospitalUpload
// @Description 医疗行业医院数据上传接口 https://opendocs.alipay.com/pre-apis/02x1v3
// @receiver  c
// @param  param
// @return  result
// @return  err
func (c *Client) HospitalUpload(param HospitalUploadReq) (result *HospitalUploadResp, err error) {
	err = c.doRequest("POST", param, &result)
	return result, err
}

// InquiryDoctorUpload
// @Description 专家问诊医生数据同步接口 https://opendocs.alipay.com/pre-apis/067y9k
// @receiver  c
// @param  param
// @return  result
// @return  err
func (c *Client) InquiryDoctorUpload(param InquiryDoctorUploadReq) (result *InquiryDoctorUploadResp, err error) {
	err = c.doRequest("POST", param, &result)
	return result, err
}

// InquiryDoctorStatusUpload
// @Description 专家问诊医生状态实时同步接口 https://opendocs.alipay.com/pre-apis/067w4m
// @receiver  c
// @param  param
// @return  result
// @return  err
func (c *Client) InquiryDoctorStatusUpload(param InquiryDoctorStatusUploadReq) (result *InquiryDoctorStatusUploadResp, err error) {
	err = c.doRequest("POST", param, &result)
	return result, err
}

// InquiryOrderUpload
// @Description 问诊订单回流接口 https://opendocs.alipay.com/pre-apis/067slc
// @receiver  c
// @param  param
// @return  result
// @return  err
func (c *Client) InquiryOrderUpload(param InquiryOrderUploadReq) (result *InquiryOrderUploadResp, err error) {
	err = c.doRequest("POST", param, &result)
	return result, err
}
