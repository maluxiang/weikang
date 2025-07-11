package models

import (
	"gorm.io/gorm"
	"time"
)

type Order struct {
	ID              uint64         `gorm:"column:id;type:bigint UNSIGNED;primaryKey;autoIncrement;not null"`
	OrderNumber     string         `gorm:"column:order_number;type:varchar(50);uniqueIndex;not null;comment:'订单号'"`
	UserID          uint64         `gorm:"column:user_id;type:bigint UNSIGNED;not null;comment:'用户ID'"`
	TotalAmount     float64        `gorm:"column:total_amount;type:decimal(10,2);not null;comment:'订单总金额'"`
	PaymentMethod   string         `gorm:"column:payment_method;type:varchar(20);comment:'支付方式(支付宝/微信/银行卡)'"`
	PaymentStatus   int            `gorm:"column:payment_status;type:int(11);not null;default:0;comment:'支付状态(0-未支付,1-已支付,2-支付失败,3-已退款)'"`
	ShippingAddress string         `gorm:"column:shipping_address;type:text;comment:'收货地址'"`
	ContactPhone    string         `gorm:"column:contact_phone;type:varchar(20);comment:'联系电话'"`
	Remark          string         `gorm:"column:remark;type:text;comment:'订单备注'"`
	CreatedAt       time.Time      `gorm:"column:created_at;type:datetime;not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt       time.Time      `gorm:"column:updated_at;type:datetime;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
	DeletedAt       gorm.DeletedAt `gorm:"column:deleted_at;type:datetime;index"`
}

type OrderItem struct {
	ID              uint64         `gorm:"column:id;type:bigint UNSIGNED;primaryKey;autoIncrement;not null"`
	OrderID         uint64         `gorm:"column:order_id;type:bigint UNSIGNED;not null;comment:'订单ID'"`
	MedicalDeviceID uint64         `gorm:"column:medical_device_id;type:bigint UNSIGNED;not null;comment:'医疗器械ID'"`
	DeviceName      string         `gorm:"column:device_name;type:varchar(255);not null;comment:'商品名称'"`
	DeviceModel     string         `gorm:"column:device_model;type:varchar(100);comment:'商品型号'"`
	Quantity        int            `gorm:"column:quantity;type:int(11);not null;comment:'购买数量'"`
	UnitPrice       float64        `gorm:"column:unit_price;type:decimal(10,2);not null;comment:'商品单价'"`
	SubtotalAmount  float64        `gorm:"column:subtotal_amount;type:decimal(10,2);not null;comment:'小计金额'"`
	CreatedAt       time.Time      `gorm:"column:created_at;type:datetime;not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt       time.Time      `gorm:"column:updated_at;type:datetime;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
	DeletedAt       gorm.DeletedAt `gorm:"column:deleted_at;type:datetime;index"`
}

func (o Order) TableName() string {
	return "order"
}

func (o OrderItem) TableName() string {
	return "order_item"
}
