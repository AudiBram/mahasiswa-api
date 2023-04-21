package controller

import (
	"github.com/julienschmidt/httprouter"
	"jobhun/dto"
	"jobhun/helper"
	"jobhun/service"
	"net/http"
	"strconv"
)

type MahasiswaControllerImpl struct {
	MahasiswaService service.MahasiswaService
}

func NewMahasiswaControllerImpl(mahasiswaService service.MahasiswaService) *MahasiswaControllerImpl {
	return &MahasiswaControllerImpl{MahasiswaService: mahasiswaService}
}

func (controller *MahasiswaControllerImpl) Create(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	request := dto.MahasiswaCreateRequest{}
	helper.ReadFromRequestBody(r, &request)

	mahasiswaResponse := controller.MahasiswaService.Create(r.Context(), request)
	response := dto.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   mahasiswaResponse,
	}

	helper.WriteToResponseBody(w, response)
}

func (controller *MahasiswaControllerImpl) Update(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	mahasiswaUpdateRequest := dto.MahasiswaUpdateRequest{}
	helper.ReadFromRequestBody(r, &mahasiswaUpdateRequest)

	mahasiswaId := param.ByName("mahasiswaId")
	id, err := strconv.Atoi(mahasiswaId)
	helper.IfError(err)

	mahasiswaUpdateRequest.Id = id

	mahasiswaResponse := controller.MahasiswaService.Update(r.Context(), mahasiswaUpdateRequest)
	webResponse := dto.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   mahasiswaResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (controller *MahasiswaControllerImpl) Delete(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	mahasiswaId := param.ByName("mahasiswaId")
	id, err := strconv.Atoi(mahasiswaId)
	helper.IfError(err)

	controller.MahasiswaService.Delete(r.Context(), id)
	webResponse := dto.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (controller *MahasiswaControllerImpl) FindById(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	mahasiswaId := param.ByName("mahasiswaId")
	id, err := strconv.Atoi(mahasiswaId)
	helper.IfError(err)

	mahasiswaResponse := controller.MahasiswaService.FindById(r.Context(), id)
	webResponse := dto.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   mahasiswaResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (controller *MahasiswaControllerImpl) FindAll(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	all := controller.MahasiswaService.FindAll(r.Context())
	response := dto.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   all,
	}
	helper.WriteToResponseBody(w, response)
}
