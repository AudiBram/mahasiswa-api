package service

import (
	"context"
	"jobhun/dto"
)

type MahasiswaService interface {
	Create(ctx context.Context, request dto.MahasiswaCreateRequest) dto.MahasiswaResponse
	Update(ctx context.Context, request dto.MahasiswaUpdateRequest) dto.MahasiswaResponse
	Delete(ctx context.Context, mahasiswaId int)
	FindById(ctx context.Context, mahasiswaId int) dto.MahasiswaResponse
	FindAll(ctx context.Context) []dto.MahasiswaResponse
}
