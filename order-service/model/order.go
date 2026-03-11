package model

import (
	__ "gospacex/proto"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	OrderNo        string `gorm:"column:order_no;type:varchar(32);not null;uniqueIndex:uk_order_no" json:"orderNo"`
	UserID         uint64 `gorm:"column:user_id;not null" json:"userId"`
	OrderStatus    int8   `gorm:"column:order_status;type:tinyint;not null;default:1" json:"orderStatus"`
	PaymentStatus  int8   `gorm:"column:payment_status;type:tinyint;not null;default:0" json:"paymentStatus"`
	TotalAmount    int64  `gorm:"column:total_amount;not null;default:0" json:"totalAmount"`
	PaidAmount     int64  `gorm:"column:paid_amount;not null;default:0" json:"paidAmount"`
	ConsigneeName  string `gorm:"column:consignee_name;type:varchar(50);not null" json:"consigneeName"`
	ConsigneePhone string `gorm:"column:consignee_phone;type:varchar(20);not null" json:"consigneePhone"`
	Address        string `gorm:"column:address;type:varchar(255);not null" json:"address"`
	Remark         string `gorm:"column:remark;type:varchar(200)" json:"remark"`
	ShipmentNo     string `gorm:"column:shipment_no;type:varchar(64)" json:"shipmentNo"`
	ExpressName    string `gorm:"column:express_name;type:varchar(50)" json:"expressName"`
}

func (o *Order) FindOrder(db *gorm.DB, no string) error {
	return db.Where("order_no=?", no).First(&o).Error
}

func (o *Order) OrderCreate(db *gorm.DB) error {
	return db.Create(&o).Error
}

func (o *Order) FindOrderId(db *gorm.DB, id int64) interface{} {
	return db.Where("id=?", id).First(&o).Error
}

func (o *Order) OrderDel(db *gorm.DB, id int64) interface{} {
	return db.Delete(&o, id).Error
}

func (o *Order) OrderUpdate(db *gorm.DB, id int64) interface{} {
	return db.Where("id=?", id).Updates(&o).Error
}

func (o *Order) FindOrderList(db *gorm.DB, in *__.OrderListReq) ([]Order, error) {
	var list []Order
	if in.Page != 0 && in.Size != 0 {
		offset := (in.Page - 1) * in.Size
		err := db.Offset(int(offset)).Limit(int(in.Size)).Find(&list).Error
		return list, err
	}
	err := db.Find(&list).Error
	return list, err
}
