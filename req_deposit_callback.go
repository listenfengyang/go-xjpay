package go_xjpay

import (
	"encoding/json"
	"errors"

	"github.com/listenfengyang/go-xjpay/utils"
)

func (cli *Client) DepositCallback(req XJPayCallbackReq, processor func(XJPayCallbackReq) error) error {
	if cli.Params.Md5Key == "" {
		return errors.New("md5Key is empty")
	}

	signSource := req.OrderNo + cli.Params.PickupUrl + cli.Params.ReceiveUrl + cli.Params.Md5Key
	expectedSign := utils.Md5Hex(signSource)
	if req.Sign != expectedSign {
		data, _ := json.Marshal(req)
		cli.logger.Errorf("xjpay deposit callback sign verify fail, req: %s", string(data))
		return errors.New("sign verify error")
	}

	return processor(req)
}
