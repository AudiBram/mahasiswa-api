package dto

import "time"

type MahasiswaCreateRequest struct {
	Nama              string    `json:"nama" validate:"required, max=200, min=1"`
	Usia              int       `json:"usia" validate:"required"`
	Gender            int       `json:"gender" validate:"required"`
	TanggalRegistrasi time.Time `json:"tanggal_registrasi" validate:"required"`
}
