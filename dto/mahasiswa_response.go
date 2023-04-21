package dto

import "time"

type MahasiswaResponse struct {
	ID                int       `json:"id"`
	Nama              string    `json:"nama"`
	Usia              int       `json:"usia"`
	Gender            int       `json:"gender"`
	TanggalRegistrasi time.Time `json:"tanggal_registrasi"`
}
