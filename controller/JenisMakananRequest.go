package controller

import uuid "github.com/satori/go.uuid"

type JenisMakananUpdateRequest struct {
	Id uuid.UUID `json:"id" form:"id" query:"id"`
	NamaJenis string `json:"nama_jenis" form:"nama_jenis" query:"nama_jenis"`
	DeskripsiJenis string `json:"deskripsi_jenis" form:"deskripsi_jenis" query:"deskripsi_jenis"`
}

//func (r *JenisMakananCreateRequest)bind(c echo.Context, u *models.JenisMakanan)error  {
//	if err := c.Bind(r); err != nil{
//		return err
//	}
//	u.NamaJenis=r.JenisMakanan.NamaJenis
//	u.DeskripsiJenis=r.JenisMakanan.DeskripsiJenis
//	return nil
//}