package go_xjpay

type XJPayInitParams struct {
	GatewayUrl    string `json:"gatewayUrl" mapstructure:"gatewayUrl" config:"gatewayUrl" yaml:"gatewayUrl"`
	ReceiveUrl    string `json:"receiveUrl" mapstructure:"receiveUrl" config:"receiveUrl" yaml:"receiveUrl"`
	PickupUrl     string `json:"pickupUrl" mapstructure:"pickupUrl" config:"pickupUrl" yaml:"pickupUrl"`
	CustomerId    string `json:"customerId" mapstructure:"customerId" config:"customerId" yaml:"customerId"`
	OrderCurrency string `json:"orderCurrency" mapstructure:"orderCurrency" config:"orderCurrency" yaml:"orderCurrency"`
	SignType      string `json:"signType" mapstructure:"signType" config:"signType" yaml:"signType"`
	Md5Key        string `json:"md5Key" mapstructure:"md5Key" config:"md5Key" yaml:"md5Key"`
	DepositUrl    string `json:"depositUrl" mapstructure:"depositUrl" config:"depositUrl" yaml:"depositUrl"`
}

type XJPayDepositReq struct {
	OrderNo       string `json:"orderNo" form:"orderNo" mapstructure:"orderNo"`
	OrderAmount   string `json:"orderAmount" form:"orderAmount" mapstructure:"orderAmount"`
	ExchangeRate  string `json:"exchangeRate" form:"exchangeRate" mapstructure:"exchangeRate"`
	PayName       string `json:"payName" form:"payName" mapstructure:"payName"`
	ReceiveUrl    string `json:"receiveUrl" form:"receiveUrl" mapstructure:"receiveUrl"`
	PickupUrl     string `json:"pickupUrl" form:"pickupUrl" mapstructure:"pickupUrl"`
	OrderCurrency string `json:"orderCurrency" form:"orderCurrency" mapstructure:"orderCurrency"`
	CustomerId    string `json:"customerId" form:"customerId" mapstructure:"customerId"`
	SignType      string `json:"signType" form:"signType" mapstructure:"signType"`
}

type XJPayDepositRsp struct {
	HttpStatusCode int    `json:"httpStatusCode" mapstructure:"httpStatusCode"`
	ResponseBody   string `json:"responseBody" mapstructure:"responseBody"`
}

type XJPayCallbackReq struct {
	Type       string `json:"type" form:"type" mapstructure:"type"`
	OrderNo    string `json:"orderNo" form:"orderNo" mapstructure:"orderNo"`
	State      string `json:"state" form:"state" mapstructure:"state"`
	PickupUrl  string `json:"pickupUrl" form:"pickupUrl" mapstructure:"pickupUrl"`
	ReceiveUrl string `json:"receiveUrl" form:"receiveUrl" mapstructure:"receiveUrl"`
	SignCheck  string `json:"signCheck" form:"signCheck" mapstructure:"signCheck"`
}
