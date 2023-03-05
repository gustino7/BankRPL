package entity

type Rekening struct {
	ID            uint64   `gorm:"primaryKey" json:"id"`
	NomorRekening string   `json:"Nomor Rekening"`
	NasabahID     uint64   `gorm:"foreignKey" json:"NasabahID"`
	Nasabah       *Nasabah `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"transaksi,omitempty"`
}
