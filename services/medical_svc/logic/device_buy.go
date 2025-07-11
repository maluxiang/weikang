package logic

import (
	"context"
	"fmt"
	"time"
	"weikang/global"
	"weikang/models"
	"weikang/services/medical_svc/proto/medical"
)

func generateOrderNumber() string {
	return fmt.Sprintf("MD%v", time.Now().UnixNano())
}

func (s Server) CreateOrder(ctx context.Context, in *medical.CreateOrderRequest) (*medical.CreateOrderResponse, error) {
	//查询用户
	var user models.Users
	err := global.DB.First(&user, in.UserId).Error
	if err != nil {
		return &medical.CreateOrderResponse{Url: ""}, fmt.Errorf("用户不存在")
	}

	//查询商品
	var device models.MedicalDevice
	err = global.DB.First(&device, in.ProductId).Error
	if err != nil {
		return &medical.CreateOrderResponse{Url: ""}, fmt.Errorf("商品不存在")
	}
	if device.Stock < int(in.Quantity) {
		return &medical.CreateOrderResponse{Url: ""}, fmt.Errorf("库存不足")
	}

	//创建订单
	orderNumber := generateOrderNumber()
	totalAmount := device.Price * float64(in.Quantity)
	order := models.Order{
		OrderNumber:     orderNumber,
		UserID:          in.UserId,
		TotalAmount:     totalAmount,
		PaymentMethod:   "alipay",
		PaymentStatus:   0,
		ShippingAddress: "",
		ContactPhone:    user.Phone,
		Remark:          "",
	}
	err = global.DB.Create(&order).Error
	if err != nil {
		return &medical.CreateOrderResponse{Url: ""}, fmt.Errorf("订单创建失败")
	}
	orderitem := models.OrderItem{
		OrderID:         order.ID,
		MedicalDeviceID: in.ProductId,
		DeviceName:      device.Name,
		DeviceModel:     device.Model,
		Quantity:        int(in.Quantity),
		UnitPrice:       device.Price,
		SubtotalAmount:  totalAmount,
	}
	global.DB.Create(&orderitem)

	// 3. 扣减库存
	global.DB.Model(&device).Update("stock", device.Stock-int(in.Quantity))
	//乐观锁防止超卖，扣减库存
	//res := global.DB.Model(&device).Where("stock > 0").Update("stock", gorm.Expr("stock - ?", in.Quantity))
	//if res.RowsAffected == 0 {
	//	return &medical.CreateOrderResponse{
	//		Url: "",
	//	}, nil
	//}
	// 4. 生成支付宝支付链接
	payment := GetPayment(in.PaymentMethod)
	if payment == nil {
		return &medical.CreateOrderResponse{Message: "不支持的支付方式"}, nil
	}
	payUrl, err := payment.Pay(device.Name, orderNumber, totalAmount)
	if err != nil {
		return &medical.CreateOrderResponse{Message: "支付链接生成失败"}, nil
	}
	return &medical.CreateOrderResponse{
		Url:     payUrl,
		Message: "下单成功，请支付",
	}, nil
}
