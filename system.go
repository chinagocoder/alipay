// @Title  commercemedical.go
// @Description  医院科室信息上传接口 https://opendocs.alipay.com/pre-apis/02x1v2
// @Author  xushuai  2022/12/20 3:10 PM

package alipay

func (c *Client) AuthCode(param AuthCodeReq) (result *AuthCodeResp, err error) {
	err = c.doRequest("POST", param, &result)
	return result, err
}
