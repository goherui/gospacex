package model

import (
	__ "gospaacex/proto"

	"gorm.io/gorm"
)

type Position struct {
	gorm.Model
	Title        string `gorm:"type:varchar(30);comment:职位标题"`
	Company      string `gorm:"type:varchar(40);comment:公司名称"`
	Salary       string `gorm:"type:varchar(50);comment:薪资范围"`
	Location     string `gorm:"type:varchar(40);comment:工作地点"`
	Description  string `gorm:"type:varchar(200);comment:职位描述"`
	Requirements string `gorm:"type:varchar(200);comment:职位要求"`
}

func (p *Position) FindTitle(db *gorm.DB, title string) error {
	return db.Where("title=?", title).First(&p).Error
}

func (p *Position) PosCreate(db *gorm.DB) error {
	return db.Create(&p).Error
}

func (p *Position) PosDel(db *gorm.DB, id int64) interface{} {
	return db.Delete(&p, id).Error
}

func (p *Position) FindId(db *gorm.DB, id int64) interface{} {
	return db.Where("id=?", id).First(&p).Error
}

func (p *Position) PosUpdate(db *gorm.DB, id int64) interface{} {
	return db.Where("id=?", id).Updates(&p).Error
}

func (p *Position) PosList(db *gorm.DB, in *__.PosListReq) ([]Position, error) {
	var list []Position
	offset := (in.Page - 1) * in.Size
	if in.Page != 0 || in.Size != 0 {
		err := db.Offset(int(offset)).Limit(int(in.Size)).Find(&list).Error
		return list, err
	}
	err := db.Find(&list).Error
	return list, err
}
