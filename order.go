package dada_delivery

import (
	"encoding/json"
	"github.com/bigashes/dada_delivery/core"
	"github.com/bigashes/dada_delivery/util"
	"io"
	"net/http"
	"strconv"
)

//新增订单 https://newopen.imdada.cn/#/development/file/add
func (c *Client) AddOrder(param AddOrder) (result *AddOrderRsp, err error) {

	bodyJson, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}

	err = core.PostRequest("/api/order/addOrder", core.NewReqBody(c.config.AppKey, c.config.SourceId, string(bodyJson)), c.config.AppSecret, &result, c.Debug)
	return result, err
}

//订单重发 https://newopen.imdada.cn/#/development/file/reAdd
func (c *Client) ReAddOrder(param ReAddOrder) (result *ReAddOrderRsp, err error) {

	bodyJson, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}

	err = core.PostRequest("/api/order/reAddOrder", core.NewReqBody(c.config.AppKey, c.config.SourceId, string(bodyJson)), c.config.AppSecret, &result, c.Debug)
	return result, err
}

// 查询订单运费接口 https://newopen.imdada.cn/#/development/file/readyAdd
func (c *Client) QueryDeliverFee(param QueryDeliverFee) (result *QueryDeliverFeeRsp, err error) {

	bodyJson, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}

	err = core.PostRequest("/api/order/queryDeliverFee", core.NewReqBody(c.config.AppKey, c.config.SourceId, string(bodyJson)), c.config.AppSecret, &result, c.Debug)
	return result, err
}

// 查询运费后发单接口 https://newopen.imdada.cn/#/development/file/readyAdd
func (c *Client) AddAfterQuery(param AddAfterQuery) (result *AddAfterQueryRsp, err error) {

	bodyJson, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}

	err = core.PostRequest("/api/order/addAfterQuery", core.NewReqBody(c.config.AppKey, c.config.SourceId, string(bodyJson)), c.config.AppSecret, &result, c.Debug)
	return result, err
}

// 增加小费 https://newopen.imdada.cn/#/development/file/addTip
func (c *Client) AddTip(param AddTip) (result *AddTipRsp, err error) {

	bodyJson, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}

	err = core.PostRequest("/api/order/addTip", core.NewReqBody(c.config.AppKey, c.config.SourceId, string(bodyJson)), c.config.AppSecret, &result, c.Debug)
	return result, err
}

// 订单详情查询 https://newopen.imdada.cn/#/development/file/statusQuery
func (c *Client) QueryOrder(param QueryOrder) (result *QueryOrderRsp, err error) {

	bodyJson, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}

	err = core.PostRequest("/api/order/status/query", core.NewReqBody(c.config.AppKey, c.config.SourceId, string(bodyJson)), c.config.AppSecret, &result, c.Debug)
	return result, err
}

// 取消订单 https://newopen.imdada.cn/#/development/file/formalCancel
func (c *Client) CancelOrder(param CancelOrder) (result *CancelOrderRsp, err error) {

	bodyJson, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}

	err = core.PostRequest("/api/order/formalCancel", core.NewReqBody(c.config.AppKey, c.config.SourceId, string(bodyJson)), c.config.AppSecret, &result, c.Debug)
	return result, err
}

// 追加订单 https://newopen.imdada.cn/#/development/file/appointOrder
func (c *Client) AppointOrder(param AppointOrder) (result *AppointOrderRsp, err error) {

	bodyJson, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}

	err = core.PostRequest("/api/order/appoint/exist", core.NewReqBody(c.config.AppKey, c.config.SourceId, string(bodyJson)), c.config.AppSecret, &result, c.Debug)
	return result, err
}

// 取消追加订单 https://newopen.imdada.cn/#/development/file/appointOrderCancel
func (c *Client) AppointOrderCancel(param AppointOrderCancel) (result *AppointOrderCancelRsp, err error) {

	bodyJson, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}

	err = core.PostRequest("/api/order/appoint/cancel", core.NewReqBody(c.config.AppKey, c.config.SourceId, string(bodyJson)), c.config.AppSecret, &result, c.Debug)
	return result, err
}

// 查询追加配送员 https://newopen.imdada.cn/#/development/file/listTransportersToAppoint
func (c *Client) AppointQueryTransporter(param AppointQueryTransporter) (result []*AppointQueryTransporterRsp, err error) {

	bodyJson, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}

	err = core.PostRequest("/api/order/appoint/list/transporter", core.NewReqBody(c.config.AppKey, c.config.SourceId, string(bodyJson)), c.config.AppSecret, &result, c.Debug)
	return result, err
}

// 商家投诉达达 https://newopen.imdada.cn/#/development/file/complaintDada
func (c *Client) Complaint(param Complaint) (result *ComplaintRsp, err error) {

	bodyJson, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}

	err = core.PostRequest("/api/complaint/dada", core.NewReqBody(c.config.AppKey, c.config.SourceId, string(bodyJson)), c.config.AppSecret, &result, c.Debug)
	return result, err
}

// 妥投异常之物品返回完成 https://newopen.imdada.cn/#/development/file/abnormalConfirm
func (c *Client) ConfirmGoods(param ConfirmGoods) (result *ConfirmGoodsRsp, err error) {

	bodyJson, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}

	err = core.PostRequest("/api/order/confirm/goods", core.NewReqBody(c.config.AppKey, c.config.SourceId, string(bodyJson)), c.config.AppSecret, &result, c.Debug)
	return result, err
}

// 订单回调 https://newopen.imdada.cn/#/development/file/order
func (c *Client) OrderNotify(req *http.Request) (result *OrderNotify, err error) {
	result = &OrderNotify{}

	body, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, result)
	if err != nil {
		return nil, err
	}

	var params []string
	params = append(params, result.ClientId, result.OrderId, strconv.FormatInt(result.UpdateTime, 10))

	err = util.CallbackSign(params, result.Signature)
	if err != nil {
		return nil, err
	}

	return result, err
}

// 订单回调 https://newopen.imdada.cn/#/development/file/order
func (c *Client) OrderNotifyData(req OrderNotify) (result *OrderNotify, err error) {

	var params []string
	params = append(params, req.ClientId, req.OrderId, strconv.FormatInt(req.UpdateTime, 10))

	err = util.CallbackSign(params, result.Signature)
	if err != nil {
		return nil, err
	}

	return &req, err
}
