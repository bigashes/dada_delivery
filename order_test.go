package dada_delivery

import (
	"testing"
)

func TestAddOrder(t *testing.T) {
	client := NewClient(Config{
		AppKey:    "",
		AppSecret: "",
		SourceId:  "",
	})

	client.Debug = true

	order, err := client.AddOrder(AddOrder{
		ShopNo:          "03cd2f4b0ce442cc",
		OriginId:        "123456",
		CityCode:        "0532",
		CargoPrice:      10.1,
		IsPrepay:        0,
		ReceiverName:    "测试",
		ReceiverAddress: "青岛市市北区卓越世纪中心1号楼",
		ReceiverLat:     36.087430,
		ReceiverLng:     120.374790,
		Callback:        "https://ccdff.fyyang.com",
		CargoWeight:     1,
		ReceiverPhone:   "15906398085",
	})

	if err != nil {
		t.Error(err)
	}

	t.Logf("order %+v \n", order)
}
