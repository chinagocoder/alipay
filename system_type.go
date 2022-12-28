// @Title  commercemedical_type.go
// @Description  医院科室信息上传接口 https://opendocs.alipay.com/pre-apis/02x1v2
// @Author  xushuai  2022/12/20 3:11 PM

package alipay

type AuthCodeReq struct {
	AppAuthToken string `json:"-"` // 可选
	GrantType    string `json:"-"` // 授权方式
	Code         string `json:"-"` // 授权码
	RefreshToken string `json:"-"` // 刷新令牌
}

func (r AuthCodeReq) APIName() string {
	return "alipay.system.oauth.token"
}

func (r AuthCodeReq) Params() map[string]string {
	var m = make(map[string]string)
	m["app_auth_token"] = r.AppAuthToken
	m["grant_type"] = r.GrantType
	m["code"] = r.Code
	m["refresh_token"] = r.RefreshToken
	return m
}

type AuthCodeResp struct {
	Content struct {
		Code         Code   `json:"code"`
		Msg          string `json:"msg"`
		SubCode      string `json:"sub_code"`
		SubMsg       string `json:"sub_msg"`
		UserId       string `json:"user_id"`       // 支付宝用户的唯一标识
		AccessToken  string `json:"access_token"`  // 访问令牌
		ExpiresIn    string `json:"expires_in"`    // 访问令牌的有效时间，单位是秒
		RefreshToken string `json:"refresh_token"` // 刷新令牌
		ReExpiresIn  string `json:"re_expires_in"` // 刷新令牌的有效时间，单位是秒。
		AuthStart    string `json:"auth_start"`    // 授权token开始时间
	} `json:"alipay_commerce_app_upload_response"`
	Sign string `json:"sign"`
}
