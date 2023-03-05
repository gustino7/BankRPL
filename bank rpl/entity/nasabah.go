package entity

type Nasabah struct {
	ID            uint64     `gorm:"primaryKey" json:"id"`
	Nama          string     `json:"Nama"`
	NIK           string     `json:"Nik"`
	Alamat        string     `json:"Alamat"`
	WargaNgr      string     `json:"Kewarganegaraan"`
	Pekerjaan     string     `json:"Pekerjaan"`
	Rekening_nsbh []Rekening `json:"Rekening,omitempty"`
}
