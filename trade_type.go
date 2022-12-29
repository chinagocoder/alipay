// @Title  trade_type.go
// @Description  todo
// @Author  xushuai  2022/12/29 3:09 PM

package alipay

type TradeRefundReq struct {
	AppAuthToken string `json:"-"` // 可选

	OutTradeNo   string  `json:"out_trade_no,omitempty"`   // 商户订单号
	TradeNo      string  `json:"trade_no,omitempty"`       // 支付宝交易号。
	RefundAmount float64 `json:"refund_amount"`            // 退款金额。
	RefundReason string  `json:"refund_reason,omitempty"`  // 退款原因说明。
	OutRequestNo string  `json:"out_request_no,omitempty"` // 退款请求号。
}

func (r TradeRefundReq) APIName() string {
	return "alipay.trade.refund"
}

func (r TradeRefundReq) Params() map[string]string {
	var m = make(map[string]string)
	m["app_auth_token"] = r.AppAuthToken
	return m
}

type TradeRefundResp struct {
	Content struct {
		Code            Code    `json:"code"`
		Msg             string  `json:"msg"`
		SubCode         string  `json:"sub_code"`
		SubMsg          string  `json:"sub_msg"`
		TradeNo         string  `json:"trade_no"`
		OutTradeNo      string  `json:"out_trade_no,omitempty"` // 商户订单号
		BuyerLogonId    string  `json:"buyer_logon_id"`
		FundChange      string  `json:"fund_change"`
		RefundFee       float64 `json:"refund_fee"`
		StoreName       string  `json:"store_name"`
		BuyerUserId     string  `json:"buyer_user_id"`
		SendBackFee     string  `json:"send_back_fee"`
		RefundHybAmount string  `json:"refund_hyb_amount"`
	} `json:"alipay_merchant_order_sync_response"`
	Sign string `json:"sign"`
}
