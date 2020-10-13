package models

import (
	cgorm "github.com/rasyidmm/RumahMakan/db"
	"github.com/rasyidmm/RumahMakan/models/orm"
	uuid "github.com/satori/go.uuid"
	"time"
)

type MenuMakanan struct {
	BaseModel
	NamaMakanan string `json:"nama_makanan" form:"nama_makanan" query:"nama_makanan"`
	DeskripsiMakanan string `json:"deskripsi_makanan" form:"deskripsi_makanan" query:"deskripsi_makanan"`
	StatusMakanan bool `json:"status_makanan" form:"status_makanan" query:"status_makanan"`
	HargaMakanan float64 `json:"harga_makanan" form:"harga_makanan" query:"harga_makanan"`
	JenisMakananId string `gorm:"references:JenisMakanan;foreignKey:id" json:"jenis_makanan_id" `
}
var tx = cgorm.Init().Begin()
var table = "menu_makanan"

func (a *MenuMakanan)BeforeUpdate()(err error)  {
	a.UpdatedAt = time.Now()
	return
}
func (a *MenuMakanan)BeforeCreate()(err error)  {
	a.CreatedAt = time.Now()
	return
}
func Create(a *MenuMakanan)(*MenuMakanan,error)  {
	var err error
	a.Id = uuid.Must(uuid.NewV1())
	err = orm.Create(&a)
	return a,err
}
func (a MenuMakanan)UpdateMenuMakanan() error {
	var err error
	err = orm.Save(&a)
	return  err
}
func GetMenuMakananAll() ([]MenuMakanan, error) {
	var(
		MenuMakanan  []MenuMakanan
		err error
	)
	err = orm.FindAll(&MenuMakanan)
	if err!=nil{
		return MenuMakanan,err
	}

	return MenuMakanan, nil
}
func GetByIdMenuMakanan(id string)(MenuMakanan,error)  {
	var(
		tx = cgorm.Init().Begin()
		table = "menu_makanans"
		MenuMakanan MenuMakanan
		err error
	)
	err = tx.Table(table).Where("id = ?",id).Scan(&MenuMakanan).Error
	if err!=nil{
		return MenuMakanan,err
	}

	return MenuMakanan,err
}

func (a MenuMakanan)DeleteMenuMakanan()error  {
	var err error
	err = orm.Delete(&a)
	return  err
}