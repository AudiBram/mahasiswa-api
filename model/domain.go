package model

import (
	"time"
)

type Mahasiswa struct {
	ID                int       `json:"id"`
	Nama              string    `json:"nama"`
	Usia              int       `json:"usia"`
	Gender            int       `json:"gender"`
	TanggalRegistrasi time.Time `json:"tanggal_registrasi"`
}

type Jurusan struct {
	ID          int    `json:"id"`
	NamaJurusan string `json:"nama_jurusan"`
}

type Hobi struct {
	ID       int    `json:"id"`
	NamaHobi string `json:"nama_hobi"`
}

type MahasiswaHobi struct {
	ID          int `json:"id"`
	MahasiswaID int `json:"mahasiswa_id"`
	HobiID      int `json:"hobi_id"`
}
