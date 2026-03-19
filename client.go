package go_xjpay

import (
	"github.com/go-resty/resty/v2"
	"github.com/listenfengyang/go-xjpay/utils"
)

type Client struct {
	Params *XJPayInitParams

	ryClient  *resty.Client
	debugMode bool
	logger    utils.Logger
}

func NewClient(logger utils.Logger, params *XJPayInitParams) *Client {
	return &Client{
		Params:   params,
		ryClient: resty.New(),
		logger:   logger,
	}
}

func (cli *Client) SetDebugModel(debugModel bool) {
	cli.debugMode = debugModel
}
