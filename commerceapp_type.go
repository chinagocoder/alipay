// @Title  commercemedical_type.go
// @Description  医院科室信息上传接口 https://opendocs.alipay.com/pre-apis/02x1v2
// @Author  xushuai  2022/12/20 3:11 PM

package alipay

import "encoding/json"

type CommerceAppUploadReq struct {
	AppAuthToken string                   `json:"-"` // 可选
	AuthToken    string                   `json:"-"` // 用户授权令牌
	ServiceName  string                   `json:"-"` // 应用服务名称
	TargetId     string                   `json:"-"` // 目标用户id
	Content      CommerceAppUploadContent `json:"-"` // 服务数据
}

type CommerceAppUploadContent struct {
	TenantAppId string `json:"tenant_app_id"`   // 租户应用ID
	ActivityId  string `json:"activity_id"`     // 业务流程ID
	Query       string `json:"query,omitempty"` // 查询条件
	Body        string `json:"body"`            // 业务流程参数
}

type CommerceAppUploadBody struct {
	TenantAppId string `json:"tenant_app_id"`   // 租户应用ID
	ActivityId  string `json:"activity_id"`     // 业务流程ID
	Query       string `json:"query,omitempty"` // 查询条件
}

func (r CommerceAppUploadReq) APIName() string {
	return "alipay.commerce.app.auth.upload"
}

func (r CommerceAppUploadReq) Params() map[string]string {
	var m = make(map[string]string)
	m["app_auth_token"] = r.AppAuthToken
	m["auth_token"] = r.AuthToken
	m["service_name"] = r.ServiceName
	m["target_id"] = r.TargetId
	jsonContent, _ := json.Marshal(r.Content)
	m["content"] = string(jsonContent)
	return m
}

type CommerceAppUploadResp struct {
	Content struct {
		Code    Code   `json:"code"`
		Msg     string `json:"msg"`
		SubCode string `json:"sub_code"`
		SubMsg  string `json:"sub_msg"`
		BizCode string `json:"biz_code"`
		BizMsg  string `json:"biz_msg"`
		Data    string `json:"data"`
	} `json:"alipay_commerce_app_upload_response"`
	Sign string `json:"sign"`
}
