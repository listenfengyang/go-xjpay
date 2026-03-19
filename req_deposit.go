package go_xjpay

import (
	"fmt"
	"net/url"

	"github.com/listenfengyang/go-xjpay/utils"
)

func (cli *Client) Deposit(req XJPayDepositReq) (string, error) {
	if req.ReceiveUrl == "" {
		req.ReceiveUrl = cli.Params.ReceiveUrl
	}
	if req.PickupUrl == "" {
		req.PickupUrl = cli.Params.PickupUrl
	}
	req.SignType = "md5"

	sign := generateDepositSign(req.PickupUrl, req.ReceiveUrl, req.SignType, req.OrderNo, req.OrderAmount, req.ExchangeRate, req.OrderCurrency, cli.Params.CustomerId, cli.Params.Md5Key)

	values := url.Values{}
	values.Set("receiveUrl", req.ReceiveUrl)
	values.Set("orderNo", req.OrderNo)
	values.Set("customerId", cli.Params.CustomerId)
	values.Set("orderCurrency", req.OrderCurrency)
	values.Set("orderAmount", req.OrderAmount)
	values.Set("sign", sign)
	values.Set("signType", req.SignType)
	values.Set("pickupUrl", req.PickupUrl)
	values.Set("exchangeRate", req.ExchangeRate)
	if req.PayName != "" {
		values.Set("payName", req.PayName)
	}

	base, err := url.Parse(cli.Params.DepositUrl)
	if err != nil {
		return "", err
	}
	base.RawQuery = values.Encode()
	rawURL := base.String()

	fmt.Printf("rawURL: %s\n\n", rawURL)
	return rawURL, nil
}

func generateDepositSign(pickupURL, receiveURL, signType, orderNo, orderAmount, exchangeRate, orderCurrency, customerID, md5Key string) string {
	signSource := pickupURL + receiveURL + signType + orderNo + orderAmount + exchangeRate + orderCurrency + customerID + md5Key
	fmt.Printf("signSource: %s\n", signSource)
	return utils.Md5Hex(signSource)
}
