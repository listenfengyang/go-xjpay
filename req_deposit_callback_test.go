package go_xjpay

import (
	"testing"

	"github.com/listenfengyang/go-xjpay/utils"
)

func TestDepositCallbackVerifySign(t *testing.T) {
	params := &XJPayInitParams{
		Md5Key: "OuUaWlFQtD",
	}

	cli := NewClient(VLog{}, params)

	req := XJPayCallbackReq{
		OrderNo:    "74133698",
		PickupUrl:  "http://www.otcpay.eu.cc/customLogin",
		ReceiveUrl: "http://www.otcpay.eu.cc",
	}
	req.SignCheck = utils.Md5Hex(req.OrderNo + req.PickupUrl + req.ReceiveUrl + params.Md5Key)

	err := cli.DepositCallback(req, func(r XJPayCallbackReq) error {
		return nil
	})
	if err != nil {
		t.Fatalf("callback verify failed: %v", err)
	}

	req.SignCheck = "invalid"
	err = cli.DepositCallback(req, func(r XJPayCallbackReq) error {
		return nil
	})
	if err == nil {
		t.Fatalf("expected error for invalid sign")
	}
}

