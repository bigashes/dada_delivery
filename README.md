# dada_delivery
dada_delivery SDK for Golang 达达配送sdk

# 使用
```
//初始化
client := NewClient(Config{
    AppKey:    "",
    AppSecret: "",
    SourceId:  "",
})

//开启debug 使用达达测试接口
client.Debug = true

order, err := client.AddOrder(AddOrder{
    ShopNo:          "03cd2f4b0ce442cc",
    OriginId:        "123456",
    CityCode:        "0532",
    CargoPrice:      10.1,
    IsPrepay:        0,
    ReceiverName:    "测试",
    ReceiverAddress: "青岛市",
    ReceiverLat:     36.087430,
    ReceiverLng:     120.374790,
    Callback:        "https://",
    CargoWeight:     1,
    ReceiverPhone:   "13122223333",
})
```
## 接口
- 新增订单 AddOrder
- 订单重发 ReAddOrder
- 查询订单运费接口 QueryDeliverFee
- 查询运费后发单接口 AddAfterQuery
- 增加小费 AddTip
- 订单详情查询 QueryOrder
- 取消订单 CancelOrder
- 追加订单 AppointOrder
- 取消追加订单 AppointOrderCancel
- 查询追加配送员 AppointQueryTransporter
- 商家投诉达达 Complaint
- 妥投异常之物品返回完成 ConfirmGoods
- 获取城市信息 CityList
- 注册商户 AddMerchant
- 新增门店 AddShop
- 编辑门店 UpdateShop
- 门店详情 QueryShop
- 获取充值链接 Recharge
- 查询余额 QueryBalance

## 回调验证签名
```
client := NewClient(Config{})
//请求体
result, err := client.OrderNotify(req)
//结构体
result, err := client.OrderNotifyData(req)
```
# License
This project is licensed under the MIT License.
