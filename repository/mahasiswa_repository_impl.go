package repository

import (
	"context"
	"database/sql"
	"errors"
	"jobhun/helper"
	"jobhun/model"
)

type MahasiswaRepositoryImpl struct {
}

func NewMahasiswaRepositoryImpl() *MahasiswaRepositoryImpl {
	return &MahasiswaRepositoryImpl{}
}

func (repository *MahasiswaRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, mahasiswa model.Mahasiswa) model.Mahasiswa {
	SQL := "insert into mahasiswa(nama, usia, gender, tanggal, jurusanId, hobiId) values (?, ?, ?, ?)"
	result, err := tx.ExecContext(ctx, SQL, mahasiswa.Nama, mahasiswa.Usia, mahasiswa.Gender, mahasiswa.TanggalRegistrasi)
	helper.IfError(err)

	id, err := result.LastInsertId()
	helper.IfError(err)

	mahasiswa.ID = int(id)
	return mahasiswa
}

func (repository *MahasiswaRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, mahasiswa model.Mahasiswa) model.Mahasiswa {
	SQL := "update mahasiswa set name = ? where id = ?"
	_, err := tx.ExecContext(ctx, SQL, mahasiswa.Nama, mahasiswa.ID)
	helper.IfError(err)
	return mahasiswa
}

func (repository *MahasiswaRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, mahasiswa model.Mahasiswa) {
	SQL := "delete from mahasiswa where id = ?"
	_, err := tx.ExecContext(ctx, SQL, mahasiswa.ID)
	helper.IfError(err)
}

func (repository *MahasiswaRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, mahasiswaId int) (model.Mahasiswa, error) {
	SQL := "select id, nama, usia, gender, tanggal, jurusanId, hobiId from mahasiswa where id = ?"
	rows, err := tx.QueryContext(ctx, SQL, mahasiswaId)
	helper.IfError(err)

	mahasiswa := model.Mahasiswa{}
	if rows.Next() {
		err := rows.Scan(&mahasiswa.ID, &mahasiswa.Nama)
		helper.IfError(err)
		return mahasiswa, nil
	} else {
		return mahasiswa, errors.New("mahasiswa not found")
	}
}

func (repository *MahasiswaRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []model.Mahasiswa {
	SQL := "select id, nama, usia, gender, tanggal, jurusanId, hobiId from mahasiswa"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.IfError(err)

	var mahasiswas []model.Mahasiswa
	for rows.Next() {
		mahasiswa := model.Mahasiswa{}
		err := rows.Scan(&mahasiswa.ID, &mahasiswa.Nama)
		helper.IfError(err)
		mahasiswas = append(mahasiswas, mahasiswa)
	}
	return mahasiswas
}
