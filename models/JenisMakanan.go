package models

import (
	cgorm "github.com/rasyidmm/RumahMakan/db"
	"github.com/rasyidmm/RumahMakan/models/orm"
	uuid "github.com/satori/go.uuid"
	"time"
)


type JenisMakanan struct {
	BaseModel
	NamaJenis string `json:"nama_jenis" form:"nama_jenis" query:"nama_jenis"`
	DeskripsiJenis string `json:"deskripsi_jenis" form:"deskripsi_jenis" query:"deskripsi_jenis"`
}


func (a *JenisMakanan)BeforeUpdate()(err error)  {
	a.UpdatedAt = time.Now()
	return
}
func (a *JenisMakanan)BeforeCreate()(err error)  {
	a.CreatedAt = time.Now()
	return
}
func CreateJenisMakanan(a *JenisMakanan)(*JenisMakanan,error)  {
	var err error
	a.Id = uuid.Must(uuid.NewV1())
	err = orm.Create(&a)
	return a,err
}

func (a *JenisMakanan)UpdateJenisMakanan()error  {
	var err error
	err = orm.Save(&a)
	return err
}

func (a *JenisMakanan)DeleteJenisMakanan()error  {
	var err error
	err = orm.Delete(a)
	return  err
}

func FindJenisMakananById(id string)(JenisMakanan,error)  {
	var(
		jenisMakanan JenisMakanan
		err error
	)
	err = orm.FindOneByID(&jenisMakanan,id)
	return jenisMakanan,err
}

func FindJenisMakananAl()(Response,error)  {
	var(
		res Response
		jenisMakanan []JenisMakanan
		err error
	)
	err = orm.FindAll(&jenisMakanan)
	if err != nil{
		return res,err
	}
	res.Status= true
	res.Data=jenisMakanan
	res.Message="Sukses Get Data"
	return res,err
}
func FindJenisMakananByIdNative(id string)(JenisMakanan,error)  {
	tx := cgorm.Init().Begin()
	table := "jenis_makanans"
	var(
		jenisMakanan JenisMakanan
		err error
	)
	err = tx.Table(table).Where("id = ?",id).Scan(&jenisMakanan).Error
	if err!=nil{
		return jenisMakanan,err
	}

	return jenisMakanan,err
}