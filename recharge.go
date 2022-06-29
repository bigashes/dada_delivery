package dada_delivery

import (
	"encoding/json"
	"github.com/bigashes/dada_delivery/core"
)

//获取充值链接 https://newopen.imdada.cn/#/development/file/recharge
func (c *Client) Recharge(param Recharge) (result string, err error) {

	bodyJson, err := json.Marshal(param)
	if err != nil {
		return "", err
	}

	err = core.PostRequest("/api/recharge", core.NewReqBody(c.config.AppKey, c.config.SourceId, string(bodyJson)), c.config.AppSecret, &result, c.Debug)
	return result, err
}

//获取充值链接 https://newopen.imdada.cn/#/development/file/recharge
func (c *Client) QueryBalance(param QueryBalance) (result *QueryBalanceRsp, err error) {

	bodyJson, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}

	err = core.PostRequest("/api/balance/query", core.NewReqBody(c.config.AppKey, c.config.SourceId, string(bodyJson)), c.config.AppSecret, &result, c.Debug)
	return result, err
}
