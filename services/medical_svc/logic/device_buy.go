package logic

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"time"
	"weikang/global"
	"weikang/models"
	"weikang/services/medical_svc/proto/medical"
)

func (s Server) CreateOrder(ctx context.Context, in *medical.CreateOrderRequest) (*medical.CreateOrderResponse, error) {
	//参数校验
	if in.ProductId == 0 || in.UserId == 0 || in.Quantity <= 0 {
		return &medical.CreateOrderResponse{
			Message: "参数错误",
		}, nil
	}

	// 查询用户
	var user models.Users
	if err := global.DB.Where("id = ?", in.UserId).First(&user).Error; err != nil {
		return &medical.CreateOrderResponse{
			Message: "用户不存在",
		}, nil
	}

	//查询商品
	var device models.MedicalDevice
	if err := global.DB.Where("id = ?", in.ProductId).First(&device).Error; err != nil {
		return &medical.CreateOrderResponse{
			Message: "商品不存在",
		}, nil
	}
	if device.Stock < int(in.Quantity) {
		return &medical.CreateOrderResponse{
			Message: "库存不足",
		}, nil
	}
	//生成订单号
	orderNumber := fmt.Sprintf("ODR%d%d", in.UserId, time.Now().UnixNano())
	amount := float64(in.Quantity) * device.Price
	//事务处理
	err := global.DB.Transaction(func(tx *gorm.DB) error {
		// 3.1 扣减库存
		if err := tx.Model(&models.MedicalDevice{}).
			Where("id = ? AND stock >= ?", in.ProductId, in.Quantity).
			Update("stock", gorm.Expr("stock - ?", in.Quantity)).Error; err != nil {
			return errors.New("扣减库存失败")
		}

		//写入订单表
		order := models.Order{
			OrderNumber:   orderNumber,
			UserID:        in.UserId,
			TotalAmount:   amount,
			PaymentMethod: in.PaymentMethod,
			PaymentStatus: 0, // 未支付
			ContactPhone:  user.Phone,
		}
		if err := tx.Create(&order).Error; err != nil {
			return err
		}

		// 3.4 写入订单明细表
		orderItem := models.OrderItem{
			OrderID:         order.ID,
			MedicalDeviceID: device.ID,
			DeviceName:      device.Name,
			DeviceModel:     device.Model,
			Quantity:        int(in.Quantity),
			UnitPrice:       device.Price,
			SubtotalAmount:  float64(in.Quantity) * device.Price,
		}
		if err := tx.Create(&orderItem).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return &medical.CreateOrderResponse{
			Message: "下单失败：" + err.Error(),
		}, nil
	}

	// 4. 返回支付URL（如有支付需求，可集成支付宝/微信等）
	payment := GetPayment(in.PaymentMethod)
	pay, err := payment.Pay("支付", orderNumber, amount)
	if err != nil {
		return nil, err
	}

	s.SendSiteMessage(int64(in.UserId), "购买成功", "您的订单已购买成功，感谢您的支持！")

	// 这里只返回订单号
	return &medical.CreateOrderResponse{
		Url:     pay, // 可集成支付后返回支付链接
		Message: "下单成功",
	}, nil
}
