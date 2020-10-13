package models

import (
	cgorm "github.com/rasyidmm/RumahMakan/db"
	"github.com/rasyidmm/RumahMakan/models/orm"
	uuid "github.com/satori/go.uuid"
	"time"
)

type Pengunjung struct {
	BaseModel
	Alamat string `json:"alamat" form:"alamat" query:"alamat"`
	Email string `json:"email" form:"email" query:"email"`
	NamaPengunjung string `json:"nama_pengunjung" form:"nama_pengunjung" query:"nama_pengunjung"`
	NomerHp string `json:"nomer_hp" form:"nomer_hp" query:"nomer_hp"`
}

func (a *Pengunjung)BeforeUpdate()(err error)  {
	a.UpdatedAt = time.Now()
	return
}
func (a *Pengunjung)BeforeCreate()(err error)  {
	a.CreatedAt = time.Now()
	return
}
func CreatePengunjung(a *Pengunjung)(*Pengunjung,error)  {
	var err error
	a.Id = uuid.Must(uuid.NewV1())
	err = orm.Create(&a)
	return a,err
}
func (a Pengunjung)UpdatePengunjung() error {
	var err error
	err = orm.Save(&a)
	return  err
}
func GetPengunjungAll() ([]Pengunjung, error) {
	var(
		Pengunjung  []Pengunjung
		err error
	)
	err = orm.FindAll(&Pengunjung)
	if err!=nil{
		return Pengunjung,err
	}

	return Pengunjung, nil
}
func GetByIdPengunjung(id string)(Pengunjung,error)  {
	var(
		tx = cgorm.Init().Begin()
		table = "pengunjungs"
		Pengunjung Pengunjung
		err error
	)
	err = tx.Table(table).Where("id = ?",id).Scan(&Pengunjung).Error
	if err!=nil{
		return Pengunjung,err
	}

	return Pengunjung,err
}
func (a Pengunjung)DeletePengunjung()error  {
	var err error
	err = orm.Delete(&a)
	return  err
}