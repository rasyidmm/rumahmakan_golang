package models

type Pembayaran struct {
	BaseModel
	totalPembayaran float64
	Pengunjungs  string `gorm:"REFERENCES:Pengunjung"`
}
