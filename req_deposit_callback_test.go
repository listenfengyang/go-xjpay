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

	req := XJPayCallbackReq{
		OrderNo:       "202603191118200521",
		SignType:      "md5",
		Status:        "success", //success, fail
		OrderPayment:  "323.0",
		OrderAmount:   "45.0",
		OrderCurrency: "USDT",
		TransactionId: "202603191618X2QK",
		CustomerId:    "407",
		Sign:          "ece94497b0d53fb7b4ca15a15c795187",
	}

	err := cli.DepositCallback(req, func(XJPayCallbackReq) error {
		return nil
	})
	if err != nil {
		t.Fatalf("sign verify failed, err: %v", err)
	}

	req.Sign = "invalid-sign"
	err = cli.DepositCallback(req, func(XJPayCallbackReq) error {
		return nil
	})
	if err == nil {
		t.Fatalf("expected sign verify error")
	}
}
