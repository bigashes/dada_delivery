package dada_delivery

import (
	"encoding/json"
	"github.com/bigashes/dada_delivery/core"
)

//获取城市信息 https://newopen.imdada.cn/#/development/file/cityList
func (c *Client) CityList() (result []*CityListRsp, err error) {

	body := ""
	err = core.PostRequest("/api/cityCode/list", core.NewReqBody(c.config.AppKey, c.config.SourceId, body), c.config.AppSecret, &result, c.Debug)
	return result, err
}

//注册商户 https://newopen.imdada.cn/#/development/file/merchantAdd
func (c *Client) AddMerchant(param AddMerchant) (result *AddMerchantRsp, err error) {

	bodyJson, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}

	err = core.PostRequest("/merchantApi/merchant/add", core.NewReqBody(c.config.AppKey, c.config.SourceId, string(bodyJson)), c.config.AppSecret, &result, c.Debug)
	return result, err
}

//新增门店 https://newopen.imdada.cn/#/development/file/shopAdd
func (c *Client) AddShop(param []*AddShop) (result *AddShopResp, err error) {

	bodyJson, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}

	err = core.PostRequest("/api/shop/add", core.NewReqBody(c.config.AppKey, c.config.SourceId, string(bodyJson)), c.config.AppSecret, &result, c.Debug)
	return result, err
}

//编辑门店 https://newopen.imdada.cn/#/development/file/shopUpdate
func (c *Client) UpdateShop(param UpdateShop) (result *UpdateShopResp, err error) {

	bodyJson, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}

	err = core.PostRequest("/api/shop/update", core.NewReqBody(c.config.AppKey, c.config.SourceId, string(bodyJson)), c.config.AppSecret, &result, c.Debug)
	return result, err
}

//门店详情 https://newopen.imdada.cn/#/development/file/shopDetail
func (c *Client) QueryShop(param QueryShop) (result *QueryShopRsp, err error) {

	bodyJson, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}

	err = core.PostRequest("/api/shop/detail", core.NewReqBody(c.config.AppKey, c.config.SourceId, string(bodyJson)), c.config.AppSecret, &result, c.Debug)
	return result, err
}
