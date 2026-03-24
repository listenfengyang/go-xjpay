package go_xjpay

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/listenfengyang/go-xjpay/utils"
)

// String signCheck = ClientUtils.MD5Encode( orderNo+pickupUrl + receiveUrl+md5_key);
func (cli *Client) DepositCallback(req XJPayCallbackReq, processor func(XJPayCallbackReq) error) error {
	if cli.Params.Md5Key == "" {
		return errors.New("md5Key is empty")
	}

	signSource := req.OrderNo + cli.Params.PickupUrl + cli.Params.ReceiveUrl + cli.Params.Md5Key
	fmt.Printf("signSource: %s\n\n", signSource)
	cli.logger.Infof("signSource: %s\n\n", signSource)

	expectedSign := utils.Md5Hex(signSource)
	if !strings.EqualFold(req.SignCheck, expectedSign) {
		data, _ := json.Marshal(req)
		cli.logger.Errorf("xjpay deposit callback sign verify fail, req: %s", string(data))
		return errors.New("sign verify error")
	}

	return processor(req)
}

// func (cli *Client) DepositCallback(req XJPayCallbackReq, processor func(XJPayCallbackReq) error) error {
// 	if cli.Params.Md5Key == "" {
// 		return errors.New("md5Key is empty")
// 	}

// 	signSource := req.SignType + req.OrderNo + req.OrderPayment + req.OrderCurrency + req.TransactionId + req.Status + cli.Params.Md5Key
// 	cli.logger.Infof("signSource: %s\n\n", signSource)

// 	expectedSign := utils.Md5Hex(signSource)
// 	if !strings.EqualFold(req.Sign, expectedSign) {
// 		data, _ := json.Marshal(req)
// 		cli.logger.Errorf("xjpay deposit callback sign verify fail, req: %s", string(data))
// 		return errors.New("sign verify error")
// 	}

// 	return processor(req)
// }
