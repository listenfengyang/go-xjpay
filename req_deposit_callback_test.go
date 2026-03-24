package go_xjpay

import "testing"

func TestDepositCallbackVerifySign(t *testing.T) {
	vLog := VLog{}
	params := &XJPayInitParams{
		GatewayUrl: GATEWAY_URL,
		DepositUrl: DEPOSIT_URL,
		ReceiveUrl: RECEIVE_URL,
		PickupUrl:  PICKUP_URL,
		Md5Key:     MD5_KEY,
		CustomerId: CUSTOMER_ID,
	}

	cli := NewClient(vLog, params)
	// type=1&orderNo=741336981235&state=success&signCheck=02fed0230a0cbcd06b81bc18e1ccc248
	req := XJPayCallbackReq{
		OrderNo:   "741336981235",
		Type:      "1",
		Status:    "success", //success, fail
		SignCheck: "02fed0230a0cbcd06b81bc18e1ccc248",
	}

	err := cli.DepositCallback(req, func(XJPayCallbackReq) error {
		return nil
	})
	if err != nil {
		t.Fatalf("sign verify failed, err: %v", err)
	}
}
