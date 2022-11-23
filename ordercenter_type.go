// @Title  ordercenter.go
// @Description  小程序订单中心接口请求参数 https://opendocs.alipay.com/mini/043zb5
// @Author  xushuai  2022/11/23 10:27 AM

package alipay

type OrderSyncReq struct {
	AppAuthToken string `json:"-"` // 可选

	OutBizNo          string           `json:"out_biz_no"`                    // 外部订单号 out_biz_no唯一对应一笔订单，相同的订单需传入相同的out_biz_no
	RecordId          string           `json:"record_id,omitempty"`           // 商户订单同步记录id
	BuyerId           string           `json:"buyer_id,omitempty"`            // 买家userId,buyer_info与buyer_id必选其一
	BuyerInfo         UserInformation  `json:"buyer_info,omitempty"`          // 买家信息
	ServiceCode       string           `json:"service_code,omitempty"`        // 服务code
	SellerId          string           `json:"seller_id,omitempty"`           // 卖家userId
	PartnerId         string           `json:"partner_id,omitempty"`          // 交易对应的签约商户userId
	Amount            float64          `json:"amount,omitempty"`              // 订单金额，单位为【元】
	PayAmount         float64          `json:"pay_amount,omitempty"`          // 支付金额，单位为【元】
	OrderType         string           `json:"order_type,omitempty"`          // 订单类型，新接入商户统一传入SERVICE_ORDER(服务订单)
	DiscountAmount    float64          `json:"discount_amount,omitempty"`     // 优惠金额，单位为【元】
	TradeNo           string           `json:"trade_no,omitempty"`            // 订单所对应的支付宝交易号
	TradeType         string           `json:"trade_type,omitempty"`          // 交易号类型 1. TRADE-交易，为空默认为TRADE 2. TRANSFER-转账 3. ENTRUST-受托
	PayTimeoutExpress string           `json:"pay_timeout_express,omitempty"` // 支付超时时间，超过时间支付宝自行关闭订单
	ItemOrderList     []*ItemOrderInfo `json:"item_order_list,omitempty"`     // 商品信息列表
	OrderCreateTime   string           `json:"order_create_time,omitempty"`   // 订单创建时间 当order_type为SERVICE_ORDER时必传
	OrderModifiedTime string           `json:"order_modified_time"`           // 订单修改时间
	OrderPayTime      string           `json:"order_pay_time,omitempty"`      // 订单支付时间 当pay_channel为非ALIPAY时，且订单状态已流转到“支付”或支付后时，需要将支付时间传入
	SendMsg           string           `json:"send_msg,omitempty"`            // 是否需要小程序订单代理发送模版消息，Y代表需要发送，N代表不需要发送，不传默认不发送
	SourceApp         string           `json:"source_app,omitempty"`          // 用于区分用户下单的订单来源，如 Alipay-支付宝端内（默认） DingTalk-钉钉小程序
	ExtInfo           []*ExtInfo       `json:"ext_info,omitempty"`            // 商户订单同步记录id
}

type UserInformation struct {
	Name     string     `json:"name,omitempty"`      // 姓名
	Mobile   string     `json:"mobile,omitempty"`    // 手机号
	CertType string     `json:"cert_type,omitempty"` // 证件类型 身份证：IDENTITY_CARD、护照：PASSPORT、军官证：OFFICER_CARD、士兵证：SOLDIER_CARD、户口本：HOKOU等。如有其它类型需要支持，请与蚂蚁金服工作人员联系。
	CertNo   string     `json:"cert_no,omitempty"`   // 证件号
	UserId   string     `json:"user_id,omitempty"`   // 支付宝uid
	ExtInfo  []*ExtInfo `json:"ext_info,omitempty"`  // 扩展信息
}

type ItemOrderInfo struct {
	SkuId      string     `json:"sku_id,omitempty"`      // 商品 sku id
	ItemId     string     `json:"item_id,omitempty"`     // 商品id
	ItemName   string     `json:"item_name,omitempty"`   // 商品名称
	UnitPrice  string     `json:"unit_price,omitempty"`  // 商品单价（单位：元）
	Quantity   int64      `json:"quantity,omitempty"`    // 商品数量（单位：自拟）
	Unit       string     `json:"unit,omitempty"`        // 商品规格
	Status     string     `json:"status,omitempty"`      // 商品状态枚举，默认无需传入，如需使用请联系业务负责人
	StatusDesc string     `json:"status_desc,omitempty"` // 商品状态描述，默认无需传入，如需使用请联系业务负责人
	ExtInfo    []*ExtInfo `json:"ext_info,omitempty"`    // 扩展信息
}

type ExtInfo struct {
	ExtKey   string `json:"ext_key"`   // 订单修改时间
	ExtValue string `json:"ext_value"` // 订单修改时间
}

func (r OrderSyncReq) APIName() string {
	return "alipay.merchant.order.sync"
}

func (r OrderSyncReq) Params() map[string]string {
	var m = make(map[string]string)
	m["app_auth_token"] = r.AppAuthToken
	return m
}

type OrderSyncResp struct {
	Content struct {
		Code             Code                `json:"code"`
		Msg              string              `json:"msg"`
		SubCode          string              `json:"sub_code"`
		SubMsg           string              `json:"sub_msg"`
		RecordId         string              `json:"record_id"`
		OrderId          string              `json:"order_id"`
		OrderStatus      string              `json:"order_status"`
		DistributeResult []*DistributeResult `json:"distribute_result"`
		SyncSuggestions  []*SyncSuggestions  `json:"sync_suggestions"`
	} `json:"alipay_merchant_order_sync_response"`
	Sign string `json:"sign"`
}

type DistributeResult struct {
	SceneCode           Code `json:"scene_code"`
	SceneName           Code `json:"scene_name"`
	NotDistributeReason Code `json:"not_distribute_reason"`
}

type SyncSuggestions struct {
	Type    Code `json:"type"`
	Message Code `json:"message"`
}
