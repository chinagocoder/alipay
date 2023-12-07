// @Title  minimessage.go
// @Description  支付宝小程序消息 https://opendocs.alipay.com/mini/introduce/message?pathHash=1ff14d91
// @Author  xushuai  2023/12/07 3:08 PM

package alipay

func (c *Client) TemplateMessageSend(param TemplateMessageSendReq) (result *TemplateMessageSendResp, err error) {
	err = c.doRequest("POST", param, &result)
	return result, err
}
