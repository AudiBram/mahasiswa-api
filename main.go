package main

import (
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
	"jobhun/app"
	"jobhun/controller"
	"jobhun/exception"
	"jobhun/helper"
	"jobhun/repository"
	"jobhun/service"
	"net/http"
)

func main() {

	db := app.NewDB()
	validate := validator.New()
	mahasiswaRepository := repository.NewMahasiswaRepositoryImpl()
	mahasiswaService := service.NewMahasiswaServiceImpl(mahasiswaRepository, db, validate)
	mahasiswaController := controller.NewMahasiswaControllerImpl(mahasiswaService)

	r := httprouter.New()

	r.GET("/mahasiswa", mahasiswaController.FindAll)
	r.GET("/mahasiswa/:mahasiswaId", mahasiswaController.FindById)
	r.PUT("/mahasiswa/:mahasiswaId", mahasiswaController.Update)
	r.POST("/mahasiswa", mahasiswaController.Create)
	r.DELETE("/mahasiswa/:mahasiswaId", mahasiswaController.Delete)

	r.PanicHandler = exception.ErrorHandler

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: r,
	}

	err := server.ListenAndServe()
	helper.IfError(err)

}
