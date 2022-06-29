package dada_delivery

type OrderProduct struct {
	SkuName      string  `json:"sku_name"`       //商品名称，限制长度128
	SrcProductNo string  `json:"src_product_no"` //商品编码，限制长度64
	Count        float64 `json:"count"`          //商品数量，精确到小数点后两位
	Unit         string  `json:"unit"`           //商品单位，默认：件
}

type AddOrder struct {
	ShopNo             string         `json:"shop_no"`               //门店编号
	OriginId           string         `json:"origin_id"`             //第三方订单ID
	CityCode           string         `json:"city_code"`             //订单所在城市的code
	CargoPrice         float64        `json:"cargo_price"`           //订单金额（单位：元）
	IsPrepay           int            `json:"is_prepay"`             //是否需要垫付 1:是 0:否 (垫付订单金额，非运费)
	ReceiverName       string         `json:"receiver_name"`         //收货人姓名
	ReceiverAddress    string         `json:"receiver_address"`      //收货人地址
	ReceiverLat        float64        `json:"receiver_lat"`          //收货人地址纬度
	ReceiverLng        float64        `json:"receiver_lng"`          //收货人地址经度
	Callback           string         `json:"callback"`              //回调URL
	CargoWeight        float64        `json:"cargo_weight"`          //订单重量
	ReceiverPhone      string         `json:"receiver_phone"`        //	收货人手机号
	ReceiverTel        string         `json:"receiver_tel"`          //收货人座机号
	Tips               float64        `json:"tips"`                  //小费（单位：元，精确小数点后一位）
	Info               string         `json:"info"`                  //	订单备注
	CargoType          int64          `json:"cargo_type"`            //请选择门店真实配送品类，如无法判断可咨询您的销售经理。
	CargoNum           int64          `json:"cargo_num"`             //订单商品数量
	InvoiceTitle       string         `json:"invoice_title"`         //发票抬头
	OriginMark         string         `json:"origin_mark"`           //	订单来源标示（只支持字母，最大长度为10）
	OriginMarkNo       string         `json:"origin_mark_no"`        //	订单来源编号
	IsUseInsurance     int            `json:"is_use_insurance"`      //是否使用保价费
	IsFinishCodeNeeded int            `json:"is_finish_code_needed"` //收货码 0：不需要；1：需要。
	DelayPublishTime   int64          `json:"delay_publish_time"`    //预约发单时间
	IsDirectDelivery   int            `json:"is_direct_delivery"`    //是否选择直拿直送 0：不需要；1：需要。
	ProductList        []OrderProduct `json:"product_list"`          //订单商品明细
	PickUpPos          string         `json:"pick_up_pos"`           //货架信息,该字段可在骑士APP订单备注中展示
}

type AddOrderRsp struct {
	Distance     float64 `json:"distance"`     //距离（单位：米）
	Fee          float64 `json:"fee"`          //实际运费（单位：元），运费减去优惠券费用
	DeliverFee   float64 `json:"deliverFee"`   //运费(单位：元)
	CouponFee    float64 `json:"couponFee"`    //优惠券费用(单位：元)
	Tips         float64 `json:"tips"`         //小费(单位：元)
	InsuranceFee float64 `json:"insuranceFee"` //保价费(单位：元)
}

type ReAddOrder struct {
	ShopNo             string         `json:"shop_no"`               //门店编号
	OriginId           string         `json:"origin_id"`             //第三方订单ID
	CityCode           string         `json:"city_code"`             //订单所在城市的code
	CargoPrice         float64        `json:"cargo_price"`           //订单金额（单位：元）
	IsPrepay           int            `json:"is_prepay"`             //是否需要垫付 1:是 0:否 (垫付订单金额，非运费)
	ReceiverName       string         `json:"receiver_name"`         //收货人姓名
	ReceiverAddress    string         `json:"receiver_address"`      //收货人地址
	ReceiverLat        float64        `json:"receiver_lat"`          //收货人地址纬度
	ReceiverLng        float64        `json:"receiver_lng"`          //收货人地址经度
	Callback           string         `json:"callback"`              //回调URL
	CargoWeight        float64        `json:"cargo_weight"`          //订单重量
	ReceiverPhone      string         `json:"receiver_phone"`        //	收货人手机号
	ReceiverTel        string         `json:"receiver_tel"`          //收货人座机号
	Tips               float64        `json:"tips"`                  //小费（单位：元，精确小数点后一位）
	Info               string         `json:"info"`                  //	订单备注
	CargoType          int64          `json:"cargo_type"`            //请选择门店真实配送品类，如无法判断可咨询您的销售经理。
	CargoNum           int64          `json:"cargo_num"`             //订单商品数量
	InvoiceTitle       string         `json:"invoice_title"`         //发票抬头
	OriginMark         string         `json:"origin_mark"`           //	订单来源标示（只支持字母，最大长度为10）
	OriginMarkNo       string         `json:"origin_mark_no"`        //	订单来源编号
	IsUseInsurance     int            `json:"is_use_insurance"`      //是否使用保价费
	IsFinishCodeNeeded int            `json:"is_finish_code_needed"` //收货码 0：不需要；1：需要。
	DelayPublishTime   int64          `json:"delay_publish_time"`    //预约发单时间
	IsDirectDelivery   int            `json:"is_direct_delivery"`    //是否选择直拿直送 0：不需要；1：需要。
	ProductList        []OrderProduct `json:"product_list"`          //订单商品明细
	PickUpPos          string         `json:"pick_up_pos"`           //货架信息,该字段可在骑士APP订单备注中展示
}

type ReAddOrderRsp struct {
	Distance     float64 `json:"distance"`     //距离（单位：米）
	Fee          float64 `json:"fee"`          //实际运费（单位：元），运费减去优惠券费用
	DeliverFee   float64 `json:"deliverFee"`   //运费(单位：元)
	CouponFee    float64 `json:"couponFee"`    //优惠券费用(单位：元)
	Tips         float64 `json:"tips"`         //小费(单位：元)
	InsuranceFee float64 `json:"insuranceFee"` //保价费(单位：元)
}

type QueryDeliverFee struct {
	ShopNo             string         `json:"shop_no"`               //门店编号
	OriginId           string         `json:"origin_id"`             //第三方订单ID
	CityCode           string         `json:"city_code"`             //订单所在城市的code
	CargoPrice         float64        `json:"cargo_price"`           //订单金额（单位：元）
	IsPrepay           int            `json:"is_prepay"`             //是否需要垫付 1:是 0:否 (垫付订单金额，非运费)
	ReceiverName       string         `json:"receiver_name"`         //收货人姓名
	ReceiverAddress    string         `json:"receiver_address"`      //收货人地址
	ReceiverLat        float64        `json:"receiver_lat"`          //收货人地址纬度
	ReceiverLng        float64        `json:"receiver_lng"`          //收货人地址经度
	Callback           string         `json:"callback"`              //回调URL
	CargoWeight        float64        `json:"cargo_weight"`          //订单重量
	ReceiverPhone      string         `json:"receiver_phone"`        //	收货人手机号
	ReceiverTel        string         `json:"receiver_tel"`          //收货人座机号
	Tips               float64        `json:"tips"`                  //小费（单位：元，精确小数点后一位）
	Info               string         `json:"info"`                  //	订单备注
	CargoType          int64          `json:"cargo_type"`            //请选择门店真实配送品类，如无法判断可咨询您的销售经理。
	CargoNum           int64          `json:"cargo_num"`             //订单商品数量
	InvoiceTitle       string         `json:"invoice_title"`         //发票抬头
	OriginMark         string         `json:"origin_mark"`           //	订单来源标示（只支持字母，最大长度为10）
	OriginMarkNo       string         `json:"origin_mark_no"`        //	订单来源编号
	IsUseInsurance     int            `json:"is_use_insurance"`      //是否使用保价费
	IsFinishCodeNeeded int            `json:"is_finish_code_needed"` //收货码 0：不需要；1：需要。
	DelayPublishTime   int64          `json:"delay_publish_time"`    //预约发单时间
	IsDirectDelivery   int            `json:"is_direct_delivery"`    //是否选择直拿直送 0：不需要；1：需要。
	ProductList        []OrderProduct `json:"product_list"`          //订单商品明细
	PickUpPos          string         `json:"pick_up_pos"`           //货架信息,该字段可在骑士APP订单备注中展示
}

type QueryDeliverFeeRsp struct {
	Distance     float64 `json:"distance"`     //距离（单位：米）
	DeliveryNo   string  `json:"deliveryNo"`   //配送单号
	Fee          float64 `json:"fee"`          //实际运费（单位：元），运费减去优惠券费用
	DeliverFee   float64 `json:"deliverFee"`   //运费(单位：元)
	CouponFee    float64 `json:"couponFee"`    //优惠券费用(单位：元)
	Tips         float64 `json:"tips"`         //小费(单位：元)
	InsuranceFee float64 `json:"insuranceFee"` //保价费(单位：元)
}

type AddAfterQuery struct {
	DeliveryNo string `json:"deliveryNo"` //查询订单运费接口的平台订单编号
}

type AddAfterQueryRsp struct {
}

type AddTip struct {
	OrderId  string  `json:"order_id"`  //第三方订单编号
	Tips     float64 `json:"tips"`      //小费（单位：元）
	CityCode string  `json:"city_code"` //订单城市区号
	Info     string  `json:"info"`      //备注(字段最大长度：512)
}

type AddTipRsp struct {
}

type QueryOrder struct {
	OrderId string `json:"order_id"` //第三方订单编号
}

type QueryOrderRsp struct {
	OrderId          string  `json:"orderId"`          //第三方订单编号
	StatusCode       int     `json:"statusCode"`       //订单状态(待接单＝1,待取货＝2,配送中＝3,已完成＝4,已取消＝5, 指派单=8,妥投异常之物品返回中=9, 妥投异常之物品返回完成=10, 骑士到店=100,创建达达运单失败=1000 可参考文末的状态说明）
	StatusMsg        string  `json:"statusMsg"`        //订单状态
	TransporterName  string  `json:"transporterName"`  //骑手姓名
	TransporterPhone string  `json:"transporterPhone"` //骑手电话
	TransporterLng   string  `json:"transporterLng"`   //骑手经度
	TransporterLat   string  `json:"transporterLat"`   //骑手纬度
	DeliveryFee      float64 `json:"deliveryFee"`      //配送费,单位为元
	Tips             float64 `json:"tips"`             //小费,单位为元
	CouponFee        float64 `json:"couponFee"`        //优惠券费用,单位为元
	InsuranceFee     float64 `json:"insuranceFee"`     //保价费,单位为元
	ActualFee        float64 `json:"actualFee"`        //实际支付费用,单位为元
	Distance         float64 `json:"distance"`         //配送距离,单位为米
	CreateTime       string  `json:"createTime"`       //发单时间
	AcceptTime       string  `json:"acceptTime"`       //接单时间,若未接单,则为空
	FetchTime        string  `json:"fetchTime"`        //取货时间,若未取货,则为空
	FinishTime       string  `json:"finishTime"`       //送达时间,若未送达,则为空
	CancelTime       string  `json:"cancelTime"`       //取消时间,若未取消,则为空
	OrderFinishCode  string  `json:"orderFinishCode"`  //收货码
	DeductFee        float64 `json:"deductFee"`        //违约金
}

type CancelOrder struct {
	OrderId        string `json:"order_id"`         //第三方订单编号
	CancelReasonId int    `json:"cancel_reason_id"` //取消原因ID
	CancelReason   string `json:"cancel_reason"`    //取消原因
}

type CancelOrderRsp struct {
	DeductFee float64 `json:"deduct_fee"` //扣除的违约金(单位：元)
}

type AppointOrder struct {
	OrderId       string `json:"order_id"`       //追加的第三方订单ID
	TransporterId int64  `json:"transporter_id"` //追加的配送员ID
	ShopNo        string `json:"shop_no"`        //追加订单的门店编码
}

type AppointOrderRsp struct {
}

type AppointOrderCancel struct {
	OrderId string `json:"order_id"` //第三方订单号
}

type AppointOrderCancelRsp struct {
}

type AppointQueryTransporter struct {
	ShopNo string `json:"shop_no"` //门店编码
}

type AppointQueryTransporterRsp struct {
	Id     int64  `json:"id"`      //配送员id
	Name   string `json:"name"`    //配送员姓名
	Phone  string `json:"phone"`   //配送员电话
	CityId int    `json:"city_id"` //配送员城市
}

type Complaint struct {
	OrderId  string `json:"order_id"`  //第三方订单编号
	ReasonId int    `json:"reason_id"` //投诉原因ID
}

type ComplaintRsp struct {
}

type ConfirmGoods struct {
	OrderId string `json:"order_id"` //第三方订单编号
}

type ConfirmGoodsRsp struct {
}

type CityListRsp struct {
	CityName string `json:"cityName"` //城市名称
	CityCode string `json:"cityCode"` //城市编码
}

type AddMerchant struct {
	Mobile            string `json:"mobile"`             //注册商户手机号,用于登陆商户后台
	CityName          string `json:"city_name"`          //商户城市名称(如,上海)
	EnterpriseName    string `json:"enterprise_name"`    //企业全称
	EnterpriseAddress string `json:"enterprise_address"` //企业地址
	ContactName       string `json:"contact_name"`       //联系人姓名
	ContactPhone      string `json:"contact_phone"`      //联系人电话
	Email             string `json:"email"`              //邮箱地址
}

type AddMerchantRsp struct {
}

type AddShop struct {
	Station_name    string  `json:"station_name"`    //门店名称
	Business        int     `json:"business"`        //支持配送的物品品类见
	Station_address string  `json:"station_address"` //门店地址
	Lng             float64 `json:"lng"`             //门店经度
	Lat             float64 `json:"lat"`             //门店纬度
	Contact_name    string  `json:"contact_name"`    //联系人姓名
	Phone           string  `json:"phone"`           //联系人电话
	Origin_shop_id  string  `json:"origin_shop_id"`  //门店编码,可自定义,但必须唯一;若不填写,则系统自动生成
	Id_card         string  `json:"id_card"`         //联系人身份证
	Username        string  `json:"username"`        //达达商家app账号(若不需要登陆app,则不用设置)
	Password        string  `json:"password"`        //达达商家app密码(若不需要登陆app,则不用设置)
}

type ShopItem struct {
	Phone          string  `json:"phone"`          //联系人电话
	Business       int     `json:"business"`       //支持配送的物品品类见
	Lng            float64 `json:"lng"`            //门店经度
	Lat            float64 `json:"lat"`            //门店纬度
	StationName    string  `json:"stationName"`    //门店名称
	OriginShopId   string  `json:"originShopId"`   //门店编码
	ContactName    string  `json:"contactName"`    //联系人姓名
	StationAddress string  `json:"stationAddress"` //门店地址
	CityName       string  `json:"cityName"`       //城市名称
	AreaName       string  `json:"areaName"`       //区域名称
}

type ShopFailedItem struct {
	ShopNo   string `json:"shopNo"`   //门店编码
	Msg      string `json:"msg"`      //	错误信息
	ShopName string `json:"shopName"` //门店名称
}

type AddShopResp struct {
	Success     int              `json:"success"`     //成功导入门店的数量
	SuccessList []ShopItem       `json:"successList"` //成功导入的门店
	FailedList  []ShopFailedItem `json:"failedList"`  //导入失败门店编号以及原因
}

type UpdateShop struct {
	OriginShopId   string  `json:"origin_shop_id"`  //门店编码
	NewShopId      string  `json:"new_shop_id"`     //新的门店编码
	StationName    string  `json:"station_name"`    //门店名称
	Business       int     `json:"business"`        //支持配送的物品品类见链接。品类需按门店真实配送品类选择，如无法判断可咨询您的销售经理。
	StationAddress string  `json:"station_address"` //门店地址
	Lng            float64 `json:"lng"`             //门店经度
	Lat            float64 `json:"lat"`             //门店纬度
	ContactName    string  `json:"contact_name"`    //联系人姓名
	Phone          string  `json:"phone"`           //联系人电话
	Status         int     `json:"status"`          //门店状态（1-门店激活，0-门店下线）
}

type UpdateShopResp struct {
}

type QueryShop struct {
	OriginShopId string `json:"origin_shop_id"` //门店编码
}

type QueryShopRsp struct {
	OriginShopId   string  `json:"origin_shop_id"`  //门店编码
	StationName    string  `json:"station_name"`    //门店名称
	Business       int     `json:"business"`        //支持配送的物品品类
	CityName       string  `json:"city_name"`       //城市名称
	AreaName       string  `json:"area_name"`       //区域名称
	StationAddress string  `json:"station_address"` //门店地址
	Lng            float64 `json:"lng"`             //门店经度
	Lat            float64 `json:"lat"`             //门店纬度
	ContactName    string  `json:"contact_name"`    //联系人姓名
	Phone          string  `json:"phone"`           //联系人电话
	IdCard         string  `json:"id_card"`         //联系人身份证
	Status         int     `json:"status"`          //门店状态（1-门店激活，0-门店下线）
}

type Recharge struct {
	Amount    float64 `json:"amount"`     //充值金额（单位元，可以精确到分）
	Category  string  `json:"category"`   //生成链接适应场景（category有二种类型值：PC、H5）
	NotifyUrl string  `json:"notify_url"` //支付成功后跳转的页面（支付宝在支付成功后可以跳转到某个指定的页面，微信支付不支持）
}

type QueryBalance struct {
	Category int `json:"category"` //查询运费账户类型（1：运费账户；2：红包账户，3：所有），默认查询运费账户余额
}

type QueryBalanceRsp struct {
	DeliverBalance   float64 `json:"deliverBalance"`   //运费账户余额（单位元，可以精确到分）
	RedPacketBalance float64 `json:"redPacketBalance"` //红包账户余额（单位元，可以精确到分）
}

type OrderNotify struct {
	ClientId     string `json:"client_id"`     //返回达达运单号，默认为空
	OrderId      string `json:"order_id"`      //添加订单接口中的origin_id值
	OrderStatus  int    `json:"order_status"`  //订单状态(待接单＝1,待取货＝2,配送中＝3,已完成＝4,已取消＝5, 指派单=8,妥投异常之物品返回中=9, 妥投异常之物品返回完成=10, 骑士到店=100,创建达达运单失败=1000 可参考文末的状态说明）
	CancelReason string `json:"cancel_reason"` //订单取消原因,其他状态下默认值为空字符串
	CancelFrom   int    `json:"cancel_from"`   //订单取消原因来源(1:达达配送员取消；2:商家主动取消；3:系统或客服取消；0:默认值)
	UpdateTime   int64  `json:"update_time"`   //更新时间，时间戳除了创建达达运单失败=1000的精确毫秒，其他时间戳精确到秒
	Signature    string `json:"signature"`     //对client_id, order_id, update_time的值进行字符串升序排列，再连接字符串，取md5值
	DmId         int    `json:"dm_id"`         //达达配送员id，接单以后会传
	DmName       string `json:"dm_name"`       //配送员姓名，接单以后会传
	DmMobile     string `json:"dm_mobile"`     //配送员手机号，接单以后会传
	FinishCode   string `json:"finish_code"`   //收货码
}

type ConfirmMessage struct {
	MessageType int    `json:"messageType"` //	消息类型（1：骑士取消订单推送消息）
	MessageBody string `json:"messageBody"` //	消息内容（json字符串）
}

type TransporterCancelOrderMessage struct {
	OrderId      string `json:"order_id"`      //商家第三方订单号
	DadaOrderId  string `json:"dada_order_id"` //达达订单号
	CancelReason string `json:"cancel_reason"` //	骑士取消原因
}

type ConfirmTransporterCancelOrder struct {
	OrderId      string   `json:"order_id"`     //商家第三方订单号
	DadaOrderId  int64    `json:"dadaOrderId"`  //达达订单号
	IsConfirm    int      `json:"isConfirm"`    //0:不同意，1:表示同意
	Imgs         []string `json:"imgs"`         //审核不通过的图片列表
	RejectReason string   `json:"rejectReason"` //拒绝原因
}

type ConfirmTransporterCancelOrderRsp struct {
}
