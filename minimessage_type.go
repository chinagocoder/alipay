// @Title  minimessage.go
// @Description  支付宝小程序消息 https://opendocs.alipay.com/mini/introduce/message?pathHash=1ff14d91
// @Author  xushuai  2023/12/07 3:08 PM

package alipay

type TemplateMessageSendReq struct {
	AppAuthToken   string `json:"-"`                          // 可选
	ToUserId       string `json:"to_user_id,omitempty"`       // 支付宝userId
	FormId         string `json:"form_id,omitempty"`          // 支付宝交易号/表单号/ftoken
	UserTemplateId string `json:"user_template_id,omitempty"` // 消息模版ID
	Page           string `json:"page,omitempty"`             // 小程序的跳转页面地址
	Data           string `json:"data,omitempty"`             // 消息模板变量内容
}

func (r TemplateMessageSendReq) APIName() string {
	return "alipay.open.app.mini.templatemessage.send"
}

func (r TemplateMessageSendReq) Params() map[string]string {
	var m = make(map[string]string)
	m["app_auth_token"] = r.AppAuthToken
	return m
}

type TemplateMessageSendResp struct {
	Content struct {
		Code    Code   `json:"code"`
		Msg     string `json:"msg"`
		SubCode string `json:"sub_code"`
		SubMsg  string `json:"sub_msg"`
	} `json:"alipay_open_app_mini_templatemessage_send_response"`
	Sign string `json:"sign"`
}
