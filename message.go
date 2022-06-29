package dada_delivery

import (
	"encoding/json"
	"github.com/bigashes/dada_delivery/core"
	"io"
	"net/http"
)

// 订单回调 https://newopen.imdada.cn/#/development/file/transporterCancelOrder
func (c Client) MessageNotify(req *http.Request) (result interface{}, err error) {
	result = &OrderNotify{}

	message := &ConfirmMessage{}
	body, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, message)
	if err != nil {
		return nil, err
	}

	switch message.MessageType {
	case 1:
		messageBody := &TransporterCancelOrderMessage{}
		err = json.Unmarshal([]byte(message.MessageBody), messageBody)
		if err != nil {
			return nil, err
		}

		return messageBody, nil
	default:
		return result, err
	}
}

// 确认骑士取消订单 https://newopen.imdada.cn/#/development/file/transporterCancelOrder
func (c *Client) ConfirmTransporterCancelOrder(param ConfirmTransporterCancelOrder) (result *ConfirmTransporterCancelOrderRsp, err error) {

	bodyJson, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}

	confirmMessage := ConfirmMessage{
		MessageType: 1,
		MessageBody: string(bodyJson),
	}

	confirmMessageJson, err := json.Marshal(confirmMessage)
	if err != nil {
		return nil, err
	}

	err = core.PostRequest("/api/message/confirm", core.NewReqBody(c.config.AppKey, c.config.SourceId, string(confirmMessageJson)), c.config.AppSecret, &result, c.Debug)
	return result, err
}
