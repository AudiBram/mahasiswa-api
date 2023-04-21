package helper

import (
	"jobhun/dto"
	"jobhun/model"
	"time"
)

func ToMahasiswaResponse(mahasiswa model.Mahasiswa) dto.MahasiswaResponse {
	return dto.MahasiswaResponse{
		ID:                mahasiswa.ID,
		Nama:              mahasiswa.Nama,
		Usia:              mahasiswa.Usia,
		Gender:            mahasiswa.Gender,
		TanggalRegistrasi: time.Time{},
	}
}

func ToMahasiswaAll(all []model.Mahasiswa) []dto.MahasiswaResponse {
	var mahasiswaResponses []dto.MahasiswaResponse
	for _, mahasiswa := range all {
		mahasiswaResponses = append(mahasiswaResponses, ToMahasiswaResponse(mahasiswa))
	}
	return mahasiswaResponses
}
