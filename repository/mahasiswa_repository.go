package repository

import (
	"context"
	"database/sql"
	"jobhun/model"
)

type MahasiswaRepository interface {
	Save(ctx context.Context, tx *sql.Tx, mahasiswa model.Mahasiswa) model.Mahasiswa
	Update(ctx context.Context, tx *sql.Tx, mahasiswa model.Mahasiswa) model.Mahasiswa
	Delete(ctx context.Context, tx *sql.Tx, mahasiswa model.Mahasiswa)
	FindById(ctx context.Context, tx *sql.Tx, mahasiswaId int) (model.Mahasiswa, error)
	FindAll(ctx context.Context, tx *sql.Tx) []model.Mahasiswa
}
