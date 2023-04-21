package service

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
	"jobhun/dto"
	"jobhun/helper"
	"jobhun/model"
	"jobhun/repository"
)

type MahasiswaServiceImpl struct {
	MahasiswaRepository repository.MahasiswaRepository
	DB                  *sql.DB
	Validate            *validator.Validate
}

func NewMahasiswaServiceImpl(mahasiswaRepository repository.MahasiswaRepository, DB *sql.DB, validate *validator.Validate) *MahasiswaServiceImpl {
	return &MahasiswaServiceImpl{
		MahasiswaRepository: mahasiswaRepository,
		DB:                  DB,
		Validate:            validate}
}

func (service *MahasiswaServiceImpl) Create(ctx context.Context, request dto.MahasiswaCreateRequest) dto.MahasiswaResponse {
	err := service.Validate.Struct(request)
	helper.IfError(err)

	tx, err := service.DB.Begin()
	helper.IfError(err)
	defer helper.CommitOrRollback(tx)

	mahasiswa := model.Mahasiswa{
		Nama:              request.Nama,
		Usia:              request.Usia,
		Gender:            request.Gender,
		TanggalRegistrasi: request.TanggalRegistrasi,
	}

	mahasiswa = service.MahasiswaRepository.Save(ctx, tx, mahasiswa)
	return helper.ToMahasiswaResponse(mahasiswa)
}

func (service *MahasiswaServiceImpl) Update(ctx context.Context, request dto.MahasiswaUpdateRequest) dto.MahasiswaResponse {
	err := service.Validate.Struct(request)
	helper.IfError(err)

	tx, err := service.DB.Begin()
	helper.IfError(err)
	defer helper.CommitOrRollback(tx)

	mahasiswa, err := service.MahasiswaRepository.FindById(ctx, tx, request.Id)
	helper.IfError(err)

	mahasiswa.Nama = request.Nama

	mahasiswa = service.MahasiswaRepository.Update(ctx, tx, mahasiswa)

	return helper.ToMahasiswaResponse(mahasiswa)
}

func (service *MahasiswaServiceImpl) Delete(ctx context.Context, mahasiswaId int) {
	tx, err := service.DB.Begin()
	helper.IfError(err)
	defer helper.CommitOrRollback(tx)

	mahasiswa, err := service.MahasiswaRepository.FindById(ctx, tx, mahasiswaId)
	helper.IfError(err)

	service.MahasiswaRepository.Delete(ctx, tx, mahasiswa)
}

func (service *MahasiswaServiceImpl) FindById(ctx context.Context, mahasiswaId int) dto.MahasiswaResponse {
	tx, err := service.DB.Begin()
	helper.IfError(err)
	defer helper.CommitOrRollback(tx)

	mahasiswa, err := service.MahasiswaRepository.FindById(ctx, tx, mahasiswaId)
	helper.IfError(err)

	return helper.ToMahasiswaResponse(mahasiswa)
}

func (service *MahasiswaServiceImpl) FindAll(ctx context.Context) []dto.MahasiswaResponse {
	tx, err := service.DB.Begin()
	helper.IfError(err)
	defer helper.CommitOrRollback(tx)

	all := service.MahasiswaRepository.FindAll(ctx, tx)

	return helper.ToMahasiswaAll(all)
}
