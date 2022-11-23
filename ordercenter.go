// @Title  ordercenter.go
// @Description  小程序订单中心
// @Author  xushuai  2022/11/23 10:27 AM

package alipay

// OrderSync
// @Description  订单同步
// @receiver  c 客户端实例
// @param  param 请求参数
// @return  result 响应结果
// @return  err 错误信息
func (c *Client) OrderSync(param OrderSyncReq) (result *OrderSyncResp, err error) {
	err = c.doRequest("POST", param, &result)
	return result, err
}
