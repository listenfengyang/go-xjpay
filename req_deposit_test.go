package go_xjpay

import (
	"fmt"
	"testing"
)

type VLog struct{}

func (l VLog) Debugf(format string, args ...interface{}) {}
func (l VLog) Infof(format string, args ...interface{})  {}
func (l VLog) Warnf(format string, args ...interface{})  {}
func (l VLog) Errorf(format string, args ...interface{}) {}

func TestDeposit(t *testing.T) {

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

	linkUrl, err := cli.Deposit(GenDepositRequestDemo())
	if err != nil {
		cli.logger.Errorf("err:%s\n", err.Error())
		t.Fatalf("deposit error: %v", err)
	}

	fmt.Printf("linkUrl: %s\n", linkUrl)
	// return linkUrl
}

func GenDepositRequestDemo() XJPayDepositReq {
	return XJPayDepositReq{
		OrderNo:       "7413369812",
		OrderAmount:   "500",
		ExchangeRate:  "7.13",
		PayName:       "李四",
		OrderCurrency: "USDT",
		OrderCny:      "1000",
	}
}
