package models

type PesananMakanan struct {
	BaseModel
	jumlahPesanan uint64
	MenuMakanans string `gorm:"REFERENCES:MenuMakanan"`
	Pengunjungs string `gorm:"REFERENCES:Pengunjung"`
}
