// @Title  trade.go
// @Description  todo
// @Author  xushuai  2022/12/29 3:08 PM

package alipay

func (c *Client) TradeRefund(param TradeRefundReq) (result *TradeRefundResp, err error) {
	err = c.doRequest("POST", param, &result)
	return result, err
}
